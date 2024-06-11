package datakeuanganmodels

import "github.com/shopspring/decimal"

type DataKeuangan struct {
	Id              	int64   		 `json:"id" from:"id"`
	Bulan           	string  		 `json:"Bulan" from:"Bulan"`
	Nominal         	decimal.Decimal  `json:"Nominal" from:"Nominal"`
	Jumlah_Customer 	float64 		 `json:"Jumlah Customer" from:" Jummlah Customer"`
	Total_Pemasukan 		decimal.Decimal  `json:"Total Masuk" from:"Total Masuk"`
	Total_Pengeluaran 	decimal.Decimal  `json:"Total Pengeluaran" from:"Total Pengeluaran"`
}