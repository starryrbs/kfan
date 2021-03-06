// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// protoc-gen-go-http v2.1.3

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

type AccountHTTPServer interface {
	CreateAccount(context.Context, *CreateAccountRequest) (*CreateAccountReply, error)
	GetAccount(context.Context, *GetAccountRequest) (*GetAccountReply, error)
	ListAccount(context.Context, *ListAccountRequest) (*ListAccountReply, error)
	Login(context.Context, *LoginRequest) (*LoginReply, error)
}

func RegisterAccountHTTPServer(s *http.Server, srv AccountHTTPServer) {
	r := s.Route("/")
	r.POST("/v1/accounts", _Account_CreateAccount0_HTTP_Handler(srv))
	r.GET("/v1/accounts/{id}", _Account_GetAccount0_HTTP_Handler(srv))
	r.GET("/v1/accounts", _Account_ListAccount0_HTTP_Handler(srv))
	r.POST("/v1/login", _Account_Login0_HTTP_Handler(srv))
}

func _Account_CreateAccount0_HTTP_Handler(srv AccountHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateAccountRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.account.service.v1.Account/CreateAccount")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateAccount(ctx, req.(*CreateAccountRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreateAccountReply)
		return ctx.Result(200, reply)
	}
}

func _Account_GetAccount0_HTTP_Handler(srv AccountHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetAccountRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.account.service.v1.Account/GetAccount")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetAccount(ctx, req.(*GetAccountRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetAccountReply)
		return ctx.Result(200, reply)
	}
}

func _Account_ListAccount0_HTTP_Handler(srv AccountHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListAccountRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.account.service.v1.Account/ListAccount")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListAccount(ctx, req.(*ListAccountRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListAccountReply)
		return ctx.Result(200, reply)
	}
}

func _Account_Login0_HTTP_Handler(srv AccountHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in LoginRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.account.service.v1.Account/Login")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Login(ctx, req.(*LoginRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*LoginReply)
		return ctx.Result(200, reply)
	}
}

type AccountHTTPClient interface {
	CreateAccount(ctx context.Context, req *CreateAccountRequest, opts ...http.CallOption) (rsp *CreateAccountReply, err error)
	GetAccount(ctx context.Context, req *GetAccountRequest, opts ...http.CallOption) (rsp *GetAccountReply, err error)
	ListAccount(ctx context.Context, req *ListAccountRequest, opts ...http.CallOption) (rsp *ListAccountReply, err error)
	Login(ctx context.Context, req *LoginRequest, opts ...http.CallOption) (rsp *LoginReply, err error)
}

type AccountHTTPClientImpl struct {
	cc *http.Client
}

func NewAccountHTTPClient(client *http.Client) AccountHTTPClient {
	return &AccountHTTPClientImpl{client}
}

func (c *AccountHTTPClientImpl) CreateAccount(ctx context.Context, in *CreateAccountRequest, opts ...http.CallOption) (*CreateAccountReply, error) {
	var out CreateAccountReply
	pattern := "/v1/accounts"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/api.account.service.v1.Account/CreateAccount"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AccountHTTPClientImpl) GetAccount(ctx context.Context, in *GetAccountRequest, opts ...http.CallOption) (*GetAccountReply, error) {
	var out GetAccountReply
	pattern := "/v1/accounts/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/api.account.service.v1.Account/GetAccount"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AccountHTTPClientImpl) ListAccount(ctx context.Context, in *ListAccountRequest, opts ...http.CallOption) (*ListAccountReply, error) {
	var out ListAccountReply
	pattern := "/v1/accounts"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/api.account.service.v1.Account/ListAccount"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AccountHTTPClientImpl) Login(ctx context.Context, in *LoginRequest, opts ...http.CallOption) (*LoginReply, error) {
	var out LoginReply
	pattern := "/v1/login"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/api.account.service.v1.Account/Login"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
