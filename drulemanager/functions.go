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
	"strings"

	"github.com/idcsource/insight00-lib/drule2/drule"
	"github.com/idcsource/insight00-lib/drule2/operator"
	"github.com/idcsource/insight00-lib/webs2"
)

type UserInfo struct {
	UserName  string
	Unid      string
	Email     string
	Authority operator.UserAuthority
}

func getUserInfo(drun *drule.DRule, w http.ResponseWriter, r *http.Request, b *webs2.Web, rt webs2.Runtime) (userinfo UserInfo, err error) {
	user_cookie, err := r.Cookie("DRuleCookie")
	if err != nil {
		fmt.Println(err)
		http.Redirect(w, r, "/DRuleManager/login", 303)
		return
	}
	cookie_a := strings.Split(user_cookie.Value, "|")
	if len(cookie_a) != 2 {
		http.Redirect(w, r, "/DRuleManager/login", 303)
		err = fmt.Errorf("no login1.")
		fmt.Println(err)
		return
	}
	userinfo = UserInfo{
		UserName: cookie_a[1],
		Unid:     cookie_a[0],
	}
	var login bool
	userinfo.Authority, login = drun.GetUserAuthority(userinfo.UserName, userinfo.Unid)
	if login == false {
		http.Redirect(w, r, "/DRuleManager/login", 303)
		err = fmt.Errorf("no login2.")
		fmt.Println(err)
		return
	}
	useri, errd := drun.UserNow(userinfo.UserName)
	if errd.IsError() != nil {
		http.Redirect(w, r, "/DRuleManager/login", 303)
		err = fmt.Errorf("no login2.")
		fmt.Println(err)
		return
	}
	userinfo.Email = useri.Email
	return
}
