package net

import (
	"fmt"
	"net"
	"strings"

	"gopkg.in/op/go-logging.v1"
)

var (
	logger    = logging.MustGetLogger("exchange.net")
	QueueSize = 1000
)

type HandlerFunc func(c *Context)

type Engine struct {
	handlerFunc   map[string]HandlerFunc
	handlers      []HandlerFunc
	groupHandlers map[string]*Group
	connPool      chan net.Conn
}

func New(quit chan bool) *Engine {
	e := &Engine{
		handlerFunc:   make(map[string]HandlerFunc),
		groupHandlers: make(map[string]*Group),
		connPool:      make(chan net.Conn, QueueSize),
	}

	for i := 0; i < QueueSize; i++ {
		w := &Worker{
			ID:   i,
			Enge: e,
		}
		w.Start(quit)
	}
	return e
}

// add middleware
func (engine *Engine) Use(handler HandlerFunc) {
	engine.handlers = append(engine.handlers, handler)
}

func (engine *Engine) Register(path string, handler HandlerFunc) {
	if _, ok := engine.handlerFunc[path]; ok {
		panic(fmt.Sprintf("duplicate router %s", path))
	}
	engine.handlerFunc[path] = handler
}

func (engine *Engine) Group(path string, handlers ...HandlerFunc) *Group {
	// check if the group path conflict.
	ps := strings.Split(path, "/")
	if len(ps) == 0 {
		panic("empty path")
	}

	root := ps[0]
	for p := range engine.groupHandlers {
		if strings.HasPrefix(p, root) {
			panic(fmt.Sprintf("conflict group path name:%s with %s", path, p))
		}
	}

	gp := &Group{
		Path:        path,
		midHandlers: handlers,
		handlers:    make(map[string]HandlerFunc),
	}

	engine.groupHandlers[path] = gp
	return gp
}

func (engine *Engine) Run(port int) {
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		panic(err)
	}
	for {
		c, err := l.Accept()
		if err != nil {
			panic(err)
		}
		logger.Debug("new connection:%s", c.RemoteAddr())
		engine.connPool <- c
	}
}

// Recovery is middleware for catching panic.
func Recovery() HandlerFunc {
	return func(c *Context) {
		defer func() {
			if r := recover(); r != nil {
				logger.Critical("%s", r)
			}
		}()
		c.Next()
	}
}
