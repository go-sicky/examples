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
 * @file web.grpc.go
 * @package handler
 * @author Dr.NP <np@herewe.tech>
 * @since 08/10/2024
 */

package handler

import (
	"context"

	"github.com/go-sicky/examples/web/proto"
	"google.golang.org/grpc"
)

type WebGRPC struct {
	proto.UnimplementedWebServer
}

func NewWebGRPC() *WebGRPC {
	h := &WebGRPC{}

	return h
}

func (h *WebGRPC) Register(app *grpc.Server) {
	proto.RegisterWebServer(app, h)
}

func (h *WebGRPC) Name() string {
	return "web.grpc"
}

func (h *WebGRPC) Type() string {
	return "grpc"
}

/* {{{ [Protobuf methods] */
func (h *WebGRPC) Call(ctx context.Context, req *proto.WebCallReq) (*proto.WebCallResp, error) {
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
