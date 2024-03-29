package handler

import (
	"strings"

	"github.com/RadhaGeethikaKandala/MovieRental/internal/app/dto/request"
	"github.com/RadhaGeethikaKandala/MovieRental/internal/app/dto/response"
	"github.com/RadhaGeethikaKandala/MovieRental/internal/app/service"
	"github.com/gin-gonic/gin"
)

type MovieHandler interface {
	SayHello(ctx *gin.Context)
	GetMoviesFromDb(ctx *gin.Context)
	GetMovieDetails(ctx *gin.Context)
	AddMovieToCart(ctx *gin.Context)
}

type movieHandler struct {
	service service.MovieService
}

func NewHandler(service service.MovieService) MovieHandler {
	return &movieHandler{
		service: service,
	}
}

func (mh movieHandler) SayHello(ctx *gin.Context) {
	ctx.String(200, "hello world!")
}

func (mh movieHandler) GetMoviesFromDb(ctx *gin.Context) {
	var moviesRequest request.MoviesRequest
	ctx.ShouldBindQuery(&moviesRequest)

	truncatedMovieResponse := mh.service.GetMoviesFromDb(&moviesRequest)
	ctx.JSON(200, truncatedMovieResponse)
}

func (mh movieHandler) GetMovieDetails(ctx *gin.Context) {
	imdbid := ctx.Param("imdbid")
	if strings.TrimSpace(imdbid) == "" {
		buildErrorResponse(ctx, "name cannot be empty or have whitespaces")
		return
	}

	movieReponse, err := mh.service.GetMovieDetails(imdbid)
	if err != nil {
		buildErrorResponse(ctx, err.Error())
		return
	}
	ctx.JSON(200, movieReponse)

}

func (mh *movieHandler) AddMovieToCart(ctx *gin.Context) {
	var request request.AddToCartRequest
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		buildErrorResponse(ctx, err.Error())
		return
	}

	err = mh.service.AddMovieToCart(&request)
	if err != nil {
		buildErrorResponse(ctx, err.Error())
		return
	}

	ctx.JSON(200, response.ApiResponse{
		Status:  "success",
		Message: "Added movie to cart successfully",
		Code:    "200",
	})
}

func buildErrorResponse(ctx *gin.Context, message string) {
	ctx.JSON(400, response.ApiResponse{
		Status:  "failure",
		Message: message,
		Code:    "400",
	})
}
