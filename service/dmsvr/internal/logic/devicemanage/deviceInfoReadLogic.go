package devicemanagelogic

import (
	"context"
	"gitee.com/i-Things/share/def"
	"gitee.com/i-Things/share/devices"
	"github.com/i-Things/things/service/dmsvr/internal/logic"
	"github.com/i-Things/things/service/dmsvr/internal/repo/relationDB"
	"github.com/i-Things/things/service/dmsvr/internal/svc"
	"github.com/i-Things/things/service/dmsvr/pb/dm"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeviceInfoReadLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	DiDB *relationDB.DeviceInfoRepo
}

func NewDeviceInfoReadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeviceInfoReadLogic {
	return &DeviceInfoReadLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
		DiDB:   relationDB.NewDeviceInfoRepo(ctx),
	}
}

// 获取设备信息详情
func (l *DeviceInfoReadLogic) DeviceInfoRead(in *dm.DeviceInfoReadReq) (*dm.DeviceInfo, error) {
	di, err := l.DiDB.FindOneByFilter(l.ctx,
		relationDB.DeviceFilter{ProductID: in.ProductID, DeviceNames: []string{in.DeviceName},
			SharedType: def.SelectTypeAll, SharedDevices: []*devices.Core{{ProductID: in.ProductID, DeviceName: in.DeviceName}},
			WithProduct: true, WithManufacturer: in.WithManufacturer})
	if err != nil {
		return nil, err
	}
	return logic.ToDeviceInfo(l.ctx, di, l.svcCtx.ProductCache), nil
}
