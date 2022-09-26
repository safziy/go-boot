package web

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"reflect"
)

var router *gin.Engine

func InitWeb(conf *Config) error {
	router = gin.Default()

	return router.Run()
}

func buildAdapter() {

}

func AddRouter(method, path string, handleFunc interface{}) error {
	funcValue, reqType, respType, err := checkAndParseHandleFunc(handleFunc)
	if err != nil {
		return err
	}
	fmt.Println(reqType)
	fmt.Println(respType)
	ginHandler := func(c *gin.Context) {
		cc := context.Background()
		req := reflect.New(reqType)
		values := funcValue.Call([]reflect.Value{reflect.ValueOf(cc), req})
		err := values[1].Interface()
		if err != nil {
			return
		}
		resp := values[0].Interface()
		c.JSON(http.StatusOK, resp)
	}
	switch method {
	case http.MethodGet:
		router.GET(path, ginHandler)
	case http.MethodHead:
		router.HEAD(path, ginHandler)
	case http.MethodPost:
		router.POST(path, ginHandler)
	case http.MethodPut:
		router.PUT(path, ginHandler)
	case http.MethodPatch:
		router.PATCH(path, ginHandler)
	case http.MethodDelete:
		router.DELETE(path, ginHandler)
	//case http.MethodConnect: router.Connect(path, ginHandler)
	case http.MethodOptions:
		router.OPTIONS(path, ginHandler)
		//case http.MethodTrace: router.Trace(path, ginHandler)
	}
	return nil
}

func checkAndParseHandleFunc(handleFunc interface{}) (funcValue reflect.Value, req, resp reflect.Type, err error) {
	funcType := reflect.TypeOf(handleFunc)
	funcValue = reflect.ValueOf(handleFunc)
	if funcType.Kind() != reflect.Func {
		return funcValue, nil, nil, errors.New("invalid handleFunc: handleFunc not func")
	}
	if funcType.NumIn() != 2 {
		return funcValue, nil, nil, errors.New("invalid handleFunc: params count not 2")
	}
	if funcType.NumOut() != 2 {
		return funcValue, nil, nil, errors.New("invalid handleFunc return count not 2")
	}
	param1 := funcType.In(0)
	if !param1.Implements(reflect.TypeOf((*context.Context)(nil)).Elem()) {
		return funcValue, nil, nil, errors.New("invalid handleFunc: the one param not context")
	}
	param2 := funcType.In(1)
	return1 := funcType.Out(0)
	return2 := funcType.Out(1)
	if !return2.Implements(reflect.TypeOf((*error)(nil)).Elem()) {
		return funcValue, nil, nil, errors.New("invalid handleFunc: the last return not error")
	}
	return funcValue, param2, return1, nil
}

func AddRouter1(method string, path string, handler Handler) {
	ginHandler := func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
		req1 := 4
		handler(context.Background(), req1)
	}
	switch method {
	case http.MethodGet:
		router.GET(path, ginHandler)
	case http.MethodHead:
		router.HEAD(path, ginHandler)
	case http.MethodPost:
		router.POST(path, ginHandler)
	case http.MethodPut:
		router.PUT(path, ginHandler)
	case http.MethodPatch:
		router.PATCH(path, ginHandler)
	case http.MethodDelete:
		router.DELETE(path, ginHandler)
	//case http.MethodConnect: router.Connect(path, ginHandler)
	case http.MethodOptions:
		router.OPTIONS(path, ginHandler)
		//case http.MethodTrace: router.Trace(path, ginHandler)
	}
}
