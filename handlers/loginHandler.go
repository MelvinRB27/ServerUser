package handlers

import (
	"errors"
	"net/http"
	"time"

	"github.com/MelvinRB27/server-user/authorization"
	H "github.com/MelvinRB27/server-user/helpers"
	"github.com/MelvinRB27/server-user/models"
	"github.com/MelvinRB27/server-user/storage"
)

type Handlers interface {
	Login(models.Login) (models.User, error)
	Register(*models.User) error
	Update(*models.User) (models.User, error)
}

type user struct {
	handlers Handlers
}

func NewData(h Handlers) user {
	return user{h}
}

var (
	errorInvalidMethod   = errors.New("invalid method")
	errorDataEmpty       = errors.New("error data empty")
	errorPasswordNoMatch = errors.New("the password no match")
)

//allow cors
func setupCORS(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Content-Type", "text/html; charset=utf-8")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-length, Accept-Encoding, X-CSRFToken, Authorization")
	(*w).Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	(*w).Header().Set("Access-Control-Allow-Headers", "Access-control-allow-origin,Access-Control-Allow-Headers, Origin,Accept, X-Requested-With, Content-Type, Access-Control-Request-Method, Access-Control-Request-Headers")

}

//used for login user
func (d *user) login(w http.ResponseWriter, r *http.Request) {
	setupCORS(&w)

	if r.Method != http.MethodPost {
		response := newResponse(Error, errorInvalidMethod.Error(), nil)
		responseJson(w, http.StatusBadRequest, response)
		return
	}

	UserName := r.FormValue("usernameLogin")
	Password := r.FormValue("passwordLogin")

	//check if data is empty or not
	userNameCheck := H.IsEmpty(UserName)
	passwordCheck := H.IsEmpty(Password)

	if userNameCheck || passwordCheck {
		response := newResponse(Error, errorDataEmpty.Error(), nil)
		responseJson(w, http.StatusBadRequest, response)
		return
	}

	//send data to method login
	data, err := storage.Login(UserName, Password)
	if err != nil {
		response := newResponse(Error, err.Error(), nil)
		responseJson(w, http.StatusBadRequest, response)
		return
	}

	//geting the token
	token, err := authorization.GenerateToken(UserName)
	if err != nil {
		response := newResponse(Error, err.Error(), nil)
		responseJson(w, http.StatusBadRequest, response)
		return
	}
	cookie, err := r.Cookie(token)
	if err != nil {
		cookie = &http.Cookie{
			Name:    "jwt-token",
			Value:   token,
			Expires: time.Time{},
			Path:    "/v1/login",
		}
	}

	http.SetCookie(w, cookie)
	response := newResponseWithCookie(Message, "OK", data, cookie.Value)
	responseJson(w, http.StatusOK, response)
}

//used for register user account
func (d *user) register(w http.ResponseWriter, r *http.Request) {
	setupCORS(&w)

	if r.Method != http.MethodPost {
		response := newResponse(Error, errorInvalidMethod.Error(), nil)
		responseJson(w, http.StatusBadRequest, response)
		return
	}

	//get data  from form frontend
	Name := r.FormValue("nameRegister")
	LastName := r.FormValue("lastNameRegister")
	UserName := r.FormValue("usernameRegister")
	Gender := r.FormValue("genderRegister")
	Rol := r.FormValue("rolRegister")
	Password := r.FormValue("passwordRegister")
	PasswordConfirm := r.FormValue("passwordConfirmRegister")

	//check if data is empty or not
	nameCheck := H.IsEmpty(Name)
	lastNameCheck := H.IsEmpty(LastName)
	userNameCheck := H.IsEmpty(UserName)
	genderNameCheck := H.IsEmpty(Gender)
	rolCheck := H.IsEmpty(Rol)
	passwordCheck := H.IsEmpty(Password)
	passwordConfirmCheck := H.IsEmpty(PasswordConfirm)

	if nameCheck || lastNameCheck || userNameCheck || genderNameCheck || rolCheck || passwordCheck || passwordConfirmCheck {
		response := newResponse(Error, errorDataEmpty.Error(), nil)
		responseJson(w, http.StatusBadRequest, response)
		return
	}

	//check if passwords match
	if Password != PasswordConfirm {
		response := newResponse(Error, errorPasswordNoMatch.Error(), nil)
		responseJson(w, http.StatusBadRequest, response)
		return
	}

	//register user
	err := storage.Register(&models.User{
		Name:            Name,
		LastName:        LastName,
		UserName:        UserName,
		Gender:          Gender,
		Rol:             Rol,
		Password:        Password,
		PasswordConfirm: PasswordConfirm,
	})

	if err != nil {
		response := newResponse(Error, err.Error(), nil)
		responseJson(w, http.StatusBadRequest, response)
		return
	}

	response := newResponse(Message, "Registration successful", nil)
	responseJson(w, http.StatusOK, response)
}

//used for update user update
func (d *user) update(w http.ResponseWriter, r *http.Request) {
	setupCORS(&w)
	w.WriteHeader(http.StatusOK)

	if r.Method != http.MethodPut {
		response := newResponse(Error, errorInvalidMethod.Error(), nil)
		responseJson(w, http.StatusBadRequest, response)
		return
	}

	//get data  from form frontend
	Name := r.FormValue("nameUpdate")
	LastName := r.FormValue("lastNameUpdate")
	UserName := r.FormValue("usernameUpdate")
	Gender := r.FormValue("genderUpdate")
	Rol := r.FormValue("rolUpdate")
	Password := r.FormValue("passwordUpdate")
	PasswordConfirm := r.FormValue("passwordConfirmUpdate")

	//check if data is empty or not
	nameCheck := H.IsEmpty(Name)
	lastNameCheck := H.IsEmpty(LastName)
	userNameCheck := H.IsEmpty(UserName)
	genderCheck := H.IsEmpty(Gender)
	rolCheck := H.IsEmpty(Rol)
	passwordCheck := H.IsEmpty(Password)
	passwordConfirmCheck := H.IsEmpty(PasswordConfirm)

	if nameCheck || lastNameCheck || userNameCheck || genderCheck || rolCheck || passwordCheck || passwordConfirmCheck {
		response := newResponse(Error, errorDataEmpty.Error(), nil)
		responseJson(w, http.StatusBadRequest, response)
		return
	}

	//check if passwords match
	if Password != PasswordConfirm {
		response := newResponse(Error, errorPasswordNoMatch.Error(), nil)
		responseJson(w, http.StatusBadRequest, response)
		return
	}

	//update user
	data, err := storage.Update(&models.User{
		Name:            Name,
		LastName:        LastName,
		UserName:        UserName,
		Gender:          Gender,
		Rol:             Rol,
		Password:        Password,
		PasswordConfirm: PasswordConfirm,
	})

	if err != nil {
		response := newResponse(Error, err.Error(), nil)
		responseJson(w, http.StatusBadRequest, response)
		return
	}

	token, err := authorization.GenerateToken(UserName)
	if err != nil {
		response := newResponse(Error, err.Error(), nil)
		responseJson(w, http.StatusBadRequest, response)
		return
	}
	cookie, err := r.Cookie(token)
	if err != nil {
		cookie = &http.Cookie{
			Name:    "jwt-token",
			Value:   token,
			Expires: time.Time{},
			Path:    "/v1/login",
		}
	}

	http.SetCookie(w, cookie)
	response := newResponseWithCookie(Message, "Successful update", data, cookie.Value)
	responseJson(w, http.StatusOK, response)
}