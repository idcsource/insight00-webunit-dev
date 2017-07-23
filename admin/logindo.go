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

	"github.com/idcsource/insight00-lib/random"
	"github.com/idcsource/insight00-lib/webs2"
)

// 登陆执行，url为logindo
type Logindo struct {
	webs2.Floor
}

func (i *Logindo) ExecHTTP() {
	var err error
	// 获取到trule
	trule, err := i.B.GetTRule2()
	if err != nil {
		fmt.Fprint(i.W, "Role Store Wrong!")
		return
	}

	/* 从配置文件中获取一些东西 */
	// 获取用户名前缀
	admin_prefix, err := i.Rt.MyConfig.GetConfig("main.admin_prefix")
	if err != nil {
		fmt.Fprint(i.W, "Configure error.")
		return
	}
	// 获取cookiename
	cookie_name, err := i.Rt.MyConfig.GetConfig("main.cookie_name")
	if err != nil {
		fmt.Fprint(i.W, "Configure error.")
		return
	}
	// 获取生存期
	lifetime, err := i.Rt.MyConfig.TranInt64("main.lifetime")
	if err != nil {
		lifetime = 60
		err = nil
	}
	// 获取运行时名称
	runtime_name, err := i.Rt.MyConfig.GetConfig("main.ext_name")
	if err != nil {
		fmt.Fprint(i.W, "Configure error.")
		return
	}
	// 获取保存的area
	areaname, err := i.Rt.MyConfig.GetConfig("main.areaname")
	if err != nil {
		fmt.Fprint(i.W, "Configure error.")
		return
	}

	// 获取运行时
	adminruntime_ext, err := i.B.GetExt(runtime_name)
	if err != nil {
		fmt.Fprint(i.W, "Configure error.")
		return
	}
	adminruntime := adminruntime_ext.(*AdminRuntime)

	i.R.ParseForm()
	username := i.R.PostForm["username"][0]
	password := i.R.PostForm["password"][0]
	check := i.R.PostForm["check"][0]

	// 获取保存的密码
	var self_password string
	err = trule.ReadData(areaname, admin_prefix+username, "Password", &self_password)
	if err != nil {
		fmt.Fprint(i.W, "Username or Password wrong.")
		return
	}
	// 对保存密码进行处理，与check合并，并与输入的password进行比对
	self_password = check + self_password
	self_password = random.GetSha1Sum(self_password)
	if self_password != password {
		fmt.Fprint(i.W, "Username or Password wrong.")
		return
	}
	// 如果一切比对正常，则生成随机ID
	unid := random.Unid(1, password, check, username)

	// 建立登录运行时
	loginruntime := &loginadmin{
		adminname:  username,
		unid:       unid,
		logintime:  time.Now(),
		activetime: time.Now(),
		lifetime:   lifetime * 60,
	}
	// 加入运行时
	adminruntime.loginlist[unid] = loginruntime

	// 写入cookie
	cookie := &http.Cookie{
		Name:   cookie_name,
		Value:  unid,
		Path:   "/",
		MaxAge: 0,
	}
	http.SetCookie(i.W, cookie)

	// 发送登录成功
	fmt.Fprint(i.W, "ok")
}
