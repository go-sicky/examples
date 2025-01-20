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
 * @file hybrid.tcp.go
 * @package handler
 * @author Dr.NP <np@herewe.tech>
 * @since 01/20/2025
 */

package handler

import (
	"bytes"
	"fmt"
	"net"
)

type TCPHybrid struct{}

func NewTCPHybrid() *TCPHybrid {
	h := &TCPHybrid{}

	return h
}

func (h *TCPHybrid) Name() string {
	return "hybrid.tcp"
}

func (h *TCPHybrid) Type() string {
	return "tcp"
}

func (h *TCPHybrid) OnConnect(c net.Conn) error {
	fmt.Println(c.RemoteAddr().String(), "connected")

	return nil
}

func (h *TCPHybrid) OnClose(c net.Conn) error {
	fmt.Println(c.RemoteAddr().String(), "closed")

	return nil
}

func (h *TCPHybrid) OnError(c net.Conn, err error) error {
	fmt.Println(c.RemoteAddr().String(), "error", err.Error())

	return err
}

func (h *TCPHybrid) OnData(c net.Conn, data []byte) error {
	fmt.Println(c.RemoteAddr().String(), "say", string(data))

	var resp bytes.Buffer

	resp.WriteString("Yes! ")
	resp.Write(data)
	_, err := c.Write(resp.Bytes())

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
