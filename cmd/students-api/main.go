package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/bishal05das/students-api/internal/config"
	"github.com/bishal05das/students-api/internal/http/handlers/student"
)

func main() {
	//load config
	cfg := config.MustLoad()
	// database setup

	//router setup
	router := http.NewServeMux()

	router.HandleFunc("POST /api/students", student.New())

	//setup server

	server := http.Server{
		Addr:    cfg.Addr,
		Handler: router,
	}
	fmt.Println("server started")
	slog.Info("server started",slog.String("address",cfg.Addr))

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
    //graceful shutdown
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Fatal("failed to start server")
		}
	}()
	<-done
	slog.Info("shutting down server gracefully")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	if err:= server.Shutdown(ctx); err != nil {
		slog.Error("failed to shutdown server", slog.String("error",err.Error()))
	}

	slog.Info(("server shutdown gracefully"))

}
