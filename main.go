package main

import (
	"github.com/julienschmidt/httprouter"
	"github.com/mgutz/logxi/v1"
	"github.com/namsral/flag"

	"net/http"
	"os"
)

var logger log.Logger

var (
	Host   string
	Port   string
	DBPath string
)

var broker *Broker

func init() {

	logger = log.NewLogger(os.Stdout, "botyard:"+VERSION)

	flag.StringVar(&Host, "host", "", "host address")
	flag.StringVar(&Port, "port", "7000", "host port")
	flag.StringVar(&DBPath, "dbpath", "/var/data/botyardq", "db path")

}

func main() {
	flag.Parse()
	addr := Host + ":" + Port

	logger.Info("botyardq", "path", DBPath, "host", Host, "port", Port)

	broker = NewBroker(DBPath)
	err := broker.Init()
	if err != nil {
		logger.Error("broker", "err", err)
		os.Exit(1)
	}

	router := httprouter.New()
	router.POST("/v1/queues/:queue", PushHandler)
	router.GET("/v1/queues/:queue", PopHandler)
	router.DELETE("/v1/queues/:queue/:id", DeleteHandler)

	logger.Info("Serving at " + addr)
	err = http.ListenAndServe(addr, router)
	if err != nil {
		log.Error("listen", "err", err)
		os.Exit(1)
	}

	logger.Info("bye")
}
