package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/parthahere001/go-bookstore-api/models"
	"github.com/parthahere001/go-bookstore-api/pkg/models"
	"github.com/parthahere001/go-bookstore-api/utils"
)

var NewBook models.Book 

func GetBooks(w http.ResponseWriter, r *http.Request){
	newBooks := models.GetAllBooks()
	res, _ := json.Marshal(newBooks)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

