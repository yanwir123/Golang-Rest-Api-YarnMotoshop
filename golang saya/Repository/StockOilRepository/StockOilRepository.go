package stockoilrepository

import (
	"golang-saya/Models"
	connections "golang-saya/Models/Connections"
	stockoilmodels "golang-saya/Models/StockOilModels"
)

func InsertStockOil(StokOil stockoilmodels.StockOil) Models.BaseResponseModels {
    var query string
    var result Models.BaseResponseModels
    db := connections.DB
    query = "INSERT INTO StockOil ( `Id`,`Nama_Product`, `Code_Product`, `Stok_Masuk`, `Stok_Ready`, `Product_Terjual`) VALUES (?, ?, ?, ?, ?, ?)"

    tempResult := db.Exec(query, StokOil.Id , StokOil.Nama_Product, StokOil.Code_Product, StokOil.Stok_Masuk, StokOil.Stok_Ready, StokOil.Product_Terjual)

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


func UpdateStockOil(StokOil stockoilmodels.StockOil) Models.BaseResponseModels {
	var query string
	var result Models.BaseResponseModels
	db := connections.DB
	query = "UPDATE StockOil SET Nama Product = ?, Code Product = ?, Stok ProductM = ?, Stok ProductP = ?, Product Terjual = ? WHERE id = ?"

	tempResult := db.Exec(query, StokOil.Id, StokOil.Nama_Product, StokOil.Code_Product, StokOil.Stok_Masuk, StokOil.Stok_Ready, StokOil.Product_Terjual)
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

func DeleteStokOil(request stockoilmodels.StockOil) Models.BaseResponseModels {
	var query string
	var result Models.BaseResponseModels
	db := connections.DB
	query = "DELETE FROM StokeOil WHERE id = ?"

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

func GetStokOilByID(request stockoilmodels.StockOil) Models.BaseResponseModels {
	var query string
	var result Models.BaseResponseModels
	var StokOilYarnmotoshop []stockoilmodels.StockOil
	db := connections.DB

	if request.Id != "" {
		query = "SELECT * FROM StockOil WHERE id = ?"
	} else {
		query = "SELECT * FROM StockOil"
	}

	tempResult := db.Raw(query).Find(& StokOilYarnmotoshop)

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
			Data:          StokOilYarnmotoshop,
		}
	}

	return result
}
