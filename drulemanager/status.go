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

	"github.com/idcsource/insight00-lib/drule2/drule"
	"github.com/idcsource/insight00-lib/webs2"
)

type Status struct {
	webs2.Floor
}

func (f *Status) ExecHTTP() {

	drule_ext, _ := f.B.GetExt("DRule")
	drun := drule_ext.(*drule.DRule)

	userinfo, err := getUserInfo(drun, f.W, f.R, f.B, f.Rt)
	if err != nil {
		return
	}

	type pageData struct {
		UserInfo   UserInfo
		WorkStatus bool
		ServerName string
	}

	page_data := pageData{
		UserInfo: userinfo,
	}
	page_data.WorkStatus = drun.WorkStatus()
	page_data.ServerName, err = f.Rt.MyConfig.GetConfig("main.name")
	if err != nil {
		fmt.Println(err)
		return
	}

	templ, err := template.ParseFiles(f.B.GetStaticPath() + "template/status.tmpl")
	if err != nil {
		fmt.Println(err)
		return
	}
	templ.Execute(f.W, page_data)
}
