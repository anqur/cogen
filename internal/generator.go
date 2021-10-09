package internal

import (
	"context"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

type cogen struct {
	DB        *gorm.DB
	Data      interface{}
	TableName string
	RawSQL    string
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

type fixedTableNamer struct {
	tableName string
	schema.NamingStrategy
}

func (n fixedTableNamer) TableName(string) string {
	return n.tableName
}

func MySQL(data interface{}, s *Setting) (*cogen, error) {
	mock, _, err := sqlmock.New()
	if err != nil {
		return nil, err
	}
	g := &cogen{
		Data:      data,
		TableName: s.TableName,
	}
	cfg := &gorm.Config{
		DryRun: true,
		Logger: g,
	}
	if s.TableName != "" {
		cfg.NamingStrategy = fixedTableNamer{tableName: s.TableName}
	}
	db, err := gorm.Open(mysql.New(
		mysql.Config{
			Conn:                      mock,
			SkipInitializeWithVersion: true,
		}),
		cfg,
	)
	if err != nil {
		return nil, err
	}
	g.DB = db
	return g, nil
}

func (c *cogen) String() string {
	m := c.DB.Migrator()
	if name := c.TableName; name != "" {
		m = c.DB.Table(name).Migrator()
	}
	if err := m.CreateTable(c.Data); err != nil {
		panic(err)
	}
	return c.RawSQL
}
