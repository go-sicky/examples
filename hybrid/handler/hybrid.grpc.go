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
 * @file hybrid.grpc.go
 * @package handler
 * @author Dr.NP <np@herewe.tech>
 * @since 01/20/2025
 */

package handler

import (
	"context"

	"github.com/go-sicky/example/hybrid/proto"
	"google.golang.org/grpc"
)

type GRPCHyprid struct {
	proto.UnimplementedHybridServer
}

func NewGRPCHybrid() *GRPCHyprid {
	h := &GRPCHyprid{}

	return h
}

func (h *GRPCHyprid) Name() string {
	return "hybrid.grpc"
}

func (h *GRPCHyprid) Type() string {
	return "grpc"
}

func (h *GRPCHyprid) Register(app *grpc.Server) {
	proto.RegisterHybridServer(app, h)
}

/* {{{ [Methods] */
func (h *GRPCHyprid) Hybrid(ctx context.Context, req *proto.HybridRequest) (*proto.HybridResponse, error) {
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
