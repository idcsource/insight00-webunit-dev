// Copyright 2016-2017
// CoderG the 2016 project
// Insight 0+0 [ 洞悉 0+0 ]
// InDimensions Construct Source [ 忆黛蒙逝·建造源 ]
// Stephen Fire Meditation Qin [ 火志溟 ] -> firemeditation@gmail.com
// Use of this source code is governed by GNU LGPL v3 license

package admin

import (
	"fmt"
	"text/template"

	"github.com/idcsource/Insight-0-0-lib/pubfunc"
	"github.com/idcsource/Insight-0-0-lib/webs2"
)

// 登陆界面，url为login
type LoginFloor struct {
	webs2.Floor
}

func (i *LoginFloor) ExecHTTP() {

	type returndata struct {
		Title string
	}

	/* 从配置文件中获取一些东西 */
	template_path, err := i.Rt.MyConfig.GetConfig("main.template_path")
	if err != nil {
		fmt.Fprint(i.W, "Configure error.")
		return
	}

	s1, err := template.ParseFiles(pubfunc.AbsolutePath(template_path, i.B.GetStaticPath()) + "login.tmpl")
	if err != nil {
		fmt.Println(err)
		return
	}
	s1.Execute(i.W, returndata{Title: "系统管理平台登录界面"})
}
