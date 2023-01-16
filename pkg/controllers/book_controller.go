package controllers

import (
	"fmt"
	"github.com/charles-co/go_crud_app/pkg/models"
	"github.com/charles-co/go_crud_app/pkg/utils"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

var Book models.Book

func GetBooks(w http.ResponseWriter, r *http.Request) {
	books, err := Book.GetBooks()
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, books)
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	book, err := Book.GetBook(fmt.Sprint(id))
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid Book ID")
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, book)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	book := models.Book{}
	err := utils.DecodeBody(r, &book)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	obj, err := book.CreateBook()
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithJSON(w, http.StatusCreated, obj)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	book := models.Book{}
	err := utils.DecodeBody(r, &book)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	obj, err := book.UpdateBook(fmt.Sprint(book.ID))
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, obj)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	book := models.Book{}
	err := utils.DecodeBody(r, &book)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	_, err = book.DeleteBook(fmt.Sprint(book.ID))
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithJSON(w, http.StatusNoContent, map[string]string{"result": "success"})
}
