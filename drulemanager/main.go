// Copyright 2016-2017
// CoderG the 2016 project
// Insight 0+0 [ 洞悉 0+0 ]
// InDimensions Construct Source [ 忆黛蒙逝·建造源 ] -> idcsource@gmail.com
// Stephen Fire Meditation Qin [ 火志溟 ] -> firemeditation@gmail.com
// Use of this source code is governed by GNU LGPL v3 license

package drulemanager

import (
	"fmt"
	"text/template"

	"github.com/idcsource/Insight-0-0-lib/webs2"
)

type Main struct {
	webs2.Floor
}

func (f *Main) ExecHTTP() {
	templ, err := template.ParseFiles(f.B.GetStaticPath() + "template/main.tmpl")
	if err != nil {
		fmt.Println(err)
		return
	}
	templ.Execute(f.W, nil)
}
