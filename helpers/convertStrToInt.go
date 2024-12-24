package helpers

import (
	"log/slog"
	"os"
	"strconv"
)

func ConvertStrToInt(s string) int {
	num, err := strconv.Atoi(s)
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	return num
}
