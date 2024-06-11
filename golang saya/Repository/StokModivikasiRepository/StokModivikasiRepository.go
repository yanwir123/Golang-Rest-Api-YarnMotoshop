package stockoilrepository

import (
	"golang-saya/Models"
	connections "golang-saya/Models/Connections"
	stockmodivikasi "golang-saya/Models/StockModivikasi"
)

func InsertStokModifikasi(StokModivikasi stockmodivikasi.StokModivikasi) Models.BaseResponseModels {
	var query string
	var result Models.BaseResponseModels
	db := connections.DB
	query = "INSERT StockOil SET Nama Product = ?, Code Product = ?, Stok Masuk = ?, Stok Ready = ?, Product Terjual = ? WHERE id = ?"

	tempResult := db.Exec(query, StokModivikasi.Id, StokModivikasi.Nama_Product, StokModivikasi.Code_Product, StokModivikasi.Stok_Masuk, StokModivikasi.Stok_Ready, StokModivikasi.Product_Terjual)

	if tempResult.Error != nil {
		result = Models.BaseResponseModels{
			CodeResponse:  400,
			HeaderMessage: "Error",
			Message:       tempResult.Error.Error(),
			Data:          nil,
		}
	} else {
		result = Models.BaseResponseModels{
			CodeResponse:  200,
			HeaderMessage: "Success",
			Message:       "Stok berhasil ditambahkan.",
			Data:          nil,
		}
	}

	return result
}

func UpdateStokModivikasi(StokModivikasi stockmodivikasi.StokModivikasi) Models.BaseResponseModels {
	var query string
	var result Models.BaseResponseModels
	db := connections.DB
	query = "UPDATE StockModivikasi SET Nama Product = ?, Code Product = ?, Stok Masuk = ?, Stok Ready = ?, Product Terjual = ? WHERE id = ?"

	tempResult := db.Exec(query, StokModivikasi.Id, StokModivikasi.Nama_Product, StokModivikasi.Code_Product, StokModivikasi.Stok_Masuk, StokModivikasi.Stok_Ready, StokModivikasi.Product_Terjual)
	if tempResult.Error != nil {
		result = Models.BaseResponseModels{
			CodeResponse:  400,
			HeaderMessage: "Error",
			Message:       tempResult.Error.Error(),
			Data:          nil,
		}
	} else {
		rowsAffected := tempResult.RowsAffected
		if rowsAffected == 0 {
			result = Models.BaseResponseModels{
				CodeResponse:  404,
				HeaderMessage: "Not Found",
				Message:       "Stok dengan ID tersebut tidak ditemukan.",
				Data:          nil,
			}
		} else {
			result = Models.BaseResponseModels{
				CodeResponse:  200,
				HeaderMessage: "Success",
				Message:       "Stok Berhasil Diperbaharui.",
				Data:          nil,
			}
		}
	}

	return result
}

func DeleteStokModivikasi(request stockmodivikasi.StokModivikasi) Models.BaseResponseModels {
	var query string
	var result Models.BaseResponseModels
	db := connections.DB
	query = "DELETE FROM StokeModivikasi WHERE id = ?"

	tempResult := db.Exec(query, request.Id)

	if tempResult.Error != nil {
		result = Models.BaseResponseModels{
			CodeResponse:  400,
			HeaderMessage: "Error",
			Message:       tempResult.Error.Error(),
			Data:          nil,
		}
	} else {
		// Periksa apakah ada baris yang terpengaruh oleh perintah DELETE
		rowsAffected := tempResult.RowsAffected
		if rowsAffected == 0 {
			result = Models.BaseResponseModels{
				CodeResponse:  404,
				HeaderMessage: "Not Found",
				Message:       "Stok dengan ID tersebut tidak ditemukan.",
				Data:          nil,
			}
		} else {
			result = Models.BaseResponseModels{
				CodeResponse:  200,
				HeaderMessage: "Success",
				Message:       "Stok Berhasil Dihapus.",
				Data:          nil,
			}
		}
	}

	return result
}

func GetStokModivikasiByID(request stockmodivikasi.StokModivikasi) Models.BaseResponseModels {
	var query string
	var result Models.BaseResponseModels
	var StokModivikasiYarnmotoshop []stockmodivikasi.StokModivikasi
	db := connections.DB

	if request.Id != "" {
		query = "SELECT * FROM StokModivikasi WHERE id = ?"
	} else {
		query = "SELECT * FROM StokModivikasi"
	}

	tempResult := db.Raw(query).Find(& StokModivikasiYarnmotoshop)

	if tempResult.Error != nil {
		result = Models.BaseResponseModels{
			CodeResponse:  400,
			HeaderMessage: "Error",
			Message:       tempResult.Error.Error(),
			Data:          nil,
		}
	} else {	
		result = Models.BaseResponseModels{
			CodeResponse:  200,
			HeaderMessage: "Success",
			Message:       "Data retrieved successfully",
			Data:          StokModivikasiYarnmotoshop,
		}
	}

	return result
}
