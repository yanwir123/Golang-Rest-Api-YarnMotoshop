package main

import (
	controller "golang-saya/Controller"
	connection "golang-saya/Models/Connections"

	"github.com/gin-gonic/gin"
)

func main() {

	port := ":7070"
	r := gin.Default()
	connection.ConnectDatabase()

	 r.Use(func(c *gin.Context) {
        c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
        c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
        if c.Request.Method == "OPTIONS" {
            return
        }
        c.Next()
    })



	//###BEGIN WEB API Keterangan
	// Get data
	r.GET("/api/YarnMotoshop/GetKeterangan", controller.GetKeterangan)

	//Insert data
	r.POST("/api/YarnMotoshop/InsertKeterangan", controller.InsertKeterangan)

	// Update data
	r.PUT("/api/YarnMotoshop/UpdateKeterangan", controller.UpdateKeterangan)

	//Delete data
	r.DELETE("/api/Yarmotoshop/DeleteKeterangan", controller.DeleteKeterangan)
	//###END WEB API Keterangan

	// BEGIN WEP API DATAPRODUCT
	// Get data
	r.GET("/api/YarnMotoshop/GetDataProduct", controller.GetDataProduct)

	//Insert data
	r.POST("/api/YarnMotoshop/InsertDataProduct", controller.InsertDataProduct)

	// Update data
	r.PUT("/api/YarnMotoshop/UpdateDataProduct", controller.UpdateDataProduct)

	//Delete data
	r.DELETE("/api/YarnMotoshop/DeleteDataProduct", controller.DeleteDataProduct)
	//###END WEB API DATA PRODUCT

	
	// BEGIN WEP API DATA KEUANGAN
	// Get data
	r.GET("/api/YarnMotoshop/GetDataKeuangan", controller.GetDataKeuanganByID)

	//Insert data
	r.POST("/api/YarnMotoshop/InsertDataKeuangan", controller.InsertDataKeuangan)

	// Update data
	r.PUT("/api/YarnMotoshop/UpdateDataKeuangan", controller.UpdateDataKeuangan)

	//Delete data
	r.DELETE("/api/YarnMotoshop/DeleteDataKeuangan", controller.DeleteDataKeuangan)
	//###END WEB API DATA KEUANGAN

	// BEGIN WEP API STOCK OIL
	// Get data
	r.GET("/api/YarnMotoshop/GetStokOil", controller.GetStokOilByID)

	//Insert data
	r.POST("/api/YarnMotoshop/InsertStokOil", controller.InsertStockOil)

	// Update data
	r.PUT("/api/YarnMotoshop/UpdateStokOil", controller.UpdateStockOil)

	//Delete data
	r.DELETE("/api/YarnMotoshop/DeleteStokOil", controller.DeleteStokOil)
	//###END WEB API STOCK OIL

	// BEGIN WEP API STOCK BAN
	// Get data
	r.GET("/api/YarnMotoshop/GetStokBan", controller.GetStokBanByID)

	//Insert data
	r.POST("/api/YarnMotoshop/InsertStokBan", controller.InsertStockBan)

	// Update data
	r.PUT("/api/YarnMotoshop/UpdateStokBan", controller.UpdateStockBan)

	//Delete data
	r.DELETE("/api/YarnMotoshop/DeleteStokBan", controller.DeleteStokBan)
	//###END WEB API STOCK BAN

	// BEGIN WEP API STOCK MODIVIKASI
	// Get data
	r.GET("/api/YarnMotoshop/GetStokModivikasi", controller.GetStokModivikasiByID)

	//Insert data
	r.POST("/api/YarnMotoshop/InsertStokModivikasi", controller.InsertStokModivikasi)

	// Update data
	r.PUT("/api/YarnMotoshop/UpdateStokModivikasi", controller.UpdateStokModivikasi)

	//Delete data
	r.DELETE("/api/YarnMotoshop/DeleteStokModivikasi", controller.DeleteStokModivikasi)
	//###END WEB API STOCK MODIVIKASI

	// BEGIN WEP API STOCK CARBON PART
	// Get data
	r.GET("/api/YarnMotoshop/GetStokCarbonPart", controller.GetStokCarbonPartByID)

	//Insert data
	r.POST("/api/YarnMotoshop/InsertStokCarbonPart", controller.InsertStokCarbonPart)

	// Update data
	r.PUT("/api/YarnMotoshop/UpdateStokCarbonPart", controller.UpdateStokCarbonPart)

	//Delete data
	r.DELETE("/api/YarnMotoshop/DeleteStokCarbonPart", controller.DeleteStokCarbonPart)
	//###END WEB API STOCK CARBON PART

	// BEGIN WEP API RATING MEKANIK
	// Get data
	r.GET("/api/YarnMotoshop/GetStokRatingMekanik", controller.GetRatingMekanikByID)

	//Insert data
	r.POST("/api/YarnMotoshop/InsertStokRatiingMekanik", controller.InsertRatingMekanik)

	// Update data
	r.PUT("/api/YarnMotoshop/UpdateStokRatingMekanik", controller.UpdateRatingMekanik)

	//Delete data
	r.DELETE("/api/YarnMotoshop/DeleteStokRatingMekanik", controller.DeleteRatingMekanik)
	//###END WEB API RATING MEKANIK

	
	// BEGIN WEP API RATING MEKANIK
	// Get data
	r.GET("/api/YarnMotoshop/GetProductTeratas", controller.GetProductTeratasByID)

	//Insert data
	r.POST("/api/YarnMotoshop/InsertProductTeratas", controller.InsertProductTeratas)

	// Update data
	r.PUT("/api/YarnMotoshop/UpdateProductTeratas", controller.UpdateProductTeratas)

	//Delete data
	r.DELETE("/api/YarnMotoshop/DeleteProductTeratas", controller.DeleteProductTeratas)
	//###END WEB API RATING MEKANIK

	r.Run(port)
}