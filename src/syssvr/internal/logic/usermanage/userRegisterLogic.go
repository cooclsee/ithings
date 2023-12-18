package usermanagelogic

import (
	"context"
	"database/sql"
	"github.com/i-Things/things/shared/ctxs"
	"github.com/i-Things/things/shared/def"
	"github.com/i-Things/things/shared/errors"
	"github.com/i-Things/things/shared/stores"
	"github.com/i-Things/things/shared/users"
	"github.com/i-Things/things/shared/utils"
	"github.com/i-Things/things/src/syssvr/internal/repo/relationDB"
	"github.com/i-Things/things/src/syssvr/internal/svc"
	"github.com/i-Things/things/src/syssvr/pb/sys"
	"gorm.io/gorm"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserRegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	uc *ctxs.UserCtx
}

func NewUserRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRegisterLogic {
	return &UserRegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
		uc:     ctxs.GetUserCtx(ctx),
	}
}

func (l *UserRegisterLogic) UserRegister(in *sys.UserRegisterReq) (*sys.UserRegisterResp, error) {
	switch in.RegType {
	case users.RegWxMiniP:
		return l.handleWxminip(in)
	case users.RegEmail, users.RegPhone:
		return l.handleEmailOrPhone(in)
	default:
		return nil, errors.NotRealize.AddMsgf(in.RegType)
	}
	return &sys.UserRegisterResp{}, nil
}

func (l *UserRegisterLogic) handleEmailOrPhone(in *sys.UserRegisterReq) (*sys.UserRegisterResp, error) {
	ui := relationDB.SysUserInfo{
		UserID: l.svcCtx.UserID.GetSnowflakeId(),
	}
	ui.Password = utils.MakePwd(in.Info.Password, ui.UserID, false)

	switch in.RegType {
	case users.RegEmail:
		email := l.svcCtx.Captcha.Verify(l.ctx, def.CaptchaTypeEmail, in.CodeID, in.Code)
		if email == "" || email != in.Account {
			return nil, errors.Captcha
		}
		ui.Email = utils.AnyToNullString(in.Account)
	}
	err := CheckPwd(l.svcCtx, in.Password)
	if err != nil {
		return nil, err
	}
	conn := stores.GetTenantConn(l.ctx)
	err = l.FillUserInfo(&ui, conn)
	if err != nil {
		return nil, err
	}
	return &sys.UserRegisterResp{
		UserID: ui.UserID,
	}, nil
}

func (l *UserRegisterLogic) handleWxminip(in *sys.UserRegisterReq) (*sys.UserRegisterResp, error) {
	auth := l.svcCtx.WxMiniProgram.GetAuth()
	ret, err := auth.Code2SessionContext(l.ctx, in.Code)
	if err != nil {
		l.Errorf("%v.Code2SessionContext err:%v", err)
		if ret.ErrCode != 0 {
			return nil, errors.System.AddDetail(ret.ErrMsg)
		}
		return nil, errors.System.AddDetail(err)
	} else if ret.ErrCode != 0 {
		return nil, errors.Parameter.AddDetail(ret.ErrMsg)
	}
	userID := l.svcCtx.UserID.GetSnowflakeId()

	conn := stores.GetTenantConn(l.ctx)
	err = conn.Transaction(func(tx *gorm.DB) error {
		uidb := relationDB.NewUserInfoRepo(tx)
		_, err = uidb.FindOneByFilter(l.ctx, relationDB.UserInfoFilter{WechatUnionID: ret.UnionID})
		if err == nil { //已经注册过
			return errors.DuplicateRegister
		}
		if !errors.Cmp(err, errors.NotFind) {
			return err
		}
		ui := relationDB.SysUserInfo{
			UserID:        userID,
			WechatUnionID: sql.NullString{Valid: true, String: ret.UnionID},
		}
		err = l.FillUserInfo(&ui, tx)
		return err
	})

	return &sys.UserRegisterResp{UserID: userID}, err
}
func (l *UserRegisterLogic) FillUserInfo(in *relationDB.SysUserInfo, tx *gorm.DB) error {
	err := tx.Transaction(func(tx *gorm.DB) error {
		cfg, err := relationDB.NewTenantConfigRepo(tx).FindOne(l.ctx)
		if err != nil {
			return err
		}
		in.RegIP = l.uc.IP
		in.Role = cfg.RegisterRoleID
		uidb := relationDB.NewUserInfoRepo(tx)
		err = uidb.Insert(l.ctx, in)
		if err != nil {
			return err
		}
		err = relationDB.NewUserRoleRepo(tx).Insert(l.ctx, &relationDB.SysUserRole{
			UserID: in.UserID,
			RoleID: cfg.RegisterRoleID,
		})
		return err
	})
	return err
}