package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

const baseURL string = "https://api.themoviedb.org/3"

func GetToken() (string, error) {
	err := godotenv.Load(".env.local")
	if err != nil {
		return "", fmt.Errorf("Error loading .env file")
	}

	return os.Getenv("TMDB_API"), nil
}

func getMovieId(arg string) (Movie, error) {
	url := fmt.Sprintf("%s/search/movie?query=%s", baseURL, arg)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Movie{}, fmt.Errorf("Error sending request: %w", err)
	}

	token, err := GetToken()
	if err != nil {
		return Movie{}, fmt.Errorf("Error loading .env file")
	}

	req.Header.Add("accept", "application/json")
	req.Header.Add("Authorization", "Bearer "+token)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return Movie{}, fmt.Errorf("error sending request: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return Movie{}, fmt.Errorf("request failed with status: %d", res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return Movie{}, fmt.Errorf("error reading response: %w", err)
	}

	var searchResp SearchResponse
	err = json.Unmarshal(body, &searchResp)
	if err != nil {
		return Movie{}, fmt.Errorf("error parsing response: %w", err)
	}

	if len(searchResp.Results) == 0 {
		return Movie{}, fmt.Errorf("no movies found for query: %s", arg)
	}

	return searchResp.Results[0], nil
}

func GetCast(arg string) ([]Actor, error) {
	movie, err := getMovieId(arg)
	if err != nil {
		return nil, fmt.Errorf("error fetching movie id: %w", err)
	}

	url := fmt.Sprintf("%s/movie/%d/credits", baseURL, movie.ID)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	token, err := GetToken()
	if err != nil {
		return nil, fmt.Errorf("Error loading .env file")
	}

	req.Header.Add("accept", "application/json")
	req.Header.Add("Authorization", "Bearer "+token)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}

	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("request failed with status: %d", res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %w", err)
	}

	var result struct {
		Cast []Actor `json:"cast"`
	}

	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, fmt.Errorf("Error parsing response: %w", err)
	}

	return result.Cast, nil
}

type Actor struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Character string `json:"character"`
}

type Movie struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

type SearchResponse struct {
	Page    int     `json:"page"`
	Results []Movie `json:"results"`
}
