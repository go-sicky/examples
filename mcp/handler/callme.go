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
 * @file callme.go
 * @package handler
 * @author Dr.NP <np@herewe.tech>
 * @since 07/08/2026
 */

package handler

import "github.com/go-sicky/sicky/service/mcp"

type CallMeHandler struct{}

func (h *CallMeHandler) Name() string {
	return "callme"
}

func (h *CallMeHandler) Description() string {
	return "CallMeHandler is a handler for call me."
}

func (h *CallMeHandler) Tools() []mcp.Tool {
	return []mcp.Tool{}
}

func (h *CallMeHandler) CallTool(name string, args map[string]interface{}) (*mcp.ToolsCallResult, error) {
	return nil, nil
}

func (h *CallMeHandler) Resources() []mcp.Resource {
	return []mcp.Resource{}
}

func (h *CallMeHandler) ReadResource(uri string) (*mcp.ResourcesReadResult, error) {
	return nil, nil
}

func (h *CallMeHandler) Prompts() []mcp.Prompt {
	return []mcp.Prompt{}
}

func (h *CallMeHandler) GetPrompt(name string, args map[string]string) (*mcp.PromptsGetResult, error) {
	return nil, nil
}

/*
 * Local variables:
 * tab-width: 4
 * c-basic-offset: 4
 * End:
 * vim600: sw=4 ts=4 fdm=marker
 * vim<600: sw=4 ts=4
 */
