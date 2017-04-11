// Copyright 2016-2017
// CoderG the 2016 project
// Insight 0+0 [ 洞悉 0+0 ]
// InDimensions Construct Source [ 忆黛蒙逝·建造源 ] -> idcsource@gmail.com
// Stephen Fire Meditation Qin [ 火志溟 ] -> firemeditation@gmail.com
// Use of this source code is governed by GNU LGPL v3 license

// 这是符合Insight 0+0项目webs2包标准的Web Unit组件。
//
// 本组件提供基本的管理员后台功能。
//
// 本组件使用Insight-0-0-lib中的*drule.TRule作为信息的存储。
//
// 本组件并没有设计成单独使用，需要配合其他Web Unit组件，并为Insight 0+0的节点监控中心功能进行了定制。
//
// 按照webs2标准，本组件接收*cpool.Block配置信息，例子如下：
//
//	{admindoor}
//
//	[main]
//		#admin用户名的前缀
//		admin_prefix = Admin_
//
//		#admin扩展的名称
//		ext_name = AdminRuntime
//
//		#用户权限的名称
//		admin_athority = AdminAdmin_Athority
//
//		#活跃生存期，单位分钟
//		lifetime = 60
//
//		#Cookie的名称
//		cookie_name = insight_admin
//
//		# 模板所在路径
//		template_path = template/admin/
//
//		# 登录的入口
//		login_url = /Admin/login
//
// 其中，admin_prefix和admin_athority在初次设置之后就不要再改变，否则会造成信息失效。
//
// 模板所在路径template_path，如果是相对路径，则是相对于webs2配置中的静态路径。
//
// 下面是使用方法：
//
// 1、注册运行时扩展，webs2提供了RegExt方法进行扩展注册，使用admin的RuntimeInit方法则可以进行这个注册，类似于：
//
//	web.RegExt(ext_name, admin.RuntimeInit())
//
// 其中的ext_name必须与配置信息中的ext_name统一。
//
// 2、将AdminDoor注册进webs2的路由管理中，并加入配置信息，代码类似于：
//
//	route_tree.AddDoor("管理平台", "Admin", &admin.AdminDoor{}, admin_config)
//
// 其中的"Admin"与配置中的login_url中的“Admin”对应。
//
// 经过上述步骤，你访问/Admin的时候，就可以自动跳转到登录界面/Admin/login，并进行登录操作。
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
