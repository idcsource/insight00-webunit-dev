// Copyright 2016-2017
// CoderG the 2016 project
// Insight 0+0 [ 洞悉 0+0 ]
// InDimensions Construct Source [ 忆黛蒙逝·建造源 ] -> idcsource@gmail.com
// Stephen Fire Meditation Qin [ 火志溟 ] -> firemeditation@gmail.com
// Use of this source code is governed by GNU LGPL v3 license

package drulemanager

import (
	"fmt"

	"github.com/idcsource/Insight-0-0-lib/drule2/drule"
	"github.com/idcsource/Insight-0-0-lib/drule2/operator"
	"github.com/idcsource/Insight-0-0-lib/pubfunc"
	"github.com/idcsource/Insight-0-0-lib/webs2"
)

type UsersDo struct {
	webs2.Floor
}

func (f *UsersDo) ExecHTTP() {

	drule_ext, _ := f.B.GetExt("DRule")
	drun := drule_ext.(*drule.DRule)

	selfinfo, err := getUserInfo(drun, f.W, f.R, f.B, f.Rt)
	if err != nil {
		fmt.Fprint(f.W, err)
		return
	}

	operatetype, find := f.Rt.UrlRequest["type"]
	if find == false {
		fmt.Fprint(f.W, "url wrong.")
		return
	}
	switch operatetype {
	case "changeself":
		f.changeself(&selfinfo, drun)
	case "adduser":
		f.adduser(&selfinfo, drun)
	case "edituser":
		f.edituser(&selfinfo, drun)
	case "userdelete":
		f.userdelete(&selfinfo, drun)
	default:
		fmt.Fprint(f.W, "url wrong.")
		return
	}
	return
}

func (f *UsersDo) changeself(selfinfo *UserInfo, drun *drule.DRule) {

	var err error
	defer func() {
		if e := recover(); e != nil {
			fmt.Fprint(f.W, e)
			return
		}
	}()

	// 获取输入
	input := pubfunc.NewInputProcessor()
	err = f.R.ParseForm()
	if err != nil {
		fmt.Fprint(f.W, err)
		return
	}
	email := f.R.PostForm["email"][0]
	password := f.R.PostForm["password"][0]
	password2 := f.R.PostForm["password2"][0]

	password, erri := input.PasswordTwo(password, password2, true)
	if erri != 0 {
		fmt.Fprint(f.W, "please check input.")
		return
	}
	email, erri = input.Email(email, true, 1, 255)
	if erri != 0 {
		fmt.Fprint(f.W, "please check input.")
		return
	}
	userinfo := operator.O_DRuleUser{
		UserName: selfinfo.UserName,
		Email:    email,
		Password: password,
	}

	if len(email) != 0 {
		errd := drun.UserEmail(&userinfo)
		if errd.IsError() != nil {
			fmt.Fprint(f.W, errd.String())
			return
		}
	}
	if len(password) != 0 {
		errd := drun.UserPassword(&userinfo)
		if errd.IsError() != nil {
			fmt.Fprint(f.W, errd.String())
			return
		}
	}
	fmt.Fprint(f.W, "ok")
	return
}

