package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/xlzpm/internal/author"
	db_author "github.com/xlzpm/internal/author/db/postgresql"
	"github.com/xlzpm/internal/author/service"
	"github.com/xlzpm/internal/config"
	"github.com/xlzpm/pkg/client/postgresql"
	"github.com/xlzpm/pkg/logging"
)

func main() {
	logger := logging.GetLogger()
	logger.Info("create router")
	router := gin.Default()

	cfg := config.GetGonfig()

	postgreSQLClient, err := postgresql.NewClient(context.TODO(), 3, cfg.Storage)
	if err != nil {
		logger.Fatalf("%v", err)
	}

	repository := db_author.NewRepository(postgreSQLClient, logger)
	logger.Info("register author handler")
	authorService := service.NewService(repository, logger)
	authorHandler := author.NewHandler(authorService, logger)
	authorHandler.Register(router)

	start(router, cfg)
}

func start(router *gin.Engine, cfg *config.Config) {
	logger := logging.GetLogger()
	logger.Info("start app")

	var listener net.Listener
	var listenErr error

	if cfg.Listen.Type == "sock" {
		logger.Info("detect app path")
		appDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			logger.Fatal(err)
		}

		logger.Info("create socket")
		sockerPath := path.Join(appDir, "app.sock")
		logger.Debugf("socket path: %s", sockerPath)

		logger.Info("listen unix socket")
		listener, listenErr = net.Listen("unix", sockerPath)
		logger.Infof("server is listening unix socket: %s", sockerPath)
	} else {
		logger.Info("listen tcp")

		listener, listenErr = net.Listen("tcp",
			fmt.Sprintf("%s:%s", cfg.Listen.BindIP, cfg.Listen.Port))

		logger.Infof("server is listening port %s:%s", cfg.Listen.BindIP, cfg.Listen.Port)
	}

	if listenErr != nil {
		logger.Fatal(listenErr)
	}

	server := &http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	logger.Fatal(server.Serve(listener))
}
