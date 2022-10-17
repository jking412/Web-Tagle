package test

import (
	"github.com/stretchr/testify/assert"
	"go-tagle/pkg/viperlib"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

func GetAllHabit(session string, t *testing.T, a *assert.Assertions) {
	req, _ := http.NewRequest("GET", "http://localhost:"+viperlib.GetString("server.port")+"/habit/all", nil)
	req.Header.Set("Cookie", session)
	resp, _ := http.DefaultClient.Do(req)
	a.Equal(http.StatusOK, resp.StatusCode)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	t.Log(string(body))
}

//func CreateHabit(session string, t *testing.T, a *assert.Assertions) {
//	req, _ := http.NewRequest("POST", "http://localhost:"+viperlib.GetString("server.port")+"/habit/create", strings.NewReader(`{"name":"test"}`))
//	req.Header.Set("Content-Type", "application/json")
//	req.Header.Set("Cookie", session)
//	resp, _ := http.DefaultClient.Do(req)
//	a.Equal(http.StatusOK, resp.StatusCode)
//	defer resp.Body.Close()
//	body, _ := ioutil.ReadAll(resp.Body)
//	t.Log(string(body))
//}

func ErrorCreateHabit(session string, t *testing.T, a *assert.Assertions) {
	req, _ := http.NewRequest("POST", "http://localhost:"+viperlib.GetString("server.port")+"/habit/create", strings.NewReader(`{"nam":"test"}`))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Cookie", session)
	resp, _ := http.DefaultClient.Do(req)
	a.Equal(http.StatusBadRequest, resp.StatusCode)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	t.Log(string(body))
}

func UpdateHabit(session string, t *testing.T, a *assert.Assertions) {
	req, _ := http.NewRequest("POST", "http://localhost:"+viperlib.GetString("server.port")+"/habit/update", strings.NewReader(`{"id":10,"name":"testUpdate"}`))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Cookie", session)
	resp, _ := http.DefaultClient.Do(req)
	a.Equal(http.StatusOK, resp.StatusCode)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	t.Log(string(body))
}

func UpdateHabitFinishedTime(session string, t *testing.T, a *assert.Assertions) {
	req, _ := http.NewRequest("POST", "http://localhost:"+viperlib.GetString("server.port")+"/habit/finish", strings.NewReader(`{"Id":10,"finishedTime":1}`))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Cookie", session)
	resp, _ := http.DefaultClient.Do(req)
	a.Equal(http.StatusOK, resp.StatusCode)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	t.Log(string(body))
}

func UpdateHabitUnfinishedTime(session string, t *testing.T, a *assert.Assertions) {
	req, _ := http.NewRequest("POST", "http://localhost:"+viperlib.GetString("server.port")+"/habit/unfinish", strings.NewReader(`{"Id":10,"unfinishedTime":1}`))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Cookie", session)
	resp, _ := http.DefaultClient.Do(req)
	a.Equal(http.StatusOK, resp.StatusCode)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	t.Log(string(body))
}

//func DeleteHabit(session string, t *testing.T, a *assert.Assertions) {
//	req, _ := http.NewRequest("POST", "http://localhost:"+viperlib.GetString("server.port")+"/habit/delete", strings.NewReader(`{"id":1}`))
//	req.Header.Set("Content-Type", "application/json")
//	req.Header.Set("Cookie", session)
//	resp, _ := http.DefaultClient.Do(req)
//	a.Equal(http.StatusOK, resp.StatusCode)
//	defer resp.Body.Close()
//	body, _ := ioutil.ReadAll(resp.Body)
//	t.Log(string(body))
//}
