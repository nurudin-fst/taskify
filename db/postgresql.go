package db

import (
	"log"
	"math/rand"
	"os"
	"strconv"
	"sync"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	Pg   *gorm.DB
	once sync.Once
)

func Init() *gorm.DB {
	once.Do(func() {
		username := os.Getenv("DB_USER")
		password := os.Getenv("DB_PASSWORD")
		host := os.Getenv("DB_HOST")
		port := os.Getenv("DB_PORT")
		dbName := os.Getenv("DB_NAME")
		maxPoolOpen, _ := strconv.Atoi(os.Getenv("POOLMAXCONN"))
		maxPoolIdle, _ := strconv.Atoi(os.Getenv("POOLIDLECONN"))
		maxPollLifeTime, _ := strconv.Atoi(os.Getenv("POOLLIFETIME"))

		var logLevel logger.LogLevel
		logLevel = logger.Info

		loggerDb := logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			logger.Config{
				SlowThreshold:             2 * time.Second,
				LogLevel:                  logLevel,
				IgnoreRecordNotFoundError: true,
				Colorful:                  true,
			},
		)
		dsn := "host=" + host + " user=" + username + " password=" + password + " dbname=" + dbName + " port=" + port + " sslmode=disable TimeZone=UTC"
		pG, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
			Logger: loggerDb,
		})
		if err != nil {
			log.Fatal(err)
		}

		sqlDB, err := pG.DB()
		if err != nil {
			log.Fatal(err)
		}

		err = sqlDB.Ping()
		if err != nil {
			log.Fatal(err)
		}

		sqlDB.SetMaxOpenConns(maxPoolOpen)
		sqlDB.SetMaxIdleConns(maxPoolIdle)
		sqlDB.SetConnMaxLifetime(time.Duration(rand.Int31n(int32(maxPollLifeTime))) * time.Millisecond)
		Pg = pG
	})
	return Pg
}
