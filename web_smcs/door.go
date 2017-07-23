// Copyright 2016-2017
// CoderG the 2016 project
// Insight 0+0 [ 洞悉 0+0 ]
// InDimensions Construct Source [ 忆黛蒙逝·建造源 ] -> idcsource@gmail.com
// Stephen Fire Meditation Qin [ 火志溟 ] -> firemeditation@gmail.com
// Use of this source code is governed by GNU LGPL v3 license

package web_smcs

import (
	"github.com/idcsource/insight00-lib/webs2"
)

// web_smcs组件的入口
type Door struct {
}

func (d *Door) FloorList() (floors webs2.FloorDoor) {
	floors = make(map[string]webs2.FloorInterface)
	floors["status_show"] = &StatusShow{}
	floors["status_json"] = &StatusJson{}
	floors["operator_node"] = &OperatorNode{}
	return
}
