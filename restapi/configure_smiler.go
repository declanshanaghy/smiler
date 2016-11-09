package restapi

import (
	"crypto/tls"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/declanshanaghy/smiler/restapi/operations"
	"github.com/declanshanaghy/smiler/restapi/operations/flash"
	"github.com/declanshanaghy/smiler/restimpl"
	"github.com/go-openapi/swag"
	"github.com/declanshanaghy/smiler/log"
)

// This file is safe to edit. Once it exists it will not be overwritten

type CmdOptions struct {
	LogFile     string `short:"l" long:"logfile" description:"Specify the log file" default:""`
	Verbose     bool   `short:"v" long:"verbose" description:"Show verbose debug information"`                    // always defaults to false
	VeryVerbose bool   `short:"V" long:"very-verbose" description:"Show verbose debug information including aws"` // always defaults to false
	StaticDir   string `short:"s" long:"static" description:"The path to the static dirs" default:""`             // default auto
}

var CmdOptionsValues CmdOptions // export for testing

func configureFlags(api *operations.SmilerAPI) {
	_ = api
	api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{
		swag.CommandLineOptionsGroup{
			ShortDescription: "Smiler Server Flags",
			LongDescription:  "Smiler Server Flags",
			Options:          &CmdOptionsValues,
		},
	}
}

func configureAPI(api *operations.SmilerAPI) http.Handler {
	log.SetDebug(CmdOptionsValues.Verbose || CmdOptionsValues.VeryVerbose)
	if CmdOptionsValues.LogFile != "" {
		log.SetOutput(CmdOptionsValues.LogFile)
	}

	// configure the api here
	api.ServeError = errors.ServeError
	api.JSONConsumer = runtime.JSONConsumer()
	api.JSONProducer = runtime.JSONProducer()

	api.FlashGetFlashStateHandler = flash.GetFlashStateHandlerFunc(func(params flash.GetFlashStateParams) middleware.Responder {
		return restimpl.HandleApiRequestWithError(restimpl.NewFlashMgr().GetFlash())
	})
	api.FlashSetFlashStateHandler = flash.SetFlashStateHandlerFunc(func(params flash.SetFlashStateParams) middleware.Responder {
		return middleware.NotImplemented("operation flash.SetFlashState has not yet been implemented")
	})

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	_ = tlsConfig
	// Make all necessary changes to the TLS configuration here.
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
