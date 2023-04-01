package helper

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

func ErrorBadReq(ctx *gin.Context, err string) {
	var resp model.WebResponse
	resp.Message = err
	resp.Data = nil
	ctx.JSON(http.StatusBadRequest, resp)
	ctx.Abort()
}

func ErrorInternal(ctx *gin.Context, err string) {
	var resp model.WebResponse
	resp.Message = err
	resp.Data = nil
	ctx.JSON(http.StatusInternalServerError, resp)
	ctx.Abort()
}

func CatchError(ctx *gin.Context, err string) {
	switch err {
	case "NF":
		ErrorNotFound(ctx, "Data Tidak Di Temukan")
	case "BR":
		ErrorBadReq(ctx, "Data Yang Di Masukkan Tidak Valid")
	case "ISE":
		ErrorInternal(ctx, "Gangguan Pada Sisi Server")
	}

}
