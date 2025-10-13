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
 * @file config.go
 * @package main
 * @author Dr.NP <np@herewe.tech>
 * @since 01/20/2025
 */

package main

import (
	brkJetstream "github.com/go-sicky/sicky/broker/jetstream"
	brkNats "github.com/go-sicky/sicky/broker/nats"
	brkNsq "github.com/go-sicky/sicky/broker/nsq"
	rgConsul "github.com/go-sicky/sicky/registry/consul"
	"github.com/go-sicky/sicky/runtime"
	srvGRPC "github.com/go-sicky/sicky/server/grpc"
	srvHTTP "github.com/go-sicky/sicky/server/http"
	srvOriginal "github.com/go-sicky/sicky/server/original"
	srvTCP "github.com/go-sicky/sicky/server/tcp"
	srvUDP "github.com/go-sicky/sicky/server/udp"
	srvWS "github.com/go-sicky/sicky/server/websocket"
	"github.com/go-sicky/sicky/service/sicky"
)

type ConfigDef struct {
	Server struct {
		HTTP      *srvHTTP.Config     `json:"http" yaml:"http" mapstructure:"http"`
		Original  *srvOriginal.Config `json:"original" yaml:"original" mapstructure:"original"`
		GRPC      *srvGRPC.Config     `json:"grpc" yaml:"grpc" mapstructure:"grpc"`
		Websocket *srvWS.Config       `json:"websocket" yaml:"websocket" mapstructure:"websocket"`
		UDP       *srvUDP.Config      `json:"udp" yaml:"udp" mapstructure:"udp"`
		TCP       *srvTCP.Config      `json:"tcp" yaml:"tcp" mapstructure:"tcp"`
	}
	Broker struct {
		Nats      *brkNats.Config      `json:"nats" yaml:"nats" mapstructure:"nats"`
		Nsq       *brkNsq.Config       `json:"nsq" yaml:"nsq" mapstructure:"nsq"`
		Jetstream *brkJetstream.Config `json:"jetstream" yaml:"jetstream" mapstructure:"jetstream"`
	}
	Registry struct {
		Consul *rgConsul.Config `json:"consul" yaml:"consul" mapstructure:"consul"`
	}
	Runtime *runtime.Config `json:"runtime" yaml:"runtime" mapstructure:"runtime"`
	Service *sicky.Config   `json:"service" yaml:"service" mapstructure:"service"`
}

var (
	config ConfigDef
)

/*
 * Local variables:
 * tab-width: 4
 * c-basic-offset: 4
 * End:
 * vim600: sw=4 ts=4 fdm=marker
 * vim<600: sw=4 ts=4
 */
