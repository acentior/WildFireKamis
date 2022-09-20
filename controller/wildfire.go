package controller

import (
	"WildFireTest/model"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

type wildfireController struct {
	generate_count   int // fetch count
	concurrent_limit int // maximum count of goroutines
}

func NewWildFireController(gn_cnt int, cc_limit int) Controller {
	return &wildfireController{
		generate_count:   gn_cnt,
		concurrent_limit: cc_limit,
	}
}

func (ctrl *wildfireController) Handler() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		waitChan := make(chan struct{}, ctrl.concurrent_limit)
		defer close(waitChan)
		out := make(chan error)

		var wg sync.WaitGroup
		wg.Add(ctrl.generate_count)

		var jokes []model.JokeModel // Joke Array that will be fetched
		var errors []string         // Error Array
		curIndex := 0

		for {
			waitChan <- struct{}{}

			curIndex += 1

			// if last fetch started, wait till last routine finishes
			if curIndex > ctrl.generate_count {
				wg.Wait()
				break
			}

			go func(cur int) {
				name := getName(out)
				joke := getJoke(name, out)
				jokes = append(jokes, joke)

				wg.Done()
				<-waitChan
			}(curIndex)

			// if some comes in out channel, print error on console and add it to error list
			go func() {
				if err := (<-out); err != nil {
					errors = append(errors, err.Error())
					fmt.Println(err.Error())
				}
			}()
		}

		// check whether fetched jokes or not. If fetched correctly, send responses with fetched data, if not response with errors
		if len(jokes) > 0 {
			ctx.JSON(
				http.StatusOK,
				jokes,
			)
		} else {
			ctx.JSON(
				http.StatusInternalServerError,
				errors,
			)
		}
	}
}

// Route for WildFireHandler
func (ctrl *wildfireController) Route() string {
	return ""
}

// Method for WildFireHandler
func (ctrl *wildfireController) Method() string {
	return http.MethodGet
}

/*
	fetch random name from "https://names.mcquay.me/api/v0/"

example : {“first_name”:“Hasina”,“last_name”:“Tanweer”}
*/
func getName(out chan error) model.NameModel {
	requestURL := "https://names.mcquay.me/api/v0/"

	wCtx, cancel := context.WithCancel(context.Background())
	defer cancel()

	req, err := http.NewRequestWithContext(
		wCtx,
		"GET",
		requestURL,
		nil,
	)
	if err != nil {
		out <- err
		return model.NameModel{}
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		out <- err
		return model.NameModel{}
	}

	defer res.Body.Close()

	var name model.NameModel
	if err = json.NewDecoder(res.Body).Decode(&name); err != nil {
		out <- err
		cancel()
		return model.NameModel{}
	}

	fmt.Println("----------------- Name --------------------")
	fmt.Println(name)
	fmt.Println("-------------------------------------------")

	return name
}

/*
	generate requestURL to fetch random Joke with FirstName and LastName

example : "http://api.icndb.com/jokes/random?firstName=John&lastName=Doe&limitTo=nerdy"
*/
func generateJokeFetchURL(name model.NameModel) string {
	url := "http://api.icndb.com/jokes/random?"
	first_name := fmt.Sprintf("firstName=%s", name.First_Name)
	second_name := fmt.Sprintf("lastName=%s", name.Last_Name)

	result := url + "&" + first_name + "&" + second_name + "&limitTo=nerdy"
	return result
}

/*
	fetch Joke with FirstName and LastName from http://api.icndb.com/jokes/random

example : "http://api.icndb.com/jokes/random?firstName=John&lastName=Doe&limitTo=nerdy"
{ “type”: “success”, “value”: { “id”: 181, “joke”: “John Doe’s OSI network model has only one layer - Physical.“, “categories”: [“nerdy”] } }
*/
func getJoke(name model.NameModel, out chan error) model.JokeModel { // ! API is not working

	// TODO: Fetch using Correct API. Current API for fetching random joke with name is not working.

	// generate a url with name to fetch random joke
	requestURL := generateJokeFetchURL(name)

	wCtx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// make new request
	req, err := http.NewRequestWithContext(
		wCtx,
		"GET",
		requestURL,
		nil,
	)
	if err != nil {
		out <- err
		return model.JokeModel{}
	}

	// get response
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		out <- err
		return model.JokeModel{}
	}

	defer res.Body.Close()

	// Decode response to JokeModel
	var joke model.JokeModel
	if err = json.NewDecoder(res.Body).Decode(&joke); err != nil {
		out <- err
		cancel()
		return model.JokeModel{}
	}

	return joke
}
