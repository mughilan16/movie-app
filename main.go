package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type Movie struct {
	id   string
	name string
}

func main() {
	Home := func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, searchQuery("Avenger"))
	}
	http.HandleFunc("/", Home)

	fmt.Println("Server started at port 8000")
	http.ListenAndServe(":8000", nil)
}

func searchQuery(movieName string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading the env file")
	}
	apiKey := os.Getenv("MOVIE_API")
	if apiKey == "" {
		fmt.Println("ERROr")
		os.Exit(1)
	}
	newUrl := fmt.Sprintf("https://api.themoviedb.org/3/search/movie?api_key=%s&query=%s", apiKey, movieName)
	fmt.Println(newUrl)

	req, err := http.NewRequest("GET", newUrl, nil)
	if err != nil {
		log.Fatal("Error in the request")
		os.Exit(1)
	}
	req.Header.Add("accept", "application/json")

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		log.Fatal("Error in the response")
		os.Exit(1)
	}

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	fmt.Println(body[0])
	return string(body)
}
