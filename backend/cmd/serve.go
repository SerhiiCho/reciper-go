package cmd

import (
	"context"
	"fmt"
	"github.com/SerhiiCho/reciper/backend/api"
	"github.com/SerhiiCho/reciper/backend/api/middleware"
	"github.com/SerhiiCho/reciper/backend/app"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"
)

func serveAPI(ctx context.Context, api *api.API) {
	middlewares := []gin.HandlerFunc{
		middleware.AppMiddle(),
	}

	router := gin.Default().SecureJsonPrefix("/api")
	router.Use(middlewares...)

	api.Init(router)

	s := &http.Server{
		Addr:        fmt.Sprintf(":%d", api.Config.Port),
		Handler:     cors(router),
		ReadTimeout: 2 * time.Minute,
	}

	done := make(chan struct{})
	go func() {
		<-ctx.Done()
		if err := s.Shutdown(context.Background()); err != nil {
			logrus.Error(err)
		}
		close(done)
	}()

	logrus.Infof("serving api at http://127.0.0.1:%d", api.Config.Port)
	if err := s.ListenAndServe(); err != http.ErrServerClosed {
		logrus.Error(err)
	}
	<-done
}

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "serves the api",
	Run: func(cmd *cobra.Command, args []string) {

		app := app.NewApp()
		api := api.NewAPI(app)

		defer app.Close()

		ctx, cancel := context.WithCancel(context.Background())

		go func() {
			ch := make(chan os.Signal, 1)
			signal.Notify(ch, os.Interrupt)
			<-ch
			logrus.Info("signal caught. shutting down...")
			cancel()
		}()

		var wg sync.WaitGroup

		wg.Add(1)

		go func() {
			defer wg.Done()
			defer cancel()
			serveAPI(ctx, api)
		}()

		wg.Wait()
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
