package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"

	_ "github.com/lib/pq"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"

	"net/http"
	"os"
	"os/signal"
	"service-trnng/account"
	"syscall"
)

<<<<<<< HEAD
func main() {
	var httpAdr = flag.String("http", ":8084", "http listen address")
	var dbPass = flag.String("dbpass", "", "Enter DB Pass")
	flag.Parse()
	var dbSource = fmt.Sprintf("postgres://jddbjrse:%s@chunee.db.elephantsql.com/jddbjrse", *dbPass)
=======
const dbSource = "postgres://jddbjrse:gbNgFSw3wOF74Kj1SookyAQ21eJWr81W@chunee.db.elephantsql.com/jddbjrse"

func main() {
	var httpAdr = flag.String("http", ":8084", "http listen address")
>>>>>>> 8c5efbfe95ad9853411449a9a7011f5c81b86b2f
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.NewSyncLogger(logger)
		logger = log.With(logger,
			"service", "account",
			"time: ", log.DefaultTimestampUTC,
			"caller", log.DefaultCaller,
		)
	}

	level.Info(logger).Log("msg", "service started")
	defer level.Info(logger).Log("msg", "service ended")

	var db *sql.DB
	{
		var err error
		db, err = sql.Open("postgres", dbSource)
		if err != nil {
			level.Error(logger).Log("exit", err)
			os.Exit(-1)
		}
	}

<<<<<<< HEAD
=======
	flag.Parse()
>>>>>>> 8c5efbfe95ad9853411449a9a7011f5c81b86b2f
	ctx := context.Background()
	var srv account.Service
	{
		repository := account.NewRepo(db, logger)
		srv = account.NewService(repository, logger)
	}

	errs := make(chan error)

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	endpoints := account.MakeEndpoints(srv)

	go func() {
		fmt.Println("listening on port", *httpAdr)
		handler := account.NewHTTPServer(ctx, endpoints)
		errs <- http.ListenAndServe(*httpAdr, handler)
	}()

	level.Error(logger).Log("exit", <-errs)
}
