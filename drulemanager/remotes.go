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
	"github.com/idcsource/insight00-lib/drule2/operator"
	"github.com/idcsource/insight00-lib/webs2"
)

type Remotes struct {
	webs2.Floor
}

func (f *Remotes) ExecHTTP() {

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
		IsWorking  bool
		IsMaster   bool
		List       []operator.O_DRuleOperator
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
		if drun.WorkMode() == operator.DRULE_OPERATE_MODE_MASTER {
			page_data.IsMaster = true
		} else {
			page_data.IsMaster = false
		}
		page_data.IsWorking = drun.WorkStatus()
		if page_data.IsMaster == true {
			var errd operator.DRuleError
			page_data.List, errd = drun.OperatorList()
			if errd.IsError() != nil {
				fmt.Fprint(f.W, errd.String())
				return
			}
		}
	}

	templ, err := template.ParseFiles(f.B.GetStaticPath() + "template/remotes.tmpl")
	if err != nil {
		fmt.Fprint(f.W, err)
		return
	}
	templ.Execute(f.W, page_data)
}
