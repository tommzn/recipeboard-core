// Contains core components for the reipe board project.
package core

import "gitlab.com/tommzn-go/utils/log"

// Returns a new stdout looger with log level error.
func newLogger() log.Logger {
	return log.NewLogger(log.Error, "recipeboard-core")
}
