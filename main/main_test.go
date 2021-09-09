package main

import (
	"encoding/json"
	"finansial-service/main/dao"
	"finansial-service/main/endpoints"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

//Configuring before tests
func BeforeTests(m *testing.M) {
	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())
}

func CreateRouter() *gin.Engine { return gin.Default() }

func assertHTTPresponse(t *testing.T, r *gin.Engine, req *http.Request, foo func(w *httptest.ResponseRecorder) bool) {

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	if !foo(w) {
		t.Fail()
	}
}

//Checking the correct HTTP code from health-check endpoint
func TestHealthCheck(t *testing.T) {

	router := CreateRouter()
	router.GET("/", HealthCheck)

	response, _ := http.NewRequest("GET", "/", nil)

	assertHTTPresponse(t, router, response, func(w *httptest.ResponseRecorder) bool {

		statusOK := w.Code == http.StatusOK
		return statusOK
	})
}

func TestCreateUser(t *testing.T) {

	dao.Connect()

	router := CreateRouter()
	router.POST("/createUser", endpoints.CreateUser)

	response, _ := http.NewRequest("POST", "/createUser", strings.NewReader("{}"))

	assertHTTPresponse(t, router, response, func(w *httptest.ResponseRecorder) bool {
		statusOk := w.Code == http.StatusOK
		userId := w.Body.String()
		log.Printf(userId)
		return statusOk
	})

}

//Utils func created to be used in more complex tests
func GetIdOfCreatedUser() (string, *gin.Engine) {
	dao.Connect()

	router := CreateRouter()
	router.POST("/createUser", endpoints.CreateUser)

	response, _ := http.NewRequest("POST", "/createUser", strings.NewReader("{}"))
	var userId string
	assertHTTPresponse(nil, router, response, func(w *httptest.ResponseRecorder) bool {

		userId = w.Body.String()
		return true
	})
	return userId, router
}

type Response struct {
	UserId string `json:"userid"`
}

//Simple addition test
func TestAddition(t *testing.T) {

	//Create new user and get his userId
	newUserId, router := GetIdOfCreatedUser()
	log.Printf(newUserId)
	var response1 Response
	json.Unmarshal([]byte(newUserId), &response1)
	newUserId = response1.UserId

	//Make an addition for created user
	router.GET("/addition", endpoints.MakeAddition)
	var body = "{\"id\":\"" + newUserId + "\", \"amount\":\"100\"}"
	response, _ := http.NewRequest("GET", "/addition", strings.NewReader(body))
	assertHTTPresponse(t, router, response, func(w *httptest.ResponseRecorder) bool {

		statusOk := w.Code == http.StatusOK
		log.Printf(w.Body.String())
		return statusOk
	})

}
