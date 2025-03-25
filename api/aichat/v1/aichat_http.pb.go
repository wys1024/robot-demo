// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.8.2
// - protoc             v5.28.3
// source: aichat/v1/aichat.proto

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

const OperationAichatSendMessage = "/aichat.v1.Aichat/SendMessage"

type AichatHTTPServer interface {
	// SendMessage SendMessage 发送聊天消息
	SendMessage(context.Context, *SendMessageRequest) (*SendMessageResponse, error)
}

func RegisterAichatHTTPServer(s *http.Server, srv AichatHTTPServer) {
	r := s.Route("/")
	r.POST("/api/chat", _Aichat_SendMessage0_HTTP_Handler(srv))
}

func _Aichat_SendMessage0_HTTP_Handler(srv AichatHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in SendMessageRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAichatSendMessage)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.SendMessage(ctx, req.(*SendMessageRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*SendMessageResponse)
		return ctx.Result(200, reply)
	}
}

type AichatHTTPClient interface {
	SendMessage(ctx context.Context, req *SendMessageRequest, opts ...http.CallOption) (rsp *SendMessageResponse, err error)
}

type AichatHTTPClientImpl struct {
	cc *http.Client
}

func NewAichatHTTPClient(client *http.Client) AichatHTTPClient {
	return &AichatHTTPClientImpl{client}
}

func (c *AichatHTTPClientImpl) SendMessage(ctx context.Context, in *SendMessageRequest, opts ...http.CallOption) (*SendMessageResponse, error) {
	var out SendMessageResponse
	pattern := "/api/chat"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationAichatSendMessage))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
