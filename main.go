package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
)

var mySigningKey = []byte("super")

func GenerateJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["user"] = "Parth Aggarwal"
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		fmt.Errorf("Something went wrong: %s", err.Error())
		return "", err
	}

	return tokenString, nil
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("request from ", r.Header.Get("X-FORWARDED-FOR"), " on route /")
	validToken := "Frazile Server Online"

	// client := &http.Client{}
	// req, _ := http.NewRequest("GET", "http://127.0.0.1:8081", nil)
	// req.Header.Set("Authorized", validToken)
	// res, err := client.Do(req)
	// if err != nil {
	// 	fmt.Fprintf(w, "Error: %s", err.Error())
	// }

	// body, err := ioutil.ReadAll(res.Body)
	// if err != nil {
	// 	fmt.Fprintf(w, err.Error())
	// }

	// fmt.Fprintf(w, string(body))

	fmt.Fprintf(w, validToken)
}

func login(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "Login POST Endpoint worked")

	validToken, err := GenerateJWT()
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	loggedin := validToken
	fmt.Println("Endpoint Hit: Login Endpoint")
	w.Header().Set("Content-Type", "application/json") // Setting the Content Type Header
	w.WriteHeader(http.StatusOK)                       // Setting the Status Code
	json.NewEncoder(w).Encode(loggedin)
}

func handleRequests() {

	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/api/login", login).Methods("POST")

	log.Fatal(http.ListenAndServe(":9001", myRouter))
}

func main() {
	fmt.Println("Prisma Server")

	handleRequests()
}
