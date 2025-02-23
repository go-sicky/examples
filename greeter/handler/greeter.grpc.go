/*
 * The MIT License (MIT)
 *
 * Copyright (c) 2021 HereweTech Co.LTD
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy of
 * this software and associated documentation files (the "Software"), to deal in
 * the Software without restriction, including without limitation the rights to
 * use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of
 * the Software, and to permit persons to whom the Software is furnished to do so,
 * subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS
 * FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
 * COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER
 * IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
 * CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
 */

/**
 * @file greeter.grpc.go
 * @package handler
 * @author Dr.NP <np@herewe.tech>
 * @since 08/10/2024
 */

package handler

import (
	"context"

	"github.com/go-sicky/examples/greeter/proto"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type GreeterGRPC struct {
	proto.UnimplementedGreeterServer
}

func NewGreeterGRPC() *GreeterGRPC {
	h := &GreeterGRPC{}

	return h
}

func (h *GreeterGRPC) Register(app *grpc.Server) {
	proto.RegisterGreeterServer(app, h)
}

func (h *GreeterGRPC) Name() string {
	return "greeter.grpc"
}

func (h *GreeterGRPC) Type() string {
	return "grpc"
}

/* {{{ [Methods] */
func (h *GreeterGRPC) Hello(ctx context.Context, req *proto.HelloRequest) (*proto.HelloResponse, error) {
	return nil, nil
}

func (h *GreeterGRPC) Timestamp(ctx context.Context, req *emptypb.Empty) (*proto.TimestampResponse, error) {
	return nil, nil
}

/* }}} */

/*
 * Local variables:
 * tab-width: 4
 * c-basic-offset: 4
 * End:
 * vim600: sw=4 ts=4 fdm=marker
 * vim<600: sw=4 ts=4
 */
