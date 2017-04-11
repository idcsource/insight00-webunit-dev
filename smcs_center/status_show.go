// Copyright 2016-2017
// CoderG the 2016 project
// Insight 0+0 [ 洞悉 0+0 ]
// InDimensions Construct Source [ 忆黛蒙逝·建造源 ] -> idcsource@gmail.com
// Stephen Fire Meditation Qin [ 火志溟 ] -> firemeditation@gmail.com
// Use of this source code is governed by GNU LGPL v3 license

package smcs_center

import (
	"fmt"

	"github.com/idcsource/Insight-0-0-lib/webs2"
)

type StatusShow struct {
	webs2.Floor
}

func (s *StatusShow) ExecHTTP() {
	// 获取执行点名称
	point_name, err := s.Rt.MyConfig.GetConfig("main.admin")
	if err != nil {
		fmt.Println(err)
		return
	}
	// 执行执行点
	err = s.B.ExecPoint(point_name, s.W, s.R, s.B, s.Rt)
	if err != nil {
		fmt.Println(err)
		return
	}
}
