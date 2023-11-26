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
 * @file cfilter.go
 * @package handler
 * @author Dr.NP <np@herewe.tech>
 * @since 11/24/2023
 */

package handler

import (
	"context"

	pb "github.com/go-sicky/examples/cfilter/proto"
	"google.golang.org/grpc"
)

type CFilter struct {
	name string
	desc *grpc.ServiceDesc
	pb.UnimplementedCFilterServer
}

func NewCFilter(name string) *CFilter {
	return &CFilter{
		name: name,
		desc: &pb.CFilter_ServiceDesc,
	}
}

func (h *CFilter) Register(app *grpc.Server) {
	app.RegisterService(h.desc, h)
}

func (h *CFilter) Name() string {
	return h.name
}

func (h *CFilter) Type() string {
	return "grpc"
}

func (h *CFilter) Filter(ctx context.Context, req *pb.FilterRequest) (*pb.FilterResponse, error) {
	resp := new(pb.FilterResponse)
	resp.Output = "Hello " + req.Input

	return resp, nil
}

/*
 * Local variables:
 * tab-width: 4
 * c-basic-offset: 4
 * End:
 * vim600: sw=4 ts=4 fdm=marker
 * vim<600: sw=4 ts=4
 */
