package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"misGastos3Backend/domain"
)

func GetDb() (*gorm.DB, error) {
	dsn := "root:nOMAGCFzrF8KPezb1MlY@tcp(containers-us-west-44.railway.app:6421)/railway"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return db, err
}

func SaveProduct(db *gorm.DB, product domain.Product) error {
	db.Create(&product)
	return nil
}

func GetProducts(db *gorm.DB) ([]domain.Product, error) {
	var products []domain.Product
	db.Find(&products)
	return products, nil
}

func GetProductsByDate(db *gorm.DB, date string) ([]domain.Product, error) {
	var products []domain.Product
	db.Where("date = ?", date).Find(&products)
	return products, nil
}
func DeleteProduct(db *gorm.DB, id int) error {
	db.Delete(&domain.Product{}, id)
	return nil
}
