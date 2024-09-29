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
 * @since 09/18/2024
 */

package main

import (
	brkNats "github.com/go-sicky/sicky/broker/nats"
	brkNsq "github.com/go-sicky/sicky/broker/nsq"
	rgConsul "github.com/go-sicky/sicky/registry/consul"
	rgMdns "github.com/go-sicky/sicky/registry/mdns"
	srvGRPC "github.com/go-sicky/sicky/server/grpc"
	"github.com/go-sicky/sicky/service/sicky"
)

type ConfigDef struct {
	Server struct {
		GRPC *srvGRPC.Config `json:"grpc" yaml:"grpc"`
	} `json:"server" yaml:"server"`
	Broker struct {
		Nats *brkNats.Config `json:"nats" yaml:"nats"`
		Nsq  *brkNsq.Config  `json:"nsq" yaml:"nsq"`
	} `json:"broker" yaml:"broker"`
	Registry struct {
		Consul *rgConsul.Config `json:"consul" yaml:"consul"`
		Mdns   *rgMdns.Config   `json:"mdns" yaml:"mdns"`
	} `json:"registry" yaml:"registry"`
	Service *sicky.Config `json:"service" yaml:"service"`
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
