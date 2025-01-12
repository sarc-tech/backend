package main

import (
	"context"
	"flag"
	"net/http"
	"os"
	"time"

	"github.com/go-faster/errors"
	"github.com/go-faster/sdk/app"
	"github.com/go-faster/sdk/zctx"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"

	"github.com/sarc-tech/backend/internal/api"
	"github.com/sarc-tech/backend/internal/httpmiddleware"
	"github.com/sarc-tech/backend/internal/oas"
	incidentRepo "github.com/sarc-tech/backend/internal/repositories/incident"
	incidentUsecase "github.com/sarc-tech/backend/internal/usecases/incident"
	"github.com/sarc-tech/backend/pkg/database"
)

const shutdownTimeout = 15 * time.Second

func main() {
	app.Run(func(ctx context.Context, lg *zap.Logger, m *app.Metrics) error {

		portStr := os.Getenv("PORT")
		strConn := os.Getenv("STRCON")

		if portStr == "" {
			portStr = "8082"
		}

		if strConn == "" {
			strConn = "postgres://user:password@localhost:5543/test?sslmode=disable"
		}

		var arg struct {
			Addr string
		}
		flag.StringVar(&arg.Addr, "addr", "0.0.0.0:"+portStr, "listen address")
		flag.Parse()

		lg.Info("Initializing",
			zap.String("http.addr", arg.Addr),
		)

		db, err := sqlx.Connect("postgres", strConn)
		if err != nil {
			return errors.Wrap(err, "connect to database")
		}

		defer db.Close()

		repo := incidentRepo.NewRepository(database.NewRepository[incidentRepo.Model](db))
		uc := incidentUsecase.NewUsecase(repo)

		oasServer, err := oas.NewServer(api.NewHandler(uc))
		if err != nil {
			return errors.Wrap(err, "server init")
		}

		// Using OpenTelemetry instrumentation for HTTP server.
		routeFinder := httpmiddleware.MakeRouteFinder(oasServer)
		httpServer := http.Server{
			ReadHeaderTimeout: time.Second,
			Addr:              arg.Addr,
			Handler: httpmiddleware.Wrap(oasServer,
				httpmiddleware.InjectLogger(zctx.From(ctx)),
				httpmiddleware.Instrument("api", routeFinder, m),
				httpmiddleware.LogRequests(routeFinder),
				httpmiddleware.Labeler(routeFinder),
			),
		}
		g, ctx := errgroup.WithContext(ctx)
		g.Go(func() error {
			// Wait until g ctx canceled, then try to shut down server.
			<-ctx.Done()

			lg.Info("Shutting down", zap.Duration("timeout", shutdownTimeout))

			shutdownCtx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
			defer cancel()
			return httpServer.Shutdown(shutdownCtx)
		})
		g.Go(func() error {
			defer lg.Info("Server stopped")
			if err := httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
				return errors.Wrap(err, "http")
			}
			return nil
		})

		return g.Wait()
	})
}
