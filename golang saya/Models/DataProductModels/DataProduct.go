package DataProductModels

type DataProduct struct {
	Id                    string `json:"id" from:"id"`
	Nama_Product          string `json:"Nama Product" from:"Nama Product"`
	Code_Product          string `json:"Code Product" from:"Code Product"`
	Jumlah_Product        string `json:"Jumlah Product" from:"Jumlah Product"`
	Nama_Distributor      string `josn:"Nama Distributor" from:"Nama Distributor"`
	Nama_Considnee        string `json:"Nama Considnee" from:"Nama Considnee"`
	Tanggal_Masuk_Product string `json:"Tanggal Masuk Product" from:"TAnggal Masuk Product"`
}
