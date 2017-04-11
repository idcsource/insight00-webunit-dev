// Copyright 2016-2017
// CoderG the 2016 project
// Insight 0+0 [ 洞悉 0+0 ]
// InDimensions Construct Source [ 忆黛蒙逝·建造源 ] -> idcsource@gmail.com
// Stephen Fire Meditation Qin [ 火志溟 ] -> firemeditation@gmail.com
// Use of this source code is governed by GNU LGPL v3 license

package web_smcs

import (
	"fmt"

	"github.com/idcsource/Insight-0-0-lib/nst"
	"github.com/idcsource/Insight-0-0-lib/smcs2"
	"github.com/idcsource/Insight-0-0-lib/webs2"
)

// 返回节点树和状态，是从smcs2.NodeTree生成的json，url为status_json
type StatusJson struct {
	webs2.Floor
}

func (s *StatusJson) ExecHTTP() {
	// 获取所有节点的列表
	ext_name, err := s.Rt.MyConfig.GetConfig("main.ext_name")
	if err != nil {
		fmt.Fprint(s.W, "Configure error.")
		return
	}
	ext, err := s.B.GetExt(ext_name)
	if err != nil {
		fmt.Fprint(s.W, "Configure error.")
		return
	}
	smcs_runtime := ext.(*smcs2.CenterSmcs)
	nodetree, err := smcs_runtime.GetNodeTree()
	if err != nil {
		fmt.Fprint(s.W, "Configure error.")
		return
	}

	// 转换成为json
	json, err := nst.StructToJson(nodetree)
	if err != nil {
		fmt.Fprint(s.W, "Json err.")
		return
	}
	fmt.Fprint(s.W, json)
}
