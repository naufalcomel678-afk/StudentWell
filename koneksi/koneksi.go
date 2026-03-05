package koneksi

import (
	//"stress-app/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func KoneksiDatabase() {
	dsn := "root:@tcp(127.0.0.1:3306)/db_stress_baru?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Gagal koneksi database")
	}

	// db.AutoMigrate(
	// 	&models.Role{},
	// 	&models.Jurusan{},
	// 	&models.Kategori{},
	// 	&models.RangeStres{},
	// 	&models.User{},
	// 	&models.Mahasiswa{},
	// 	&models.Pertanyaan{},
	// 	&models.Kuesioner{},
	// 	&models.DetailJawaban{},
	// 	&models.HasilStres{},
	// )
	DB = db
}
