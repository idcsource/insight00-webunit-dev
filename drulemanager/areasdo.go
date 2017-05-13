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
	"github.com/idcsource/Insight-0-0-lib/iendecode"
	"github.com/idcsource/Insight-0-0-lib/pubfunc"
	"github.com/idcsource/Insight-0-0-lib/webs2"
)

type AreasDo struct {
	webs2.Floor
}

func (f *AreasDo) ExecHTTP() {

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
	case "addarea":
		f.addarea(&selfinfo, drun)
	case "delete":
		f.delarea(&selfinfo, drun)
	case "rename":
		f.rename(&selfinfo, drun)
	case "list":
		f.listareas(&selfinfo, drun)
	default:
		fmt.Fprint(f.W, "url wrong.")
		return
	}
	return
}

func (f *AreasDo) addarea(selfinfo *UserInfo, drun *drule.DRule) {
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
	areaname := f.R.PostForm["areaname"][0]

	var erri int
	areaname, erri = input.Mark(areaname, false, 1, 50)
	if erri != 0 {
		fmt.Fprint(f.W, "please check input.")
		return
	}

	errd := drun.AreaAdd(areaname)
	if errd.IsError() != nil {
		fmt.Fprint(f.W, errd.String())
		return
	}
	fmt.Fprint(f.W, "ok")
	return
}

func (f *AreasDo) delarea(selfinfo *UserInfo, drun *drule.DRule) {
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
	areaname := f.R.PostForm["areaname"][0]

	var erri int
	areaname, erri = input.Mark(areaname, false, 1, 50)
	if erri != 0 {
		fmt.Fprint(f.W, "please check input.")
		return
	}

	errd := drun.AreaDelete(areaname)
	if errd.IsError() != nil {
		fmt.Fprint(f.W, errd.String())
		return
	}
	fmt.Fprint(f.W, "ok")
	return
}

func (f *AreasDo) rename(selfinfo *UserInfo, drun *drule.DRule) {
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
	areaname := f.R.PostForm["areaname"][0]
	oldname := f.R.PostForm["oldname"][0]

	var erri int
	areaname, erri = input.Mark(areaname, false, 1, 50)
	if erri != 0 {
		fmt.Fprint(f.W, "please check input.")
		return
	}
	oldname, erri = input.Mark(oldname, false, 1, 50)
	if erri != 0 {
		fmt.Fprint(f.W, "please check input.")
		return
	}

	errd := drun.AreaRename(oldname, areaname)
	if errd.IsError() != nil {
		fmt.Fprint(f.W, errd.String())
		return
	}
	fmt.Fprint(f.W, "ok")
	return
}

func (f *AreasDo) listareas(selfinfo *UserInfo, drun *drule.DRule) {

	if selfinfo.Authority != operator.USER_AUTHORITY_ROOT {
		fmt.Fprint(f.W, "no authority")
		return
	}

	list, errd := drun.AreaList()
	if errd.IsError() != nil {
		fmt.Fprint(f.W, errd.String())
		return
	}
	list_j, err := iendecode.StructToJson(list)
	if err != nil {
		fmt.Fprint(f.W, err)
		return
	}
	fmt.Fprint(f.W, list_j)
	return
}
