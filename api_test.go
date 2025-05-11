package main

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/obanlatomiwa/go-inventory-api/database"
	"github.com/obanlatomiwa/go-inventory-api/models"
	"github.com/obanlatomiwa/go-inventory-api/utils"
	"github.com/steinfletcher/apitest"
	"io"
	"net/http"
	"testing"
)

func testApp() *fiber.App {
	app := NewApp()
	database.InitialiseDatabase(utils.GetValueFromConfigFile("DB_TEST_NAME"))
	return app
}

func getItem() models.Item {
	database.InitialiseDatabase(utils.GetValueFromConfigFile("DB_TEST_NAME"))
	item, err := database.CreateFakeItemsForTesting()
	if err != nil {
		panic(err)
	}
	return item
}

func cleanUp(res *http.Response, r *http.Request, apiTest *apitest.APITest) {
	if http.StatusOK == res.StatusCode || http.StatusCreated == res.StatusCode {
		database.CleanTestData()
	}
}

func FiberToHandlerFunc(app *fiber.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res, err := app.Test(r)
		if err != nil {
			panic(err)
		}

		// copy the headers
		for name, values := range res.Header {
			for _, value := range values {
				w.Header().Add(name, value)
			}
		}
		w.WriteHeader(res.StatusCode)

		if _, err := io.Copy(w, res.Body); err != nil {
			panic(err)
		}
	}
}

func getJWTToken(t *testing.T) string {
	database.InitialiseDatabase(utils.GetValueFromConfigFile("DB_TEST_NAME"))
	user, err := database.CreateFakeUsersForTesting()
	if err != nil {
		panic(err)
	}

	// create a login request
	var userReq *models.UserRequest = &models.UserRequest{
		Email:    user.Email,
		Password: user.Password,
	}

	// get the response from the login request
	var res *http.Response = apitest.New().
		HandlerFunc(FiberToHandlerFunc(testApp())).
		Post("/api/v1/login").
		JSON(userReq).
		Expect(t).
		Status(http.StatusOK).
		End().Response

	// store the response body
	var response *models.Response[string] = &models.Response[string]{}

	// decode the response body
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return ""
	}

	var token string = response.Data

	// create a bearer token
	var JwtToken = "Bearer " + token

	return JwtToken

}

func TestSignUp_Success(t *testing.T) {
	// create a sample data for user
	user, err := utils.CreateFaker[models.User]()
	if err != nil {
		panic(err)
	}

	// create a signup request
	var userReq *models.UserRequest = &models.UserRequest{
		Email:    user.Email,
		Password: user.Password,
	}

	// create a test
	apitest.New().
		Observe(cleanUp).
		HandlerFunc(FiberToHandlerFunc(testApp())).
		Post("/api/v1/signup").
		JSON(userReq).
		Expect(t).
		Status(http.StatusCreated).
		End()
}

func TestSignUp_ValidationFailed(t *testing.T) {
	// create an invalid signup request
	var userReq *models.UserRequest = &models.UserRequest{
		Email:    "",
		Password: "",
	}

	// create a test
	apitest.New().
		HandlerFunc(FiberToHandlerFunc(testApp())).
		Post("/api/v1/signup").
		JSON(userReq).
		Expect(t).
		Status(http.StatusBadRequest).
		End()
}

func TestLogin_Success(t *testing.T) {
	// connect to database
	database.InitialiseDatabase(utils.GetValueFromConfigFile("DB_TEST_NAME"))

	// create a sample data for user
	user, err := database.CreateFakeUsersForTesting()
	if err != nil {
		panic(err)
	}

	// create a login request
	var userReq *models.UserRequest = &models.UserRequest{
		Email:    user.Email,
		Password: user.Password,
	}

	// create a test
	apitest.New().
		Observe(cleanUp).
		HandlerFunc(FiberToHandlerFunc(testApp())).
		Post("/api/v1/login").
		JSON(userReq).
		Expect(t).
		Status(http.StatusOK).
		End()
}

func TestLogin_ValidationFailed(t *testing.T) {
	// create an invalid signup request
	var userReq *models.UserRequest = &models.UserRequest{
		Email:    "",
		Password: "",
	}

	// create a test
	apitest.New().
		HandlerFunc(FiberToHandlerFunc(testApp())).
		Post("/api/v1/login").
		JSON(userReq).
		Expect(t).
		Status(http.StatusBadRequest).
		End()
}

