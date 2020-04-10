package server

import (
	"encoding/base64"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/securecookie"
)

type User struct {
	User     string `json:"user"`
	Password string `json:"password"`
}

type Auth struct {
	Token string `json:"token"`
}

func NewHandle() {
	http.HandleFunc("/auth", auth)
	http.HandleFunc("/reg", register)
	http.HandleFunc("/isAuth", isAuth)
}

func auth(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Fatal(err)
	}
	var user User
	json.Unmarshal([]byte(r.FormValue("login")), &user)
	conn := NewConnection()
	defer conn.Close()
	var password string
	conn.QueryRow("select password from users where login=$1", user.User).Scan(&password)
	var jsonSession []byte
	if password == user.Password && password != "" {
		var token = securecookie.GenerateRandomKey(32)
		conn.Exec("update users set token=$1 where login=$2", base64.StdEncoding.EncodeToString(token), user.User)
		jsonSession, _ = json.Marshal(map[string]interface{}{"datetime": time.Now().Format("01-02-2006 15:04:05"), "token": token, "status": "OK"})
	} else {
		jsonSession, _ = json.Marshal(map[string]interface{}{"datetime": time.Now().Format("01-02-2006 15:04:05"), "status": "NotAuthorize"})
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonSession)
}

func register(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Fatal(err)
	}
	var user User
	json.Unmarshal([]byte(r.FormValue("register")), &user)
	conn := NewConnection()
	defer conn.Close()
	var count int
	conn.QueryRow("select count(*) from users where login=$1", user.User).Scan(&count)
	var jsonSession []byte
	if count == 0 {
		_, err := conn.Exec("insert into users(login, password) values ($1, $2)", user.User, user.Password)
		if err != nil {
			log.Fatal("Connection DB error: ", err)
		}
		jsonSession, _ = json.Marshal(map[string]interface{}{"datetime": time.Now().Format("01-02-2006 15:04:05"), "status": "OK"})
	} else {
		jsonSession, _ = json.Marshal(map[string]interface{}{"datetime": time.Now().Format("01-02-2006 15:04:05"), "status": "NotRegister"})
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonSession)
}

func isAuth(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Fatal(err)
	}
	var auth Auth
	json.Unmarshal([]byte(r.FormValue("token")), &auth)
	var jsonSession []byte
	isAuth, premium := isAuthorized(auth.Token)
	if isAuth {
		jsonSession, _ = json.Marshal(map[string]interface{}{"datetime": time.Now().Format("01-02-2006 15:04:05"), "status": "OK", "premium": premium})
	} else {
		jsonSession, _ = json.Marshal(map[string]interface{}{"datetime": time.Now().Format("01-02-2006 15:04:05"), "status": "NotAuthorize"})
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonSession)
}
