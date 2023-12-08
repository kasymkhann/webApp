package routes

import (
	"net/http"
	"webAPI_lesson/middleware"
	"webAPI_lesson/repository"
	"webAPI_lesson/session"
	"webAPI_lesson/utils"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
    
	r := mux.NewRouter()
	r.HandleFunc("/", middleware.AuthRequired(indexHandlerGet)).Methods("GET")
	r.HandleFunc("/", middleware.AuthRequired(indexHandlerPost)).Methods("POST")
	r.HandleFunc("/login", loginHandlerGet).Methods("GET")
	r.HandleFunc("/login", loginHandlerPost).Methods("POST")
	r.HandleFunc("/register", registerHandlerGet).Methods("GET")
	r.HandleFunc("/register", registerHandlerPost).Methods("POST")
	fs := http.FileServer(http.Dir("./static/"))
	r.PathPrefix("/static").Handler(http.StripPrefix("/static/", fs))
	return r
}

func indexHandlerGet(w http.ResponseWriter, r *http.Request) {
	session, _ := session.Store.Get(r, "session")
	_, ok := session.Values["username"]
	if !ok {
		http.Redirect(w, r, "/login", 302)
		return

	}
	comments, err := repository.GetComments()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal server error"))
		return
	}

	utils.ExecuteTemplate(w, "icon.html", comments)
}

func indexHandlerPost(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	comment := r.PostForm.Get("comment")
	err := repository.PostComments(comment)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal server error"))
		return
	}
	http.Redirect(w, r, "/", 302)
}

func loginHandlerGet(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "login.html", nil)
}

func loginHandlerPost(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	username := r.PostForm.Get("username")
	password := r.PostForm.Get("password")
	err := repository.Authenticate(username, password)
	if err != nil {
		switch err {
		case repository.ErrUserNotFound:
			utils.ExecuteTemplate(w, "login.html", "Unkown user")
		case repository.ErrInvalidLogin:
			utils.ExecuteTemplate(w, "login.html", "Invalid login")
		default:
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Internal server error"))
		}
		return

	}

	session, _ := session.Store.Get(r, "session")
	session.Values["username"] = username
	session.Save(r, w)
	http.Redirect(w, r, "/", 302)
}

func registerHandlerGet(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "register.html", nil)
}

func registerHandlerPost(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	username := r.PostForm.Get("username")
	password := r.PostForm.Get("password")
	err := repository.Register(username, password)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal server error"))
		return
	}
	http.Redirect(w, r, "/login", 302)
}
