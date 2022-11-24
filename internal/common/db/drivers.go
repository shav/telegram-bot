package db

import (
	"database/sql"
	"database/sql/driver"
)

// RegisterDriverIfNotExists регистрирует драйвер БД, если он еще не зарегистрирован ранее.
func RegisterDriverIfNotExists(name string, driver driver.Driver) {
	existingDrivers := sql.Drivers()
	for _, existingDriver := range existingDrivers {
		if name == existingDriver {
			return
		}
	}
	sql.Register(name, driver)
}
