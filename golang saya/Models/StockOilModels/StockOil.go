package stockoilmodels

type StockOil struct {
	Id              string  `json:"id" from:"id"`
	Nama_Product    string  `json:"Nama Product" from:"Nama Product"`
	Code_Product    string  `json:"Code Product" from:"Code Product"`
	Stok_Masuk      float64 `json:"Stok Masuk" from:"Stock Masuk"`
	Stok_Ready      float64 `json:"Stok Ready" from:"Stock Ready"`
	Product_Terjual float64 `json:"Product Terjual" from:"Product Terjual"`
}