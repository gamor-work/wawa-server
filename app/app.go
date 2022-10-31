package app

import (
	config "./config"
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Options struct {
	Configfile string
	ModelFile  string
	MenuFile   string
	WWWDir     string
	Version    string
}

type Option func(*Options)

func InitHTTPServer(ctx context.Context) func() {
	cfg := config.C.HTTP
	addr := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)
	srv := &http.Server{
		Addr: addr,
		//Handler:      handler,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  30 * time.Second,
	}

	fmt.Println(srv)

	go func() {
		var err error
		err = srv.ListenAndServe()

		if err != nil && err != http.ErrServerClosed {
			//panic(err)
		}
	}()

	return func() {

	}
}

func Init(ctx context.Context, opts ...Option) (func(), error) {
	var o Options
	for _, opt := range opts {
		opt(&o)
	}

	httpServerCleanFunc := InitHTTPServer(ctx)

	return func() {
		httpServerCleanFunc()
	}, nil
}

// 运行后台服务
func Run(ctx context.Context, opts ...Option) error {
	state := 1
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	cleanFunc, err := Init(ctx, opts...)

	// 如果有错误则报出异常
	if err != nil {
		return err
	}

EXIT:
	for {
		sig := <-sc
		switch sig {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			state = 0
			break EXIT
		case syscall.SIGHUP:
		default:
			break EXIT
		}
	}

	cleanFunc()
	time.Sleep(time.Second)
	os.Exit(state)
	return nil
}
