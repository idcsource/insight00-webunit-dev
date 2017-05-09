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

	"github.com/idcsource/Insight-0-0-lib/drule2/drule"
	"github.com/idcsource/Insight-0-0-lib/webs2"
)

type LoginDo struct {
	webs2.Floor
}

func (f *LoginDo) ExecHTTP() {

	drule_ext, _ := f.B.GetExt("DRule")
	drun := drule_ext.(*drule.DRule)

	f.R.ParseForm()
	username := f.R.PostForm["username"][0]
	password := f.R.PostForm["password"][0]

	unid, _, errd := drun.UserLogin(username, password)
	if errd.IsError() != nil {
		fmt.Fprint(f.W, errd.String())
		return
	}

	// 写入cookie
	cookie := &http.Cookie{
		Name:   "DruleCookie",
		Value:  unid + "|" + username,
		Path:   "/",
		MaxAge: 0,
	}
	http.SetCookie(f.W, cookie)

	// 发送登录成功
	fmt.Fprint(f.W, "ok")
}
