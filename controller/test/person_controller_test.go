package test

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"project-pertama/controller"
	"project-pertama/model"
	"project-pertama/repository/mocks"
	"testing"

	"github.com/bmizerany/assert"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
)

type TestCase struct {
	Name            string
	Request         string
	MockFunc        func()
	ExpectedSuccess bool
}

func TestCreate(t *testing.T) {
	personRepositoryMock := mocks.NewPersonRepositoryMock()

	testCases := []TestCase{
		TestCase{
			Name: `
			Given request body is valid
			and person repository return some error
			When create new person
			Then should return error response
			`,
			Request: `
			{
				"name": "saka",
				"address": "suku air"
			}
			`,
			MockFunc: func() {
				mockedPerson := model.Person{
					Name: "saka",
				}
				personRepositoryMock.On("Create", mock.Anything).Return(mockedPerson, errors.New("some error")).Once()
			},
			ExpectedSuccess: false,
		},
		TestCase{
			Name: `
			Given request body is valid
			and person repository return success
			When create new person
			Then should return success response
			`,
			Request: `
			{
				"name": "saka",
				"address": "suku air"
			}
			`,
			MockFunc: func() {
				mockedPerson := model.Person{
					Name: "saka",
				}
				personRepositoryMock.On("Create", mock.Anything).Return(mockedPerson, nil).Once()
			},
			ExpectedSuccess: true,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			gin.SetMode(gin.TestMode)

			personController := controller.NewPersonController(personRepositoryMock)
			testCase.MockFunc()

			ginEngine := gin.Default()
			ginEngine.POST("/person", personController.Create)

			req, _ := http.NewRequest("POST", "/person", bytes.NewBuffer([]byte(testCase.Request)))
			rr := httptest.NewRecorder()

			ginEngine.ServeHTTP(rr, req)

			resultByte, _ := io.ReadAll(rr.Body)
			var resultResponse model.Response

			json.Unmarshal(resultByte, &resultResponse)

			assert.Equal(t, testCase.ExpectedSuccess, resultResponse.Success)
			personRepositoryMock.AssertExpectations(t)
		})
	}
}
