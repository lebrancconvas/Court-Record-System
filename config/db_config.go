package config

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type SqlLogger struct{
	logger.Interface
}

func (l SqlLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	sql, _ := fc();
	fmt.Printf("%v\n=====================\n", sql);
}

var Config *gorm.Config = &gorm.Config{
	DryRun: true,
	Logger: &SqlLogger{},
}