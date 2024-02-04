// Code generated by MockGen. DO NOT EDIT.
// Source: movie_rental_repository.go

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	reflect "reflect"

	dto "github.com/RadhaGeethikaKandala/MovieRental/internal/app/dto"
	request "github.com/RadhaGeethikaKandala/MovieRental/internal/app/dto/request"
	gomock "github.com/golang/mock/gomock"
)

// MockRepository is a mock of Repository interface.
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository.
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance.
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// AddMovieToCart mocks base method.
func (m *MockRepository) AddMovieToCart(custId, imdbId string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddMovieToCart", custId, imdbId)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddMovieToCart indicates an expected call of AddMovieToCart.
func (mr *MockRepositoryMockRecorder) AddMovieToCart(custId, imdbId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddMovieToCart", reflect.TypeOf((*MockRepository)(nil).AddMovieToCart), custId, imdbId)
}

// GetCustomer mocks base method.
func (m *MockRepository) GetCustomer(custId string) dto.Customer {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCustomer", custId)
	ret0, _ := ret[0].(dto.Customer)
	return ret0
}

// GetCustomer indicates an expected call of GetCustomer.
func (mr *MockRepositoryMockRecorder) GetCustomer(custId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCustomer", reflect.TypeOf((*MockRepository)(nil).GetCustomer), custId)
}

// GetMovie mocks base method.
func (m *MockRepository) GetMovie(imdbId string) dto.Movie {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMovie", imdbId)
	ret0, _ := ret[0].(dto.Movie)
	return ret0
}

// GetMovie indicates an expected call of GetMovie.
func (mr *MockRepositoryMockRecorder) GetMovie(imdbId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMovie", reflect.TypeOf((*MockRepository)(nil).GetMovie), imdbId)
}

// GetMovies mocks base method.
func (m *MockRepository) GetMovies(qmovieRequest *request.MoviesRequest) []dto.Movie {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMovies", qmovieRequest)
	ret0, _ := ret[0].([]dto.Movie)
	return ret0
}

// GetMovies indicates an expected call of GetMovies.
func (mr *MockRepositoryMockRecorder) GetMovies(qmovieRequest interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMovies", reflect.TypeOf((*MockRepository)(nil).GetMovies), qmovieRequest)
}

// GetRatingsFor mocks base method.
func (m *MockRepository) GetRatingsFor(movieId int) []dto.Rating {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRatingsFor", movieId)
	ret0, _ := ret[0].([]dto.Rating)
	return ret0
}

// GetRatingsFor indicates an expected call of GetRatingsFor.
func (mr *MockRepositoryMockRecorder) GetRatingsFor(movieId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRatingsFor", reflect.TypeOf((*MockRepository)(nil).GetRatingsFor), movieId)
}
