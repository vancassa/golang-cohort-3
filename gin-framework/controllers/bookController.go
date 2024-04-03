package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Book struct {
	BookID string `json:"book_id"`
	Title string `json:"title"`
	Stock int `json:"stock"`
	Author string `json:"author"`

}

var BookDatas = []Book{}

func CreateBook(ctx *gin.Context) {
	var newBook Book

	if err := ctx.ShouldBindJSON(&newBook); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	newBook.BookID = fmt.Sprintf("%d", len(BookDatas)+1)
	BookDatas = append(BookDatas, newBook)

	ctx.JSON(http.StatusOK, gin.H{
		"data": newBook,
		"message": "Succeeded creating a new book",
	})
}

func GetBook(ctx *gin.Context) {
	bookID := ctx.Param("bookID")
	var bookData Book
	condition := false

	for i, book := range BookDatas {
		if bookID == book.BookID {
			condition = true
			bookData = BookDatas[i]
			break
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"data": nil,
			"message": fmt.Sprintf("book with id %v not found", bookID),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": bookData,
		"message": "Book found",
	})
}

func UpdateBook(ctx *gin.Context) {
	bookID := ctx.Param("bookID")
	var updatedBook Book
	condition := false

	if err := ctx.ShouldBindJSON(&updatedBook); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	for i, book := range BookDatas {
		if bookID == book.BookID {
			condition = true
			BookDatas[i] = updatedBook
			break
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"data": nil,
			"message": fmt.Sprintf("book with id %v not found", bookID),
		})
		return
	}

	convertBook, _ := strconv.Atoi(bookID)

	ctx.JSON(http.StatusOK, gin.H{
		"data": BookDatas[convertBook-1],
		"message": "Book updated",
	})
}

func DeleteBook(ctx *gin.Context) {
	bookID := ctx.Param("bookID")
	condition := false
	var bookIndex int

	for i, book := range BookDatas {
		if bookID == book.BookID {
			condition = true
			bookIndex = i
			break
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"data": nil,
			"message": fmt.Sprintf("book with id %v not found", bookID),
		})
		return
	}

	copy(BookDatas[bookIndex:], BookDatas[bookIndex+1:])
	BookDatas[len(BookDatas)-1] = Book{}
	BookDatas = BookDatas[:len(BookDatas)-1]

	ctx.JSON(http.StatusOK, gin.H{
		"data": nil,
		"message": "Book deleted",
	})
}