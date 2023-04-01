package handler

import (
	"net/http"

	"github.com/arudji1211/go-dts-07/module/model"
	"github.com/gin-gonic/gin"
)

func ErrorNotFound(ctx *gin.Context, err string) {
	var resp model.WebResponse
	resp.Message = err
	resp.Data = nil
	ctx.JSON(http.StatusNotFound, resp)
	ctx.Abort()
}

func ErrorBadReq(ctx *gin.Context) {
	var resp model.WebResponse
	resp.Message = "Data Yang Di Masukkan Tidak Valid"
	resp.Data = nil
	ctx.JSON(http.StatusBadRequest, resp)
	ctx.Abort()
}
