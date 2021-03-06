// Code generated by protoc-gen-bm v0.1, DO NOT EDIT.
// source: api/http/v1/titans.proto

/*
Package v1 is a generated blademaster stub package.
This code was generated with go-common/app/tool/bmgen/protoc-gen-bm v0.1.

It is generated from these files:
	api/http/v1/titans.proto
*/
package v1

import (
	"context"

	bm "go-common/library/net/http/blademaster"
	"go-common/library/net/http/blademaster/binding"
)

// to suppressed 'imported but not used warning'
var _ *bm.Context
var _ context.Context
var _ binding.StructValidator

// ================
// Titans Interface
// ================

type Titans interface {
	// 获取配置 请求参数逗号隔开的字符串 返回`internal:"true"`
	GetMultiConfigs(ctx context.Context, req *MultiConfigReq) (resp *MultiConfigResp, err error)

	// 获取服务级配置 `internal:"true"`
	GetServiceConfig(ctx context.Context, req *ServiceConfigReq) (resp *ServiceConfigResp, err error)

	// 插入服务配置 `method:"POST" internal:"true"`
	SetServiceConfig(ctx context.Context, req *SetReq) (resp *SetResp, err error)

	// 获取服务级配置 `internal:"true"`
	GetServiceConfigList(ctx context.Context, req *ServiceListReq) (resp *ServiceListResp, err error)

	// 获取已配置的discoveryId `internal:"true"`
	GetTreeIds(ctx context.Context, req *TreeIdsReq) (resp *TreeIdsResp, err error)

	// 获取运营数据列表 `internal:"true"`
	GetEasyList(ctx context.Context, req *EasyGetReq) (resp *EasyGetResp, err error)

	// 设置运营列表 `method:"POST" internal:"true"`
	SetEasyList(ctx context.Context, req *EasySetReq) (resp *EasySetResp, err error)
}

var v1TitansSvc Titans

// @params MultiConfigReq
// @router GET /xlive/internal/resource/v1/titans/getMultiConfigs
// @response MultiConfigResp
func titansGetMultiConfigs(c *bm.Context) {
	p := new(MultiConfigReq)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := v1TitansSvc.GetMultiConfigs(c, p)
	c.JSON(resp, err)
}

// @params ServiceConfigReq
// @router GET /xlive/internal/resource/v1/titans/getServiceConfig
// @response ServiceConfigResp
func titansGetServiceConfig(c *bm.Context) {
	p := new(ServiceConfigReq)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := v1TitansSvc.GetServiceConfig(c, p)
	c.JSON(resp, err)
}

// @params SetReq
// @router POST /xlive/internal/resource/v1/titans/setServiceConfig
// @response SetResp
func titansSetServiceConfig(c *bm.Context) {
	p := new(SetReq)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := v1TitansSvc.SetServiceConfig(c, p)
	c.JSON(resp, err)
}

// @params ServiceListReq
// @router GET /xlive/internal/resource/v1/titans/getServiceConfigList
// @response ServiceListResp
func titansGetServiceConfigList(c *bm.Context) {
	p := new(ServiceListReq)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := v1TitansSvc.GetServiceConfigList(c, p)
	c.JSON(resp, err)
}

// @params TreeIdsReq
// @router GET /xlive/internal/resource/v1/titans/getTreeIds
// @response TreeIdsResp
func titansGetTreeIds(c *bm.Context) {
	p := new(TreeIdsReq)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := v1TitansSvc.GetTreeIds(c, p)
	c.JSON(resp, err)
}

// @params EasyGetReq
// @router GET /xlive/internal/resource/v1/titans/getEasyList
// @response EasyGetResp
func titansGetEasyList(c *bm.Context) {
	p := new(EasyGetReq)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := v1TitansSvc.GetEasyList(c, p)
	c.JSON(resp, err)
}

// @params EasySetReq
// @router POST /xlive/internal/resource/v1/titans/setEasyList
// @response EasySetResp
func titansSetEasyList(c *bm.Context) {
	p := new(EasySetReq)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := v1TitansSvc.SetEasyList(c, p)
	c.JSON(resp, err)
}

// RegisterV1TitansService Register the blademaster route with middleware map
// midMap is the middleware map, the key is defined in proto
func RegisterV1TitansService(e *bm.Engine, svc Titans, midMap map[string]bm.HandlerFunc) {
	v1TitansSvc = svc
	e.GET("/xlive/internal/resource/v1/titans/getMultiConfigs", titansGetMultiConfigs)
	e.GET("/xlive/internal/resource/v1/titans/getServiceConfig", titansGetServiceConfig)
	e.POST("/xlive/internal/resource/v1/titans/setServiceConfig", titansSetServiceConfig)
	e.GET("/xlive/internal/resource/v1/titans/getServiceConfigList", titansGetServiceConfigList)
	e.GET("/xlive/internal/resource/v1/titans/getTreeIds", titansGetTreeIds)
	e.GET("/xlive/internal/resource/v1/titans/getEasyList", titansGetEasyList)
	e.POST("/xlive/internal/resource/v1/titans/setEasyList", titansSetEasyList)
}
