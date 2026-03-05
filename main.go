package main

import (
	"net/http"
	"stress-app/koneksi"
	"stress-app/models"

	"github.com/gin-gonic/gin"
)

func main() {
	// 1. Hubungkan Database
	koneksi.KoneksiDatabase()

	r := gin.Default()

	// 2. Endpoint Select Data (Melihat isi tabel Kategori)
	r.GET("/kategori", func(c *gin.Context) {
		var kategori []models.Kategori
		if err := koneksi.DB.Find(&kategori).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, kategori)
	})

	// 3. Endpoint Tambah Data (Tambah Kategori via POST)
	r.POST("/kategori", func(c *gin.Context) {
		var input models.Kategori
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := koneksi.DB.Create(&input).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, gin.H{"message": "Data berhasil ditambahkan", "data": input})
	})

	r.Run(":8080") // Aplikasi jalan di localhost:8080
}