func TestLogin_LoginFailed(t *testing.T) {

	// create a signup request
	var userReq *models.UserRequest = &models.UserRequest{
		Email:    "error@gmail.com",
		Password: "123123",
	}

	// create a test
	apitest.New().
		Observe(cleanUp).
		HandlerFunc(FiberToHandlerFunc(testApp())).
		Post("/api/v1/login").
		JSON(userReq).
		Expect(t).
		Status(http.StatusInternalServerError).
		End()
}

func TestGetItems_Success(t *testing.T) {

	// get jwt token
	token := getJWTToken(t)

	// create a test
	apitest.New().
		HandlerFunc(FiberToHandlerFunc(testApp())).
		Get("/api/v1/items").
		Header("Authorization", token).
		Expect(t).
		Status(http.StatusOK).
		End()
}

func TestGetItem_Success(t *testing.T) {
	// get sample data
	var item models.Item = getItem()

	token := getJWTToken(t)

	// create a test
	apitest.New().
		HandlerFunc(FiberToHandlerFunc(testApp())).
		Get("/api/v1/items/"+item.ID).
		Header("Authorization", token).
		Expect(t).
		Status(http.StatusOK).
		End()
}

func TestGetItem_NotFound(t *testing.T) {

	token := getJWTToken(t)

	// create a test
	apitest.New().
		HandlerFunc(FiberToHandlerFunc(testApp())).
		Get("/api/v1/items/0").
		Header("Authorization", token).
		Expect(t).
		Status(http.StatusNotFound).
		End()
}

func TestCreateItem_Success(t *testing.T) {
	item, err := utils.CreateFaker[models.Item]()
	if err != nil {
		panic(err)
	}

	var itemReq *models.ItemRequest = &models.ItemRequest{
		Name:     item.Name,
		Price:    item.Price,
		Quantity: item.Quantity,
	}

	token := getJWTToken(t)

	apitest.New().
		Observe(cleanUp).
		HandlerFunc(FiberToHandlerFunc(testApp())).
		Post("/api/v1/items").
		Header("Authorization", token).
		JSON(itemReq).
		Expect(t).
		Status(http.StatusCreated).
		End()
}

func TestCreateItem_ValidationFailed(t *testing.T) {
	var itemReq *models.ItemRequest = &models.ItemRequest{
		Name:     "",
		Price:    0,
		Quantity: 0,
	}

	token := getJWTToken(t)

	apitest.New().
		HandlerFunc(FiberToHandlerFunc(testApp())).
		Post("/api/v1/items").
		Header("Authorization", token).
		JSON(itemReq).
		Expect(t).
		Status(http.StatusBadRequest).
		End()
}

func TestUpdateItem_Success(t *testing.T) {
	item := getItem()

	var itemReq *models.ItemRequest = &models.ItemRequest{
		Name:     item.Name,
		Price:    item.Price,
		Quantity: item.Quantity,
	}

	token := getJWTToken(t)

	apitest.New().
		Observe(cleanUp).
		HandlerFunc(FiberToHandlerFunc(testApp())).
		Put("/api/v1/items/"+item.ID).
		Header("Authorization", token).
		JSON(itemReq).
		Expect(t).
		Status(http.StatusOK).
		End()
}

func TestUpdateItem_Failed(t *testing.T) {
	var itemReq *models.ItemRequest = &models.ItemRequest{
		Name:     "changed",
		Price:    10,
		Quantity: 10,
	}

	token := getJWTToken(t)

	apitest.New().
		HandlerFunc(FiberToHandlerFunc(testApp())).
		Put("/api/v1/items/0").
		Header("Authorization", token).
		JSON(itemReq).
		Expect(t).
		Status(http.StatusNotFound).
		End()

}

func TestDeleteItem_Success(t *testing.T) {
	var item models.Item = getItem()
	token := getJWTToken(t)

	apitest.New().
		HandlerFunc(FiberToHandlerFunc(testApp())).
		Delete("/api/v1/items/"+item.ID).
		Header("Authorization", token).
		Expect(t).
		Status(http.StatusOK).
		End()
}

func TestDeleteItem_Failed(t *testing.T) {
	token := getJWTToken(t)

	apitest.New().
		HandlerFunc(FiberToHandlerFunc(testApp())).
		Delete("/api/v1/items/0").
		Header("Authorization", token).
		Expect(t).
		Status(http.StatusNotFound).
		End()
}
