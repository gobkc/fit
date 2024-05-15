package middleware

import (
	"context"
	"fmt"
	"gorm.io/gorm/logger"
	"log/slog"
	"time"
)

type GLog struct {
}

func (g *GLog) LogMode(level logger.LogLevel) logger.Interface {
	return g
}

func (g *GLog) Info(ctx context.Context, s string, i ...interface{}) {
	fmt.Println(s, i)
}

func (g *GLog) Warn(ctx context.Context, s string, i ...interface{}) {
	fmt.Println(s, i)
}

func (g *GLog) Error(ctx context.Context, s string, i ...interface{}) {
	fmt.Println(s, i)
}

func (g *GLog) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	sql, affected := fc()
	if err != nil {
		slog.Default().Error(`Sql Execute:`, slog.String(`Begin`, begin.String()), slog.Any(`SQL`, sql), slog.Any(`Affected`, affected), slog.String(`Error`, err.Error()))
	} else {
		slog.Default().Info(`Sql Execute:`, slog.String(`Begin`, begin.String()), slog.Any(`Affected`, affected))
		fmt.Println(sql)
	}
}
