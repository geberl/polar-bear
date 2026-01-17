package cmd

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/lmittmann/tint"
)

func MakeLogger(level string, format string) (*slog.Logger, error) {
	lv := slog.LevelInfo
	err := lv.UnmarshalText([]byte(level))
	if err != nil {
		return nil, fmt.Errorf("unrecognized loglevel flag or env var, must be one of debug/info/warn/error if given")
	}

	switch format {
	case "human":
		return slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.Level(lv)})), nil
	case "json":
		return slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.Level(lv)})), nil
	case "color":
		return slog.New(tint.NewHandler(os.Stderr, nil)), nil
	default:
		return nil, fmt.Errorf("unrecognized logformat flag or env var, must be one of human/json if given")
	}
}
