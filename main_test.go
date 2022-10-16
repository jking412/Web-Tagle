package main

import (
	"github.com/stretchr/testify/assert"
	"go-tagle/pkg/encrypt"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

func TestPing(t *testing.T) {
	a := assert.New(t)
	req, err := http.NewRequest("GET", "http://localhost:8000/ping", nil)
	client := &http.Client{}
	resp, err := client.Do(req)
	a.Nil(err)
	a.Equal(http.StatusOK, resp.StatusCode)
}

func TestRegister(t *testing.T) {
	a := assert.New(t)
	req, _ := http.NewRequest("POST", "http://localhost:8000/user/register", strings.NewReader(`{"username":"test","password":"123456","email":"3220293029@163.com"}`))
	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	a.Nil(err)
	a.Equal(http.StatusOK, resp.StatusCode)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	t.Log(string(body))
}

func TestLogin(t *testing.T) {
	a := assert.New(t)
	req, _ := http.NewRequest("POST", "http://localhost:8000/user/login", strings.NewReader(`{"account":"test","password":"123456"}`))
	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	a.Nil(err)
	a.Equal(http.StatusOK, resp.StatusCode)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	t.Log(string(body))
}

func TestEncrypt(t *testing.T) {
	t.Log(encrypt.EncryptPassword("123456h7ZsdKO6WEtvHIEqRFHHnhJ5X9sNRe0z"))
}
