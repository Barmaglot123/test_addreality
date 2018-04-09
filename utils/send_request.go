package utils

import (
    "github.com/gin-gonic/gin"
    "net/http/httptest"
    "bytes"
    "net/http"
)

func SendPostRequest(url string, body []byte, s *gin.Engine) *httptest.ResponseRecorder {
    buf := bytes.NewBuffer(body)
    req, _ := http.NewRequest("POST", url, buf)
    req.Header.Add("Content-Type","application/json")
    w := httptest.NewRecorder()
    s.ServeHTTP(w, req)
    return w
}