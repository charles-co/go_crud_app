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

// @Summary get all books
// @ID get-books
// @Produce json
// @Success 200 {array} models.Book
// @Router /books [get]
func GetBooks(w http.ResponseWriter, r *http.Request) {
	books, err := Book.GetBooks()
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, books)
}

// @Summary get a book item by ID
// @ID get-book-by-id
// @Produce json
// @Param id path string true "book ID"
// @Success 200 {object} models.Book
// @Router /books/{id} [get]
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

// @Summary create a book item
// @ID create-book
// @Produce json
// @Param book body models.Book true "book item"
// @Success 201 {object} models.Book
// @Router /books [post]
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

// @Summary update a book item
// @ID update-book
// @Produce json
// @Param book body models.Book true "book item"
// @Success 200 {object} models.Book
// @Router /books [put]
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

// @Summary delete a book item
// @ID delete-book
// @Produce json
// @Param book body models.Book true "book item"
// @Success 204 {object} models.Book
// @Router /books [delete]
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
