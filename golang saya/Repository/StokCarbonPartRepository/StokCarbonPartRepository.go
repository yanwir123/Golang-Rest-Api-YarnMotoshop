package stockoilrepository

import (
	"golang-saya/Models"
	connections "golang-saya/Models/Connections"
	stockcarbonpart "golang-saya/Models/StockCarbonPart"
)

func InsertStokCarbonPart(StokCarbonPart stockcarbonpart.StokCarbonPark) Models.BaseResponseModels {
	var query string
	var result Models.BaseResponseModels
	db := connections.DB
	query = "INSERT StockCarbonPart SET Nama Product = ?, Code Product = ?, Stok ProductM = ?, Stok ProductP = ?, Product Terjual =? WHERE id = ?"

	tempResult := db.Exec(query, StokCarbonPart.Id, StokCarbonPart.Nama_Product, StokCarbonPart.Code_Product, StokCarbonPart.Stok_Masuk, StokCarbonPart.Stok_Ready, StokCarbonPart.Product_Terjual)

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

func UpdateStokCarbonPart(StokCarbonPart stockcarbonpart.StokCarbonPark) Models.BaseResponseModels {
	var query string
	var result Models.BaseResponseModels
	db := connections.DB
	query = "UPDATE StockCarbonPart SET Nama Product = ?, Code Product = ?, Stok ProductM = ?, Stok ProductP = ?, Product Terjual =? WHERE id = ?"

	tempResult := db.Exec(query, StokCarbonPart.Id, StokCarbonPart.Nama_Product, StokCarbonPart.Code_Product, StokCarbonPart.Stok_Masuk, StokCarbonPart.Stok_Ready, StokCarbonPart.Product_Terjual)
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

func DeleteStokCarbonPart(request stockcarbonpart.StokCarbonPark) Models.BaseResponseModels {
	var query string
	var result Models.BaseResponseModels
	db := connections.DB
	query = "DELETE FROM StokeCarbonPart WHERE id = ?"

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

func GetStokCarbonPartByID(request stockcarbonpart.StokCarbonPark) Models.BaseResponseModels {
	var query string
	var result Models.BaseResponseModels
	var StokCarbonPartYarnmotoshop []stockcarbonpart.StokCarbonPark
	db := connections.DB

	if request.Id != "" {
		query = "SELECT * FROM StokCarbonPart WHERE Id = ?"
	} else {
		query = "SELECT * FROM StokCarbonPart"
	}

	tempResult := db.Raw(query).Find(& StokCarbonPartYarnmotoshop)

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
			Data:          StokCarbonPartYarnmotoshop,
		}
	}

	return result
}
