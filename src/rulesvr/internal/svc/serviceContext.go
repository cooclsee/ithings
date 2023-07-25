package svc

import (
	"context"
	"github.com/i-Things/things/shared/conf"
	"github.com/i-Things/things/shared/domain/schema"
	"github.com/i-Things/things/shared/eventBus"
	"github.com/i-Things/things/shared/stores"
	deviceinteract "github.com/i-Things/things/src/disvr/client/deviceinteract"
	devicemsg "github.com/i-Things/things/src/disvr/client/devicemsg"
	"github.com/i-Things/things/src/disvr/didirect"
	devicemanage "github.com/i-Things/things/src/dmsvr/client/devicemanage"
	productmanage "github.com/i-Things/things/src/dmsvr/client/productmanage"
	"github.com/i-Things/things/src/dmsvr/dmdirect"
	"github.com/i-Things/things/src/dmsvr/pb/dm"
	"github.com/i-Things/things/src/rulesvr/internal/config"
	"github.com/i-Things/things/src/rulesvr/internal/domain/scene"
	"github.com/i-Things/things/src/rulesvr/internal/repo/cache"
	"github.com/i-Things/things/src/rulesvr/internal/repo/event/dataUpdate"
	"github.com/i-Things/things/src/rulesvr/internal/repo/relationDB"
	"github.com/i-Things/things/src/rulesvr/internal/timer"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/kv"
	"github.com/zeromicro/go-zero/zrpc"
	"os"
)

type ServiceContext struct {
	Config config.Config
	Repo
	SvrClient
	SceneTimerControl timer.SceneControl
	Bus               eventBus.Bus
	DataUpdate        dataUpdate.DataUpdate
}
type Repo struct {
	Store           kv.Store
	SceneDeviceRepo scene.DeviceRepo
	SchemaRepo      schema.ReadRepo
}
type SvrClient struct {
	ProductM       productmanage.ProductManage
	DeviceInteract deviceinteract.DeviceInteract
	DeviceMsg      devicemsg.DeviceMsg
	DeviceM        devicemanage.DeviceManage
}

func NewServiceContext(c config.Config) *ServiceContext {
	var (
		deviceM        devicemanage.DeviceManage
		productM       productmanage.ProductManage
		deviceInteract deviceinteract.DeviceInteract
		deviceMsg      devicemsg.DeviceMsg
	)
	stores.InitConn(c.Database)

	// 自动迁移数据库
	db := stores.GetCommonConn(context.Background())
	errdb := db.AutoMigrate(&relationDB.RuleAlarmLog{})
	if errdb != nil {
		logx.Error("failed to migrate database: %v", errdb)
		os.Exit(-1)
	}
	errdb = db.AutoMigrate(&relationDB.RuleAlarmDealRecord{})
	if errdb != nil {
		logx.Error("failed to migrate database: %v", errdb)
		os.Exit(-1)
	}
	errdb = db.AutoMigrate(&relationDB.RuleAlarmRecord{})
	if errdb != nil {
		logx.Error("failed to migrate database: %v", errdb)
		os.Exit(-1)
	}
	errdb = db.AutoMigrate(&relationDB.RuleAlarmInfo{})
	if errdb != nil {
		logx.Error("failed to migrate database: %v", errdb)
		os.Exit(-1)
	}
	errdb = db.AutoMigrate(&relationDB.RuleAlarmScene{})
	if errdb != nil {
		logx.Error("failed to migrate database: %v", errdb)
		os.Exit(-1)
	}
	errdb = db.AutoMigrate(&relationDB.RuleSceneInfo{})
	if errdb != nil {
		logx.Error("failed to migrate database: %v", errdb)
		os.Exit(-1)
	}

	logx.Info("NewPubDev db.AutoMigrate!")

	store := kv.NewStore(c.CacheRedis)
	sceneDevice := cache.NewSceneDeviceRepo(relationDB.NewSceneInfoRepo(context.TODO()))
	err := sceneDevice.Init(context.TODO())
	if err != nil {
		logx.Error("设备场景数据初始化失败 err:", err)
		os.Exit(-1)
	}
	if c.DmRpc.Mode == conf.ClientModeGrpc {
		productM = productmanage.NewProductManage(zrpc.MustNewClient(c.DmRpc.Conf))
		deviceM = devicemanage.NewDeviceManage(zrpc.MustNewClient(c.DmRpc.Conf))
	} else {
		productM = dmdirect.NewProductManage(c.DmRpc.RunProxy)
		deviceM = dmdirect.NewDeviceManage(c.DmRpc.RunProxy)
	}

	tr := schema.NewReadRepo(func(ctx context.Context, productID string) (*schema.Model, error) {
		info, err := productM.ProductSchemaTslRead(ctx, &dm.ProductSchemaTslReadReq{ProductID: productID})
		if err != nil {
			return nil, err
		}
		return schema.ValidateWithFmt([]byte(info.Tsl))
	})
	if c.DiRpc.Mode == conf.ClientModeGrpc {
		deviceMsg = devicemsg.NewDeviceMsg(zrpc.MustNewClient(c.DiRpc.Conf))
		deviceInteract = deviceinteract.NewDeviceInteract(zrpc.MustNewClient(c.DiRpc.Conf))
	} else {
		deviceMsg = didirect.NewDeviceMsg(c.DiRpc.RunProxy)
		deviceInteract = didirect.NewDeviceInteract(c.DiRpc.RunProxy)
	}
	bus := eventBus.NewEventBus()
	du, err := dataUpdate.NewDataUpdate(c.Event)
	logx.Must(err)
	return &ServiceContext{
		Bus:        bus,
		Config:     c,
		DataUpdate: du,
		SvrClient: SvrClient{
			ProductM:       productM,
			DeviceInteract: deviceInteract,
			DeviceMsg:      deviceMsg,
			DeviceM:        deviceM,
		},
		Repo: Repo{
			Store:           store,
			SceneDeviceRepo: sceneDevice,
			SchemaRepo:      tr,
		},
	}
}
