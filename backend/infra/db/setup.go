package db

import (
	"log"
	"net"
	"os"
	"time"

	gomysql "github.com/go-sql-driver/mysql"
	"github.com/iChemy/MyKnoQ/backend/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// DatabaseRepository is an interface that defines the methods for interacting with the database.
type DatabaseRepository interface{}

// KnoQDatabaseAPI represents the database API for KnoQ.
// It holds a gorm.DB instance to interact with the database.
type KnoQDatabaseAPI struct {
	db *gorm.DB
}

// Setup initializes the database connection using environment variables.
// It returns a pointer to KnoQDatabaseAPI and an error if any required environment variable is missing or if the connection fails.
func Setup() (*KnoQDatabaseAPI, error) {
	user, err := utils.GetRequiredEnv("MARIADB_USER")
	if err != nil {
		return nil, err
	}

	password, err := utils.GetRequiredEnv("MARIADB_PASSWORD")
	if err != nil {
		return nil, err
	}

	host, err := utils.GetRequiredEnv("MARIADB_HOST")
	if err != nil {
		return nil, err
	}

	database, err := utils.GetRequiredEnv("MARIADB_DATABASE")
	if err != nil {
		return nil, err
	}

	port, err := utils.GetRequiredEnv("MARIADB_PORT")
	if err != nil {
		return nil, err
	}

	addr := net.JoinHostPort(host, port)

	// TZ は必須じゃなくても良い
	tz := os.Getenv("TZ")
	loc, err := time.LoadLocation(tz)
	if err != nil {
		// Log the error and set location to UTC
		log.Printf("invalid TZ environment variable: %v, defaulting to UTC\n", err)
		loc = time.UTC
	}

	c := gomysql.Config{
		User:                 user,
		Passwd:               password,
		Net:                  "tcp",
		Addr:                 addr,
		DBName:               database,
		Loc:                  loc,
		AllowNativePasswords: true,
		ParseTime:            true,
	}

	db, err := gorm.Open(mysql.New(
		mysql.Config{
			DSNConfig:                 &c,
			DefaultStringSize:         256,   // default size for string fields
			DisableDatetimePrecision:  true,  // disable datetime precision, which not supported before MySQL 5.6
			DontSupportRenameIndex:    true,  // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
			DontSupportRenameColumn:   false, // `change` when rename column, rename column not supported before MySQL 8, MariaDB
			SkipInitializeWithVersion: false, // auto configure based on currently MySQL version
		},
	), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		// Log the error before returning
		log.Printf("failed to connect to database: %v", err)
		return nil, err
	}

	return &KnoQDatabaseAPI{
		db: db,
	}, nil

}
