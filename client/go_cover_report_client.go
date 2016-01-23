package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
)

type message struct {
	Name string
	Body string
}

func main() {
	cmd := exec.Command("go", "test", "-cover")
	out, err := cmd.Output()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	m := message{"GoCoverReporter", string(out)}
	b, err := json.Marshal(m)

	url := "http://" + os.Args[1] + "/receiver"

	var jsonStr = b
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("X-Custom-Header", "GoCoverReporter")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
}
