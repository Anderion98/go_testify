package main

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestMainHandlerWhenCountMoreThanOk(t *testing.T) {

	req := httptest.NewRequest("GET", "/cafe?count=2&city=moscow", nil) // здесь нужно создать запрос к сервису

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.NoError(t, nil)
	require.NotEmpty(t, responseRecorder.Body)
	assert.Equal(t, responseRecorder.Code, http.StatusOK)

	// здесь нужно добавить необходимые проверки
}

func TestMainHandlerWhenCountMoreBadCity(t *testing.T) {

	req := httptest.NewRequest("GET", "/cafe?count=2&city=novosibirsk", nil) // здесь нужно создать запрос к сервису

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, responseRecorder.Code, http.StatusBadRequest)
	assert.Equal(t, "wrong city value", responseRecorder.Body.String())
	// здесь нужно добавить необходимые проверки
}

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 4
	req := httptest.NewRequest("GET", "/cafe?count=6&city=moscow", nil) // здесь нужно создать запрос к сервису

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.NoError(t, nil)
	require.Equal(t, responseRecorder.Code, http.StatusOK)

	value := strings.Split(responseRecorder.Body.String(), `,`)
	require.LessOrEqual(t, totalCount, len(value))
	assert.Len(t, value, totalCount)

	// здесь нужно добавить необходимые проверки
}
