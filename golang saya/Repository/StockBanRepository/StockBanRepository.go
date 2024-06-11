package stockbanrepository

import (
	"golang-saya/Models"
	connections "golang-saya/Models/Connections"
	stockbanmodels "golang-saya/Models/StockBanModels"
)

func InsertStockBan(StockBan stockbanmodels.StockBan) Models.BaseResponseModels {
	var query string
	var result Models.BaseResponseModels
	db := connections.DB
	query =  "INSERT INTO StockBan ( `Id`,`Nama_Product`, `Code_Product`, `Stok_Masuk`, `Stok_Ready`, `Product_Terjual`) VALUES (?, ?, ?, ?, ?, ?)"

	tempResult := db.Exec(query, StockBan.Id, StockBan.Nama_Product, StockBan.Code_Product, StockBan.Stok_Masuk, StockBan.Stok_Ready, StockBan.Product_Terjual)

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

func UpdateStockBan(StockBan stockbanmodels.StockBan) Models.BaseResponseModels {
	var query string
	var result Models.BaseResponseModels
	db := connections.DB
	query = "UPDATE StockBan SET Nama_Product = ?, Code_Product = ?, Stok_Masuk = ?, Stok_Ready = ?, Product_Terjual = ? WHERE id = ?"


	tempResult := db.Exec(query, StockBan.Id, StockBan.Nama_Product, StockBan.Code_Product, StockBan.Stok_Masuk, StockBan.Stok_Ready, StockBan.Product_Terjual)
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

func DeleteStokBan(request stockbanmodels.StockBan) Models.BaseResponseModels {
	var query string
	var result Models.BaseResponseModels
	db := connections.DB
	query = "DELETE FROM StockBan WHERE id = ?"

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

func GetStokBanByID(request stockbanmodels.StockBan) Models.BaseResponseModels {
	var query string
	var result Models.BaseResponseModels
	var StokBanYarnmotoshop []stockbanmodels.StockBan
	db := connections.DB

	if request.Id != 0 {
		query = "SELECT * FROM StockBan WHERE Id = ?"
	} else {
		query = "SELECT * FROM StockBan"
	}

	tempResult := db.Raw(query).Find(& StokBanYarnmotoshop)

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
			Data:          StokBanYarnmotoshop,
		}
	}

	return result
}
