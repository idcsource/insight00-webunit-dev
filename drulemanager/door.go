// Copyright 2016-2017
// CoderG the 2016 project
// Insight 0+0 [ 洞悉 0+0 ]
// InDimensions Construct Source [ 忆黛蒙逝·建造源 ] -> idcsource@gmail.com
// Stephen Fire Meditation Qin [ 火志溟 ] -> firemeditation@gmail.com
// Use of this source code is governed by GNU LGPL v3 license

// Drule Manager
//
// 这是一个符合Insight 0+0 webs2规范的站点单元，但不包括js、css、template等静态文件
package drulemanager

import (
	"github.com/idcsource/Insight-0-0-lib/webs2"
)

// drulemanager的入口
type Door struct {
}

func (d *Door) FloorList() (floors webs2.FloorDoor) {
	floors = make(map[string]webs2.FloorInterface)
	floors["login"] = &Login{}
	floors["logindo"] = &LoginDo{}
	floors["logout"] = &LogOut{}
	floors["main"] = &Main{}
	floors["status"] = &Status{}
	floors["statusdo"] = &StatusDo{}
	floors["users"] = &Users{}
	floors["usersdo"] = &UsersDo{}
	floors["areas"] = &Areas{}
	floors["areasdo"] = &AreasDo{}
	floors["remotes"] = &Remotes{}
	floors["remotesdo"] = &RemotesDo{}
	floors["routers"] = &Routers{}
	return
}
