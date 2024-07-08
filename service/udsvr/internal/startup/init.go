package startup

import (
	"context"
	"fmt"
	"gitee.com/i-Things/core/service/timed/timedjobsvr/client/timedmanage"
	"gitee.com/i-Things/share/ctxs"
	"gitee.com/i-Things/share/def"
	"gitee.com/i-Things/share/domain/application"
	"gitee.com/i-Things/share/errors"
	"gitee.com/i-Things/share/eventBus"
	"gitee.com/i-Things/share/events/topics"
	"gitee.com/i-Things/share/tools"
	"gitee.com/i-Things/share/utils"
	"github.com/i-Things/things/service/udsvr/internal/event/timerEvent"
	"github.com/i-Things/things/service/udsvr/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
	"sync"
	"time"
)

func Init(svcCtx *svc.ServiceContext) {
	tools.InitStore(svcCtx.Config.CacheRedis)
	TimerInit(svcCtx)
	InitEventBus(svcCtx)
}

func InitEventBus(svcCtx *svc.ServiceContext) {
	//err := svcCtx.FastEvent.QueueSubscribe(eventBus.UdRuleTimer, func(ctx context.Context, t time.Time, body []byte) error {
	//	if t.Add(2 * time.Second).Before(time.Now()) { //2秒过期时间
	//		return nil
	//	}
	//	th := timerEvent.NewSceneHandle(ctx, svcCtx)
	//	return th.DeviceTimer()
	//})
	//logx.Must(err)
	err := svcCtx.FastEvent.QueueSubscribe(eventBus.UdRuleTimer, func(ctx context.Context, t time.Time, body []byte) error {
		if t.Before(time.Now().Add(-time.Second * 2)) { //2秒之前的跳过
			return nil
		}
		th := timerEvent.NewSceneHandle(ctxs.WithRoot(ctx), svcCtx)
		var wait sync.WaitGroup
		wait.Add(2)
		utils.Go(ctx, func() {
			err := th.SceneTiming()
			if err != nil {
				logx.WithContext(ctx).Error(err)
			}
		})
		utils.Go(ctx, func() {
			err := th.DeviceTriggerCheck()
			if err != nil {
				logx.WithContext(ctx).Error(err)
			}
		})
		return nil
	})
	logx.Must(err)
	err = svcCtx.FastEvent.QueueSubscribe(topics.ApplicationDeviceReportThingPropertyAllDevice, func(ctx context.Context, t time.Time, body []byte) error {
		if t.Before(time.Now().Add(-time.Second * 2)) { //2秒之前的跳过
			return nil
		}
		th := timerEvent.NewSceneHandle(ctxs.WithRoot(ctx), svcCtx)
		var stu application.PropertyReport
		err := utils.Unmarshal(body, &stu)
		if err != nil {
			logx.WithContext(ctx).Errorf("Subscribe.QueueSubscribe.Unmarshal body:%v err:%v", string(body), err)
			return err
		}
		return th.SceneThingPropertyReport(stu)
	})
	logx.Must(err)
	err = svcCtx.FastEvent.Start()
	logx.Must(err)
}

func TimerInit(svcCtx *svc.ServiceContext) {
	ctx := context.Background()
	_, err := svcCtx.TimedM.TaskInfoCreate(ctx, &timedmanage.TaskInfo{
		GroupCode: def.TimedIThingsQueueGroupCode,                                   //组编码
		Type:      1,                                                                //任务类型 1 定时任务 2 延时任务
		Name:      "iThings规则引擎定时任务",                                                // 任务名称
		Code:      "iThingsRuleTimer",                                               //任务编码
		Params:    fmt.Sprintf(`{"topic":"%s","payload":""}`, eventBus.UdRuleTimer), // 任务参数,延时任务如果没有传任务参数会拿数据库的参数来执行
		CronExpr:  "@every 1s",                                                      // cron执行表达式
		Status:    def.StatusWaitRun,                                                // 状态
		Priority:  3,                                                                //优先级: 10:critical 最高优先级  3: default 普通优先级 1:low 低优先级
	})
	if err != nil && !errors.Cmp(errors.Fmt(err), errors.Duplicate) {
		logx.Must(err)
	}
}
