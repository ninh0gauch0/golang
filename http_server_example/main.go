package main

import (
	"context"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/JulianCobas/golang/http_server_example/server"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var (
	contextLogger *log.Entry
	baseContext   = context.Background()
	reloadChan    chan bool
	workerContext context.Context
	vip           *viper.Viper
)

func init() {

	pflag.BoolP("verbose", "v", false, "If set, the log level is set to DEBUG")
	pflag.BoolP("quiet", "q", false, "If set, the log level is set to FATAL")
	pflag.String("ip", "0.0.0.0", "The default IP where we're going to be listening")
	pflag.Int("port", 8080, "The port where the application is going to run")
	pflag.String("log.level", "error", "The log level we're going to be using")

	vip = viper.New()
	vip.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	vip.SetDefault("ip", "0.0.0.0")
	vip.SetDefault("port", "8080")
	vip.SetDefault("log.level", "error")
	vip.AutomaticEnv()
	vip.SetEnvPrefix("SAMPLE")
	vip.BindEnv()

	pflag.Parse()
	vip.BindPFlags(pflag.CommandLine)
	rand.Seed(time.Now().UTC().UnixNano())

}

func main() {
	log.Infof("Welcome to http server example!")

	if vip.GetBool("verbose") {
		vip.Set("log.level", "debug")
	}

	if vip.GetBool("quiet") {
		vip.Set("log.level", "error")
	}

	level, err := log.ParseLevel(vip.GetString("log.level"))
	if err != nil {
		log.Fatalf("Failed during config parsing : %#v", err)
	}

	logger := log.New()
	logger.SetLevel(level)
	formatter := &log.TextFormatter{
		FullTimestamp: true,
	}
	logger.Formatter = formatter
	logger.Out = os.Stdout
	contextLogger = logger.WithFields(log.Fields{
		"service": "sample",
	})
	contextLogger.Debugf("Starting server...")

	serverContext, cancelFunc := context.WithCancel(baseContext)
	server := server.Server{
		Viper: vip,
		Ctx:   &serverContext,
	}
	defer cancelFunc()

	server.SetLogger(logger)

	if pflag.CommandLine.Arg(0) == "server" {
		log.Infof("Starting server...")

		exitChan, err := server.Start()
		if err != nil {
			logger.Errorf("error starting server: %#v", err)
			os.Exit(1)
			return
		}
		result := <-exitChan
		log.Infof("Result: %v", result)
		log.Infof("Exiting...")
		ctx, cancel := context.WithTimeout(baseContext, 5*time.Second)
		defer cancel()
		server.Shutdown(ctx)
		log.Infof("exited with a lot of grace")
	}
}
