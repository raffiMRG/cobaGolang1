package ListMahasiswaRepository

import (
	"ListMahasiswa/Model"
	connection "ListMahasiswa/Model/Connection"
	"ListMahasiswa/Model/ListMahasiswaModel"
	"fmt"
)

func InsertListMahasiswa(ListMahasiswa ListMahasiswaModel.ListMahasiswa) Model.BaseResponseModel {
	var query string
	var result Model.BaseResponseModel
	db := connection.DB
	query = "INSERT INTO tb_mahasiswa (nim, nama) VALUES (?, ?)"

	tempResult := db.Exec(query, ListMahasiswa.NIM, ListMahasiswa.Nama)

	if tempResult.Error != nil {
		result = Model.BaseResponseModel{
			CodeResponse:  400,
			HeaderMessage: "Error",
			Message:       tempResult.Error.Error(),
			Data:          nil,
		}
	} else {
		result = Model.BaseResponseModel{
			CodeResponse:  200,
			HeaderMessage: "Success",
			Message:       "Data berhasil ditambahkan.",
			Data:          ListMahasiswa.NIM,
		}
	}

	return result
}

func UpdateListMahasiswa(ListMahasiswa ListMahasiswaModel.ListMahasiswa) Model.BaseResponseModel {
	var query string
	var result Model.BaseResponseModel
	db := connection.DB
	query = "UPDATE tb_mahasiswa SET nama = ? WHERE nim = ?"

	tempResult := db.Exec(query, ListMahasiswa.Nama, ListMahasiswa.NIM)

	if tempResult.Error != nil {
		result = Model.BaseResponseModel{
			CodeResponse:  400,
			HeaderMessage: "Error",
			Message:       tempResult.Error.Error(),
			Data:          nil,
		}
	} else {
		rowsAffected := tempResult.RowsAffected
		if rowsAffected == 0 {
			result = Model.BaseResponseModel{
				CodeResponse:  404,
				HeaderMessage: "Not Found",
				Message:       "Data dengan nim tersebut tidak ditemukan.",
				Data:          nil,
			}
		} else {
			result = Model.BaseResponseModel{
				CodeResponse:  200,
				HeaderMessage: "Success",
				Message:       "Data berhasil diubah.",
				Data:          ListMahasiswa.NIM,
			}
		}
	}

	return result
}

func DeleteListMahasiswa(request ListMahasiswaModel.ListMahasiswa) Model.BaseResponseModel {
	var query string
	var result Model.BaseResponseModel
	db := connection.DB
	query = "DELETE FROM tb_mahasiswa WHERE nim = ?"

	tempResult := db.Exec(query, request.NIM)

	if tempResult.Error != nil {
		result = Model.BaseResponseModel{
			CodeResponse:  400,
			HeaderMessage: "Error",
			Message:       tempResult.Error.Error(),
			Data:          nil,
		}
	} else {
		// Periksa apakah ada baris yang terpengaruh oleh perintah DELETE
		rowsAffected := tempResult.RowsAffected
		if rowsAffected == 0 {
			result = Model.BaseResponseModel{
				CodeResponse:  404,
				HeaderMessage: "Not Found",
				Message:       "Data dengan NIM tersebut tidak ditemukan.",
				Data:          nil,
			}
		} else {
			result = Model.BaseResponseModel{
				CodeResponse:  200,
				HeaderMessage: "Success",
				Message:       "Data berhasil dihapus.",
				Data:          request.NIM,
			}
		}
	}

	return result
}

func GetListMahasiswaByID(request ListMahasiswaModel.ListMahasiswa) Model.BaseResponseModel {
	var query string
	var result Model.BaseResponseModel
	var ListMahasiswaList []ListMahasiswaModel.ListMahasiswa
	db := connection.DB
	tempResult := connection.DB

	fmt.Printf("Received data repository: %+v\n", request)

	if request.NIM != "" {
		query = "SELECT * FROM tb_mahasiswa WHERE nim = ?"
		tempResult = db.Where("nim = ?", request.NIM).Find(&ListMahasiswaList)
		fmt.Println("nim query")
	} else {
		query = "SELECT * FROM tb_mahasiswa"
		tempResult = db.Raw(query).Find(&ListMahasiswaList)
		fmt.Println(tempResult)
	}

	// tempResult = db.Raw(query).Find(&ListMahasiswaList)
	// tempResult := db.Exec(query, request.NIM)

	if tempResult.Error != nil {
		result = Model.BaseResponseModel{
			CodeResponse:  400,
			HeaderMessage: "Error",
			Message:       tempResult.Error.Error(),
			Data:          nil,
		}
	} else {
		result = Model.BaseResponseModel{
			CodeResponse:  200,
			HeaderMessage: "Success",
			Message:       "Data retrieved successfully",
			Data:          ListMahasiswaList,
		}
	}

	return result
}
