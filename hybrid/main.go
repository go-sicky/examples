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
 * @since 01/20/2025
 */

package main

import (
	"github.com/go-sicky/example/hybrid/handler"
	"github.com/go-sicky/sicky/broker"
	brkNats "github.com/go-sicky/sicky/broker/nats"
	brkNsq "github.com/go-sicky/sicky/broker/nsq"
	"github.com/go-sicky/sicky/logger"
	"github.com/go-sicky/sicky/runtime"
	"github.com/go-sicky/sicky/server"
	srvGRPC "github.com/go-sicky/sicky/server/grpc"
	srvHTTP "github.com/go-sicky/sicky/server/http"
	srvTCP "github.com/go-sicky/sicky/server/tcp"
	srvUDP "github.com/go-sicky/sicky/server/udp"
	srvWebsocket "github.com/go-sicky/sicky/server/websocket"
	"github.com/go-sicky/sicky/service"
	"github.com/go-sicky/sicky/service/sicky"
)

const (
	AppName = "hybrid.examnples.sicky"
	Version = "latest"
)

func main() {
	// Runtime
	runtime.Init(AppName)
	runtime.LoadConfig(&config)
	runtime.Start(config.Runtime)

	// HTTP server
	httpSrv := srvHTTP.New(&server.Options{Name: AppName + "@http"}, config.Server.HTTP)
	httpSrv.Handle(
		handler.NewHTTPHybrid(),
	)

	// GRPC server
	grpcSrv := srvGRPC.New(&server.Options{Name: AppName + "@grpc"}, config.Server.GRPC)
	grpcSrv.Handle(
		handler.NewGRPCHybrid(),
	)

	// Websocket server
	wsSrv := srvWebsocket.New(&server.Options{Name: AppName + "@websocket"}, config.Server.Websocket)
	wsSrv.Handle(
		handler.NewWSHybrid(),
	)

	// UDP server
	udpSrv := srvUDP.New(&server.Options{Name: AppName + "@udp"}, config.Server.UDP)
	udpSrv.Handle(
		handler.NewUDPHybrid(),
	)

	// TCP server
	tcpSrv := srvTCP.New(&server.Options{Name: AppName + "@tcp"}, config.Server.TCP)
	tcpSrv.Handle(
		handler.NewTCPHybrid(),
	)

	// Nats broker
	natsBrk := brkNats.New(&broker.Options{Name: AppName + "@nats"}, config.Broker.Nats)
	natsBrk.Handle(
		handler.NewNatsHybrid(),
	)

	// Nsq broker
	nsqBrk := brkNsq.New(&broker.Options{Name: AppName + "@nsq"}, config.Broker.Nsq)
	nsqBrk.Handle(
		handler.NewNsqHybrid(),
	)

	// Service
	svc := sicky.New(&service.Options{Name: AppName}, config.Service)
	svc.Brokers(natsBrk, nsqBrk)
	svc.Servers(
		httpSrv,
		grpcSrv,
		wsSrv,
		udpSrv,
		tcpSrv,
	)
	err := service.Run()
	if err != nil {
		logger.Fatal(err.Error())
	}
}

/*
 * Local variables:
 * tab-width: 4
 * c-basic-offset: 4
 * End:
 * vim600: sw=4 ts=4 fdm=marker
 * vim<600: sw=4 ts=4
 */
