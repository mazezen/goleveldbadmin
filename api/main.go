package main

import (
	"flag"
	"fmt"
	"github.com/mazezen/goleveldbadmin/framework"
	"github.com/mazezen/goleveldbadmin/internel/controller"
	"github.com/mazezen/goleveldbadmin/sdk"
	"github.com/mazezen/goleveldbadmin/utils"
	"log"
	"net/http"
	"time"
)

var c string

func init() {
	flag.StringVar(&c, "c", "config.yaml", "Start server with provided configuration file")
	flag.Parse()
}

func main() {
	sdk.ParseConfig(c)
	utils.Expire = time.Duration(sdk.Cf.Jwt.ExpireTime) * time.Second
	utils.Secret = sdk.Cf.Jwt.Secret
	var err error
	sdk.CreateLevelDB(sdk.Cf.Source.Dir)

	server := &http.Server{
		Addr:        ":9091",
		Handler:     framework.Router,
		ReadTimeout: 30 * time.Second,
	}
	RegisterRouter(framework.Router)
	fmt.Println("start server success")
	err = server.ListenAndServe()
	if err != nil {
		log.Fatal("start server error")
	}
}

func RegisterRouter(handler *framework.RouterHandler) {
	new(controller.AuthHandler).Router(handler)
	new(controller.DbController).Router(handler)
}
