package server

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/adityatresnobudi/go-restapi-gin/config"
	"github.com/adityatresnobudi/go-restapi-gin/docs"
	accHandler "github.com/adityatresnobudi/go-restapi-gin/internal/domain/account/handler"
	accService "github.com/adityatresnobudi/go-restapi-gin/internal/domain/account/service"
	trxHandler "github.com/adityatresnobudi/go-restapi-gin/internal/domain/transaction/handler"
	trxService "github.com/adityatresnobudi/go-restapi-gin/internal/domain/transaction/service"
	"github.com/adityatresnobudi/go-restapi-gin/internal/repositories/account_repo/account_pg"
	"github.com/adityatresnobudi/go-restapi-gin/internal/repositories/transaction_repo/transaction_pg"
	"github.com/adityatresnobudi/go-restapi-gin/pkg/postgres"
	"github.com/gin-gonic/gin"
)

type server struct {
	cfg config.Config
	r   *gin.Engine
}

func NewServer(cfg config.Config) *server {
	return &server{
		cfg: cfg,
		r:   gin.Default(),
	}
}

func (s *server) Run() {
	db, err := postgres.NewDB(
		s.cfg.Postgres.Host,
		s.cfg.Postgres.Port,
		s.cfg.Postgres.User,
		s.cfg.Postgres.Password,
		s.cfg.Postgres.DBName,
	)

	if err != nil {
		log.Fatal(err.Error())
	}

	if err = postgres.InitializeTable(db); err != nil {
		if err = db.Close(); err != nil {
			log.Printf("db graceful shutdown: %s\n", err.Error())
		} else {
			fmt.Printf("db graceful shutdown succeeded\n")
		}
		return
	}

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	defer cancel()
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	signal.Notify(ch, syscall.SIGTERM)

	docs.SwaggerInfo.Host = fmt.Sprintf("%s%s", s.cfg.Http.Host, s.cfg.Http.Port)

	accountRepo := account_pg.NewRepo(db)
	transactionRepo := transaction_pg.NewRepo(db)

	accountService := accService.NewAccountService(accountRepo)
	transactionService := trxService.NewTransactionService(transactionRepo, accountRepo)

	accountHandler := accHandler.NewAccountHandler(s.r, ctx, accountService)
	transactionHandler := trxHandler.NewTransactionHandler(s.r, ctx, transactionService)

	accountHandler.MapRoutes()
	transactionHandler.MapRoutes()

	go func() {
		log.Printf("Listening on PORT: %s\n", s.cfg.Http.Port)
		if err := s.runGinServer(); err != nil {
			log.Printf("s.r.Run: %s\n", err.Error())
		}

	}()

	oscall := <-ch

	if err = db.Close(); err != nil {
		log.Printf("db graceful shutdown: %s\n", err.Error())
	} else {
		fmt.Printf("db graceful shutdown succeeded\n")
	}

	fmt.Printf("system call: %+v\n", oscall)
}
