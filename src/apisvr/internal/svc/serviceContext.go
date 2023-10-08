package svc

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/i-Things/things/shared/caches"
	"github.com/i-Things/things/shared/conf"
	"github.com/i-Things/things/shared/oss"
	"github.com/i-Things/things/shared/verify"
	"github.com/i-Things/things/src/apisvr/internal/config"
	"github.com/i-Things/things/src/apisvr/internal/middleware"
	"github.com/i-Things/things/src/disvr/client/deviceinteract"
	"github.com/i-Things/things/src/disvr/client/devicemsg"
	"github.com/i-Things/things/src/disvr/didirect"
	"github.com/i-Things/things/src/dmsvr/client/deviceauth"
	"github.com/i-Things/things/src/dmsvr/client/devicegroup"
	"github.com/i-Things/things/src/dmsvr/client/devicemanage"
	firmwaremanage "github.com/i-Things/things/src/dmsvr/client/firmwaremanage"
	otataskmanage "github.com/i-Things/things/src/dmsvr/client/otataskmanage"
	"github.com/i-Things/things/src/dmsvr/client/productmanage"
	"github.com/i-Things/things/src/dmsvr/client/remoteconfig"
	"github.com/i-Things/things/src/dmsvr/dmdirect"
	alarmcenter "github.com/i-Things/things/src/rulesvr/client/alarmcenter"
	scenelinkage "github.com/i-Things/things/src/rulesvr/client/scenelinkage"
	"github.com/i-Things/things/src/rulesvr/ruledirect"
	api "github.com/i-Things/things/src/syssvr/client/api"
	"github.com/i-Things/things/src/syssvr/client/areamanage"
	common "github.com/i-Things/things/src/syssvr/client/common"
	log "github.com/i-Things/things/src/syssvr/client/log"
	menu "github.com/i-Things/things/src/syssvr/client/menu"
	"github.com/i-Things/things/src/syssvr/client/projectmanage"
	role "github.com/i-Things/things/src/syssvr/client/role"
	user "github.com/i-Things/things/src/syssvr/client/user"
	"github.com/i-Things/things/src/syssvr/sysdirect"
	"github.com/i-Things/things/src/timedjobsvr/client/timedjob"
	"github.com/i-Things/things/src/timedjobsvr/timedjobdirect"
	"github.com/i-Things/things/src/timedschedulersvr/client/timedscheduler"
	"github.com/i-Things/things/src/timedschedulersvr/timedschedulerdirect"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"os"
	"time"
)

func init() {
	jwt.TimeFunc = func() time.Time {
		return time.Now()
	}
}

type SvrClient struct {
	UserRpc user.User
	RoleRpc role.Role
	MenuRpc menu.Menu
	LogRpc  log.Log
	ApiRpc  api.Api

	ProjectM projectmanage.ProjectManage
	AreaM    areamanage.AreaManage
	ProductM productmanage.ProductManage
	DeviceM  devicemanage.DeviceManage
	DeviceA  deviceauth.DeviceAuth
	DeviceG  devicegroup.DeviceGroup

	DeviceMsg      devicemsg.DeviceMsg
	DeviceInteract deviceinteract.DeviceInteract
	RemoteConfig   remoteconfig.RemoteConfig
	Common         common.Common
	Scene          scenelinkage.SceneLinkage
	Alarm          alarmcenter.AlarmCenter
	Timedscheduler timedscheduler.Timedscheduler
	TimedJob       timedjob.TimedJob
}

type ServiceContext struct {
	SvrClient
	Config         config.Config
	SetupWare      rest.Middleware
	CheckTokenWare rest.Middleware
	DataAuthWare   rest.Middleware
	TeardownWare   rest.Middleware
	Captcha        *verify.Captcha
	OssClient      *oss.Client
	FirmwareM      firmwaremanage.FirmwareManage
	OtaTaskM       otataskmanage.OtaTaskManage
	FileChan       chan int64
}

