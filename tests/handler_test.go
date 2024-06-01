package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetData_Success(t *testing.T) {
	// Подготовка фиктивного обработчика
	handlerFunc := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Success"))
	}

	// Создание нового запроса GET
	req, err := http.NewRequest("GET", "/", nil)
	assert.NoError(t, err)

	// Создание рекордера ответов
	rr := httptest.NewRecorder()

	// Выполнение запроса
	handlerFunc(rr, req)

	// Проверка кода состояния
	assert.Equal(t, http.StatusOK, rr.Code)

	// Проверка содержимого ответа
	assert.Equal(t, "Success", rr.Body.String())
}

func TestGetData_NotFound(t *testing.T) {
	// Подготовка фиктивного обработчика
	handlerFunc := func(w http.ResponseWriter, r *http.Request) {
		http.NotFound(w, r)
	}

	// Создание нового запроса GET
	req, err := http.NewRequest("GET", "/", nil)
	assert.NoError(t, err)

	// Создание рекордера ответов
	rr := httptest.NewRecorder()

	// Выполнение запроса
	handlerFunc(rr, req)

	// Проверка кода состояния
	assert.Equal(t, http.StatusNotFound, rr.Code)
}

func TestPostData_Success(t *testing.T) {
	// Подготовка фиктивного обработчика
	handlerFunc := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("Data created successfully"))
	}

	// Создание нового запроса POST
	req, err := http.NewRequest("POST", "/", nil)
	assert.NoError(t, err)

	// Создание рекордера ответов
	rr := httptest.NewRecorder()

	// Выполнение запроса
	handlerFunc(rr, req)

	// Проверка кода состояния
	assert.Equal(t, http.StatusCreated, rr.Code)

	// Проверка содержимого ответа
	assert.Equal(t, "Data created successfully", rr.Body.String())
}
