package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/fatih/color"
)

func main() {
	// confirm that user provided input
	if len(os.Args) != 2 {
		fmt.Println("usage: ./hind-search file")
		return
	}
	// check if input is a file
	file, err := os.Open(os.Args[1])
	if err != nil {
		color.Red(err.Error())
		return
	}
	// attempt to get authToken. Retry on failure.
	for err = getOAuthToken(); err != nil; err = getOAuthToken() {
		color.Red("API authentication failed.\n - %v\n", err)
		fmt.Printf("Retrying in 5 seconds...\n\n")
		time.Sleep(5 * time.Second)
	}
	// iterate input file for users and get the locations
	fmt.Println("----------------------")
	fmt.Printf("FIND-THAT-HIND v0.1.3\n")
	fscanner := bufio.NewScanner(file)
	for fscanner.Scan() {
		line := fscanner.Text()
		fmt.Println("----------------------")
		fmt.Println(line)
		getUserLocations(line)
		time.Sleep(500 * time.Millisecond) // API only allows 2 calls per second
	}
	fmt.Println("----------------------")
}

// OAuthToken : this is the token the 42 API returns
var OAuthToken struct {
	Token      string `json:"access_token"`
	Type       string `json:"token_type"`
	Expires    int    `json:"expires_in"`
	Scope      string `json:"scope"`
	Created    int    `json:"created_at"`
	Expiration time.Time
}

// Locations : this is the json the 42 API returns
type Locations []struct {
	ID       int         `json:"id"`
	BeginAt  string      `json:"begin_at"`
	EndAt    string      `json:"end_at"`
	Primary  bool        `json:"primary"`
	Floor    interface{} `json:"floor"`
	Row      interface{} `json:"row"`
	Post     interface{} `json:"post"`
	Host     string      `json:"host"`
	CampusID int         `json:"campus_id"`
	User     struct {
		ID    int    `json:"id"`
		Login string `json:"login"`
		URL   string `json:"url"`
	} `json:"user"`
}

var client = &http.Client{Timeout: time.Second * 5}

func getOAuthToken() error {
	data := url.Values{
		"grant_type":    {"client_credentials"},
		"client_id":     {os.Getenv("CLIENT_ID_42API")},
		"client_secret": {os.Getenv("SECRET_ID_42API")},
	}
	// request authentication token
	res, err := client.PostForm("https://api.intra.42.fr/oauth/token", data)
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("%v", res.Status)
	}
	defer res.Body.Close()
	// save token
	err = json.NewDecoder(res.Body).Decode(&OAuthToken)
	if err != nil {
		return err
	}
	OAuthToken.Expiration = time.Now().Add(time.Duration(OAuthToken.Expires) * time.Second)
	return nil
}

func getUserLocations(userID string) {
	URL := fmt.Sprintf("https://api.intra.42.fr/v2/users/%v/locations", userID)
	var locations Locations

	// confirm token is still valid
	if OAuthToken.Expiration.Sub(time.Now()) < 1 {
		if err := getOAuthToken(); err != nil {
			color.Red("- OAuth Token expired and re-authentication failed.")
			return
		}
	}
	// make request for location
	req, _ := http.NewRequest("GET", URL, nil)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", OAuthToken.Token))
	res, err := client.Do(req)
	if err != nil {
		color.Red(fmt.Sprintln("-", err))
		return
	}
	if res.StatusCode != http.StatusOK {
		if res.StatusCode == http.StatusNotFound {
			color.Red("- not a valid userId")
		} else {
			color.Red("- %v", res.Status)
		}
		return
	}
	defer res.Body.Close()

	json.NewDecoder(res.Body).Decode(&locations)
	if locations[0].EndAt == "" {
		color.Green(fmt.Sprintln("- location:", locations[0].Host))
	} else {
		color.Blue("- offline")
	}
}

/*
	- GET TOKEN
	curl -X POST --data "grant_type=client_credentials&client_id=297c991e52306c9b2de9b470e169e853b84046356c666c74dbe1f89ae8006e10&client_secret=80bb6892c0baf733f64846167c3e09554581fca08d3d02299c5cbff429945797" https://api.intra.42.fr/oauth/token

	- MAKE REQUEST
	curl  -H "Authorization: Bearer 6a918760be133b4cc62de83d3f58a83ed46df2e3514b5cbc38f124dc71137275" "https://api.intra.42.fr/v2/users/lilam/locations"

	- CHECK TIME ON TOKEN
	fmt.Println(OAuthToken.Expiration)
	fmt.Println(time.Now())
	fmt.Println(OAuthToken.Expiration.Sub(time.Now()))
	fmt.Println(OAuthToken.Expires)
*/
