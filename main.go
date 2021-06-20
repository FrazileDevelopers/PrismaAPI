package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
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
	fmt.Println("request from ", r.Header.Get("X-FORWARDED-FOR") ," on route /")
	validToken, err := GenerateJWT()
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

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

func handleRequests() {
	http.HandleFunc("/", homePage)

	log.Fatal(http.ListenAndServe(":9001", nil))
}

func main() {
	fmt.Println("My Client")

	handleRequests()
}
