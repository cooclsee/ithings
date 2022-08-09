package product

import (
	"context"
	"github.com/i-Things/things/shared/errors"
	"github.com/i-Things/things/shared/utils"
	"github.com/i-Things/things/src/dmsvr/pb/dm"

	"github.com/i-Things/things/src/apisvr/internal/svc"
	"github.com/i-Things/things/src/apisvr/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type InfoIndexLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewInfoIndexLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InfoIndexLogic {
	return &InfoIndexLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *InfoIndexLogic) InfoIndex(req *types.ProductInfoIndexReq) (resp *types.ProductInfoIndexResp, err error) {
	dmReq := &dm.ProductInfoIndexReq{
		DeviceType:  req.DeviceType, //产品id
		ProductName: req.ProductName,
		Page: &dm.PageInfo{
			Page: req.Page.Page,
			Size: req.Page.Size,
		},
	}
	dmResp, err := l.svcCtx.DmRpc.ProductInfoIndex(l.ctx, dmReq)
	if err != nil {
		er := errors.Fmt(err)
		l.Errorf("%s|rpc.GetDeviceInfo|req=%v|err=%+v", utils.FuncName(), req, er)
		return nil, er
	}
	pis := make([]*types.ProductInfo, 0, len(dmResp.List))
	for _, v := range dmResp.List {
		pi := productInfoToApi(v)
		pis = append(pis, pi)
	}
	return &types.ProductInfoIndexResp{
		Total: dmResp.Total,
		List:  pis,
		Num:   int64(len(pis)),
	}, nil

}
