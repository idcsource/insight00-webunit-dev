// Copyright 2016-2017
// CoderG the 2016 project
// Insight 0+0 [ 洞悉 0+0 ]
// InDimensions Construct Source [ 忆黛蒙逝·建造源 ]
// Stephen Fire Meditation Qin [ 火志溟 ] -> firemeditation@gmail.com
// Use of this source code is governed by GNU LGPL v3 license

package admin

import (
	"fmt"
	"text/template"

	"github.com/idcsource/Insight-0-0-lib/pubfunc"
	"github.com/idcsource/Insight-0-0-lib/webs2"
)

type Index struct {
	webs2.Floor
}

func (i *Index) ExecHTTP() {
	err := CheckLogin(i.W, i.R, i.B, i.Rt)
	if err != nil {
		fmt.Println(err)
		return
	}
	template_path, err := i.Rt.MyConfig.GetConfig("main.template_path")
	if err != nil {
		fmt.Fprint(i.W, "Configure error.")
		return
	}

	templ, err := template.ParseFiles(pubfunc.AbsolutePath(template_path, i.B.GetStaticPath()) + "admin.tmpl")
	if err != nil {
		fmt.Println(err)
		return
	}
	templ.Execute(i.W, nil)
}
