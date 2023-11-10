// Code generated by goctl. DO NOT EDIT.
// Source: dm.proto

package server

import (
	"context"

	"github.com/i-Things/things/src/dmsvr/internal/logic/otataskmanage"
	"github.com/i-Things/things/src/dmsvr/internal/svc"
	"github.com/i-Things/things/src/dmsvr/pb/dm"
)

type OtaTaskManageServer struct {
	svcCtx *svc.ServiceContext
	dm.UnimplementedOtaTaskManageServer
}

func NewOtaTaskManageServer(svcCtx *svc.ServiceContext) *OtaTaskManageServer {
	return &OtaTaskManageServer{
		svcCtx: svcCtx,
	}
}

// 创建批量升级任务
func (s *OtaTaskManageServer) OtaTaskCreate(ctx context.Context, in *dm.OtaTaskCreateReq) (*dm.OtaTaskCreatResp, error) {
	l := otataskmanagelogic.NewOtaTaskCreateLogic(ctx, s.svcCtx)
	return l.OtaTaskCreate(in)
}

func (s *OtaTaskManageServer) OtaTaskUpdate(ctx context.Context, in *dm.OtaTaskInfo) (*dm.OtaCommonResp, error) {
	l := otataskmanagelogic.NewOtaTaskUpdateLogic(ctx, s.svcCtx)
	return l.OtaTaskUpdate(in)
}

// 批量取消升级任务
func (s *OtaTaskManageServer) OtaTaskCancle(ctx context.Context, in *dm.OtaTaskCancleReq) (*dm.OtaCommonResp, error) {
	l := otataskmanagelogic.NewOtaTaskCancleLogic(ctx, s.svcCtx)
	return l.OtaTaskCancle(in)
}

func (s *OtaTaskManageServer) OtaTaskIndex(ctx context.Context, in *dm.OtaTaskIndexReq) (*dm.OtaTaskIndexResp, error) {
	l := otataskmanagelogic.NewOtaTaskIndexLogic(ctx, s.svcCtx)
	return l.OtaTaskIndex(in)
}

// 升级任务详情
func (s *OtaTaskManageServer) OtaTaskRead(ctx context.Context, in *dm.OtaTaskReadReq) (*dm.OtaTaskReadResp, error) {
	l := otataskmanagelogic.NewOtaTaskReadLogic(ctx, s.svcCtx)
	return l.OtaTaskRead(in)
}

// 升级批次详情列表
func (s *OtaTaskManageServer) OtaTaskDeviceIndex(ctx context.Context, in *dm.OtaTaskDeviceIndexReq) (*dm.OtaTaskDeviceIndexResp, error) {
	l := otataskmanagelogic.NewOtaTaskDeviceIndexLogic(ctx, s.svcCtx)
	return l.OtaTaskDeviceIndex(in)
}

// 设备升级状态详情
func (s *OtaTaskManageServer) OtaTaskDeviceRead(ctx context.Context, in *dm.OtaTaskDeviceReadReq) (*dm.OtaTaskDeviceInfo, error) {
	l := otataskmanagelogic.NewOtaTaskDeviceReadLogic(ctx, s.svcCtx)
	return l.OtaTaskDeviceRead(in)
}

// 获取当前可执行批次信息
func (s *OtaTaskManageServer) OtaTaskDeviceEnableBatch(ctx context.Context, in *dm.OtaTaskBatchReq) (*dm.OtaTaskBatchResp, error) {
	l := otataskmanagelogic.NewOtaTaskDeviceEnableBatchLogic(ctx, s.svcCtx)
	return l.OtaTaskDeviceEnableBatch(in)
}

// 升级进度上报
func (s *OtaTaskManageServer) OtaTaskDeviceProcess(ctx context.Context, in *dm.OtaTaskDeviceProcessReq) (*dm.OtaCommonResp, error) {
	l := otataskmanagelogic.NewOtaTaskDeviceProcessLogic(ctx, s.svcCtx)
	return l.OtaTaskDeviceProcess(in)
}

// 取消单个设备的升级
func (s *OtaTaskManageServer) OtaTaskDeviceCancle(ctx context.Context, in *dm.OtaTaskDeviceCancleReq) (*dm.OtaCommonResp, error) {
	l := otataskmanagelogic.NewOtaTaskDeviceCancleLogic(ctx, s.svcCtx)
	return l.OtaTaskDeviceCancle(in)
}
