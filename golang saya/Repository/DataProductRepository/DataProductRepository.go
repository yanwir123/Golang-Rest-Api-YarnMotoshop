package dataproductrepository

import (
	"golang-saya/Models"
	connections "golang-saya/Models/Connections"

	"golang-saya/Models/DataProductModels"
)



func InsertDataProduct(DataProduct DataProductModels.DataProduct) Models.BaseResponseModels {
	var query string
	var result Models.BaseResponseModels
	db := connections.DB
	query = "INSERT INTO DataProduct (Nama Product, Code Product, Jumlah Product, Nama Distributor, Nama Considnee, Tanggal Masuk Product) VALUES (?, ?, ?, ?, ?, ?)" // Perhatikan penggunaan tanda kutip dan nama kolom yang diubah.

	tempResult := db.Exec(query, DataProduct.Id ,DataProduct.Nama_Product, DataProduct.Code_Product, DataProduct.Jumlah_Product, DataProduct.Nama_Distributor, DataProduct.Nama_Considnee, DataProduct.Tanggal_Masuk_Product) // Menyesuaikan jumlah placeholder dengan nilai yang diberikan.

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
			Message:       "Data berhasil ditambahkan.",
			Data:          nil,
		}
	}

	return result
}


func UpdateDataProduct(DataProduct DataProductModels.DataProduct) Models.BaseResponseModels {
	var query string
	var result Models.BaseResponseModels
	db := connections.DB
	query = "UPDATE DataProduct SET Nama Product = ?, Code Product = ?, Jumlah Product = ?, Nama Distributor = ?, Nama Considnee = ?, Tanggal Masuk Product = ? WHERE id = ?"

	tempResult := db.Exec(query, DataProduct.Id, DataProduct.Nama_Product, DataProduct.Code_Product, DataProduct.Jumlah_Product, DataProduct.Nama_Distributor, DataProduct.Nama_Considnee, DataProduct.Tanggal_Masuk_Product)

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
				Message:       "Data dengan ID tersebut tidak ditemukan.",
				Data:          nil,
			}
		} else {
			result = Models.BaseResponseModels{
				CodeResponse:  200,
				HeaderMessage: "Success",
				Message:       "Data berhasil diubah.",
				Data:          nil,
			}
		}
	}

	return result
}

func DeleteDataProduct(request DataProductModels.DataProduct) Models.BaseResponseModels {
	var query string
	var result Models.BaseResponseModels
	db := connections.DB
	query = "DELETE FROM DataProduct WHERE id = ?"

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
				Message:       "Data dengan ID tersebut tidak ditemukan.",
				Data:          nil,
			}
		} else {
			result = Models.BaseResponseModels{
				CodeResponse:  200,
				HeaderMessage: "Success",
				Message:       "Data berhasil dihapus.",
				Data:          nil,
			}
		}
	}

	return result
}


func GetDataProductByID(request DataProductModels.DataProduct) Models.BaseResponseModels {
	var query string
	var result Models.BaseResponseModels
	var DataProductYarnmotoshop []DataProductModels.DataProduct
	db := connections.DB

	if request.Id != "" {
		query = "SELECT * FROM DataProduct WHERE Id = ?"
	} else {
		query = "SELECT * FROM DataProduct"
	}

	tempResult := db.Raw(query).Find(& DataProductYarnmotoshop)

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
			Data:          DataProductYarnmotoshop,
		}
	}

	return result
}
