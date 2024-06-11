package KeteranganModels

type KeteranganBengkel struct {
	Id            string  `json:"id" from:"id"`
	Nama_Staff    string  `json:"Nama Staff" from:"Nama Staff"`
	Tanggal_Masuk string  `json:"Tanggal Masuk" from:"Tanggal Masuk"`
	Jabatan       string  `json:"Jabatam" from:"JabatanS"`
	Gender        string  `json:"Gender" from:"Gender"`
	Keterangan    string  `json:"Keterangan" from:"Keterangan"`
	Baik          float64 `jsom:"Baik" from:"Baik"`
	Buruk         float64 `json:"Buruk" from:"Buruk"`
}
