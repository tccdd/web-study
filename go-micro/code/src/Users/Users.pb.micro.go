// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: Users.proto

package Users

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Uservice service

type UserviceService interface {
	Test(ctx context.Context, in *UserRequest, opts ...client.CallOption) (*UserResponce, error)
}

type userviceService struct {
	c    client.Client
	name string
}

func NewUserviceService(name string, c client.Client) UserviceService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "Users"
	}
	return &userviceService{
		c:    c,
		name: name,
	}
}

func (c *userviceService) Test(ctx context.Context, in *UserRequest, opts ...client.CallOption) (*UserResponce, error) {
	req := c.c.NewRequest(c.name, "Uservice.Test", in)
	out := new(UserResponce)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Uservice service

type UserviceHandler interface {
	Test(context.Context, *UserRequest, *UserResponce) error
}

func RegisterUserviceHandler(s server.Server, hdlr UserviceHandler, opts ...server.HandlerOption) error {
	type uservice interface {
		Test(ctx context.Context, in *UserRequest, out *UserResponce) error
	}
	type Uservice struct {
		uservice
	}
	h := &userviceHandler{hdlr}
	return s.Handle(s.NewHandler(&Uservice{h}, opts...))
}

type userviceHandler struct {
	UserviceHandler
}

func (h *userviceHandler) Test(ctx context.Context, in *UserRequest, out *UserResponce) error {
	return h.UserviceHandler.Test(ctx, in, out)
}