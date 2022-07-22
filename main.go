package main

import (
	"encoding/json"
	"fmt"
	"github.com/gole-dev/gole-layout/internal/server"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	goleApp "github.com/gole-dev/gole/pkg/app"
	"github.com/gole-dev/gole/pkg/config"
	logger "github.com/gole-dev/gole/pkg/log"
	"github.com/gole-dev/gole/pkg/redis"
	v "github.com/gole-dev/gole/pkg/version"
	"github.com/spf13/pflag"
	_ "go.uber.org/automaxprocs"

	"github.com/gole-dev/gole-layout/internal/model"
)

var (
	cfgDir  = pflag.StringP("config dir", "c", "config", "config path.")
	env     = pflag.StringP("env name", "e", "local", "env var name.")
	version = pflag.BoolP("version", "v", false, "show version info.")
)

// @title API Doc
// @version 1.0
// @description API Doc of Project

// @host localhost:8080
// @BasePath /v1
func main() {
	pflag.Parse()
	if *version {
		ver := v.Get()
		marshaled, err := json.MarshalIndent(&ver, "", "  ")
		if err != nil {
			fmt.Printf("%v\n", err)
			os.Exit(1)
		}

		fmt.Println(string(marshaled))
		return
	}

	// init config
	c := config.New(*cfgDir, config.WithEnv(*env))
	var cfg goleApp.Config
	if err := c.Load("app", &cfg); err != nil {
		fmt.Printf("load config error: %v \nyou should copy config/sample as config/local for local dev\n", err)
		return
	}
	// set global
	goleApp.Conf = &cfg

	// -------------- init resource -------------
	logger.Init()
	// init db
	model.Init()
	// init redis
	redis.Init()

	gin.SetMode(cfg.Mode)

	// init pprof server
	go func() {
		fmt.Printf("Listening and serving PProf HTTP on %s\n", cfg.PprofPort)
		if err := http.ListenAndServe(cfg.PprofPort, http.DefaultServeMux); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen ListenAndServe for PProf, err: %s", err.Error())
		}
	}()

	// start app
	app := goleApp.New(
		goleApp.WithName(cfg.Name),
		goleApp.WithVersion(cfg.Version),
		goleApp.WithLogger(logger.GetLogger()),
		goleApp.WithServer(
			// init http server
			server.NewHTTPServer(&cfg.HTTP),
			// init gRPC server
			server.NewGRPCServer(&cfg.GRPC),
		),
	)

	if err := app.Run(); err != nil {
		panic(err)
	}

}
