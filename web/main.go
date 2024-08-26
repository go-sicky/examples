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
 * @file main.go
 * @package main
 * @author Dr.NP <np@herewe.tech>
 * @since 08/10/2024
 */

package main

import (
	"github.com/go-sicky/examples/web/handler"
	brkNats "github.com/go-sicky/sicky/broker/nats"

	//rgConsul "github.com/go-sicky/sicky/registry/consul"
	rgMdns "github.com/go-sicky/sicky/registry/mdns"
	srvGRPC "github.com/go-sicky/sicky/server/grpc"
	srvHTTP "github.com/go-sicky/sicky/server/http"
	"github.com/go-sicky/sicky/service"
	"github.com/go-sicky/sicky/service/sicky"
)

const (
	AppName = "web.examples.sicky"
	Version = "latest"
)

func main() {
	// HTTP server
	httpSrv := srvHTTP.New(AppName+"@http", nil, nil)
	httpSrv.Handle(handler.NewCallHTTP())
	httpSrv.Handle(handler.NewBrokerHTTP())

	// GRPC server
	grpcSrv := srvGRPC.New(AppName+"@grpc", nil, nil)
	grpcSrv.Handle(handler.NewWebGRPC())

	// Broker
	brk := brkNats.New(nil, nil)

	// Registry
	//rg := rgConsul.New(nil, nil)
	rg := rgMdns.New(nil, nil)

	// Service
	svc := sicky.New(nil, nil)
	svc.Servers(httpSrv, grpcSrv)
	svc.Brokers(brk)
	svc.Registries(rg)

	service.Run()
	// cfg, err := sicky.LoadConfig(AppName, Version)
	// if err != nil {
	// 	logger.Logger.Errorf("Load config failed : %s", err)
	// }

	// logger.Logger.Level(logger.LogLevel(cfg.Sicky.LogLevel))

	// // Server
	// httpSrv := shttp.NewServer(
	// 	cfg.HTTPServer(AppName),
	// 	server.Handle(handler.NewCFilter("bff")),
	// )

	// // Client
	// grpcClt := cgrpc.NewClient(
	// 	cfg.GRPCClient(CFilterName),
	// )

	// svc := sicky.NewService(
	// 	cfg,
	// 	sicky.Server(httpSrv),
	// 	sicky.Client(grpcClt),
	// )

	// err = svc.Run()
	// if err != nil {
	// 	logger.Logger.Error(err.Error())
	// }
}

/*
 * Local variables:
 * tab-width: 4
 * c-basic-offset: 4
 * End:
 * vim600: sw=4 ts=4 fdm=marker
 * vim<600: sw=4 ts=4
 */
