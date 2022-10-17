package main

import (
	"github.com/stretchr/testify/assert"
	"go-tagle/boot"
	"go-tagle/model"
	"go-tagle/pkg/test"
	"go-tagle/pkg/viperlib"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

func TestMain(m *testing.M) {
	boot.Initialize()
	m.Run()
}

func TestPing(t *testing.T) {
	a := assert.New(t)
	req, err := http.NewRequest("GET", "http://localhost:8000/ping", nil)
	client := &http.Client{}
	resp, err := client.Do(req)
	a.Nil(err)
	a.Equal(http.StatusOK, resp.StatusCode)
}

func TestRegister(t *testing.T) {
	user := &model.User{Username: "test"}

	if user.IsExistUsername() {
		user.DeleteUserByUsername()
	}

	a := assert.New(t)
	req, _ := http.NewRequest("POST", "http://localhost:"+viperlib.GetString("server.port")+"/user/register", strings.NewReader(`{"username":"test","password":"123456","email":"3220293029@163.com"}`))
	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	a.Nil(err)
	a.Equal(http.StatusOK, resp.StatusCode)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	t.Log(string(body))
}

func TestHabit(t *testing.T) {
	a := assert.New(t)
	req, _ := http.NewRequest("POST", "http://localhost:"+viperlib.GetString("server.port")+"/user/login", strings.NewReader(`{"account":"test","password":"123456"}`))
	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	a.Nil(err)
	a.Equal(http.StatusOK, resp.StatusCode)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	t.Log(string(body))
	session := resp.Header.Get("Set-Cookie")
	t.Log(session)

	{
		test.GetAllHabit(session, t, a)
		//test.CreateHabit(session, t, a)
		test.UpdateHabit(session, t, a)
		test.UpdateHabitFinishedTime(session, t, a)
		test.UpdateHabitUnfinishedTime(session, t, a)
		//test.DeleteHabit(session, t, a)
	}
}

func TestTask(t *testing.T) {
	a := assert.New(t)
	req, _ := http.NewRequest("POST", "http://localhost:"+viperlib.GetString("server.port")+"/user/login", strings.NewReader(`{"account":"test","password":"123456"}`))
	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	a.Nil(err)
	a.Equal(http.StatusOK, resp.StatusCode)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	t.Log(string(body))
	session := resp.Header.Get("Set-Cookie")
	t.Log(session)
	{
		test.GetAllTask(session, t, a)
		//test.CreateTask(session, t, a)
		test.UpdateTask(session, t, a)
		test.UpdateTaskFinishedTime(session, t, a)
	}

}

func TestTemp(t *testing.T) {
}
