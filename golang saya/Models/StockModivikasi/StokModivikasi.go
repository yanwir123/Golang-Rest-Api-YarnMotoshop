package stockmodivikasi

type StokModivikasi struct {
	Id              string  `json:"id" from:"id"`
	Nama_Product    string  `json:"Nama Product" from:"Nama Product"`
	Code_Product    string  `json:"Code Product" from:"Code_Product"`
	Stok_Masuk      float64 `json:"Stok Product" from:"Stock_Product"`
	Stok_Ready      float64 `json:"Stok ProductM" from:"Stock_ProductM"`
	Product_Terjual float64 `json:"Product Terjual" from:"Product_Terjual"`
}