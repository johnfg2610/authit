package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"gocloud.dev/runtimevar"
	"gocloud.dev/secrets"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/johnfg2610/authit/pkg/password"
	"gocloud.dev/server"
	"gocloud.dev/server/requestlog"
)

type Config struct {
	GormConnectURL []byte
}

var (
	serverOptions = &server.Options{
		//TODO customize logger maybe using zerolog?
		RequestLogger: requestlog.NewNCSALogger(os.Stdout, func(error) {}),
	}
)

func NewVariable(ctx context.Context) (*runtimevar.Variable, error) {
	v, err := runtimevar.OpenVariable(ctx, "constant://?val=hello+world&decoder=string")
	if err != nil {
		return nil, err
	}
	return v, nil
}

func NewSecrets(ctx context.Context) (*secrets.Keeper, error) {
	keyKeeper, err := secrets.OpenKeeper(ctx, "base64key://")
	if err != nil {
		return nil, err
	}
	return keyKeeper, nil
}

func NewGormDB(ctx context.Context, vari *runtimevar.Variable, secr *secrets.Keeper) (*gorm.DB, error) {
	snapshot, err := vari.Latest(ctx)
	if err != nil {
		return nil, err
	}

	cnf, ok := snapshot.Value.(Config)

	if !ok {
		return nil, fmt.Errorf("Snapshot value was not of expected type")
	}

	plain, err := secr.Decrypt(ctx, cnf.GormConnectURL)
	if err != nil {
		return nil, err
	}

	return gorm.Open(sqlite.Open(string(plain)), &gorm.Config{})
}

func NewPasswordHasher() (password.Hasher, error) {
	return password.NewBcryptHasher(), nil
}

func main() {
	//e, err := InitializeEvent()

	router := chi.NewRouter()
	srv := server.New(router, serverOptions)

	router.Use(middleware.RequestID)
	//router.Use(middleware.RealIP)
	//router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	if err := srv.ListenAndServe(":8080"); err != nil {
		log.Fatalf("%v", err)
	}
}
