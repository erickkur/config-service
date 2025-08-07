package main

import (
	"flag"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/configservice/adapter/middleware"
	"github.com/configservice/adapter/router"
	"github.com/configservice/infra"
	"github.com/configservice/internal/env"
	"github.com/configservice/internal/logger"

	sl "github.com/configservice/adapter/database/sqllite"
	cModel "github.com/configservice/adapter/models/configuration"
	cService "github.com/configservice/application/services/configuration"
	rs "github.com/configservice/application/services/router"
	cDomain "github.com/configservice/domain/configuration"
)

func main() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	if err := run(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func run() error {
	addr := flag.String("addr", env.AppPort(), "http service address")

	// Internal shared package
	log := logger.New()

	// Infra layer initialization
	// ++++++++++++++++++++++++++++++++++++++++++
	infraObj := infra.Init()
	infraObj.Database.Connect()
	// ++++++++++++++++++++++++++++++++++++++++++

	// Adapter layer initialization
	// ++++++++++++++++++++++++++++++++++++++++++
	routerAdapter := router.NewAdapter()
	middlewareAdapter := middleware.NewAdapter()
	postgresAdapter := sl.NewAdapter(infraObj.Database)

	configurationModel := cModel.NewModel()
	// ++++++++++++++++++++++++++++++++++++++++++

	// Service layer initialization
	// ++++++++++++++++++++++++++++++++++++++++++
	routerService := rs.NewService(
		routerAdapter,
		middlewareAdapter,
		"/api",
	)

	configurationService := cService.NewConfigurationService(
		cService.Dependency{
			ConfigurationModel: configurationModel,
			DBClient:           postgresAdapter,
		},
	)
	// ++++++++++++++++++++++++++++++++++++++++++

	// Domain layer initialization
	// ++++++++++++++++++++++++++++++++++++++++++
	cDomain.NewDomain(cDomain.RouteDependency{
		Context:              routerService,
		Logger:               log,
		ConfigurationService: configurationService,
	})
	// ++++++++++++++++++++++++++++++++++++++++++

	routerAdapter.ReArrange()

	s := &http.Server{
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 10 * time.Second,
		Addr:         *addr,
		Handler:      routerAdapter,
	}

	log.Info(fmt.Sprint("Configuration service started on port", env.AppPort()))

	return s.ListenAndServe()
}
