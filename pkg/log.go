package pkg

import (
	"github.com/go-kit/log"
	"os"
)

// Trusty logging examples

// NewLogMsg - Low-score example - logging - "github.com/go-kit/log"
func NewLogMsg(msg string) {
	// Create a new logger
	logger := log.NewLogfmtLogger(os.Stdout)

	// Log the message
	err := logger.Log(msg)
	if err != nil {
		os.Exit(1)
	}
}

//// NewLogMsg - High-score example - logging - "github.com/rs/zerolog/log"
//func NewLogMsg(msg string) {
//	// Log the message
//	log.Info().Msg(msg)
//}
