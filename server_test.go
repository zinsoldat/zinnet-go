package main_test

import (
	"fmt"
	"net/http"
	"testing"

	zinnet "github.com/zinsoldat/zinnet-go"
)

var serverConfig = &zinnet.ServerConfig{
	Port: 3000,
	Host: "127.0.0.1",
}

func TestCreateServer(t *testing.T) {
	zinnet.NewServer(serverConfig)
	// server should be creatable.
}

func TestServerStart(t *testing.T) {
	server := zinnet.NewServer(serverConfig)
	go server.Start()
	defer server.Stop()

	request, _ := http.NewRequest("GET", fmt.Sprintf("http://%s:%d", serverConfig.Host, serverConfig.Port), nil)
	client := http.Client{}
	_, err := client.Do(request)
	if err != nil {
		t.Errorf("could not request root path `/`: %s", err.Error())
	}
}
