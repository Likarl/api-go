package main

import (
	"api-go/config"
	"api-go/model"
	v "api-go/pkg/version"
	"api-go/router"
	"api-go/router/middleware"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/lexkong/log"

	"github.com/gin-gonic/gin"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var (
	cfg     = pflag.StringP("config", "c", "", "apiserver config file path.")
	version = pflag.BoolP("version", "v", false, "show version info")
)

func main() {
	pflag.Parse()
	// read version info
	if *version {
		v := v.Get()
		marshlled, err := json.MarshalIndent(&v, "", "  ")
		if err != nil {
			fmt.Printf("%v\n", err)
			os.Exit(2)
		}

		fmt.Println(string(marshlled))
		return
	}

	// init config
	if err := config.Init(*cfg); err != nil {
		log.Fatal("Fatal init config", err)
	}

	// init db
	model.DB.Init()
	defer model.DB.Close()

	// running mode
	gin.SetMode(viper.GetString("runmode"))

	// Create the Gin engine
	g := gin.New()

	// // gin middlewares
	// middlewares := []gin.HandlerFunc{}

	//Routers
	router.Load(
		g,
		// middlewares
		middleware.Logging(),
		middleware.RequestId(),
	)

	// Ping the server to make sure the router is working.
	go func() {
		if err := pingServer(); err != nil {
			log.Fatal("The router has no response, or it might took too long to start up.", err)
		}
		log.Info("The router has been deployed successfully.")
	}()

	// Https
	cert := viper.GetString("tls.cert")
	key := viper.GetString("tls.key")
	if cert != "" && key != "" {
		go func() {
			log.Infof("Start to listening the incoming requests on https address: %s", viper.GetString("tls.addr"))
			log.Info(http.ListenAndServeTLS(viper.GetString("tls.addr"), cert, key, g).Error())
		}()
	}

	log.Infof("Start to listening the incoming requests on http address: %s", viper.GetString("addr"))
	log.Infof(http.ListenAndServe(viper.GetString("addr"), g).Error())
}

// pingServer pings the http server to make sure the router is working.
func pingServer() error {
	for i := 0; i < viper.GetInt("max_ping_count"); i++ {
		// Ping the server by sending a GET request to `/health`.
		resp, err := http.Get(viper.GetString("url") + "/sd/health")
		if err == nil && resp.StatusCode == 200 {
			return nil
		}

		// Sleep for a second to continue the next ping.
		log.Info("Waiting for the router, retry in 1 second.")
		time.Sleep(time.Second)
	}
	return errors.New("Cannot connect to the router.")
}
