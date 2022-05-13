package main

import (
	"context"
	"flag"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"go.uber.org/zap"

	"go-template/config"
	"go-template/internal/code"
	"go-template/internal/router"
	"go-template/internal/util/v"
	"go-template/pkg/logger"
	"go-template/pkg/routine"
)

var configFile = flag.String("f", "./conf/template.yaml", "the config file")

func main() {
	flag.Parse()
	// 初始化配置
	if err := config.LoadConfig(*configFile); err != nil {
		log.Fatal(err)
	}
	if err := code.Loading(); err != nil {
		log.Fatal(err)
	}
	if mode := strings.ToLower(viper.GetString("GIN_MODE")); mode != "" {
		gin.SetMode(mode)
	}
	// 初始化系统日志
	zap.ReplaceGlobals(logger.New(
		logger.WithFields(zap.String("service", v.ServiceName)),
		logger.WithLevel(viper.GetString("level")),
		logger.WithWriter(logger.SetWriter(viper.GetString("path")))))
	if err := run(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatal(err)
	}
}

func run() error {
	ctx := context.Background()
	g := routine.NewGroup(ctx)
	srv := &http.Server{
		Addr:    ":8080",
		Handler: router.New(),
		BaseContext: func(_ net.Listener) context.Context {
			return ctx
		},
	}
	zap.S().Debugf("listen on %s", srv.Addr)
	// 服务启动流程
	g.Go(func(ctx context.Context) error {
		return startAction(ctx, srv)
	})
	// 服务关闭流程
	g.Go(func(ctx context.Context) error {
		return shutdownAction(ctx, srv)
	})
	// 启动mq
	if err := g.Wait(); err != nil && !errors.Is(err, context.Canceled) {
		return err
	}
	return nil
}

func startAction(ctx context.Context, srv *http.Server) error {

	zap.S().Infof("run on %s", gin.Mode())
	return srv.ListenAndServe()
}

func shutdownAction(ctx context.Context, srv *http.Server) error {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	select {
	case <-ctx.Done():
	case <-quit:
	}
	zap.L().Info("shutting down server...")
	return srv.Shutdown(ctx)
}
