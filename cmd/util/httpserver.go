package util

import (
	"github.com/golang/glog"
	"github.com/labstack/echo/v4"
	"strings"
)

type Handler func(c echo.Context) error

var s *httpServer

type httpServer struct {
	getHandlers  map[string]Handler
	postHandlers map[string]Handler
	port         string
}

func init() {
	s = &httpServer{
		getHandlers:  make(map[string]Handler),
		postHandlers: make(map[string]Handler),
	}
}

func AddHandler(method string, key string, h Handler) {
	m := strings.ToUpper(method)
	switch m {
	case "GET":
		s.getHandlers[key] = h
	case "POST":
		s.postHandlers[key] = h
	default:
		glog.Error("addHandler: not support http method")
	}
}

func SetPort(port string) {
	s.port = port
}

func Start() error {
	e := echo.New()

	for pk, pv := range s.postHandlers {
		e.POST(pk, echo.HandlerFunc(pv))
	}

	for gk, gv := range s.getHandlers {
		e.GET(gk, echo.HandlerFunc(gv))
	}

	if err := e.Start(s.port); err != nil {
		glog.Errorf("http server started on err:%+v", err)
		return err
	}
	return nil
}
