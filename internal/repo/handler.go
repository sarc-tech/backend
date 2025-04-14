package repo

import (
	"fmt"
	"strconv"

	"github.com/sarc-tech/backend/internal/utils"
	"go.uber.org/zap"

	"github.com/jmoiron/sqlx"
)

// Compile-time check for Handler.
type RepoHandler struct {
	Db *sqlx.DB
	sessionTTL *sessionTTLMap
}

func NewCommonDBHandler(lg *zap.Logger) *RepoHandler {
	port, _ := strconv.ParseInt(utils.POSTGRES_PORT, 10, 16)
	connStr := fmt.Sprintf("postgresql://%s:%s@%s:%d/%s?sslmode=disable", utils.POSTGRES_USER, utils.POSTGRES_PASSWORD, utils.POSTGRES_HOST, uint16(port), utils.POSTGRES_DB)
	db, err := sqlx.Connect("postgres", connStr)
	sessionTTL := newSessionTTLMap(7200)
	if err != nil {
		lg.Fatal("error connecting to DB", zap.Error(err))
	}
	return &RepoHandler{
		Db: db,
		sessionTTL: sessionTTL,
	}
}

func (r *RepoHandler) Close() {
	r.Db.Close()
}