func (f *UsersDo) adduser(selfinfo *UserInfo, drun *drule.DRule) {
	var err error

	if selfinfo.Authority != operator.USER_AUTHORITY_ROOT {
		fmt.Fprint(f.W, "no authority")
		return
	}

	defer func() {
		if e := recover(); e != nil {
			fmt.Fprint(f.W, e)
			return
		}
	}()

	// 获取输入
	input := pubfunc.NewInputProcessor()
	err = f.R.ParseForm()
	if err != nil {
		fmt.Fprint(f.W, err)
		return
	}
	username := f.R.PostForm["username"][0]
	email := f.R.PostForm["email"][0]
	password := f.R.PostForm["password"][0]
	autority := f.R.PostForm["autority"][0]

	var erri int
	username, erri = input.Text(username, false, 2, 50)
	if erri != 0 {
		fmt.Fprint(f.W, "please check input.")
		return
	}
	email, erri = input.Email(email, false, 2, 255)
	if erri != 0 {
		fmt.Fprint(f.W, "please check input.")
		return
	}
	password, erri = input.Password(password, false)
	if erri != 0 {
		fmt.Fprint(f.W, "please check input.")
		return
	}
	var ai operator.UserAuthority = operator.USER_AUTHORITY_NO
	if autority == "1" {
		ai = operator.USER_AUTHORITY_ROOT
	} else if autority == "2" {
		ai = operator.USER_AUTHORITY_DRULE
	} else if autority == "3" {
		ai = operator.USER_AUTHORITY_NORMAL
	}

	userinfo := operator.O_DRuleUser{
		UserName:  username,
		Email:     email,
		Password:  password,
		Authority: ai,
	}

	errd := drun.UserAdd(&userinfo)
	if errd.IsError() != nil {
		fmt.Fprint(f.W, errd.String())
		return
	}
	fmt.Fprint(f.W, "ok")
	return
}

func (f *UsersDo) edituser(selfinfo *UserInfo, drun *drule.DRule) {
	var err error

	if selfinfo.Authority != operator.USER_AUTHORITY_ROOT {
		fmt.Fprint(f.W, "no authority")
		return
	}

	defer func() {
		if e := recover(); e != nil {
			fmt.Fprint(f.W, e)
			return
		}
	}()

	// 获取输入
	input := pubfunc.NewInputProcessor()
	err = f.R.ParseForm()
	if err != nil {
		fmt.Fprint(f.W, err)
		return
	}
	username := f.R.PostForm["username"][0]
	email := f.R.PostForm["email"][0]
	password := f.R.PostForm["password"][0]

	var erri int
	username, erri = input.Text(username, false, 2, 50)
	if erri != 0 {
		fmt.Fprint(f.W, "please check input.")
		return
	}
	email, erri = input.Email(email, false, 2, 255)
	if erri != 0 {
		fmt.Fprint(f.W, "please check input.")
		return
	}
	password, erri = input.Password(password, true)
	if erri != 0 {
		fmt.Fprint(f.W, "please check input.")
		return
	}

	userinfo := operator.O_DRuleUser{
		UserName: username,
		Email:    email,
		Password: password,
	}

	errd := drun.UserEmail(&userinfo)
	if errd.IsError() != nil {
		fmt.Fprint(f.W, errd.String())
		return
	}
	if len(password) != 0 {
		errd = drun.UserPassword(&userinfo)
		if errd.IsError() != nil {
			fmt.Fprint(f.W, errd.String())
			return
		}
	}
	fmt.Fprint(f.W, "ok")
	return
}

func (f *UsersDo) userdelete(selfinfo *UserInfo, drun *drule.DRule) {
	var err error

	if selfinfo.Authority != operator.USER_AUTHORITY_ROOT {
		fmt.Fprint(f.W, "no authority")
		return
	}

	defer func() {
		if e := recover(); e != nil {
			fmt.Fprint(f.W, e)
			return
		}
	}()

	// 获取输入
	input := pubfunc.NewInputProcessor()
	err = f.R.ParseForm()
	if err != nil {
		fmt.Fprint(f.W, err)
		return
	}
	username := f.R.PostForm["username"][0]

	if username == selfinfo.UserName {
		fmt.Fprint(f.W, "can not delete yourself.")
		return
	}

	var erri int
	username, erri = input.Text(username, false, 2, 50)
	if erri != 0 {
		fmt.Fprint(f.W, "please check input.")
		return
	}
	errd := drun.UserDelete(username)
	if errd.IsError() != nil {
		fmt.Fprint(f.W, errd.String())
		return
	}
	fmt.Fprint(f.W, "ok")
	return
}
