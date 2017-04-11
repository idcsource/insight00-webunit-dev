// Copyright 2016-2017
// CoderG the 2016 project
// Insight 0+0 [ 洞悉 0+0 ]
// InDimensions Construct Source [ 忆黛蒙逝·建造源 ]
// Stephen Fire Meditation Qin [ 火志溟 ] -> firemeditation@gmail.com
// Use of this source code is governed by GNU LGPL v3 license

package admin

import (
	"fmt"
	"net/http"
	"time"

	"github.com/idcsource/Insight-0-0-lib/cpool"
	"github.com/idcsource/Insight-0-0-lib/webs2"
)

// 运行时生成，用这个注册web的扩展
func RuntimeInit(config *cpool.Block) (admin *AdminRuntime) {
	admin = &AdminRuntime{
		loginlist: make(map[string]*loginadmin),
		config:    config,
	}
	return
}

// 检查登录状态，如果没有登录，则会直接将页面跳转回login，返回的错误处理直接返回调用
func (admin *AdminRuntime) ExecPoint(w http.ResponseWriter, r *http.Request, b *webs2.Web, rt webs2.Runtime) (err error) {
	/* 从配置文件中获取一些东西 */
	// 获取cookiename
	cookie_name, err := admin.config.GetConfig("main.cookie_name")
	if err != nil {
		return
	}
	// 获取登录地址
	login_url, err := admin.config.GetConfig("main.login_url")
	if err != nil {
		return
	}

	// 获取cookie
	user_cookie, err := r.Cookie(cookie_name)
	if err != nil {
		http.Redirect(w, r, login_url, 303)
		return
	}
	unid := user_cookie.Value

	userlogin, found := admin.loginlist[unid]
	if found == false {
		http.Redirect(w, r, login_url, 303)
		err = fmt.Errorf("not login")
		return
	}
	// 看活跃期
	if userlogin.activetime.Unix()+userlogin.lifetime > time.Now().Unix() {
		userlogin.activetime = time.Now()
	} else {
		delete(admin.loginlist, unid)
		http.Redirect(w, r, login_url, 303)
		err = fmt.Errorf("not login")
		return
	}
	return
}
