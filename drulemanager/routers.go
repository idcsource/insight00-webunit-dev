// Copyright 2016-2017
// CoderG the 2016 project
// Insight 0+0 [ 洞悉 0+0 ]
// InDimensions Construct Source [ 忆黛蒙逝·建造源 ] -> idcsource@gmail.com
// Stephen Fire Meditation Qin [ 火志溟 ] -> firemeditation@gmail.com
// Use of this source code is governed by GNU LGPL v3 license

package drulemanager

import (
	"fmt"
	"strings"
	"text/template"

	"github.com/idcsource/Insight-0-0-lib/drule2/drule"
	"github.com/idcsource/Insight-0-0-lib/drule2/operator"
	"github.com/idcsource/Insight-0-0-lib/iendecode"
	"github.com/idcsource/Insight-0-0-lib/webs2"
)

type Routers struct {
	webs2.Floor
}

func (f *Routers) ExecHTTP() {

	drule_ext, _ := f.B.GetExt("DRule")
	drun := drule_ext.(*drule.DRule)

	userinfo, err := getUserInfo(drun, f.W, f.R, f.B, f.Rt)
	if err != nil {
		return
	}

	type RouterSet struct {
		Set  operator.O_AreasRouter
		Json string
	}
	type pageData struct {
		UserInfo  UserInfo
		IsRoot    bool
		IsWorking bool
		IsMaster  bool
		List      []RouterSet
	}

	page_data := pageData{
		UserInfo: userinfo,
		List:     make([]RouterSet, 0),
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
			list, errd := drun.AreaRouterList()
			if errd.IsError() != nil {
				fmt.Fprint(f.W, errd.String())
				return
			}
			for i, _ := range list {
				rs := RouterSet{
					Set: list[i],
				}
				rs.Json, err = iendecode.StructToJson(list[i])
				if err != nil {
					fmt.Fprint(f.W, err)
					return
				}
				rs.Json = strings.TrimSpace(rs.Json)
				page_data.List = append(page_data.List, rs)
			}
		}
	}

	templ, err := template.ParseFiles(f.B.GetStaticPath() + "template/routers.tmpl")
	if err != nil {
		fmt.Fprint(f.W, err)
		return
	}
	templ.Execute(f.W, page_data)
}
