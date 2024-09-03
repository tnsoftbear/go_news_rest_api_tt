package storage

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"fiber-reform-rest/internal/infra/config"

	"github.com/go-sql-driver/mysql"
	"gopkg.in/reform.v1"
	dialectsMysql "gopkg.in/reform.v1/dialects/mysql"
)

func Setup(cfg *config.MysqlStorage) *reform.DB {
	dbAddr := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)
	mysqlConfig := mysql.Config{
		User:   cfg.Username,
		Passwd: cfg.Password,
		Net:    "tcp",
		Addr:   dbAddr,
		DBName: cfg.Database,
	}

	// Get a database handle.
	db, err := sql.Open("mysql", mysqlConfig.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	db.SetMaxOpenConns(10) // sets the maximum number of open connections to the database.
	db.SetMaxIdleConns(10) // sets the maximum number of connections in the idle connection pool
	db.SetConnMaxLifetime(time.Minute * 3)

	reformDB := reform.NewDB(db, dialectsMysql.Dialect, reform.NewPrintfLogger(log.Printf))
	return reformDB
}
