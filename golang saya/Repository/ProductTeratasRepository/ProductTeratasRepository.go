package productteratasrepository

import (
	"golang-saya/Models"
	connections "golang-saya/Models/Connections"
	ProductTeratasModels "golang-saya/Models/ProductTeratasModels"
)

func InsertProductTeratas(ProductTeratas ProductTeratasModels.ProductTeratas) Models.BaseResponseModels {
    var query string
    var result Models.BaseResponseModels
    db := connections.DB
    query = "INSERT INTO ProductTeratas ( `Id`,`Nama_Product`, `Code_Product`, `Stok_Masuk`, `Stok_Ready`, `Product_Terjual`) VALUES (?, ?, ?, ?, ?, ?)"

    tempResult := db.Exec(query, 	ProductTeratas.Id , ProductTeratas.Nama_Product, ProductTeratas.Code_Product, ProductTeratas.Stok_Masuk, ProductTeratas.Stok_Ready, ProductTeratas.Product_Terjual)

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


func UpdateProductTeratas(ProductTeratas ProductTeratasModels.ProductTeratas) Models.BaseResponseModels {
	var query string
	var result Models.BaseResponseModels
	db := connections.DB
	query = "UPDATE ProductTeratas SET Nama Product = ?, Code Product = ?, Stok ProductM = ?, Stok ProductP = ?, Product Terjual = ? WHERE id = ?"

	tempResult := db.Exec(query, ProductTeratas.Id, ProductTeratas.Nama_Product, ProductTeratas.Code_Product, ProductTeratas.Stok_Masuk, ProductTeratas.Stok_Ready, ProductTeratas.Product_Terjual)
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

func DeleteProductTeratas(request ProductTeratasModels.ProductTeratas) Models.BaseResponseModels {
	var query string
	var result Models.BaseResponseModels
	db := connections.DB
	query = "DELETE FROM ProductTeratas WHERE id = ?"

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

func GetProductTeratasByID(request ProductTeratasModels.ProductTeratas) Models.BaseResponseModels {
	var query string
	var result Models.BaseResponseModels
	var ProductTeratasYarnmotoshop []ProductTeratasModels.ProductTeratas
	db := connections.DB

	if request.Id != "" {
		query = "SELECT * FROM ProductTeratas WHERE id = ?"
	} else {
		query = "SELECT * FROM ProductTeratas"
	}

	tempResult := db.Raw(query).Find(& ProductTeratasYarnmotoshop)

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
			Data:          ProductTeratasYarnmotoshop,
		}
	}

	return result
}
