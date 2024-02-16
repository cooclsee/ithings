// Code generated by goctl. DO NOT EDIT.
// Source: dm.proto

package remoteconfig

import (
	"context"

	"github.com/i-Things/things/service/dmsvr/internal/svc"
	"github.com/i-Things/things/service/dmsvr/pb/dm"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	CommonSchemaCreateReq              = dm.CommonSchemaCreateReq
	CommonSchemaIndexReq               = dm.CommonSchemaIndexReq
	CommonSchemaIndexResp              = dm.CommonSchemaIndexResp
	CommonSchemaInfo                   = dm.CommonSchemaInfo
	CommonSchemaUpdateReq              = dm.CommonSchemaUpdateReq
	CustomTopic                        = dm.CustomTopic
	DeviceCore                         = dm.DeviceCore
	DeviceCountInfo                    = dm.DeviceCountInfo
	DeviceCountReq                     = dm.DeviceCountReq
	DeviceCountResp                    = dm.DeviceCountResp
	DeviceGatewayBindDevice            = dm.DeviceGatewayBindDevice
	DeviceGatewayIndexReq              = dm.DeviceGatewayIndexReq
	DeviceGatewayIndexResp             = dm.DeviceGatewayIndexResp
	DeviceGatewayMultiCreateReq        = dm.DeviceGatewayMultiCreateReq
	DeviceGatewayMultiDeleteReq        = dm.DeviceGatewayMultiDeleteReq
	DeviceGatewaySign                  = dm.DeviceGatewaySign
	DeviceInfo                         = dm.DeviceInfo
	DeviceInfoCount                    = dm.DeviceInfoCount
	DeviceInfoCountReq                 = dm.DeviceInfoCountReq
	DeviceInfoDeleteReq                = dm.DeviceInfoDeleteReq
	DeviceInfoIndexReq                 = dm.DeviceInfoIndexReq
	DeviceInfoIndexResp                = dm.DeviceInfoIndexResp
	DeviceInfoMultiUpdateReq           = dm.DeviceInfoMultiUpdateReq
	DeviceInfoReadReq                  = dm.DeviceInfoReadReq
	DeviceTypeCountReq                 = dm.DeviceTypeCountReq
	DeviceTypeCountResp                = dm.DeviceTypeCountResp
	DynamicUpgradeJobReq               = dm.DynamicUpgradeJobReq
	EventIndex                         = dm.EventIndex
	EventIndexResp                     = dm.EventIndexResp
	EventLogIndexReq                   = dm.EventLogIndexReq
	Firmware                           = dm.Firmware
	FirmwareFile                       = dm.FirmwareFile
	FirmwareInfo                       = dm.FirmwareInfo
	FirmwareInfoDeleteReq              = dm.FirmwareInfoDeleteReq
	FirmwareInfoDeleteResp             = dm.FirmwareInfoDeleteResp
	FirmwareInfoIndexReq               = dm.FirmwareInfoIndexReq
	FirmwareInfoIndexResp              = dm.FirmwareInfoIndexResp
	FirmwareInfoReadReq                = dm.FirmwareInfoReadReq
	FirmwareInfoReadResp               = dm.FirmwareInfoReadResp
	FirmwareResp                       = dm.FirmwareResp
	GetPropertyReplyReq                = dm.GetPropertyReplyReq
	GetPropertyReplyResp               = dm.GetPropertyReplyResp
	GroupDeviceIndexReq                = dm.GroupDeviceIndexReq
	GroupDeviceIndexResp               = dm.GroupDeviceIndexResp
	GroupDeviceMultiDeleteReq          = dm.GroupDeviceMultiDeleteReq
	GroupDeviceMultiSaveReq            = dm.GroupDeviceMultiSaveReq
	GroupInfo                          = dm.GroupInfo
	GroupInfoCreateReq                 = dm.GroupInfoCreateReq
	GroupInfoIndexReq                  = dm.GroupInfoIndexReq
	GroupInfoIndexResp                 = dm.GroupInfoIndexResp
	GroupInfoUpdateReq                 = dm.GroupInfoUpdateReq
	HubLogIndex                        = dm.HubLogIndex
	HubLogIndexReq                     = dm.HubLogIndexReq
	HubLogIndexResp                    = dm.HubLogIndexResp
	JobReq                             = dm.JobReq
	MultiSendPropertyReq               = dm.MultiSendPropertyReq
	MultiSendPropertyResp              = dm.MultiSendPropertyResp
	OTAModuleDeleteReq                 = dm.OTAModuleDeleteReq
	OTAModuleDetail                    = dm.OTAModuleDetail
	OTAModuleIndexReq                  = dm.OTAModuleIndexReq
	OTAModuleIndexResp                 = dm.OTAModuleIndexResp
	OTAModuleReq                       = dm.OTAModuleReq
	OTAModuleVersionsIndexResp         = dm.OTAModuleVersionsIndexResp
	OTATaskByDeviceCancelReq           = dm.OTATaskByDeviceCancelReq
	OTATaskByDeviceNameReq             = dm.OTATaskByDeviceNameReq
	OTATaskByJobCancelReq              = dm.OTATaskByJobCancelReq
	OTATaskByJobIndexReq               = dm.OTATaskByJobIndexReq
	OTATaskConfirmReq                  = dm.OTATaskConfirmReq
	OTATaskReUpgradeReq                = dm.OTATaskReUpgradeReq
	OTAUnfinishedTaskByDeviceIndexReq  = dm.OTAUnfinishedTaskByDeviceIndexReq
	OTAUnfinishedTaskByDeviceIndexResp = dm.OTAUnfinishedTaskByDeviceIndexResp
	OtaCommonResp                      = dm.OtaCommonResp
	OtaFirmwareCreateReq               = dm.OtaFirmwareCreateReq
	OtaFirmwareDeleteReq               = dm.OtaFirmwareDeleteReq
	OtaFirmwareDeviceInfoReq           = dm.OtaFirmwareDeviceInfoReq
	OtaFirmwareDeviceInfoResp          = dm.OtaFirmwareDeviceInfoResp
	OtaFirmwareFile                    = dm.OtaFirmwareFile
	OtaFirmwareFileIndexReq            = dm.OtaFirmwareFileIndexReq
	OtaFirmwareFileIndexResp           = dm.OtaFirmwareFileIndexResp
	OtaFirmwareFileInfo                = dm.OtaFirmwareFileInfo
	OtaFirmwareFileReq                 = dm.OtaFirmwareFileReq
	OtaFirmwareFileResp                = dm.OtaFirmwareFileResp
	OtaFirmwareIndexReq                = dm.OtaFirmwareIndexReq
	OtaFirmwareIndexResp               = dm.OtaFirmwareIndexResp
	OtaFirmwareInfo                    = dm.OtaFirmwareInfo
	OtaFirmwareReadReq                 = dm.OtaFirmwareReadReq
	OtaFirmwareReadResp                = dm.OtaFirmwareReadResp
	OtaFirmwareResp                    = dm.OtaFirmwareResp
	OtaFirmwareUpdateReq               = dm.OtaFirmwareUpdateReq
	OtaFirmwareVerifyReq               = dm.OtaFirmwareVerifyReq
	OtaJobByDeviceIndexReq             = dm.OtaJobByDeviceIndexReq
	OtaJobByFirmwareIndexReq           = dm.OtaJobByFirmwareIndexReq
	OtaJobInfo                         = dm.OtaJobInfo
	OtaJobInfoIndexResp                = dm.OtaJobInfoIndexResp
	OtaModuleInfo                      = dm.OtaModuleInfo
	OtaPageInfo                        = dm.OtaPageInfo
	OtaPromptIndexReq                  = dm.OtaPromptIndexReq
	OtaPromptIndexResp                 = dm.OtaPromptIndexResp
	OtaTaskBatchReq                    = dm.OtaTaskBatchReq
	OtaTaskBatchResp                   = dm.OtaTaskBatchResp
	OtaTaskByJobIndexResp              = dm.OtaTaskByJobIndexResp
	OtaTaskCancleReq                   = dm.OtaTaskCancleReq
	OtaTaskCreatResp                   = dm.OtaTaskCreatResp
	OtaTaskCreateReq                   = dm.OtaTaskCreateReq
	OtaTaskDeviceCancleReq             = dm.OtaTaskDeviceCancleReq
	OtaTaskDeviceIndexReq              = dm.OtaTaskDeviceIndexReq
	OtaTaskDeviceIndexResp             = dm.OtaTaskDeviceIndexResp
	OtaTaskDeviceInfo                  = dm.OtaTaskDeviceInfo
	OtaTaskDeviceProcessReq            = dm.OtaTaskDeviceProcessReq
	OtaTaskDeviceReadReq               = dm.OtaTaskDeviceReadReq
	OtaTaskIndexReq                    = dm.OtaTaskIndexReq
	OtaTaskIndexResp                   = dm.OtaTaskIndexResp
	OtaTaskInfo                        = dm.OtaTaskInfo
	OtaTaskReadReq                     = dm.OtaTaskReadReq
	OtaTaskReadResp                    = dm.OtaTaskReadResp
	OtaUpTaskInfo                      = dm.OtaUpTaskInfo
	PageInfo                           = dm.PageInfo
	PageInfo_OrderBy                   = dm.PageInfo_OrderBy
	Point                              = dm.Point
	ProductCategory                    = dm.ProductCategory
	ProductCategoryIndexReq            = dm.ProductCategoryIndexReq
	ProductCategoryIndexResp           = dm.ProductCategoryIndexResp
	ProductCustom                      = dm.ProductCustom
	ProductCustomReadReq               = dm.ProductCustomReadReq
	ProductInfo                        = dm.ProductInfo
	ProductInfoDeleteReq               = dm.ProductInfoDeleteReq
	ProductInfoIndexReq                = dm.ProductInfoIndexReq
	ProductInfoIndexResp               = dm.ProductInfoIndexResp
	ProductInfoReadReq                 = dm.ProductInfoReadReq
	ProductRemoteConfig                = dm.ProductRemoteConfig
	ProductSchemaCreateReq             = dm.ProductSchemaCreateReq
	ProductSchemaDeleteReq             = dm.ProductSchemaDeleteReq
	ProductSchemaIndexReq              = dm.ProductSchemaIndexReq
	ProductSchemaIndexResp             = dm.ProductSchemaIndexResp
	ProductSchemaInfo                  = dm.ProductSchemaInfo
	ProductSchemaTslImportReq          = dm.ProductSchemaTslImportReq
	ProductSchemaTslReadReq            = dm.ProductSchemaTslReadReq
	ProductSchemaTslReadResp           = dm.ProductSchemaTslReadResp
	ProductSchemaUpdateReq             = dm.ProductSchemaUpdateReq
	PropertyIndex                      = dm.PropertyIndex
	PropertyIndexResp                  = dm.PropertyIndexResp
	PropertyLatestIndexReq             = dm.PropertyLatestIndexReq
	PropertyLogIndexReq                = dm.PropertyLogIndexReq
	ProtocolInfo                       = dm.ProtocolInfo
	ProtocolInfoIndexReq               = dm.ProtocolInfoIndexReq
	ProtocolInfoIndexResp              = dm.ProtocolInfoIndexResp
	RemoteConfigCreateReq              = dm.RemoteConfigCreateReq
	RemoteConfigIndexReq               = dm.RemoteConfigIndexReq
	RemoteConfigIndexResp              = dm.RemoteConfigIndexResp
	RemoteConfigLastReadReq            = dm.RemoteConfigLastReadReq
	RemoteConfigLastReadResp           = dm.RemoteConfigLastReadResp
	RemoteConfigPushAllReq             = dm.RemoteConfigPushAllReq
	RespActionReq                      = dm.RespActionReq
	RespReadReq                        = dm.RespReadReq
	Response                           = dm.Response
	RootCheckReq                       = dm.RootCheckReq
	SdkLogIndex                        = dm.SdkLogIndex
	SdkLogIndexReq                     = dm.SdkLogIndexReq
	SdkLogIndexResp                    = dm.SdkLogIndexResp
	SendActionReq                      = dm.SendActionReq
	SendActionResp                     = dm.SendActionResp
	SendMsgReq                         = dm.SendMsgReq
	SendMsgResp                        = dm.SendMsgResp
	SendOption                         = dm.SendOption
	SendPropertyMsg                    = dm.SendPropertyMsg
	SendPropertyReq                    = dm.SendPropertyReq
	SendPropertyResp                   = dm.SendPropertyResp
	ShadowIndex                        = dm.ShadowIndex
	ShadowIndexResp                    = dm.ShadowIndexResp
	StaticUpgradeDeviceInfo            = dm.StaticUpgradeDeviceInfo
	StaticUpgradeJobReq                = dm.StaticUpgradeJobReq
	Tag                                = dm.Tag
	TimeRange                          = dm.TimeRange
	UpgradeJobResp                     = dm.UpgradeJobResp
	VerifyOtaFirmwareReq               = dm.VerifyOtaFirmwareReq
	WithID                             = dm.WithID

	RemoteConfig interface {
		RemoteConfigCreate(ctx context.Context, in *RemoteConfigCreateReq, opts ...grpc.CallOption) (*Response, error)
		RemoteConfigIndex(ctx context.Context, in *RemoteConfigIndexReq, opts ...grpc.CallOption) (*RemoteConfigIndexResp, error)
		RemoteConfigPushAll(ctx context.Context, in *RemoteConfigPushAllReq, opts ...grpc.CallOption) (*Response, error)
		RemoteConfigLastRead(ctx context.Context, in *RemoteConfigLastReadReq, opts ...grpc.CallOption) (*RemoteConfigLastReadResp, error)
	}

	defaultRemoteConfig struct {
		cli zrpc.Client
	}

	directRemoteConfig struct {
		svcCtx *svc.ServiceContext
		svr    dm.RemoteConfigServer
	}
)

