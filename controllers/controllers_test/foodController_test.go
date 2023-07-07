package controllers_test

import (
	"mongotest/controllers"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
)

type FoodSuite struct {
	suite.Suite
}

func TestFoodSuit(t *testing.T) {
	suite.Run(t, &FoodSuite{})
}

func (fs *FoodSuite) TestGetFoods() {
	fs.T().Log("runnng test get foods")
	router := gin.Default()

	router.GET("/foods", controllers.GetFoods())

	req, err := http.NewRequest("GET","/foods", nil)
	fs.Nil(err)

	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)
	fs.Equal(http.StatusOK, rr.Code)
}

func (fs *FoodSuite) SetupSuite() {
	fs.T().Log("setup suite...")
}

func (fs *FoodSuite) SetupTest() {
	fs.T().Log("setup test...")
}

func (fs *FoodSuite) TearDownSuite(suiteName, testName string) {
	fs.T().Log("teardown suite...")
}

func (fs *FoodSuite) TearDownTest(suiteName, testName string) {
	fs.T().Log("teardown suite...")
}

func (fs *FoodSuite) BeforeTest() {
	fs.T().Log("before suite...")
}

func (fs *FoodSuite) AfterTest() {
	fs.T().Log("after suite...")
}
