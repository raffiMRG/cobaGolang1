package Controller

import (
	"ListMahasiswa/Model"
	"ListMahasiswa/Model/ListMahasiswaModel"
	"ListMahasiswa/Repository/ListMahasiswaRepository"
	"net/http"

	"encoding/json"
	"fmt"
	"io"

	"github.com/gin-gonic/gin"
)

func GetData(c *gin.Context) {
	var requestData ListMahasiswaModel.ListMahasiswa
	var response Model.BaseResponseModel

	// Read the body of the request
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unable to read request body"})
		return
	}

	// Unmarshal the JSON data
	json.Unmarshal(body, &requestData)
	// if err := json.Unmarshal(body, &requestData); err != nil {
	// c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
	// return
	// }

	response = ListMahasiswaRepository.GetListMahasiswaByID(requestData)
	if response.CodeResponse != 200 {
		c.JSON(response.CodeResponse, response)
		return
	}

	// Use the parsed data
	fmt.Printf("Received data: %+v\n", requestData)

	// Respond with the received data
	c.JSON(http.StatusOK, response)
	// c.JSON(http.StatusOK, gin.H{
	// 	"id":   requestData.ID,
	// 	"nim":  requestData.NIM,
	// 	"nama": requestData.Nama,
	// })
}

// func GetData(c *gin.Context) {
// 	var request ListMahasiswaModel.ListMahasiswa
// 	var response Model.BaseResponseModel

// 	response = ListMahasiswaRepository.GetListMahasiswaByID(request)
// 	if response.CodeResponse != 200 {
// 		c.JSON(response.CodeResponse, response)
// 		return
// 	}

// 	c.JSON(http.StatusOK, response)
// }

func InsertData(c *gin.Context) {
	var request ListMahasiswaModel.ListMahasiswa
	var response Model.BaseResponseModel

	if err := c.ShouldBindJSON(&request); err != nil {
		response = Model.BaseResponseModel{
			CodeResponse:  400,
			HeaderMessage: "Bad Request",
			Message:       err.Error(),
			Data:          nil,
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response = ListMahasiswaRepository.InsertListMahasiswa(request)
	if response.CodeResponse != 200 {
		c.JSON(response.CodeResponse, response)
		return
	}

	c.JSON(http.StatusOK, response)
}

func UpdateData(c *gin.Context) {
	var request ListMahasiswaModel.ListMahasiswa
	var response Model.BaseResponseModel

	if err := c.ShouldBindJSON(&request); err != nil {
		response = Model.BaseResponseModel{
			CodeResponse:  400,
			HeaderMessage: "Bad Request",
			Message:       err.Error(),
			Data:          nil,
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response = ListMahasiswaRepository.UpdateListMahasiswa(request)
	if response.CodeResponse != 200 {
		c.JSON(response.CodeResponse, response)
		return
	}

	c.JSON(http.StatusOK, response)
}

func DeleteData(c *gin.Context) {
	var request ListMahasiswaModel.ListMahasiswa
	var response Model.BaseResponseModel

	if err := c.ShouldBindJSON(&request); err != nil {
		response = Model.BaseResponseModel{
			CodeResponse:  400,
			HeaderMessage: "Bad Request",
			Message:       err.Error(),
			Data:          nil,
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response = ListMahasiswaRepository.DeleteListMahasiswa(request)
	if response.CodeResponse != 200 {
		c.JSON(response.CodeResponse, response)
		return
	}

	c.JSON(http.StatusOK, response)
}
