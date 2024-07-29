package main

import (
  "log/slog"
  "flag"
  "net/http"
  "os"
  "time"
  "fmt"
)

// Application version number
const version = "1.0.0"

// Holds all configurations settings for application
type config struct {
  port int
  env string
}
// Holds application dependencies
type application struct {
  config config
  logger *slog.Logger
}

func main() {
  // instance of config struct 
  var cfg config
  // Config variables set via flags
  flag.IntVar(&cfg.port, "port", 4000, "API Server Port")
  flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")
  flag.Parse()
  // Initialize structured logger which writes log entries to standard stream
  logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
  
  // declare application
  app := &application{
    config: cfg,
    logger: logger,
  }
  
  srv := &http.Server{
    Addr: fmt.Sprintf(":%d", cfg.port),
    Handler: app.routes(),
    IdleTimeout: time.Minute,
    ReadTimeout: 5 * time.Second,
    WriteTimeout: 10 * time.Second,
    ErrorLog: slog.NewLogLogger(logger.Handler(), slog.LevelError),
  }

  // Start http server
  logger.Info("starting server", "addr", srv.Addr, "env", cfg.env)

  err := srv.ListenAndServe()
  logger.Error(err.Error())
  os.Exit(1)

}







