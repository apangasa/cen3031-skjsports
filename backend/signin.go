package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var jwtKey = []byte("key")

var writers = map[string]string{
	"writer": "password",
}

type Writer struct {
	password string `json:"password"`
	username string `json:"username"`
}

type Claims struct {
	username string `json:"username"`
	jwt.RegisteredClaims
}

func Signin(w http.ResponseWriter, r *http.Request) {
	var writer Writer

	err := json.NewDecoder(r.Body).Decode(&writer)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	expectedPwd, valid := writers[writer.username]

	if !valid || expectedPwd != writer.password {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	expirationTime := time.Now().Add(5 * time.Minute)

	claims := &Claims{
		username:         writer.username,
		RegisteredClaims: jwt.RegisteredClaims{ExpiresAt: jwt.NewNumericDate(expirationTime)},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})

	w.WriteHeader(http.StatusOK)
}

func parser(w http.ResponseWriter, r *http.Request, claims *Claims) bool {
	cookie, err := r.Cookie("token")

	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return false
		}

		w.WriteHeader(http.StatusBadRequest)
		return false
	}

	tokenString := cookie.Value

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) { return jwtKey, nil })

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return false
		}
		w.WriteHeader(http.StatusBadRequest)
		return false
	}

	if !token.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return false
	}

	return true
}

func Auth(w http.ResponseWriter, r *http.Request) {
	claims := &Claims{}

	if parser(w, r, claims) {
		w.Write([]byte(fmt.Sprintf("Welcome to SKJ Sports %s!", claims.username)))
	}
}

func Renew(w http.ResponseWriter, r *http.Request) {
	claims := &Claims{}

	if parser(w, r, claims) {
		if time.Until(claims.ExpiresAt.Time) > 15*time.Second {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		expirationTime := time.Now().Add(5 * time.Minute)
		claims.ExpiresAt = jwt.NewNumericDate(expirationTime)

		token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
		tokenString, err := token.SignedString(jwtKey)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		http.SetCookie(w, &http.Cookie{
			Name:    "token",
			Value:   tokenString,
			Expires: expirationTime,
		})
	}

	return

}

func addWriter(w http.ResponseWriter, r *http.Request) {
	var writer Writer
	err := json.NewDecoder(r.Body).Decode(&writer)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if _, ok := writers[writer.username]; ok {
		w.WriteHeader(http.StatusConflict)
		return
	}

	writers[writer.username] = writer.password
	w.WriteHeader(http.StatusOK)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   "",
		Expires: time.Now(),
	})

	w.WriteHeader(http.StatusOK)
}
