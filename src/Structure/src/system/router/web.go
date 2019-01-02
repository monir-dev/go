package router

import (
	model "Structure/src/Model"
	jwt "Structure/src/system/middleware"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func WebRoutes(r *mux.Router, db *gorm.DB) {
	// r.HandleFunc("/", middleware.AuthRequired(indexGetHandler)).Methods("GET")
	DB = db
	r.HandleFunc("/", HomeHandler)

	r.HandleFunc("/register", RegisterHandler).Methods("GET")
	r.HandleFunc("/login", LoginHandler).Methods("POST")

	r.Handle("/check", IsAuthenticated(CheckHandler))

	fs := http.FileServer(http.Dir("./static/"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	// token, err := jwt.CreateJwtToken()
	// if err != nil {
	// 	w.Write([]byte("err"))
	// } else {
	// 	w.Write([]byte(token))
	// }
	w.Write([]byte("Home"))
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {

	var user model.User

	var email = r.FormValue("email")
	var password = r.FormValue("password")

	// email = "lkjasdfl"

	// DB.First(&user, 1)
	DB.Where("email = ? AND password = ?", email, password).First(&user)

	if user.Email == "" {
		w.Write([]byte("Invalid Username and password"))
	} else {
		token, err := jwt.CreateJwtToken(user.ID, user.Name, user.Email)
		if err != nil {
			w.Write([]byte("Token creation error"))
		} else {
			w.Write([]byte(token))
		}
	}
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Register route"))
}

func CheckHandler(w http.ResponseWriter, r *http.Request) {
	// Read the response body
	// buf := new(bytes.Buffer)
	// io.Copy(buf, r.Body)
	// r.Body.Close()
	// tokenString := strings.TrimSpace(buf.String())

	// token, errMsg, err := jwt.PurseToken(tokenString)
	// if err != nil {
	// 	w.Write([]byte("err"))
	// } else if errMsg != "" {
	// 	w.Write([]byte(errMsg))
	// } else {
	// 	w.Write([]byte(token))
	// }
	w.Write([]byte("Finally Here"))
}
