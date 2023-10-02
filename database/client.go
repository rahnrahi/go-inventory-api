package database

import (
	"inventory-app/config"
	"inventory-app/entities"
	"inventory-app/utilities"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Instance *gorm.DB
var err error

/**
 * Check the connection and make it if it is lost or not opened yet.
 * @return	error
 */
func checkConnection() error {
	// If already connected then OK
	if Instance != nil {
		return nil
	}
	config.LoadAppConfig()
	var constring = config.AppConfig.ConnectionString
	// Else try to connect
	conn, err := Connect(constring)

	// If error then return with it
	if err != nil {
		return err
	}

	// Else set the connection
	Instance = conn

	// Tell that it is OK
	return nil
}

/**
 * openDatabase function opens a MySql database connection
 */

func Connect(connectionString string) (*gorm.DB, error) {
	Instance, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		panic("Cannot connect to DB")
	}
	log.Println("Connected to Database...")
	return Instance, nil
}

func Migrate() {
	Instance.AutoMigrate(&entities.Product{}, &entities.User{})
	log.Println("Database Migration Completed...")
}

/**
 * SqlQuery function runs a select query to MySql
 * @param 	string
 * @param	[]string
 */
func SqlQuery(sqlQuery string, sqlArgument []string) (*gorm.DB, error) {

	err := checkConnection()
	if err != nil {
		utilities.Log("MySQL: DB Query failed. " + err.Error())
		return nil, err
	}
	//	defer db.Close()

	sqlInterface := utilities.StringArrayToInterface(sqlArgument)

	// query
	rows := Instance.Exec(sqlQuery, sqlInterface...)
	if err != nil {
		utilities.Log("MySQL: SQL query failed. " + err.Error())
		return nil, err
	}

	return rows, nil
}
