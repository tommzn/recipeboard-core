// Package core contains core components for the reipe board project.
package core

import (
	"context"

	config "github.com/tommzn/go-config"
	log "github.com/tommzn/go-log"
	secrets "github.com/tommzn/go-secrets"
)

// Returns a new stdout looger with log level error.
func newLogger(conf config.Config) log.Logger {
	logger := log.NewLoggerFromConfig(conf, secrets.NewSecretsManager())
	ctxValues := make(map[string]string)
	ctxValues[log.LogCtxNamespace] = "recipeboard-core"
	logger.WithContext(log.LogContextWithValues(context.Background(), ctxValues))
	return logger
}
