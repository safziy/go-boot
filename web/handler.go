package web

import (
	"context"
)

type BootRequest struct {
}

type BootResponse struct {
}

// Handler 处理方法
type Handler func(c context.Context, req interface{}) (resp interface{}, err error)
