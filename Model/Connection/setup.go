package Connection

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() (*gorm.DB, error) {
	// dsn := "root:@tcp(localhost:3306)/db_buku_tamu?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := "root:@tcp(localhost:3306)/db_mahasiswa?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	DB = db
	return db, nil
}
