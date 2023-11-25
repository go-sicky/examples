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
 * @since 11/23/2023
 */

package main

import (
	"log/slog"
	"os"

	"github.com/go-sicky/examples/bff/handler"
	"github.com/go-sicky/sicky"
	"github.com/go-sicky/sicky/server"
	shttp "github.com/go-sicky/sicky/server/http"
)

const (
	AppName = "bff.examples.sicky"
	Version = "v0.0.1"
)

func main() {
	cfg := sicky.DefaultConfig(AppName, Version)
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	httpSrv := shttp.NewServer(
		shttp.DefaultConfig(AppName),
		server.Logger(logger),
	)
	httpSrv.RegisterHandler(handler.NewCFilter())
	svc := sicky.NewService(
		&cfg.Sicky.Service,
		sicky.Logger(logger),
		sicky.Server(httpSrv),
	)

	err := svc.Run()
	if err != nil {
		logger.Error(err.Error())
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
