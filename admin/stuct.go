// Copyright 2016-2017
// CoderG the 2016 project
// Insight 0+0 [ 洞悉 0+0 ]
// InDimensions Construct Source [ 忆黛蒙逝·建造源 ]
// Stephen Fire Meditation Qin [ 火志溟 ] -> firemeditation@gmail.com
// Use of this source code is governed by GNU LGPL v3 license

package admin

import (
	"encoding/gob"
	"time"

	"github.com/idcsource/Insight-0-0-lib/roles"
)

// Admin用户
type AdminUser struct {
	roles.Role
	AdminName string // 用户名，加上前缀就是角色的ID
	Password  string // 用户的密码
}

// Admin权限
type AdminAthority struct {
	roles.Role
}

// Admin的运行时结构
type AdminRuntime struct {
	loginlist map[string]*loginadmin
}

// 登录列表
type loginadmin struct {
	adminname  string    // 用户名
	unid       string    // 登录的id
	logintime  time.Time // 登录时间
	activetime time.Time // 活跃时间
	lifetime   int64     // 生存期限
}

// 为Gob注册角色类型
func RegInterfaceForGob() {
	gob.Register(&AdminUser{})
	gob.Register(&AdminAthority{})
}
