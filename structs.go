package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	type Post struct {
		UserId int
		Id     int
		Title  string
		Body   string
	}

	response, error := http.Get("https://jsonplaceholder.typicode.com/posts")

	var posts []Post

	if error != nil {
		fmt.Println(error)
	}

	if response.StatusCode == 200 {
		responseData, err := ioutil.ReadAll(response.Body)

		if err != nil {
			log.Fatal(err)
		}

		json.Unmarshal(responseData, &posts)

		for _, post := range posts {
			fmt.Println(post.Id, post.Title)
		}

	}
}
