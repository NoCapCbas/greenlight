package main

import (
  "context"
  "database/sql"
  "log/slog"
  "flag"
  "net/http"
  "os"
  "time"
  "fmt"

  _ "github.com/lib/pq"
)

// Application version number
const version = "1.0.0"

// Holds all configurations settings for application
type config struct {
  port int
  env string
  db struct {
    dsn string
  }
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
  // database dsn
  flag.StringVar(&cfg.db.dsn, "db-dsn", fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", os.Getenv("POSTGRES_USERNAME"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_DATABASE")), "PostgreSQL DSN")
  flag.Parse()
  // Initialize structured logger which writes log entries to standard stream
  logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
  
  db, err := openDB(cfg)
  if err != nil {
    logger.Error(err.Error())
    os.Exit(1)
  }
  defer db.Close()
  logger.Info("database connection pool established")

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

  err = srv.ListenAndServe()
  logger.Error(err.Error())
  os.Exit(1)

}

func openDB(cfg config) (*sql.DB, error) {
  // create empty pool connection
  fmt.Println(cfg.db.dsn)
  db, err := sql.Open("postgres", cfg.db.dsn)
  if err != nil {
    return nil, err
  }
  
  // create context with a 5-second timeout deadline
  ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
  defer cancel()
 
  err = db.PingContext(ctx)
  if err != nil {
    fmt.Println("failed")
    db.Close()
    return nil, err
  }

  return db, nil
}


