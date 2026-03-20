package pprof

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"

	"github.com/quangdtptit/go-cli/config"
	"github.com/quangdtptit/go-cli/pkg/logger"
)

func Start(cfg *config.Config) {
	if cfg.PPROF.Enabled {
		go func() {
			logger.Logger.Info(fmt.Sprintf("PPROF running at [::%d]", cfg.PPROF.Port))
			_ = http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", cfg.PPROF.Port), nil)
		}()
	}
}
