// Copyright 2016-2017
// CoderG the 2016 project
// Insight 0+0 [ 洞悉 0+0 ]
// InDimensions Construct Source [ 忆黛蒙逝·建造源 ] -> idcsource@gmail.com
// Stephen Fire Meditation Qin [ 火志溟 ] -> firemeditation@gmail.com
// Use of this source code is governed by GNU LGPL v3 license

package drulemanager

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/idcsource/insight00-lib/drule2/drule"
	"github.com/idcsource/insight00-lib/drule2/operator"
	"github.com/idcsource/insight00-lib/webs2"
)

type RoutersDo struct {
	webs2.Floor
}

func (f *RoutersDo) ExecHTTP() {

	drule_ext, _ := f.B.GetExt("DRule")
	drun := drule_ext.(*drule.DRule)

	if drun.WorkStatus() == true {
		fmt.Fprint(f.W, "you must pause DRule first.")
		return
	}
	if drun.WorkMode() != operator.DRULE_OPERATE_MODE_MASTER {
		fmt.Fprint(f.W, "The working mode does not support this.")
		return
	}

	selfinfo, err := getUserInfo(drun, f.W, f.R, f.B, f.Rt)
	if err != nil {
		fmt.Fprint(f.W, err)
		return
	}
	if selfinfo.Authority != operator.USER_AUTHORITY_ROOT {
		fmt.Fprint(f.W, "You have no authority to do this")
		return
	}

	operatetype, find := f.Rt.UrlRequest["type"]
	if find == false {
		fmt.Fprint(f.W, "url wrong.")
		return
	}
	switch operatetype {
	case "add":
		f.add(&selfinfo, drun)
	case "delete":
		f.del(&selfinfo, drun)
	default:
		fmt.Fprint(f.W, "url wrong.")
		return
	}
	return
}

func (f *RoutersDo) add(selfinfo *UserInfo, drun *drule.DRule) {

	var err error
	defer func() {
		if e := recover(); e != nil {
			fmt.Fprint(f.W, e)
			return
		}
	}()

	// 获取输入
	//	input := pubfunc.NewInputProcessor()
	err = f.R.ParseForm()
	if err != nil {
		fmt.Fprint(f.W, err)
		return
	}
	areaname := f.R.PostForm["areaname"][0]
	if areaname == drule.INSIDE_DMZ {
		fmt.Fprint(f.W, "This area not support to do this set.")
		return
	}

	router_s := operator.O_AreasRouter{
		AreaName: areaname,
	}

	mirror := f.R.PostForm["mirror"][0]
	if mirror == "1" {
		router_s.Mirror = true
		router_s.Mirrors = make([]string, 0)
		for _, m := range f.R.PostForm["mirrors"] {
			if have, _, _ := drun.OperatorExist(m); have == true {
				router_s.Mirrors = append(router_s.Mirrors, m)
			}
		}
	} else {
		router_s.Mirror = false
		ol, errd := drun.OperatorList()
		if errd.IsError() != nil {
			fmt.Fprint(f.W, errd.String())
			return
		}
		router_s.Chars = make(map[string][]string)
		// 查看系统里都有哪些operator
		for i, _ := range ol {
			// 如果这个系统里的operator在from中找到
			if _, ofind := f.R.PostForm[ol[i].Name]; ofind == true {
				oset := f.R.PostForm[ol[i].Name][0]
				oset = strings.TrimSpace(oset)
				if len(oset) != 0 {
					// 分割form的配置，变成一个个字母
					oset_a := strings.Split(oset, ",")
					for _, onechar := range oset_a {
						if len(onechar) != 0 {
							match, err := regexp.MatchString(onechar, `[0123456789abcdef]{1}`)
							if err != nil {
								fmt.Fprint(f.W, err)
							}
							if match == true {
								// 看看routers_s.Chars中有没有这个字母
								if _, find := router_s.Chars[onechar]; find == false {
									router_s.Chars[onechar] = make([]string, 0)
								}
								// 将operator的名字加进去
								router_s.Chars[onechar] = append(router_s.Chars[onechar], ol[i].Name)
							}
						}
					}
				}
			}
		}
	}
	// 执行router的添加
	errd := drun.AreaRouterSet(&router_s)
	if errd.IsError() != nil {
		fmt.Fprint(f.W, errd.String())
		return
	}
	fmt.Fprint(f.W, "ok")
	return
}

func (f *RoutersDo) del(selfinfo *UserInfo, drun *drule.DRule) {
	var err error
	defer func() {
		if e := recover(); e != nil {
			fmt.Fprint(f.W, e)
			return
		}
	}()

	// 获取输入
	//	input := pubfunc.NewInputProcessor()
	err = f.R.ParseForm()
	if err != nil {
		fmt.Fprint(f.W, err)
		return
	}
	name := f.R.PostForm["name"][0]
	errd := drun.AreaRouterDelete(name)
	if errd.IsError() != nil {
		fmt.Fprint(f.W, errd.String())
		return
	}

	fmt.Fprint(f.W, "ok")
	return
}
