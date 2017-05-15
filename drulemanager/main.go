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

	"github.com/idcsource/Insight-0-0-lib/drule2/drule"
	"github.com/idcsource/Insight-0-0-lib/drule2/operator"
	"github.com/idcsource/Insight-0-0-lib/webs2"
)

type Main struct {
	webs2.Floor
}

func (f *Main) ExecHTTP() {

	drule_ext, _ := f.B.GetExt("DRule")
	drun := drule_ext.(*drule.DRule)

	userinfo, err := getUserInfo(drun, f.W, f.R, f.B, f.Rt)
	if err != nil {
		return
	}

	type pageData struct {
		ServerName string
		UserInfo   UserInfo
		Authority  string
		WorkStatus bool
		WorkMode   operator.DRuleOperateMode
	}

	page_data := pageData{
		UserInfo: userinfo,
	}

	page_data.ServerName, err = f.Rt.MyConfig.GetConfig("main.name")
	if err != nil {
		fmt.Println(err)
		return
	}

	page_data.WorkStatus = drun.WorkStatus()
	page_data.WorkMode = drun.WorkMode()
	if userinfo.Authority == operator.USER_AUTHORITY_DRULE {
		page_data.Authority = "DRule"
	} else if userinfo.Authority == operator.USER_AUTHORITY_ROOT {
		page_data.Authority = "Root"
	} else {
		page_data.Authority = "Normal"
	}

	templ, err := template.ParseFiles(f.B.GetStaticPath() + "template/main.tmpl")
	if err != nil {
		fmt.Println(err)
		return
	}
	templ.Execute(f.W, page_data)
}
