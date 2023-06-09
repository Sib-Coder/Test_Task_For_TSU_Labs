package pages

import (
	"LabsTetsProject/internal/data"
	database "LabsTetsProject/internal/repository"
	"errors"

	"encoding/json"
	"fmt"
	"net/http"
)

func Hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "Привет это сервис, предоставляющий API для работы с данными пользователей.\n")
}

func UpdateDAtaHTTP(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {

		//fmt.Println("My request is: ", r.Body)
		var user model.UserInfo
		decode := json.NewDecoder(r.Body).Decode(&user)
		if decode != nil {
			DecodeJsonErr := errors.New("DECODE not Valid ")
			fmt.Println(DecodeJsonErr)
		}

		//обработка параметров через бд
		result := database.UpdateUser(user)

		fmt.Fprintln(w, result)

	}
}

// получение информации по имени
func ReceiveSend(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {

		//работа с запросом
		var user model.UserInfo
		decode := json.NewDecoder(r.Body).Decode(&user)
		if decode != nil {
			DecodeJsonErr := errors.New("DECODE not Valid ")
			fmt.Println(DecodeJsonErr) //надо найти выдачу ошибки
		}

		//обработка параметров через бд
		result := database.ExtractUserData(user.Name)

		//конвертим в Json
		b, err := json.Marshal(result)
		if err != nil {
			e1 := errors.New("Error: Marshal")
			fmt.Println(e1)
			return
		}

		fmt.Fprintln(w, string(b))

	}
}

func DeleteDAtaHTTP(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {

		//fmt.Println("My request is: ", r.Body)
		var user model.UserInfo
		decode := json.NewDecoder(r.Body).Decode(&user)
		if decode != nil {
			DecodeJsonErr := errors.New("DECODE not Valid ")
			fmt.Println(DecodeJsonErr) //надо найти выдачу ошибки
		}

		//обработка параметров через бд
		result := database.DeleteUser(user)

		fmt.Fprintln(w, result)

	}
}
func AddDAtaHTTP(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {

		//fmt.Println("My request is: ", r.Body)
		var user model.UserInfo
		decode := json.NewDecoder(r.Body).Decode(&user)
		if decode != nil {
			DecodeJsonErr := errors.New("DECODE not Valid ")
			fmt.Println(DecodeJsonErr) //надо найти выдачу ошибки
		}

		//обработка параметров через бд
		result := database.AddUserData(user)

		fmt.Fprintln(w, result)

	}
}
func ExtractAllUserHTTP(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		//обработка параметров через бд
		result := database.ExtractUserDataMas()
		//конвертим в Json
		b, err := json.Marshal(result)
		if err != nil {
			e1 := errors.New("Error: Marshal")
			fmt.Println(e1)
			return
		}

		fmt.Fprintln(w, string(b))

	}
}

func ExtractAllUserHTTPGender(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		//fmt.Println("My request is: ", r.Body)
		var user model.UserInfo
		decode := json.NewDecoder(r.Body).Decode(&user)
		if decode != nil {
			DecodeJsonErr := errors.New("DECODE not Valid ")
			fmt.Println(DecodeJsonErr) //надо найти выдачу ошибки
		}
		//обработка параметров через бд
		result := database.ExtractUserDataMasfForIdexGender(user)
		//конвертим в Json
		b, err := json.Marshal(result)
		if err != nil {
			e1 := errors.New("Error: Marshal")
			fmt.Println(e1)
			return
		}

		fmt.Fprintln(w, string(b))

	}
}
func ExtractAllUserHTTPStatus(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		//fmt.Println("My request is: ", r.Body)
		var user model.UserInfo
		decode := json.NewDecoder(r.Body).Decode(&user)
		if decode != nil {
			DecodeJsonErr := errors.New("DECODE not Valid ")
			fmt.Println(DecodeJsonErr) //надо найти выдачу ошибки
		}
		//обработка параметров через бд
		result := database.ExtractUserDataMasfForIdexGender(user)
		//конвертим в Json
		b, err := json.Marshal(result)
		if err != nil {
			e1 := errors.New("Error: Marshal")
			fmt.Println(e1)
			return
		}

		fmt.Fprintln(w, string(b))

	}
}
func ExtractAllUserHTTPDESC(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		//обработка параметров через бд
		result := database.ExtractUserDataMasDes()
		//конвертим в Json
		b, err := json.Marshal(result)
		if err != nil {
			e1 := errors.New("Error: Marshal")
			fmt.Println(e1)
			return
		}

		fmt.Fprintln(w, string(b))

	}
}
func ExtractAllUserHTTPASC(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		//обработка параметров через бд
		result := database.ExtractUserDataMasASC()
		//конвертим в Json
		b, err := json.Marshal(result)
		if err != nil {
			e1 := errors.New("Error: Marshal")
			fmt.Println(e1)
			return
		}

		fmt.Fprintln(w, string(b))

	}
}

func ExtractAllUserHTTPLimit(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		//обработка параметров через бд
		result := database.ExtractUserDataMasLimit()
		//конвертим в Json
		b, err := json.Marshal(result)
		if err != nil {
			e1 := errors.New("Error: Marshal")
			fmt.Println(e1)
			return
		}

		fmt.Fprintln(w, string(b))

	}
}
func ExtractAllUserHTTPOffset(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		//обработка параметров через бд
		result := database.ExtractUserDataMasOffset()
		//конвертим в Json
		b, err := json.Marshal(result)
		if err != nil {
			e1 := errors.New("Error: Marshal")
			fmt.Println(e1)
			return
		}

		fmt.Fprintln(w, string(b))

	}
}
