package socialmedia

import (
	"errors"
	"net/http"

	SocmedModel "github.com/aruji1211/myGram/module/model/socialmedia"
	SocmedRepo "github.com/aruji1211/myGram/module/repository/socialmedia"
	MyLog "github.com/aruji1211/myGram/pkg/logger"
	responseTemplate "github.com/aruji1211/myGram/pkg/response"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type SocialmediaServiceImpl struct {
	socmedRepo SocmedRepo.SocialmediaRepository
	Validate   *validator.Validate
}

func NewSocialmediaService(socmedrepo SocmedRepo.SocialmediaRepository, validate *validator.Validate) SocialmediaService {
	return &SocialmediaServiceImpl{
		socmedRepo: socmedrepo,
		Validate:   validate,
	}
}

func (Cs *SocialmediaServiceImpl) GetAll(ctx *gin.Context) (socMeds []SocmedModel.Socialmedia, err error) {
	//logging
	MyLog.LogMyApp("i", "Socialmedia Service Invoked", "SocialmediaService - GetAll", nil)

	socMeds, err = Cs.socmedRepo.GetAll(ctx)
	if err != nil {
		MyLog.LogMyApp("e", "Repository Returning Error", "SocialmediaService - GetAll", err)
		return
	}

	return
}

func (Cs *SocialmediaServiceImpl) GetById(ctx *gin.Context, idSoc int) (socOut SocmedModel.Socialmedia, err error) {
	// panic("not implemented") // TODO: Implement
	MyLog.LogMyApp("i", "Socialmedia Service Invoked", "SocialmediaService - GetById", nil)

	socOut, err = Cs.socmedRepo.GetById(ctx, idSoc)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, responseTemplate.WebResponseFailed{
			Message: "Data tidak di temukan",
			Error:   err.Error(),
		})
		MyLog.LogMyApp("e", "Repository Returning Error", "SociamediaService - GetById", err)
		return
	}

	return
}

func (Cs *SocialmediaServiceImpl) Create(ctx *gin.Context, socIn SocmedModel.Socialmedia) (socOut SocmedModel.Socialmedia, err error) {
	// panic("not implemented") // TODO: Implement
	MyLog.LogMyApp("i", "Socialmedia Service Invoked", "SocialmediaService - Create", nil)

	//validasi input
	MyLog.LogMyApp("i", "Validating Process invoked", "SocialmediaService - Create", nil)
	Cs.Validate = validator.New()
	err = Cs.Validate.Struct(socIn)

	if err != nil {
		MyLog.LogMyApp("e", "Validating Process Error", "SocialmediaService - Create", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, responseTemplate.WebResponseFailed{
			Message: responseTemplate.InvalidBody,
			Error:   err.Error(),
		})
		return
	}

	socOut, err = Cs.socmedRepo.Create(ctx, socIn)

	if err != nil {
		MyLog.LogMyApp("e", "Repository Returning Error", "SocialmediaService - Create", err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, responseTemplate.WebResponseFailed{
			Message: responseTemplate.InternalServer,
			Error:   err.Error(),
		})
		return
	}

	return
}

func (Cs *SocialmediaServiceImpl) Update(ctx *gin.Context, socIn SocmedModel.Socialmedia, idUser int) (socOut SocmedModel.Socialmedia, err error) {
	MyLog.LogMyApp("i", "Socialmedia Service Invoked", "SocialmediaService - Update", nil)

	//autorisasi
	MyLog.LogMyApp("i", "Autorisasi kepemilikan Socialmedia", "SocialmediaService - Update", nil)
	socOut, err = Cs.socmedRepo.GetById(ctx, socIn.Id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, responseTemplate.WebResponseFailed{
			Message: "Data tidak di temukan",
			Error:   err.Error(),
		})
		MyLog.LogMyApp("e", "Repository Returning Error", "SocialmediaService - Update", err)
		return
	}

	if socOut.UserId != idUser {
		MyLog.LogMyApp("e", "Forbidden", "PhotoService - Delete", err)
		err = errors.New("Forbidden")
		ctx.AbortWithStatusJSON(http.StatusForbidden, responseTemplate.WebResponseFailed{
			Message: responseTemplate.Forbidden,
			Error:   responseTemplate.Forbidden,
		})
		return
	}

	MyLog.LogMyApp("i", "Hit Repository For Update Proccess", "SocialmediaService - Update", nil)
	socOut, err = Cs.socmedRepo.Update(ctx, socIn)
	if err != nil {
		MyLog.LogMyApp("e", "Repository Returning Error", "SocialmediaService - Update", err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, responseTemplate.WebResponseFailed{
			Message: responseTemplate.InternalServer,
			Error:   err.Error(),
		})
		return
	}

	return
}

func (Cs *SocialmediaServiceImpl) Delete(ctx *gin.Context, idPhoto int, idUser int) (err error) {
	MyLog.LogMyApp("i", "Socialmedia Service Invoked", "Socialmedia Service - Delete", nil)

	//autorisasi
	MyLog.LogMyApp("i", "Autorisasi kepemilikan Socialmedia", "SocialmediaService - Delete", nil)
	socOut, err := Cs.socmedRepo.GetById(ctx, idPhoto)
	if err != nil {
		MyLog.LogMyApp("e", "Repository Returning Error", "SocialmediaService - Delete", err)
		ctx.AbortWithStatusJSON(http.StatusNotFound, responseTemplate.WebResponseFailed{
			Message: "Data tidak di temukan",
			Error:   err.Error(),
		})
		return
	}

	if socOut.UserId != idUser {
		MyLog.LogMyApp("e", "Forbidden", "PhotoService - Delete", err)
		err = errors.New("Forbidden")
		ctx.AbortWithStatusJSON(http.StatusForbidden, responseTemplate.WebResponseFailed{
			Message: responseTemplate.Forbidden,
			Error:   responseTemplate.Forbidden,
		})
		return
	}

	MyLog.LogMyApp("i", "hit repository", "PhotoService - Delete", nil)
	err = Cs.socmedRepo.Delete(ctx, idPhoto)
	if err != nil {
		MyLog.LogMyApp("e", "Repository Returning Error", "PhotoService - Delete", err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, responseTemplate.WebResponseFailed{
			Message: responseTemplate.InternalServer,
			Error:   err.Error(),
		})
		return
	}

	return
}
