// Copyright 2016-2017
// CoderG the 2016 project
// Insight 0+0 [ 洞悉 0+0 ]
// InDimensions Construct Source [ 忆黛蒙逝·建造源 ] -> idcsource@gmail.com
// Stephen Fire Meditation Qin [ 火志溟 ] -> firemeditation@gmail.com
// Use of this source code is governed by GNU LGPL v3 license

package web_smcs

import (
	"fmt"
	"net/http"

	"github.com/idcsource/Insight-0-0-lib/smcs2"
	"github.com/idcsource/Insight-0-0-lib/webs2"
)

// 对节点的添加删除等操作
type OperatorNode struct {
	webs2.Floor
}

type operatorNode_NodeSimpleInfo struct {
	Name    string
	Code    string
	Disname string
	Group   string
	Type    int64
}

func (s *OperatorNode) ExecHTTP() {
	/* 判断是否登录开始 */
	// 获取执行点名称
	point_name, err := s.Rt.MyConfig.GetConfig("main.admin")
	if err != nil {
		fmt.Println(err)
		return
	}
	// 执行执行点
	err = s.B.ExecPoint(point_name, s.W, s.R, s.B, s.Rt)
	if err != nil {
		fmt.Println(err)
		return
	}
	/* 判断是否登录结束 */

	operatetype, find := s.Rt.UrlRequest["type"]
	if find == false {
		fmt.Fprint(s.W, "url wrong.")
		return
	}
	switch operatetype {
	case "addnode":
		s.addnode(s.W, s.R, s.B, s.Rt)
	case "delnode":
		s.delnode(s.W, s.R, s.B, s.Rt)
	}
}

func (s *OperatorNode) addnode(w http.ResponseWriter, r *http.Request, b *webs2.Web, rt webs2.Runtime) {
	// 获取字段的配置文件
	field_config, err := s.Rt.MyConfig.GetSection("nodeinfo_field")
	if err != nil {
		fmt.Fprint(s.W, "url wrong.")
		return
	}

	// 处理POST发送的字段
	fields := webs2.NewFormData(field_config, s.R)
	allfield, check, checks := fields.GetAll(nil)
	if check == false {
		fmt.Fprint(s.W, checks)
		return
	}
	data := operatorNode_NodeSimpleInfo{
		Name:    allfield["name"].String,
		Disname: allfield["disname"].String,
		Code:    allfield["code"].String,
		Group:   allfield["groupid"].String,
		Type:    allfield["nodetype"].Int,
	}

	// 获取SMCS的扩展
	ext_name, err := s.Rt.MyConfig.GetConfig("main.ext_name")
	if err != nil {
		fmt.Fprint(s.W, "Configure error.")
		return
	}
	ext, err := s.B.GetExt(ext_name)
	if err != nil {
		fmt.Fprint(s.W, "Configure error.")
		return
	}
	smcs_runtime := ext.(*smcs2.CenterSmcs)

	// 执行添加
	err = smcs_runtime.AddNode(data.Name, data.Disname, data.Code, uint8(data.Type), data.Group)
	if err != nil {
		fmt.Fprint(s.W, err)
		return
	}
	fmt.Fprint(s.W, "ok")
}

func (s *OperatorNode) delnode(w http.ResponseWriter, r *http.Request, b *webs2.Web, rt webs2.Runtime) {
	// 获取字段的配置文件
	field_config, err := s.Rt.MyConfig.GetSection("nodeinfo_field")
	if err != nil {
		fmt.Fprint(s.W, "url wrong.")
		return
	}

	// 处理POST发送的字段
	fields := webs2.NewFormData(field_config, s.R)
	allfield, check, checks := fields.GetAll([]string{"name"})
	if check == false {
		fmt.Fprint(s.W, checks)
		return
	}
	data := operatorNode_NodeSimpleInfo{
		Name: allfield["name"].String,
	}
	// 获取SMCS的扩展
	ext_name, err := s.Rt.MyConfig.GetConfig("main.ext_name")
	if err != nil {
		fmt.Fprint(s.W, "Configure error.")
		return
	}
	ext, err := s.B.GetExt(ext_name)
	if err != nil {
		fmt.Fprint(s.W, "Configure error.")
		return
	}
	smcs_runtime := ext.(*smcs2.CenterSmcs)

	// 执行删除
	err = smcs_runtime.DelNode(data.Name)
	if err != nil {
		fmt.Fprint(s.W, err)
		return
	}
	fmt.Fprint(s.W, "ok")
}
