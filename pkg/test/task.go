package test

import (
	"github.com/stretchr/testify/assert"
	"go-tagle/pkg/config"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

func GetAllTask(session string, t *testing.T, a *assert.Assertions) {
	req, _ := http.NewRequest("GET", "http://localhost:"+config.GetString("server.port")+"/task/all", nil)
	req.Header.Set("Cookie", session)
	resp, _ := http.DefaultClient.Do(req)
	a.Equal(http.StatusOK, resp.StatusCode)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	t.Log(string(body))
}

func CreateTask(session string, t *testing.T, a *assert.Assertions) {
	req, _ := http.NewRequest("POST", "http://localhost:"+config.GetString("server.port")+"/task/create", strings.NewReader(`{"name":"test"}`))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Cookie", session)
	resp, _ := http.DefaultClient.Do(req)
	a.Equal(http.StatusOK, resp.StatusCode)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	t.Log(string(body))
}

func ErrorCreateTask(session string, t *testing.T, a *assert.Assertions) {
	req, _ := http.NewRequest("POST", "http://localhost:"+config.GetString("server.port")+"/task/create", strings.NewReader(`{"nam":"test"}`))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Cookie", session)
	resp, _ := http.DefaultClient.Do(req)
	a.Equal(http.StatusBadRequest, resp.StatusCode)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	t.Log(string(body))
}

func UpdateTask(session string, t *testing.T, a *assert.Assertions) {
	req, _ := http.NewRequest("POST", "http://localhost:"+config.GetString("server.port")+"/task/update", strings.NewReader(`{"id":10,"name":"testUpdate"}`))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Cookie", session)
	resp, _ := http.DefaultClient.Do(req)
	a.Equal(http.StatusOK, resp.StatusCode)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	t.Log(string(body))
}

func UpdateTaskFinishedTime(session string, t *testing.T, a *assert.Assertions) {
	req, _ := http.NewRequest("POST", "http://localhost:"+config.GetString("server.port")+"/task/finish", strings.NewReader(`{"id":1,"finishedTime":1}`))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Cookie", session)
	resp, _ := http.DefaultClient.Do(req)
	a.Equal(http.StatusOK, resp.StatusCode)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	t.Log(string(body))
}

func DeleteTask(session string, t *testing.T, a *assert.Assertions) {
	req, _ := http.NewRequest("POST", "http://localhost:"+config.GetString("server.port")+"/task/delete", strings.NewReader(`{"id":1}`))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Cookie", session)
	resp, _ := http.DefaultClient.Do(req)
	a.Equal(http.StatusOK, resp.StatusCode)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	t.Log(string(body))
}
