package gorm

import (
	"fmt"
	"go-todolist/utils/log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitMySQL() *gorm.DB {
	// Load .env file
	errEnv := godotenv.Load()
	if errEnv != nil {
		log.Panic("Failed to load env file")
	}

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	// Create the connection string
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPassword, dbHost, dbPort, dbName)

	// Open connection to the database
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		// Show all sql query
		// Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		fmt.Println("Failed to connect database : " + err.Error())
		log.Panic("Failed to connect database : " + err.Error())
	}

	sqlDB, err := db.DB()

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Hour)

	// SetConnIdleTime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(10 * time.Hour)

	// Migrate the schema
	// db.AutoMigrate(&model.User{})

	return db
}

func Close(db *gorm.DB) {
	// Get the underlying sql.DB instance from the gorm.DB instance.
	dbSQL, err := db.DB()

	if err != nil {
		fmt.Println("Failed to close connection form database : " + err.Error())
		log.Panic("Failed to close connection form database : " + err.Error())
	}

	// Close the database connection
	dbSQL.Close()
}
