package datakeuanganrepository

import (
	"golang-saya/Models"
	connections "golang-saya/Models/Connections"
	datakeuanganmodels "golang-saya/Models/DataKeuanganModels"
)



func InsertDataKeuangan(Keuangan datakeuanganmodels.DataKeuangan) Models.BaseResponseModels {
	var query string
	var result Models.BaseResponseModels
	db := connections.DB
	query = "INSERT DataKeuangan SET id = ?, Bulan = ?,Nominal = ?, Jumlah_Custumer = ?, Total Pemasukan = ?, Total Pengeluaran = ? WHERE id = ?"

	tempResult := db.Exec(query, Keuangan.Id, Keuangan.Bulan, Keuangan.Nominal, Keuangan.Jumlah_Customer, Keuangan.Total_Pemasukan, Keuangan.Total_Pengeluaran)

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

func UpdateDataKeuangan(Keuangan datakeuanganmodels.DataKeuangan) Models.BaseResponseModels {
	var query string
	var result Models.BaseResponseModels
	db := connections.DB
query = "Update DataKeuangan SET id = ?, Bulan = ?,Nominal = ?, Jumlah_Custumer = ?, Total Pemasukan = ?, Total Pengeluaran = ? WHERE id = ?"

	tempResult := db.Exec(query, Keuangan.Id, Keuangan.Bulan, Keuangan.Nominal, Keuangan.Jumlah_Customer, Keuangan.Total_Pemasukan, Keuangan.Total_Pengeluaran)

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

func DeleteDataKeuangan(request datakeuanganmodels.DataKeuangan) Models.BaseResponseModels {
	var query string
	var result Models.BaseResponseModels
	db := connections.DB
	query = "DELETE FROM DataKeuangan WHERE id = ?"

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


func GetDataKeuanganByID(request datakeuanganmodels.DataKeuangan) Models.BaseResponseModels {
	var query string
	var result Models.BaseResponseModels
	var DataProductYarnmotoshop []datakeuanganmodels.DataKeuangan
	db := connections.DB

	if request.Id != 0 {
		query = "SELECT * FROM DataKeuangan WHERE Id = ?"
	} else {
		query = "SELECT * FROM DataKeuangan"
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
