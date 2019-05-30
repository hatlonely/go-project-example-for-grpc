package main

import (
	"flag"
	"fmt"
	apigogrpc "github.com/hatlonely/go-project-example-for-grpc/api/gogrpc"
	"github.com/hatlonely/go-project-example-for-grpc/internal/gogrpc"
	"github.com/hatlonely/go-project-example-for-grpc/internal/logger"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"net"
	"os"
)

//func RegisterHandler(r *gin.Engine) {
//	r.GET("/hello", gohttp.GoHttpHandler)
//}

// AppVersion name
var AppVersion = "unknown"

func main() {
	version := flag.Bool("v", false, "print current version")
	configfile := flag.String("c", "configs/gohttp.json", "config file path")
	flag.Parse()
	if *version {
		fmt.Println(AppVersion)
		os.Exit(0)
	}

	config := viper.New()
	config.SetConfigType("json")
	fp, err := os.Open(*configfile)
	if err != nil {
		panic(err)
	}
	err = config.ReadConfig(fp)
	if err != nil {
		panic(err)
	}

	infoLog, err := logger.NewTextLoggerWithViper(config.Sub("logger.infoLog"))
	if err != nil {
		panic(err)
	}
	warnLog, err := logger.NewTextLoggerWithViper(config.Sub("logger.warnLog"))
	if err != nil {
		panic(err)
	}
	accessLog, err := logger.NewJsonLoggerWithViper(config.Sub("logger.accessLog"))
	gogrpc.InfoLog = infoLog
	gogrpc.WarnLog = warnLog
	gogrpc.AccessLog = accessLog

	infoLog.Infof("%v init success, port[%v]", os.Args[0], config.GetString("service.port"))

	server := grpc.NewServer()

	apigogrpc.RegisterServiceServer(server, gogrpc.NewService())
	address, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%v", config.GetInt("service.port")))
	if err != nil {
		panic(err)
	}

	if err := server.Serve(address); err != nil {
		panic(err)
	}
}
