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
 * @file call.http.go
 * @package handler
 * @author Dr.NP <np@herewe.tech>
 * @since 08/10/2024
 */

package handler

import (
	"github.com/go-sicky/examples/hybrid/proto"
	"github.com/go-sicky/examples/web/model"
	"github.com/go-sicky/sicky/broker"
	grpcClt "github.com/go-sicky/sicky/client/grpc"
	"github.com/go-sicky/sicky/utils"
	"github.com/gofiber/fiber/v2"
)

type CallHTTP struct {
	hc proto.HybridClient
}

func NewCallHTTP() *CallHTTP {
	grpcClient := grpcClt.New(
		nil,
		&grpcClt.Config{
			Service: "hybrid.examples.sicky@grpc",
		},
	)
	h := &CallHTTP{
		hc: proto.NewHybridClient(grpcClient),
	}

	return h
}

func (h *CallHTTP) Register(app *fiber.App) {
	app.Get("/", h.index).Name("CallGetIndex")
	app.Get("/pool", h.pool).Name("CallGetPool")
	app.Get("/routers", h.routers).Name("CallGetRouters")
	app.Get("/tcp", h.tcp).Name("CallGetTcp")
	app.Head("/udp", h.udp).Name("CallHeadUdp")
	app.Get("/grpc", h.grpc).Name("CallGetGrpc")
	app.Post("/broker", h.broker).Name("BrokerPostBroker")
}

func (h *CallHTTP) Name() string {
	return "call.http"
}

func (h *CallHTTP) Type() string {
	return "http"
}

/* {{{ [HTTP handlers] */
func (h *CallHTTP) index(c *fiber.Ctx) error {
	return c.JSON(utils.WrapHTTPResponse(nil))
}

func (h *CallHTTP) grpc(c *fiber.Ctx) error {
	e := utils.WrapHTTPResponse(nil)
	req := &proto.HybridRequest{
		Name: "John Doe",
	}

	resp, err := h.hc.Hybrid(c.Context(), req)
	if err != nil {
		e.Status = fiber.StatusInternalServerError
		e.Message = err.Error()
	} else {
		e.Data = resp
	}

	return c.JSON(e)
}

func (h *CallHTTP) pool(c *fiber.Ctx) error {
	// return c.JSON(utils.WrapHTTPResponse(registry.Pool))
	return nil
}

func (h *CallHTTP) routers(c *fiber.Ctx) error {
	r := c.App().GetRoutes()

	return c.JSON(utils.WrapHTTPResponse(r))
}

func (h *CallHTTP) tcp(c *fiber.Ctx) error {
	return c.JSON(utils.WrapHTTPResponse(nil))
}

func (h *CallHTTP) udp(c *fiber.Ctx) error {
	return c.JSON(utils.WrapHTTPResponse(nil))
}

func (h *CallHTTP) broker(c *fiber.Ctx) error {
	v := &model.Person{
		Name:    "John Doe",
		Age:     24,
		Address: "123 Main St",
	}
	msg := broker.NewMessage(nil)
	msg.Format(v, broker.MsgJson)
	broker.Publish("hybrid", msg)

	return c.Format(utils.WrapHTTPResponse(nil))
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
