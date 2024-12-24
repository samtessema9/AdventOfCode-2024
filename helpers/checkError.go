package helpers

import (
	"log/slog"
	"os"
)

func Check(e error) {
	if e != nil {
		slog.Error(e.Error())
		os.Exit(1)
	}
}