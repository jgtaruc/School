/*
This program continues from step8.
User can only access user edit form when they are logged-in
The form in index.html can only be accessed by users who are logged-in.

Initial username: user
Initial password: 1234
*/

/*
model.go - contains the User data type and functions related to that
			data type.
*/

package main

import (
	"net/http"
	"strings"
)


//User data type
type User struct {
	Name string
	Password string
	State bool
}


//gets the user info stored in the cookie(name, password, state)
func getModel(cookie *http.Cookie) User {
	//cookie structure = uuid|hmac|base64
	//gets base64 code
	xs := strings.Split(cookie.Value, "|")
	b64Data := xs[2]

	bs := decodeB64(b64Data)

	return unmarshalModel(bs)
}


//returns an initial model with name=user, password=1234 and state=false
func initialModel() []byte {
	model := User{
		Name: "user",
		Password: "1234",
		State: false,
	}
	return marshalModel(model)
}


//saves the values submitted in the form to the model
//to be save to the cookie.
func editModel(req *http.Request) []byte {
	model := User{
		Name: req.FormValue("user_name"),
		Password: req.FormValue("password"),
		State: true,
	}
	return marshalModel(model)
}