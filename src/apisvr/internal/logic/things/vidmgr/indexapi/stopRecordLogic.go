package indexapi

import (
	"context"
	"encoding/json"

	"github.com/i-Things/things/src/apisvr/internal/svc"
	"github.com/i-Things/things/src/apisvr/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type StopRecordLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewStopRecordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StopRecordLogic {
	return &StopRecordLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *StopRecordLogic) StopRecord(req *types.IndexApiReq) (resp *types.IndexApiStopRecordResp, err error) {
	// todo: add your logic here and delete this line
	bytetmp := make([]byte, 0)
	data, err := proxySetMediaServer(l.ctx, l.svcCtx, STOPRECORD, req.VidmgrID, bytetmp)
	dataRecv := new(types.IndexApiStopRecordResp)
	json.Unmarshal(data, dataRecv)
	return dataRecv, err
}
