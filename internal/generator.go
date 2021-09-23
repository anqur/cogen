package internal

import (
	"context"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type cogen struct {
	DB     *gorm.DB
	Data   interface{}
	RawSQL string
}

func (c *cogen) LogMode(logger.LogLevel) logger.Interface      { return c }
func (c *cogen) Info(context.Context, string, ...interface{})  {}
func (c *cogen) Warn(context.Context, string, ...interface{})  {}
func (c *cogen) Error(context.Context, string, ...interface{}) {}

func (c *cogen) Trace(
	_ context.Context,
	_ time.Time,
	f func() (sql string, rowsAffected int64),
	_ error,
) {
	sql, _ := f()
	c.RawSQL = sql
}

func MySQL(data interface{}) (*cogen, error) {
	mock, _, err := sqlmock.New()
	if err != nil {
		return nil, err
	}
	g := &cogen{Data: data}
	db, err := gorm.Open(mysql.New(
		mysql.Config{
			Conn:                      mock,
			SkipInitializeWithVersion: true,
		}),
		&gorm.Config{
			DryRun: true,
			Logger: g,
		},
	)
	if err != nil {
		return nil, err
	}
	g.DB = db
	return g, nil
}

func (c *cogen) String() string {
	if err := c.DB.Migrator().CreateTable(c.Data); err != nil {
		panic(err)
	}
	return c.RawSQL
}
