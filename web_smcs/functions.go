// Copyright 2016-2017
// CoderG the 2016 project
// Insight 0+0 [ 洞悉 0+0 ]
// InDimensions Construct Source [ 忆黛蒙逝·建造源 ] -> idcsource@gmail.com
// Stephen Fire Meditation Qin [ 火志溟 ] -> firemeditation@gmail.com
// Use of this source code is governed by GNU LGPL v3 license

package web_smcs

import (
	"github.com/idcsource/Insight-0-0-lib/webs2"
)

var UnitFields webs2.StaticFields

func init() {
	UnitFields = make(map[string]*webs2.FieldConfig)
	UnitFields["name"] = &webs2.FieldConfig{
		Name:          "name",
		DisName:       "节点名",
		DatabaseField: "name",
		RoleField:     "Name",
		UseIt:         true,
		Type:          webs2.FIELD_FORM_TEXT,
		CanNull:       false,
		Info:          "[英文字母或数字]",
		Min:           1,
		Max:           50,
	}
	UnitFields["code"] = &webs2.FieldConfig{
		Name:          "code",
		DisName:       "身份码",
		DatabaseField: "code",
		RoleField:     "Code",
		UseIt:         true,
		Type:          webs2.FIELD_FORM_TEXT,
		CanNull:       false,
		Info:          "[英文字母或数字]",
		Min:           1,
		Max:           255,
	}
	UnitFields["disname"] = &webs2.FieldConfig{
		Name:          "disname",
		DisName:       "显示名",
		DatabaseField: "disname",
		RoleField:     "DisName",
		UseIt:         true,
		Type:          webs2.FIELD_FORM_TEXT,
		CanNull:       false,
		Info:          "[用来显示的有意义的字符串]",
		Min:           1,
		Max:           50,
	}
	UnitFields["nodetype"] = &webs2.FieldConfig{
		Name:          "nodetype",
		DisName:       "节点类型",
		DatabaseField: "nodetype",
		RoleField:     "NodeType",
		UseIt:         true,
		Type:          webs2.FIELD_FORM_INT,
		CanNull:       false,
		Info:          "[选择这个节点的类型]",
		Min:           0,
		Max:           2,
	}
	UnitFields["groupid"] = &webs2.FieldConfig{
		Name:          "groupid",
		DisName:       "父ID",
		DatabaseField: "groupid",
		RoleField:     "GroupId",
		UseIt:         true,
		Type:          webs2.FIELD_FORM_TEXT,
		CanNull:       false,
		Info:          "[父ID]",
		Min:           1,
		Max:           50,
	}
}
