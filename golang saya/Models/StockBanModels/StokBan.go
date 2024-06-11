package stockbanmodels

type StockBan struct {
	Id              int64   `json:"id" from:"id"`
	Nama_Product    string  `json:"Nama Product" from:"Nama Product"`
	Code_Product    string  `json:"Code Product" from:"Code_Product"`
	Stok_Masuk      float64 `json:"Stok Masuk" from:"Stok Masuk"`
	Stok_Ready      float64 `json:"Stok Ready" from:"Stok Ready"`
	Product_Terjual float64 `json:"Product Terjual" from:"Product_Terjual"`
}