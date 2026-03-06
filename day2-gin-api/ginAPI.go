package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Customer struct {
	Name       string `json:"name"`
	Email      string `json:"email"`
	SecretCode string `json:"-"`
}

type Items struct {
	ProductName string `json:"product_name"`
	Price       int    `json:"price"`
	Quantity    int    `json:"quantity"`
}

type Orders struct {
	ID       int      `json:"id"`
	Customer Customer `json:"customer"`
	Items    []Items  `json:"items"`
}

func main() {
	r := gin.Default()
	r.POST("/terima-order", func(c *gin.Context) {
		var inputData Orders
		err := c.ShouldBindJSON(&inputData)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": "Gagal",
				"pesan":  "Data JSON berantakan atau salah ketik!",
				"error":  err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"status":          "Sukses!",
			"pesan":           "Order berhasil diterima server!",
			"data_yang_masuk": inputData,
		})
	})

	r.GET("/menu-hari-ini", func (c*gin.Context) {
		c.JSON(200,gin.H{
			"status":"Toko Buka",
			"pesan":"Silakan pilih menu dibawah ya kak",
			"daftar_menu":[]string{
				"Melon Tea - Rp.10.000",
				"Lemon Tea - Rp.10.000",
				"Kopi Susu - Rp.15.000",
			}
		})
		
	})

	r.Run(":8080")

}
