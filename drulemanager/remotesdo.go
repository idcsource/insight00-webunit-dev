// Copyright 2016-2017
// CoderG the 2016 project
// Insight 0+0 [ 洞悉 0+0 ]
// InDimensions Construct Source [ 忆黛蒙逝·建造源 ] -> idcsource@gmail.com
// Stephen Fire Meditation Qin [ 火志溟 ] -> firemeditation@gmail.com
// Use of this source code is governed by GNU LGPL v3 license

package drulemanager

import (
	"fmt"
	"strconv"

	"github.com/idcsource/Insight-0-0-lib/drule2/drule"
	"github.com/idcsource/Insight-0-0-lib/drule2/operator"
	"github.com/idcsource/Insight-0-0-lib/webs2"
)

type RemotesDo struct {
	webs2.Floor
}

func (f *RemotesDo) ExecHTTP() {

	drule_ext, _ := f.B.GetExt("DRule")
	drun := drule_ext.(*drule.DRule)

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
	case "addremote":
		f.addremote(&selfinfo, drun)
	case "delete":
		f.deleteremote(&selfinfo, drun)
	default:
		fmt.Fprint(f.W, "url wrong.")
		return
	}
	return
}
func (f *RemotesDo) deleteremote(selfinfo *UserInfo, drun *drule.DRule) {
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
	errd := drun.OperatorDelete(name)
	if errd.IsError() != nil {
		fmt.Fprint(f.W, errd.String())
		return
	}

	fmt.Fprint(f.W, "ok")
	return
}

func (f *RemotesDo) addremote(selfinfo *UserInfo, drun *drule.DRule) {

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
	address := f.R.PostForm["address"][0]
	connum := f.R.PostForm["connum"][0]
	username := f.R.PostForm["username"][0]
	password := f.R.PostForm["password"][0]
	tls := f.R.PostForm["tls"][0]

	var tls_t bool
	if tls == "1" {
		tls_t = true
	} else {
		tls_t = false
	}
	connum_i, err := strconv.Atoi(connum)
	if err != nil {
		connum_i = 4
	}

	op := operator.O_DRuleOperator{
		Name:     name,
		Address:  address,
		ConnNum:  connum_i,
		TLS:      tls_t,
		Username: username,
		Password: password,
	}
	errd := drun.OperatorSet(&op)
	if errd.IsError() != nil {
		fmt.Fprint(f.W, errd.String())
		return
	}

	fmt.Fprint(f.W, "ok")
	return
}
