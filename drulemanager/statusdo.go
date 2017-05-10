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
	"github.com/idcsource/Insight-0-0-lib/webs2"
)

type StatusDo struct {
	webs2.Floor
}

func (f *StatusDo) ExecHTTP() {

	drule_ext, _ := f.B.GetExt("DRule")
	drun := drule_ext.(*drule.DRule)

	userinfo, err := getUserInfo(drun, f.W, f.R, f.B, f.Rt)
	if err != nil {
		return
	}
	if userinfo.Authority != operator.USER_AUTHORITY_ROOT {
		fmt.Fprint(f.W, "You have no authority to do this.")
		return
	}

	workstatus := drun.WorkStatus()

	if workstatus == true {
		drun.Pause()
	} else {
		err = drun.Start()
		if err != nil {
			fmt.Fprint(f.W, err)
			return
		}
	}
	fmt.Fprint(f.W, "ok")
	return
}
