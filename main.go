package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
)

// User struct
type User struct {
	Login             string `json:"login"`
	ID                int    `json:"id"`
	NodeID            string `json:"node_id"`
	AvatarURL         string `json:"avatar_url"`
	GravatarID        string `json:"gravatar_id"`
	URL               string `json:"url"`
	HTMLURL           string `json:"html_url"`
	FollowersURL      string `json:"followers_url"`
	FollowingURL      string `json:"following_url"`
	GistsURL          string `json:"gists_url"`
	StarredURL        string `json:"starred_url"`
	SubscriptionsURL  string `json:"subscriptions_url"`
	OrganizationsURL  string `json:"organinzataions_url"`
	ReposURL          string `json:"repos_url"`
	EventsURL         string `json:"events_url"`
	ReceivedEventsURL string `json:"received_events_url"`
	Type              string `json:"type"`
	SiteAdmin         bool   `json:"site_admin"`
	Name              string `json:"name"`
	Company           string `json:"company"`
	Blog              string `json:"blog"`
	Location          string `json:"location"`
	Email             string `json:"email"`
	Hireable          bool   `json:"hireable"`
	Bio               string `json:"bio"`
	TwitterUsername   string `json:"twitter_username"`
	PublicRepos       int    `json:"public_repos"`
	PublicGists       int    `json:"public_gists"`
	Followers         int    `json:"followers"`
	Following         int    `json:"following"`
	CreatedAT         string `json:"created_at"`
	UpdatedAt         string `json:"updated_at"`
}

func worker(username string, wg *sync.WaitGroup) {
	defer wg.Done()
	// fmt.Printf("+%v\n", getUserInfo(username))
	fptr, err := os.Create(fmt.Sprint(username + ".txt"))
	if err != nil {
		log.Fatal(err)
	}
	jsonBytes, err := json.Marshal(fmt.Sprint(getUserInfo(username)))
	if err != nil {
		log.Fatal(err)
	}

	_, err = fptr.Write(jsonBytes)

}

func getUserInfo(username string) (user User) {

	fmt.Println(username)
	res, err := http.Get(fmt.Sprintf("https://api.github.com/users/" + username))
	if err != nil {

		log.Fatal(err)
	}

	err = json.NewDecoder(res.Body).Decode(&user)
	if err != nil {
		log.Fatal(err)
	}
	// userChan <- user
	return

}

func main() {
	usernames := []string{"nishantasarma", "gauravssnl"}
	var wg sync.WaitGroup

	for _, username := range usernames {
		wg.Add(1)
		go worker(username, &wg)
	}
	wg.Wait()
}
