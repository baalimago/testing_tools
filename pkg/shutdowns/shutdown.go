package shutdowns

import (
	"context"
	"os"

	"github.com/baalimago/go_away_boilerplate/pkg/ancli"
)

// MonitorShutdowns listens for a shutdown signal and cancels the context
// if the signal is received. If the signal is received again, it will
// force a shutdown.
func MonitorShutdowns(signalCh chan os.Signal, cancel context.CancelFunc) {
	amountOfCancels := 0
	for {
		select {
		case <-signalCh:
			if amountOfCancels == 0 {
				ancli.PrintWarn("initiated forceful shutdown\n")
				cancel()
			} else if amountOfCancels == 1 {
				ancli.PrintWarn("graceful shutdown ongoing, cancel again to force shutdown\n")
			} else {
				ancli.PrintWarn("forcing shutdown\n")
				os.Exit(1)
			}
			amountOfCancels++
		}
	}
}
