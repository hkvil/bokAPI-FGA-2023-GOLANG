package controllers

import(
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
)

type Book struct {
    BookID      string     `json:"book_id"`
    Title       string  `json:"title"`
    Author      string  `json:"author"`
    Price       float64 `json:"price"`
}


var BookDatas = []Book{}

//CREATE BOOK
func CreateBook(ctx *gin.Context){
	var newBook Book

	if err:= ctx.ShouldBindJSON(&newBook);err!= nil{
		ctx.AbortWithError(http.StatusBadRequest,err)
	}
	newBook.BookID =fmt.Sprintf("b%d",len(BookDatas)+1)
	BookDatas = append(BookDatas,newBook)
	ctx.JSON(http.StatusCreated,gin.H{
		"book":newBook,
	})
}

//GET ALL BOOK
func GetAllBook(ctx *gin.Context){
	// if len(BookDatas)==0 {
	// 	ctx.AbortWithStatusJSON(http.StatusNotFound,gin.H{
	// 		"error_status":"Data not found",
	// 		"error_message":fmt.Sprint("NO Data Been Created"),
	// 	})
	// }
	ctx.JSON(http.StatusOK,gin.H{
		"books":BookDatas,
	})
}
//GET BOOK BY ID
func GetBookByID(ctx *gin.Context){
	bookID:=ctx.Param("bookID")
	condition:=false
	var bookData Book

	for i,book:= range BookDatas{
		if bookID == book.BookID{
			condition=true
			bookData=BookDatas[i]
			break
		}
	}
	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound,gin.H{
			"error_status":"Data not found",
			"error_message":fmt.Sprint("book with id %v not found"),
		})
	}
	ctx.JSON(http.StatusOK,gin.H{
		"book":bookData,
	})
}
//UPDATE BOOK
func UpdateBook(ctx *gin.Context){
	bookID:=ctx.Param("bookID")
	condition:=false
	var updatedBook Book

	if err:= ctx.ShouldBindJSON(&updatedBook);err!=nil{
		ctx.AbortWithError(http.StatusBadRequest,err)
		return
	}
	for i,book:= range BookDatas{
		if bookID ==book.BookID{
			condition = true
			BookDatas[i] =updatedBook
			BookDatas[i].BookID = bookID
			break
		}
	}
	if !condition{
		ctx.AbortWithStatusJSON(http.StatusNotFound,gin.H{
			"error_status":"Data Not Found",
			"error_message":fmt.Sprint("book with id %v not found",bookID),
		})
	}
	ctx.JSON(http.StatusOK,gin.H{
		"message":fmt.Sprint("book with id %v has been succesfully updated",bookID),
	})
}
//DELETE BOOK

func DeleteBook(ctx *gin.Context){
	bookID := ctx.Param("bookID")
	condition:=false
	var bookIndex int

	for i,book:= range BookDatas{
		if bookID == book.BookID{
			condition=true
			bookIndex=i
			break
		}
	}
	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound,gin.H{
			"error_status":"Data not found",
			"error_message":fmt.Sprint("book with id %v not found"),
		})
	}
	copy(BookDatas[bookIndex:],BookDatas[bookIndex+1:])
	BookDatas[len(BookDatas)-1] = Book{}
	BookDatas = BookDatas[:len(BookDatas)-1]

	ctx.JSON(http.StatusOK,gin.H{
		"message":fmt.Sprintf("car with id %v has been succesfuly deleted",bookID),
	})
}

