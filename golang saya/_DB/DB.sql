create database YarnMotoshop;

create table KeteranganBengkel (
    Id int PRIMARY KEY
	,Nama_Staff varchar(250) not null
	,Tanggal_Masuk DATETIME
	,Posisi varchar(250) not null
	,Gender varchar(250) not null
	,Keterangan varchar(500) not null
	,Baik float
	,Buruk float
);

create table DataProduct (
	Id int PRIMARY KEY
	,Nama_Product varchar(500) not null
	,Code_Product varchar(100) not null
	,Jenis_Product varchar(200) not null
	,Jumlah_Product float not null
	,Nama_Distributor varchar(300) not null
	,Nama_Considnee varchar(150) not null
	,Tanggal_Masuk_Product DATETIME
);

create table ProductTeratas (
	Id int PRIMARY KEY
	,Nama_Product varchar(300)
	,Code_Product varchar(300)
	,Stok_Masuk float
	,Stok_Ready float
	,Product_Terjual float
);

create table DataKeuangan (
	Id int PRIMARY KEY
	,Bulan varchar(100)
	,Nominal DECIMAL(14,2)
	,Jumlah_Customer float
	,Total_Pengeluaran DECIMAL(14,2)
	,Total_Pemasukan DECIMAL(14,2)
);

create table StockOil (
	Id int PRIMARY KEY
	,Nama_Product varchar(300)
	,Code_Product varchar(300)
	,Stok_Masuk float
	,Stok_Ready float
	,Product_Terjual float
);

create table StockBan (
	Id int PRIMARY KEY
	,Nama_Product varchar(300)
	,Code_Product varchar(300)
	,Stok_Masuk float
	,Stok_Ready float
	,Product_Terjual float
);

create table StokModivikasi (
	Id int PRIMARY KEY
	,Nama_Product varchar(300)
	,Code_Product varchar(300)
	,Stok_Masuk float
	,Stok_Ready float
	,Product_Terjual float
);


create table StokCarbonPart (
	Id int PRIMARY KEY
	,Nama_Product varchar(300)
	,Code_Product varchar(300)
	,Stok_Masuk float
	,Stok_Ready float
	,Product_Terjual float
);

create table RatingMekanik (
	Id int PRIMARY KEY
	,Nama_Mekanik varchar(300)
	,Rating_Mekanik float
);



