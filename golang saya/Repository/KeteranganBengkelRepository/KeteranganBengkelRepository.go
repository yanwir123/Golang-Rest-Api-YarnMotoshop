package KeteranganBengkelRepository

import (
	"golang-saya/Models"
	connections "golang-saya/Models/Connections"
	"golang-saya/Models/KeteranganModels"
)

func InsertKeterangan(KeteranganBengkel KeteranganModels.KeteranganBengkel) Models.BaseResponseModels {
	var query string
	var result Models.BaseResponseModels
	db := connections.DB
	query = "INSERT KeteranganBengkel SET Nama Staff = ?, Tanggal Masuk = ?, Jabatan = ?, Gender = ?, Keterangan = ?, Baik = ?, Buruk = ? WHERE id = ?"

	tempResult := db.Exec(query, KeteranganBengkel.Id, KeteranganBengkel.Nama_Staff, KeteranganBengkel.Tanggal_Masuk, KeteranganBengkel.Jabatan, KeteranganBengkel.Gender, KeteranganBengkel.Keterangan, KeteranganBengkel.Baik, KeteranganBengkel.Buruk)

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

func UpdateKeterangan(KeteranganBengkel KeteranganModels.KeteranganBengkel) Models.BaseResponseModels {
	var query string
	var result Models.BaseResponseModels
	db := connections.DB
	query = "UPDATE KeteranganBengkel SET Nama Staff = ?, Tanggal Masuk = ?, Jabatan = ?, Gender = ?, Keterangan = ? WHERE id = ?"

	tempResult := db.Exec(query, KeteranganBengkel.Id, KeteranganBengkel.Nama_Staff, KeteranganBengkel.Tanggal_Masuk, KeteranganBengkel.Jabatan, KeteranganBengkel.Gender, KeteranganBengkel.Keterangan)

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

func DeleteKeterangan(request KeteranganModels.KeteranganBengkel) Models.BaseResponseModels {
	var query string
	var result Models.BaseResponseModels
	db := connections.DB
	query = "DELETE FROM Keterangan WHERE id = ?"

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

func GetKeteranganByID(request KeteranganModels.KeteranganBengkel) Models.BaseResponseModels {
	var query string
	var result Models.BaseResponseModels
	var KeteranganYarnmotoshop []KeteranganModels.KeteranganBengkel
	db := connections.DB

	if request.Id != "" {
		query = "SELECT * FROM KeteranganBengkel WHERE Id = ?"
	} else {
		query = "SELECT * FROM KeteranganBengkel"
	}

	tempResult := db.Raw(query).Find(& KeteranganYarnmotoshop)

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
			Data:          KeteranganYarnmotoshop,
		}
	}

	return result
}
