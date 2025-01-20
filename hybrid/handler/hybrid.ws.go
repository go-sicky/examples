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
 * @file hybrid.ws.go
 * @package handler
 * @author Dr.NP <np@herewe.tech>
 * @since 01/20/2025
 */

package handler

import (
	"bytes"
	"fmt"

	"github.com/gofiber/contrib/websocket"
)

type WSHybrid struct{}

func NewWSHybrid() *WSHybrid {
	h := &WSHybrid{}

	return h
}

func (h *WSHybrid) Name() string {
	return "hybrid.websocket"
}

func (h *WSHybrid) Type() string {
	return "websocket"
}

func (h *WSHybrid) OnConnect(c *websocket.Conn) error {
	fmt.Println(c.RemoteAddr().String(), "connected")

	return nil
}

func (h *WSHybrid) OnClose(c *websocket.Conn) error {
	fmt.Println(c.RemoteAddr().String(), "closed")

	return nil
}

func (h *WSHybrid) OnError(c *websocket.Conn, err error) error {
	fmt.Println(c.RemoteAddr().String(), "error", err.Error())

	return err
}

func (h *WSHybrid) OnData(c *websocket.Conn, msgType int, data []byte) error {
	fmt.Println(c.RemoteAddr().String(), "say", string(data))

	var resp bytes.Buffer

	resp.WriteString("Yes! ")
	resp.Write(data)
	err := c.WriteMessage(msgType, resp.Bytes())

	return err
}

/*
 * Local variables:
 * tab-width: 4
 * c-basic-offset: 4
 * End:
 * vim600: sw=4 ts=4 fdm=marker
 * vim<600: sw=4 ts=4
 */
