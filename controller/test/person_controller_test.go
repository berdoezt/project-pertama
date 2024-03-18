package test

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"project-pertama/controller"
	"project-pertama/model"
	"project-pertama/repository/mocks"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreate(t *testing.T) {
	var bodyString = `
	{
		"name": "saka",
		"address": "suku air"
	}
	`

	userId := uuid.NewString()

	var newPerson model.Person

	json.Unmarshal([]byte(bodyString), &newPerson)

	personRepositoryMock := mocks.NewPersonRepositoryMock()

	personRepositoryMock.On("Create", mock.Anything).Return(model.Person{
		Name: "saka",
		UUID: userId,
	})

	personController := controller.NewPersonController(personRepositoryMock)

	gin.SetMode(gin.TestMode)

	ginEngine := gin.Default()
	ginEngine.POST("/person", personController.Create)

	req, _ := http.NewRequest("POST", "/person", bytes.NewBuffer([]byte(bodyString)))
	rr := httptest.NewRecorder()

	ginEngine.ServeHTTP(rr, req)

	resultByte, _ := io.ReadAll(rr.Body)
	var resultResponse model.Response

	json.Unmarshal(resultByte, &resultResponse)

	assert.Equal(t, true, resultResponse.Success)
}
