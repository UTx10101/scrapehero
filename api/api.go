package api

import (
	// builtin
	"context"
	"net"
	"net/http"
	"os"
	"os/signal"
	"runtime/debug"
	"syscall"
	"time"
	
	// self
	"github.com/UTx10101/scrapehero/config"
	"github.com/UTx10101/scrapehero/db"
	"github.com/UTx10101/scrapehero/middlewares"
	"github.com/UTx10101/scrapehero/models"
	"github.com/UTx10101/scrapehero/routes"
	"github.com/UTx10101/scrapehero/services"
	
	// vendored
	"github.com/apex/log"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/olivere/elastic/v7"
	"github.com/spf13/viper"
)

func main() {
	app := gin.New()
	app.Use(gin.Logger(), gin.Recovery())
	
	if err := config.InitConfig(""); err != nil {
		log.Error("init config error:" + err.Error())
		panic(err)
	}
	log.Info("initialized config successfully")
	
	if err := db.InitMongo(); err != nil {
		log.Error("init mongodb error:" + err.Error())
		debug.PrintStack()
		panic(err)
	}
	log.Info("initialized mongodb successfully")

	if err := services.InitScheduler(); err != nil {
		log.Error("init scheduler error:" + err.Error())
		debug.PrintStack()
		panic(err)
	}
	log.Info("initialized schedule successfully")

	if err := services.InitTaskExecutor(); err != nil {
		log.Error("init task executor error:" + err.Error())
		debug.PrintStack()
		panic(err)
	}
	log.Info("initialized task executor successfully")

	app.Use(middlewares.CORSMiddleware())
	anonymousGroup := app.Group("/")
	{
		anonymousGroup.POST("/login", routes.Login)
	}
	authGroup := app.Group("/", middlewares.AuthMiddleware())
	{
		{
			authGroup.GET("/projects", routes.GetProjects)
			authGroup.GET("/projects/:pid", routes.GetProject)
			authGroup.GET("/projects/:pid/editor/:action", routes.EditProject)
			authGroup.PUT("/projects", routes.CreateProject)
			authGroup.POST("/projects/:pid", routes.ModProject)
			authGroup.DELETE("/projects/:pid", routes.DeleteProject)
		}
		{
			authGroup.GET("/apikeys", routes.GetAPIKeys)
			authGroup.PUT("/apikeys", routes.CreateAPIKey)
			authGroup.POST("/apikeys/:kid/:status", routes.ModAPIKeyStatus)
			authGroup.DELETE("/apikeys/:kid", routes.DeleteAPIKey)
		}
	}
	apiGroup := app.Group("/", middlewares.APIAuthMiddleware())
	{
		{
			apiGroup.POST("/projects/:pid/data", routes.GetProjectData)
		}
	}

	app.GET("/ping", routes.Ping)

	host := viper.GetString("api.host")
	port := viper.GetString("api.port")
	address := net.JoinHostPort(host, port)
	srv := &http.Server{
		Handler: app,
		Addr:    address,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			if err != http.ErrServerClosed {
				log.Error("run server error:" + err.Error())
			} else {
				log.Info("server graceful down")
			}
		}
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	ctx2, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx2); err != nil {
		log.Error("run server error:" + err.Error())
	}
}