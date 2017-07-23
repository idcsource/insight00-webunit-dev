// Copyright 2016-2017
// CoderG the 2016 project
// Insight 0+0 [ 洞悉 0+0 ]
// InDimensions Construct Source [ 忆黛蒙逝·建造源 ] -> idcsource@gmail.com
// Stephen Fire Meditation Qin [ 火志溟 ] -> firemeditation@gmail.com
// Use of this source code is governed by GNU LGPL v3 license

package drulemanager

import (
	"fmt"
	"net/http"

	"github.com/idcsource/insight00-lib/drule2/drule"
	"github.com/idcsource/insight00-lib/drule2/operator"
	"github.com/idcsource/insight00-lib/webs2"
)

type LogOut struct {
	webs2.Floor
}

func (f *LogOut) ExecHTTP() {

	drule_ext, _ := f.B.GetExt("DRule")
	drun := drule_ext.(*drule.DRule)

	selfinfo, err := getUserInfo(drun, f.W, f.R, f.B, f.Rt)
	if err != nil {
		fmt.Fprint(f.W, err)
		return
	}
	o_send := operator.O_OperatorSend{
		User: selfinfo.UserName,
		Unid: selfinfo.Unid,
	}
	errd := drun.UserLogout(&o_send)
	if errd.IsError() != nil {
		fmt.Fprint(f.W, errd.String())
		return
	}

	cookie := &http.Cookie{
		Name:   "DRuleCookie",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(f.W, cookie)
	// 发送登录成功
	fmt.Fprint(f.W, "ok")
}
