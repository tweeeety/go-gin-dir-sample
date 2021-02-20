package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tweeeety/go-gin-gae-sample/service"
)

func TodoIndex(ctx *gin.Context) {
	todoService := service.NewTodoService()
	todos := todoService.GetAll()
	ctx.HTML(http.StatusOK, "index.html", gin.H{"todos": todos})
}

func TodoNew(ctx *gin.Context) {
	text := ctx.PostForm("text")
	status := ctx.PostForm("status")

	todoService := service.NewTodoService()
	todoService.Add(text, status)
	ctx.Redirect(http.StatusFound, "/")
}

func TodoDetail(ctx *gin.Context) {
	n := ctx.Param("id")
	id, err := strconv.Atoi(n)
	if err != nil {
		panic(err)
	}

	todoService := service.NewTodoService()
	todo := todoService.GetOne(id)
	ctx.HTML(http.StatusOK, "detail.html", gin.H{"todo": todo})
}

func TodoUpdate(ctx *gin.Context) {
	n := ctx.Param("id")
	id, err := strconv.Atoi(n)
	if err != nil {
		panic("ERROR")
	}
	text := ctx.PostForm("text")
	status := ctx.PostForm("status")

	todoService := service.NewTodoService()
	todoService.Update(id, text, status)

	ctx.Redirect(http.StatusFound, "/")
}

func TodoDeleteConfirm(ctx *gin.Context) {
	n := ctx.Param("id")
	id, err := strconv.Atoi(n)
	if err != nil {
		panic("ERROR")
	}

	todoService := service.NewTodoService()
	todo := todoService.GetOne(id)
	ctx.HTML(http.StatusOK, "delete.html", gin.H{"todo": todo})
}

func TodoDelete(ctx *gin.Context) {
	n := ctx.Param("id")
	id, err := strconv.Atoi(n)
	if err != nil {
		panic("ERROR")
	}
	todoService := service.NewTodoService()

	todoService.Delete(id)
	ctx.Redirect(http.StatusFound, "/")

}
