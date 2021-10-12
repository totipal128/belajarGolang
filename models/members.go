package models

type Member struct{
	Id			int 	`json:"id"`
	Id_user 	int		`json:"id_user"`
	Code_rekrut string	`json:"code_rekrut"`
	Nama		string	`json:"nama"`
	Nik			string	`json:"nik"`
	Alamat		string	`json:"alamat"`
	Tempat_lahir 	string `json:"tempat_lahir"`
	Tanggal_lahir	string `json:"tanggal_lahir"`
	Jeniskelamin 	string `json:"jenisKelamin"`
	Nohp			string `json:"noHp"`
}