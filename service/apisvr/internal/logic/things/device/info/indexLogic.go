package info

import (
	"context"
	"gitee.com/i-Things/share/ctxs"
	"gitee.com/i-Things/share/errors"
	"gitee.com/i-Things/share/utils"
	"github.com/i-Things/things/service/apisvr/internal/logic"
	"github.com/i-Things/things/service/apisvr/internal/logic/things"
	"github.com/i-Things/things/service/dmsvr/pb/dm"
	"sync"

	"github.com/i-Things/things/service/apisvr/internal/svc"
	"github.com/i-Things/things/service/apisvr/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type IndexLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewIndexLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IndexLogic {
	return &IndexLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctxs.WithDefaultRoot(ctx),
		svcCtx: svcCtx,
	}
}

func (l *IndexLogic) Index(req *types.DeviceInfoIndexReq) (resp *types.DeviceInfoIndexResp, err error) {
	dmReq := &dm.DeviceInfoIndexReq{
		TenantCode:        req.TenantCode,
		ProductID:         req.ProductID, //产品id
		ProductIDs:        req.ProductIDs,
		AreaIDs:           req.AreaIDs, //项目区域ids
		DeviceName:        req.DeviceName,
		Tags:              logic.ToTagsMap(req.Tags),
		Page:              logic.ToDmPageRpc(req.Page),
		Range:             req.Range,
		Position:          logic.ToDmPointRpc(req.Position),
		DeviceAlias:       req.DeviceAlias,
		IsOnline:          req.IsOnline,
		ProductCategoryID: req.ProductCategoryID,
		WithShared:        req.WithShared,
		WithCollect:       req.WithCollect,
		Versions:          req.Versions,
		Gateway:           utils.Copy[dm.DeviceCore](req.Gateway),
		GroupID:           req.GroupID,
		Devices:           utils.CopySlice[dm.DeviceCore](req.Devices),
		NotGroupID:        req.NotGroupID,
		DeviceTypes:       req.DeviceTypes,
		Status:            req.Status,
		DeviceNames:       req.DeviceNames,
		NotAreaID:         req.NotAreaID,
		HasOwner:          req.HasOwner,
		UserID:            req.UserID,
	}
	dmResp, err := l.svcCtx.DeviceM.DeviceInfoIndex(l.ctx, dmReq)
	if err != nil {
		er := errors.Fmt(err)
		l.Errorf("%s.rpc.GetDeviceInfo req=%v err=%+v", utils.FuncName(), req, er)
		return nil, er
	}
	pis := make([]*types.DeviceInfo, 0, len(dmResp.List))
	var piMap = map[int64]*types.DeviceInfo{}
	wait := sync.WaitGroup{}
	mutex := sync.Mutex{}
	for _, v := range dmResp.List {
		wait.Add(1)
		info := v
		utils.Go(l.ctx, func() {
			defer wait.Done()
			pi := things.InfoToApi(l.ctx, l.svcCtx, info, req.WithProperties, req.WithProfiles, req.WithOwner)
			mutex.Lock()
			defer mutex.Unlock()
			piMap[pi.ID] = pi
		})
	}
	wait.Wait()
	for _, v := range dmResp.List {
		pis = append(pis, piMap[v.Id])
	}
	return &types.DeviceInfoIndexResp{
		Total: dmResp.Total,
		List:  pis,
	}, nil
}
