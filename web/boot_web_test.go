package web

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"testing"
)

func TestAddRouter(t *testing.T) {
	validHandler := func(c context.Context, req struct{}) (resp struct{}, err error) {
		return resp, nil
	}
	errorHandler := func(req struct{}, c context.Context) (resp struct{}, err error) {
		return resp, nil
	}
	router = gin.Default()
	type args struct {
		method     string
		path       string
		handleFunc interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "case1", args: args{method: http.MethodGet, path: "/test1", handleFunc: validHandler}, wantErr: false},
		{name: "case2", args: args{method: http.MethodGet, path: "/test2", handleFunc: errorHandler}, wantErr: true},
		{name: "case3", args: args{method: http.MethodGet, path: "/test3", handleFunc: 1}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := AddRouter(tt.args.method, tt.args.path, tt.args.handleFunc); (err != nil) != tt.wantErr {
				t.Errorf("AddRouter() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