func NewServiceContext(c config.Config) *ServiceContext {
	var (
		projectM projectmanage.ProjectManage
		areaM    areamanage.AreaManage
		productM productmanage.ProductManage
		deviceM  devicemanage.DeviceManage
		deviceA  deviceauth.DeviceAuth
		deviceG  devicegroup.DeviceGroup

		deviceMsg      devicemsg.DeviceMsg
		deviceInteract deviceinteract.DeviceInteract
		remoteConfig   remoteconfig.RemoteConfig
		sysCommon      common.Common
		scene          scenelinkage.SceneLinkage
		alarm          alarmcenter.AlarmCenter
		firmwareM      firmwaremanage.FirmwareManage
		otaTaskM       otataskmanage.OtaTaskManage
		timedSchedule  timedscheduler.Timedscheduler
		timedJob       timedjob.TimedJob
	)
	var ur user.User
	var ro role.Role
	var me menu.Menu
	var lo log.Log
	var ap api.Api

	caches.InitStore(c.CacheRedis)

	//var me menu.Menu
	if c.DmRpc.Enable {
		if c.DmRpc.Mode == conf.ClientModeGrpc { //服务模式
			productM = productmanage.NewProductManage(zrpc.MustNewClient(c.DmRpc.Conf))
			deviceM = devicemanage.NewDeviceManage(zrpc.MustNewClient(c.DmRpc.Conf))
			deviceA = deviceauth.NewDeviceAuth(zrpc.MustNewClient(c.DmRpc.Conf))
			deviceG = devicegroup.NewDeviceGroup(zrpc.MustNewClient(c.DmRpc.Conf))
			remoteConfig = remoteconfig.NewRemoteConfig(zrpc.MustNewClient(c.DmRpc.Conf))
			firmwareM = firmwaremanage.NewFirmwareManage(zrpc.MustNewClient(c.DmRpc.Conf))
			otaTaskM = otataskmanage.NewOtaTaskManage(zrpc.MustNewClient(c.DmRpc.Conf))
		} else { //直连模式
			deviceM = dmdirect.NewDeviceManage(c.DmRpc.RunProxy)
			productM = dmdirect.NewProductManage(c.DmRpc.RunProxy)
			deviceA = dmdirect.NewDeviceAuth(c.DmRpc.RunProxy)
			deviceG = dmdirect.NewDeviceGroup(c.DmRpc.RunProxy)
			remoteConfig = dmdirect.NewRemoteConfig(c.DmRpc.RunProxy)
			firmwareM = dmdirect.NewFirmwareManage(c.DmRpc.RunProxy)
			otaTaskM = dmdirect.NewOtaTaskManage(c.DmRpc.RunProxy)
		}
	}
	if c.RuleRpc.Enable {
		if c.RuleRpc.Mode == conf.ClientModeGrpc {
			scene = scenelinkage.NewSceneLinkage(zrpc.MustNewClient(c.RuleRpc.Conf))
			alarm = alarmcenter.NewAlarmCenter(zrpc.MustNewClient(c.RuleRpc.Conf))
		} else {
			scene = ruledirect.NewSceneLinkage(c.RuleRpc.RunProxy)
			alarm = ruledirect.NewAlarmCenter(c.RuleRpc.RunProxy)
		}
	}
	if c.SysRpc.Enable {
		if c.SysRpc.Mode == conf.ClientModeGrpc {
			projectM = projectmanage.NewProjectManage(zrpc.MustNewClient(c.SysRpc.Conf))
			areaM = areamanage.NewAreaManage(zrpc.MustNewClient(c.SysRpc.Conf))
			ur = user.NewUser(zrpc.MustNewClient(c.SysRpc.Conf))
			ro = role.NewRole(zrpc.MustNewClient(c.SysRpc.Conf))
			me = menu.NewMenu(zrpc.MustNewClient(c.SysRpc.Conf))
			lo = log.NewLog(zrpc.MustNewClient(c.SysRpc.Conf))
			ap = api.NewApi(zrpc.MustNewClient(c.SysRpc.Conf))
			sysCommon = common.NewCommon(zrpc.MustNewClient(c.SysRpc.Conf))
		} else {
			projectM = sysdirect.NewProjectManage(c.SysRpc.RunProxy)
			areaM = sysdirect.NewAreaManage(c.SysRpc.RunProxy)
			ur = sysdirect.NewUser(c.SysRpc.RunProxy)
			ro = sysdirect.NewRole(c.SysRpc.RunProxy)
			me = sysdirect.NewMenu(c.SysRpc.RunProxy)
			lo = sysdirect.NewLog(c.SysRpc.RunProxy)
			ap = sysdirect.NewApi(c.SysRpc.RunProxy)
			sysCommon = sysdirect.NewCommon(c.SysRpc.RunProxy)
		}
	}
	if c.DiRpc.Enable {
		if c.DiRpc.Mode == conf.ClientModeGrpc {
			deviceMsg = devicemsg.NewDeviceMsg(zrpc.MustNewClient(c.DiRpc.Conf))
			deviceInteract = deviceinteract.NewDeviceInteract(zrpc.MustNewClient(c.DiRpc.Conf))

		} else {
			deviceMsg = didirect.NewDeviceMsg(c.DiRpc.RunProxy)
			deviceInteract = didirect.NewDeviceInteract(c.DiRpc.RunProxy)
		}
	}
	if c.TimedSchedulerRpc.Enable {
		if c.TimedSchedulerRpc.Mode == conf.ClientModeGrpc {
			timedSchedule = timedscheduler.NewTimedscheduler(zrpc.MustNewClient(c.TimedSchedulerRpc.Conf))
		} else {
			timedSchedule = timedschedulerdirect.NewScheduler(c.TimedSchedulerRpc.RunProxy)
		}
	}
	if c.TimedJobRpc.Enable {
		if c.TimedJobRpc.Mode == conf.ClientModeGrpc {
			timedJob = timedjob.NewTimedJob(zrpc.MustNewClient(c.TimedJobRpc.Conf))
		} else {
			timedJob = timedjobdirect.NewTimedJob(c.TimedJobRpc.RunProxy)
		}
	}

	ossClient := oss.NewOssClient(c.OssConf)
	if ossClient == nil {
		logx.Error("NewOss err")
		os.Exit(-1)
	}

	captcha := verify.NewCaptcha(c.Captcha.ImgHeight, c.Captcha.ImgWidth,
		c.Captcha.KeyLong, c.CacheRedis, time.Duration(c.Captcha.KeepTime)*time.Second)
	return &ServiceContext{
		Config:         c,
		SetupWare:      middleware.NewSetupWareMiddleware(c, lo).Handle,
		CheckTokenWare: middleware.NewCheckTokenWareMiddleware(c, ur, ro).Handle,
		DataAuthWare:   middleware.NewDataAuthWareMiddleware(c).Handle,
		TeardownWare:   middleware.NewTeardownWareMiddleware(c, lo).Handle,
		Captcha:        captcha,
		OssClient:      ossClient,
		FirmwareM:      firmwareM,
		OtaTaskM:       otaTaskM,
		SvrClient: SvrClient{
			UserRpc:        ur,
			RoleRpc:        ro,
			MenuRpc:        me,
			LogRpc:         lo,
			ApiRpc:         ap,
			Timedscheduler: timedSchedule,
			TimedJob:       timedJob,

			ProjectM: projectM,
			AreaM:    areaM,
			ProductM: productM,
			DeviceM:  deviceM,
			DeviceA:  deviceA,
			DeviceG:  deviceG,

			DeviceMsg:      deviceMsg,
			DeviceInteract: deviceInteract,
			RemoteConfig:   remoteConfig,
			Common:         sysCommon,
			Scene:          scene,
			Alarm:          alarm,
		},
		//OSS:        ossClient,
	}
}