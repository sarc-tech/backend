package controller

import (
	"fmt"
	"strconv"

	"github.com/sarc-tech/backend/internal/utils"
	"go.uber.org/zap"

	"github.com/jmoiron/sqlx"
	"github.com/syncMutex/sarc_user_lib/lib"
)

// Compile-time check for Handler.
type Controller struct {
	Db *sqlx.DB
	authLib *lib.AuthClient
}

func NewController(lg *zap.Logger) *Controller {
	port, _ := strconv.ParseInt(utils.POSTGRES_PORT, 10, 16)
	connStr := fmt.Sprintf("postgresql://%s:%s@%s:%d/%s?sslmode=disable", utils.POSTGRES_USER, utils.POSTGRES_PASSWORD, utils.POSTGRES_HOST, uint16(port), utils.POSTGRES_DB)
	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		lg.Fatal("error connecting to DB", zap.Error(err))
	}
	authLib := lib.NewAuthClient("localhost")
	return &Controller{
		Db: db,
		authLib: authLib,
	}
}

func (r *Controller) Close() {
	r.Db.Close()
}