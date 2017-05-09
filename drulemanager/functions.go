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

	"github.com/idcsource/Insight-0-0-lib/drule2/drule"
	"github.com/idcsource/Insight-0-0-lib/drule2/operator"
	"github.com/idcsource/Insight-0-0-lib/webs2"
)

func getUserInfo(drun *drule.DRule, w http.ResponseWriter, r *http.Request, b *webs2.Web, rt webs2.Runtime) (username, unid string, auth operator.UserAuthority, err error) {
	user_cookie, err := r.Cookie("DruleCookie")
	if err != nil {
		fmt.Println(err)
		http.Redirect(w, r, "/DruleManager/login", 303)
		return
	}
	cookie_a := strings.Split(user_cookie.Value, "|")
	if len(cookie_a) != 2 {
		http.Redirect(w, r, "/DruleManager/login", 303)
		err = fmt.Errorf("no login1.")
		fmt.Println(err)
		return
	}
	unid = cookie_a[0]
	username = cookie_a[1]

	var login bool
	auth, login = drun.GetUserAuthority(username, unid)
	if login == false {
		http.Redirect(w, r, "/DruleManager/login", 303)
		err = fmt.Errorf("no login2.")
		fmt.Println(err)
		return
	}
	return
}
