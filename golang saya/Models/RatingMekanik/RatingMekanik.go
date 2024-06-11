package ratingmekanik

type RatingMekanik struct {
	Id             string  `json:"id" from:"id"`
	Nama_Mekanik   string  `json:"Nama Mekanik" from:"Nama Mekanik"`
	Rating_Mekanik float64 `json:"Rating" from:"Rating"`
}