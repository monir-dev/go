package router

import (
	model "Structure/src/Model"
	jwt "Structure/src/system/middleware"
	"fmt"
	"log"
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
	r.HandleFunc("/login", LoginHandler).Methods("GET")

	r.Handle("/check", IsAuthenticated(CheckHandler))

	fs := http.FileServer(http.Dir("./static/"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	token, err := jwt.CreateJwtToken()
	if err != nil {
		w.Write([]byte("err"))
	} else {
		w.Write([]byte(token))
	}
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {

	var users []model.User

	// DB.First(&user, 1)
	DB.Find(&users)

	log.Println(users)

	fmt.Println("here")
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
