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
	brkNsq "github.com/go-sicky/sicky/broker/nsq"
	"github.com/go-sicky/sicky/logger"
	rgConsul "github.com/go-sicky/sicky/registry/consul"
	rgMdns "github.com/go-sicky/sicky/registry/mdns"
	"github.com/go-sicky/sicky/runtime"
	"github.com/go-sicky/sicky/server"
	srvGRPC "github.com/go-sicky/sicky/server/grpc"
	srvHTTP "github.com/go-sicky/sicky/server/http"
	srvWebsocket "github.com/go-sicky/sicky/server/websocket"
	"github.com/go-sicky/sicky/service"
	"github.com/go-sicky/sicky/service/sicky"
)

const (
	AppName = "web.examples.sicky"
	Version = "latest"
)

func main() {
	// Runtime
	runtime.Init(AppName)
	runtime.LoadConfig(&config)
	runtime.Start(config.Runtime)

	// Logger
	logger.Logger.Level(logger.DebugLevel)

	// HTTP server
	httpSrv := srvHTTP.New(&server.Options{Name: AppName + "@http"}, config.Server.HTTP)
	httpSrv.Handle(handler.NewCallHTTP())
	httpSrv.Handle(handler.NewBrokerHTTP())

	// GRPC server
	grpcSrv := srvGRPC.New(&server.Options{Name: AppName + "@grpc"}, config.Server.GRPC)
	grpcSrv.Handle(handler.NewWebGRPC())

	// Websocket server
	wsSrv := srvWebsocket.New(&server.Options{Name: AppName + "@websocket"}, config.Server.Websocket)

	// Broker
	brkNats := brkNats.New(nil, config.Broker.Nats)
	brkNsq := brkNsq.New(nil, config.Broker.Nsq)

	// Registry
	rgConsul := rgConsul.New(nil, config.Registry.Consul)
	rgMdns := rgMdns.New(nil, config.Registry.Mdns)

	// Service
	svc := sicky.New(&service.Options{Name: AppName}, config.Service)
	svc.Servers(httpSrv, grpcSrv, wsSrv)
	svc.Brokers(brkNats, brkNsq)
	svc.Registries(rgMdns, rgConsul)

	service.Run()
}

/*
 * Local variables:
 * tab-width: 4
 * c-basic-offset: 4
 * End:
 * vim600: sw=4 ts=4 fdm=marker
 * vim<600: sw=4 ts=4
 */
