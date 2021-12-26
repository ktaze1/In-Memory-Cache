package handlers

import (
	"net/http"
)

// @termsOfService http://swagger.io/terms/

type notAllowedHandler struct{}

func (h notAllowedHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	MethodNotAllowedHandler(rw, r)
}

// User defines the structure for a Full User Record
//
// swagger:model
type User struct {
	// The ID for the user
	// in: body
	//
	// required: false
	// min: 1
	ID int `json:"id"`

	// The Username of the user
	// in: body
	//
	// required: true
	Username string `json:"username"`

	// The Password of the user
	//
	// required: true
	Password string `json:"password"`

	// The Last Login time of the User
	//
	// required: true
	// min: 0
	LastLogin int64 `json:"lastlogin"`

	// Is the User Admin or not
	//
	// required: true
	Admin int `json:"admin"`

	// Is the User Logged In or Not
	//
	// required: true
	Active int `json:"active"`
}

// swagger:route DELETE /delete/{id} DeleteUser deleteID
// Delete a user given their ID.
// The command should be issued by an admin user
//
// responses:
// 200: noContent
// 404: ErrorMessage

// DeleteHandler is for deleting users based on user ID
func DeleteHandler(rw http.ResponseWriter, r *http.Request) {
}

// swagger:route POST / DefaultHandler noContent
// Default Handler for everything that is not a match.
// Works with all HTTP methods
//
// responses:
// 200: noContent
// 404: ErrorMessage

// DefaultHandler is for handling everything that is not a match
func DefaultHandler(rw http.ResponseWriter, r *http.Request) {

}

// swagger:route GET /*
// Default Handler for endpoints used with incorrect HTTP request method
//
// responses:
//	404: ErrorMessage

// MethodNotAllowedHandler is executed when the HTTP method is incorrect
func MethodNotAllowedHandler(rw http.ResponseWriter, r *http.Request) {

}

// swagger:route GET /time getTime
// Return current time
//
// responses:
//	200: OK

// TimeHandler is for handling /time – it works with plain text
func TimeHandler(rw http.ResponseWriter, r *http.Request) {
}

// swagger:route PUT /add UserInput createUser
// Create a new User
//
// responses:
//	200: OK
//  400: BadRequest

// AddHandler is for adding a new user
func setKeyHandler(rw http.ResponseWriter, r *http.Request) {
}

// swagger:route GET /getid getUserId loggedInfo
// Returns the ID of a User given their username and password
//
// responses:
//	200: OK
//  400: BadRequest

// GetIDHandler returns the ID of an existing user
func GetIDHandler(rw http.ResponseWriter, r *http.Request) {
}

// swagger:route POST /login user getLoginInfo
// Login an existing user
//
// responses:
//	200: OK
//  400: BadRequest

// LoginHandler is for updating the LastLogin time of a user
// And changing the Active field to true
func LoginHandler(rw http.ResponseWriter, r *http.Request) {
}

// swagger:route GET /logged logged getUserInfo
// Returns a list of logged in users
//
// responses:
// 200: UsersResponse
// 400: BadRequest

// LoggedUsersHandler returns the list of all logged in users
func LoggedUsersHandler(rw http.ResponseWriter, r *http.Request) {

}

// Generic OK message returned as an HTTP Status Code
// swagger:response OK
type OK struct {
	// Description of the situation
	// in: body
	Body int
}

// Generic error message returned as an HTTP Status Code
// swagger:response ErrorMessage
type ErrorMessage struct {
	// Description of the situation
	// in: body
	Body int
}

// Generic BadRequest message returned as an HTTP Status Code
// swagger:response BadRequest
type BadRequest struct {
	// Description of the situation
	// in: body
	Body int
}

// Generic noContent message returned as an HTTP Status Code
// swagger:response noContent
type noContent struct {
	// Description of the situation
	// in: body
	Body int
}

// A list of Users
// swagger:response UsersResponse
type UsersResponseWrapper struct {
	// A list of users
	// in: body
	Body []User
}

// swagger:parameters deleteID
type idParamWrapper struct {
	// The user id to be deleted
	// in: path
	// required: true
	ID int `json:"id"`
}

// A User
// swagger:parameters getUserInfo loggedInfo
type UserInputWrapper struct {
	// A list of users
	// in: body
	Body User
}
