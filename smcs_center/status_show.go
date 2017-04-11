// Copyright 2016-2017
// CoderG the 2016 project
// Insight 0+0 [ 洞悉 0+0 ]
// InDimensions Construct Source [ 忆黛蒙逝·建造源 ] -> idcsource@gmail.com
// Stephen Fire Meditation Qin [ 火志溟 ] -> firemeditation@gmail.com
// Use of this source code is governed by GNU LGPL v3 license

package smcs_center

import (
	"fmt"
	"text/template"

	"github.com/idcsource/Insight-0-0-lib/pubfunc"
	"github.com/idcsource/Insight-0-0-lib/smcs2"
	"github.com/idcsource/Insight-0-0-lib/webs2"
)

type StatusShow struct {
	webs2.Floor
}

func (s *StatusShow) ExecHTTP() {
	/* 判断是否登录开始 */
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
	/* 判断是否登录结束 */

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

	// 模板文件
	template_path, err := s.Rt.MyConfig.GetConfig("main.template_path")
	if err != nil {
		fmt.Fprint(s.W, "Configure error.")
		return
	}

	templ, err := template.ParseFiles(pubfunc.AbsolutePath(template_path, s.B.GetStaticPath()) + "status_show.tmpl")
	if err != nil {
		fmt.Println(err)
		return
	}
	templ.Execute(s.W, nodetree)
}