func NewRemoteConfig(cli zrpc.Client) RemoteConfig {
	return &defaultRemoteConfig{
		cli: cli,
	}
}

func NewDirectRemoteConfig(svcCtx *svc.ServiceContext, svr dm.RemoteConfigServer) RemoteConfig {
	return &directRemoteConfig{
		svr:    svr,
		svcCtx: svcCtx,
	}
}

func (m *defaultRemoteConfig) RemoteConfigCreate(ctx context.Context, in *RemoteConfigCreateReq, opts ...grpc.CallOption) (*Response, error) {
	client := dm.NewRemoteConfigClient(m.cli.Conn())
	return client.RemoteConfigCreate(ctx, in, opts...)
}

func (d *directRemoteConfig) RemoteConfigCreate(ctx context.Context, in *RemoteConfigCreateReq, opts ...grpc.CallOption) (*Response, error) {
	return d.svr.RemoteConfigCreate(ctx, in)
}

func (m *defaultRemoteConfig) RemoteConfigIndex(ctx context.Context, in *RemoteConfigIndexReq, opts ...grpc.CallOption) (*RemoteConfigIndexResp, error) {
	client := dm.NewRemoteConfigClient(m.cli.Conn())
	return client.RemoteConfigIndex(ctx, in, opts...)
}

