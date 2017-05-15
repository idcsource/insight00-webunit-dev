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

type Users struct {
	webs2.Floor
}

func (f *Users) ExecHTTP() {

	drule_ext, _ := f.B.GetExt("DRule")
	drun := drule_ext.(*drule.DRule)

	userinfo, err := getUserInfo(drun, f.W, f.R, f.B, f.Rt)
	if err != nil {
		return
	}

	type pageData struct {
		ServerName string
		UserInfo   UserInfo
		IsRoot     bool
		UserList   []operator.O_DRuleUser
	}

	page_data := pageData{
		UserInfo: userinfo,
	}

	page_data.ServerName, err = f.Rt.MyConfig.GetConfig("main.name")
	if err != nil {
		fmt.Println(err)
		return
	}

	if userinfo.Authority != operator.USER_AUTHORITY_ROOT {
		page_data.IsRoot = false
	} else {
		page_data.IsRoot = true
		var errd operator.DRuleError
		page_data.UserList, errd = drun.UserList()
		if errd.IsError() != nil {
			fmt.Fprint(f.W, errd.String())
			return
		}
	}

	templ, err := template.ParseFiles(f.B.GetStaticPath() + "template/users.tmpl")
	if err != nil {
		fmt.Fprint(f.W, err)
		return
	}
	templ.Execute(f.W, page_data)
}
