// Copyright 2016-2017
// CoderG the 2016 project
// Insight 0+0 [ 洞悉 0+0 ]
// InDimensions Construct Source [ 忆黛蒙逝·建造源 ] -> idcsource@gmail.com
// Stephen Fire Meditation Qin [ 火志溟 ] -> firemeditation@gmail.com
// Use of this source code is governed by GNU LGPL v3 license

package drulemanager

import (
	"fmt"
	"text/template"

	"github.com/idcsource/Insight-0-0-lib/webs2"
)

type Login struct {
	webs2.Floor
}

func (f *Login) ExecHTTP() {
	var err error

	type pageData struct {
		ServerName string
	}

	pagedata := pageData{}

	pagedata.ServerName, err = f.Rt.MyConfig.GetConfig("main.name")
	if err != nil {
		fmt.Println(err)
		return
	}

	templ, err := template.ParseFiles(f.B.GetStaticPath() + "template/login.tmpl")
	if err != nil {
		fmt.Fprint(f.W, err)
		return
	}
	templ.Execute(f.W, pagedata)
}
