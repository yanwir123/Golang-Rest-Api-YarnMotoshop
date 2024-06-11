package KeteranganBengkelRepository

import (
	"golang-saya/Models"
	connections "golang-saya/Models/Connections"
	ratingmekanik "golang-saya/Models/RatingMekanik"
)

func InsertRatingMekanik(RatingMekanik ratingmekanik.RatingMekanik) Models.BaseResponseModels {
	var query string
	var result Models.BaseResponseModels
	db := connections.DB
	query = "INSERT RatingMekanik SET Nama Mekanik = ?, Rating Mkanik = ? WHERE id = ?"

	tempResult := db.Exec(query, RatingMekanik.Id, RatingMekanik.Nama_Mekanik, RatingMekanik.Rating_Mekanik)

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

func UpdateRatingMekanik(RatingMekanik ratingmekanik.RatingMekanik) Models.BaseResponseModels {
	var query string
	var result Models.BaseResponseModels
	db := connections.DB
	query = "UPDATE RatingMekanik SET Nama Mekanik = ?, Rating Mekanik = ? WHERE id = ?"

	tempResult := db.Exec(query, RatingMekanik.Id, RatingMekanik.Nama_Mekanik, RatingMekanik.Rating_Mekanik)

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

func DeleteRatingMekanik(request ratingmekanik.RatingMekanik) Models.BaseResponseModels {
	var query string
	var result Models.BaseResponseModels
	db := connections.DB
	query = "DELETE FROM RatingMekanik WHERE id = ?"

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

func GetRatingMekanikByID(request ratingmekanik.RatingMekanik) Models.BaseResponseModels {
	var query string
	var result Models.BaseResponseModels
	var RatingMekanikYarnmotoshop []ratingmekanik.RatingMekanik
	db := connections.DB

	if request.Id != "" {
		query = "SELECT * FROM RatingMekanik WHERE Id = ?"
	} else {
		query = "SELECT * FROM RatingMekanik"
	}

	tempResult := db.Raw(query).Find(& RatingMekanikYarnmotoshop)

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
			Data:          RatingMekanikYarnmotoshop,
		}
	}

	return result
}
