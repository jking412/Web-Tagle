package test

import (
	"fmt"
	"go-tagle/pkg/viperlib"
	"io/ioutil"
	"net/http"
	"strings"
)

func CreateHabit(session string) {
	req, _ := http.NewRequest("POST", "http://localhost:"+viperlib.GetString("server.port")+"/habit/create", strings.NewReader(`{"name":"test"}`))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Cookie", session)
	resp, _ := http.DefaultClient.Do(req)
	fmt.Println(resp.StatusCode)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}
