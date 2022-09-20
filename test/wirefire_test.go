package test

import (
	"WildFireTest/config"
	"WildFireTest/controller"
	"WildFireTest/server"
	"context"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func runServer() *config.Config {
	appConfiguration, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	server := server.NewServer(
		appConfiguration,
		controller.NewWildFireController(
			appConfiguration.App.Count,
			appConfiguration.App.Limit,
		),
	)

	if err = server.Run(); err != nil {
		panic(err)
	}

	return appConfiguration
}

type RandomJokes struct {
	Data []string `json:"data"`
}

func Test_Fetch_Random_Joke(t *testing.T) {
	appConfig := runServer()

	requestURL := "https://localhost:5000/"

	wCtx, cancel := context.WithCancel(context.Background())
	defer cancel()

	req, err := http.NewRequestWithContext(
		wCtx,
		"GET",
		requestURL,
		nil,
	)
	require.NoError(t, err, "Making Request Error")

	res, err := http.DefaultClient.Do(req)
	require.NoError(t, err, "Receive Response Error")

	defer res.Body.Close()

	var result RandomJokes
	err = json.NewDecoder(res.Body).Decode(&result)
	require.NoError(t, err, "Not Correct Response")

	require.Equal(t, appConfig.App.Count, len(result.Data), "Fetch Random Joke Error")
}
