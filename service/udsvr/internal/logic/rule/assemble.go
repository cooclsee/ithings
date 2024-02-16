package rulelogic

import (
	"gitee.com/i-Things/share/utils"
	"github.com/i-Things/things/service/udsvr/internal/domain/scene"
	"github.com/i-Things/things/service/udsvr/internal/repo/relationDB"
	"github.com/i-Things/things/service/udsvr/pb/ud"
)

func ToSceneInfoDo(in *ud.SceneInfo) *scene.Info {
	if in == nil {
		return nil
	}

	return &scene.Info{
		ID:      in.Id,
		Name:    in.Name,
		HeadImg: in.HeadImg,
		Tag:     in.Tag,
		Desc:    in.Desc,
		AreaIDs: in.AreaIDs,
		Trigger: utils.UnmarshalNoErr[scene.Trigger](in.Trigger),
		When:    utils.UnmarshalNoErr[scene.When](in.When),
		Then:    utils.UnmarshalNoErr[scene.Then](in.Then),
		Status:  in.Status,
	}
}

func ToSceneInfoPo(in *scene.Info) *relationDB.UdSceneInfo {
	return &relationDB.UdSceneInfo{
		ID:      in.ID,
		AreaIDs: in.AreaIDs,
		Name:    in.Name,
		Desc:    in.Desc,
		Tag:     in.Tag,
		HeadImg: in.HeadImg,
		UdSceneTrigger: relationDB.UdSceneTrigger{
			Type:    string(in.Trigger.Type),
			Devices: in.Trigger.Devices,
			Timers:  in.Trigger.Timers,
		},
		UdSceneWhen: relationDB.UdSceneWhen{
			ValidRanges:   in.When.ValidRanges,
			InvalidRanges: in.When.InvalidRanges,
			Conditions:    in.When.Conditions,
		},
		UdSceneThen: relationDB.UdSceneThen{Actions: in.Then.Actions},
	}
}

func PoToSceneInfoDo(in *relationDB.UdSceneInfo) *scene.Info {
	if in == nil {
		return nil
	}
	return &scene.Info{
		ID:          in.ID,
		AreaIDs:     in.AreaIDs,
		Name:        in.Name,
		Tag:         in.Tag,
		HeadImg:     in.HeadImg,
		Desc:        in.Desc,
		CreatedTime: in.CreatedTime,
		Trigger: scene.Trigger{
			Type:    scene.TriggerType(in.UdSceneTrigger.Type),
			Devices: in.UdSceneTrigger.Devices,
			Timers:  in.UdSceneTrigger.Timers,
		},
		When: scene.When{
			ValidRanges:   in.UdSceneWhen.ValidRanges,
			InvalidRanges: in.UdSceneWhen.InvalidRanges,
			Conditions:    in.UdSceneWhen.Conditions,
		},
		Then: scene.Then{
			Actions: in.UdSceneThen.Actions,
		},
		Status: in.Status,
	}
}

func PoToSceneInfoPb(in *relationDB.UdSceneInfo) *ud.SceneInfo {
	if in == nil {
		return nil
	}
	do := PoToSceneInfoDo(in)
	return &ud.SceneInfo{
		Id:      in.ID,
		Name:    in.Name,
		Desc:    in.Desc,
		Tag:     in.Tag,
		HeadImg: in.HeadImg,
		AreaIDs: in.AreaIDs,
		Trigger: utils.MarshalNoErr(do.Trigger),
		When:    utils.MarshalNoErr(do.When),
		Then:    utils.MarshalNoErr(do.Then),
		Status:  in.Status,
	}
}
func PoToSceneInfoPbs(in []*relationDB.UdSceneInfo) (ret []*ud.SceneInfo) {
	if in == nil {
		return nil
	}
	for _, v := range in {
		ret = append(ret, PoToSceneInfoPb(v))
	}
	return ret
}