func (d *directRemoteConfig) RemoteConfigIndex(ctx context.Context, in *RemoteConfigIndexReq, opts ...grpc.CallOption) (*RemoteConfigIndexResp, error) {
	return d.svr.RemoteConfigIndex(ctx, in)
}

func (m *defaultRemoteConfig) RemoteConfigPushAll(ctx context.Context, in *RemoteConfigPushAllReq, opts ...grpc.CallOption) (*Response, error) {
	client := dm.NewRemoteConfigClient(m.cli.Conn())
	return client.RemoteConfigPushAll(ctx, in, opts...)
}

func (d *directRemoteConfig) RemoteConfigPushAll(ctx context.Context, in *RemoteConfigPushAllReq, opts ...grpc.CallOption) (*Response, error) {
	return d.svr.RemoteConfigPushAll(ctx, in)
}

func (m *defaultRemoteConfig) RemoteConfigLastRead(ctx context.Context, in *RemoteConfigLastReadReq, opts ...grpc.CallOption) (*RemoteConfigLastReadResp, error) {
	client := dm.NewRemoteConfigClient(m.cli.Conn())
	return client.RemoteConfigLastRead(ctx, in, opts...)
}

func (d *directRemoteConfig) RemoteConfigLastRead(ctx context.Context, in *RemoteConfigLastReadReq, opts ...grpc.CallOption) (*RemoteConfigLastReadResp, error) {
	return d.svr.RemoteConfigLastRead(ctx, in)
}