package main

import (
	"context"
	"flag"
	"net/http"
	"os"
	"os/signal"
	"time"

	"net"

	"github.com/8tomat8/yetAnotherCRUD/api/HTTPHandlers"
	"github.com/8tomat8/yetAnotherCRUD/api/router"
	"github.com/8tomat8/yetAnotherCRUD/storage"
	"github.com/go-chi/chi"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

var (
	APIPort         = flag.String("apiport", "8080", "Port for API listener")
	APIHost         = flag.String("apihost", "0.0.0.0", "Host for API listener")
	DBPort          = flag.String("dbport", "3306", "MySQL port")
	DBHost          = flag.String("dbhost", "127.0.0.1", "MySQL host")
	DBUser          = flag.String("dbuser", "root", "MySQL user")
	DBPassword      = flag.String("dbpass", "root", "MySQL password")
	shutdownTimeout = flag.Duration("shutdownTimeout", 30*time.Second, "Shutdown timeout")
)

func main() {
	flag.Parse()
	ctx, cancel := context.WithCancel(context.Background())

	logger := logrus.New()
	logger.SetLevel(logrus.DebugLevel)

	store, err := storage.New(logger, *DBHost, *DBPort, *DBUser, *DBPassword)
	if err != nil {
		logger.Fatal(errors.Wrap(err, "cannot create new storage"))
	}

	APIAddr := net.JoinHostPort(*APIHost, *APIPort)
	srv := http.Server{
		Addr: APIAddr,
		Handler: chi.ServerBaseContext(ctx, router.NewRouter(
			HTTPHandlers.UsersHandler(store, logger),
		)),
	}

	done := make(chan struct{})
	stop := make(chan os.Signal)
	signal.Notify(stop, os.Interrupt)

	go func() {
		select {
		case <-stop:
		case <-ctx.Done():
		}
		cancel()

		// Creating new context with timeout for Shutdown only
		ctx, _ := context.WithTimeout(context.Background(), *shutdownTimeout*time.Second)

		if err := srv.Shutdown(ctx); err != nil {
			logger.Error(errors.Wrap(err, "cannot shutdown gracefully"))
		}

		close(done)
	}()

	logger.Infof("Starting listener on %s", APIAddr)
	err = srv.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		cancel()
		logger.Error(err)
	}

	select {
	case <-done:
		logger.Info("Application stopped gracefully")
	case <-stop:
		logger.Warn("Received second SIGINT. Stopping immediately")
	}
}
