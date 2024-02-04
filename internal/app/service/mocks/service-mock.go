// Code generated by MockGen. DO NOT EDIT.
// Source: movie_rental_service.go

// Package mock_service is a generated GoMock package.
package mock_service

import (
	reflect "reflect"

	dto "github.com/RadhaGeethikaKandala/MovieRental/internal/app/dto"
	request "github.com/RadhaGeethikaKandala/MovieRental/internal/app/dto/request"
	response "github.com/RadhaGeethikaKandala/MovieRental/internal/app/dto/response"
	gomock "github.com/golang/mock/gomock"
)

// MockMovieService is a mock of MovieService interface.
type MockMovieService struct {
	ctrl     *gomock.Controller
	recorder *MockMovieServiceMockRecorder
}

// MockMovieServiceMockRecorder is the mock recorder for MockMovieService.
type MockMovieServiceMockRecorder struct {
	mock *MockMovieService
}

// NewMockMovieService creates a new mock instance.
func NewMockMovieService(ctrl *gomock.Controller) *MockMovieService {
	mock := &MockMovieService{ctrl: ctrl}
	mock.recorder = &MockMovieServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMovieService) EXPECT() *MockMovieServiceMockRecorder {
	return m.recorder
}

// GetMovieDetails mocks base method.
func (m *MockMovieService) GetMovieDetails(imdbid string) (response.MovieResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMovieDetails", imdbid)
	ret0, _ := ret[0].(response.MovieResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMovieDetails indicates an expected call of GetMovieDetails.
func (mr *MockMovieServiceMockRecorder) GetMovieDetails(imdbid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMovieDetails", reflect.TypeOf((*MockMovieService)(nil).GetMovieDetails), imdbid)
}

// GetMovies mocks base method.
func (m *MockMovieService) GetMovies(movieName string) ([]dto.Movie, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMovies", movieName)
	ret0, _ := ret[0].([]dto.Movie)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMovies indicates an expected call of GetMovies.
func (mr *MockMovieServiceMockRecorder) GetMovies(movieName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMovies", reflect.TypeOf((*MockMovieService)(nil).GetMovies), movieName)
}

// GetMoviesFromDb mocks base method.
func (m *MockMovieService) GetMoviesFromDb(arg0 *request.MoviesRequest) response.TruncatedMovieReponse {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMoviesFromDb", arg0)
	ret0, _ := ret[0].(response.TruncatedMovieReponse)
	return ret0
}

// GetMoviesFromDb indicates an expected call of GetMoviesFromDb.
func (mr *MockMovieServiceMockRecorder) GetMoviesFromDb(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMoviesFromDb", reflect.TypeOf((*MockMovieService)(nil).GetMoviesFromDb), arg0)
}
