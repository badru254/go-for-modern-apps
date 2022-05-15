package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"
)

func main() {

	user, err := fetchUser("badru254")

	checkError(err)

	fmt.Printf("Data found : %+v", *user)

}

func fetchUser(login string) (*User, error) {

	url := fmt.Sprintf("https://api.github.com/users/%s", url.PathEscape(login))

	//-------------Using context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 3000*time.Millisecond)
	defer cancel()

	// resp, err := http.Get(url)

	// checkError(err)

	//-------------Using context with timeout
	req, err1 := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)

	if err1 != nil {
		return nil, err1
	}

	resp, err2 := http.DefaultClient.Do(req)
	if err2 != nil {
		return nil, err2
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(resp.Status)
	}

	defer resp.Body.Close()

	//Using Marshall/Unmarshal json

	//body, err1 := io.ReadAll(resp.Body)
	//checkError(err1)

	// if err1!= nil {
	// 	return nil, err1
	// }

	// //var user interface{}
	// var user User
	// err2 := json.Unmarshal(body, &user)
	// if err2 != nil {
	// 	return nil, err2
	// }
	//Using decode json
	// var user User
	// dec := json.NewDecoder(resp.Body)
	// if err3 := dec.Decode(&user); err3 != nil {
	// 	if err3 != nil {
	// 	return nil, err3
	// }
	// }

	//Using decode with timeout
	var user User
	dec := json.NewDecoder(resp.Body)
	if err3 := dec.Decode(&user); err3 != nil {
		if err3 != nil {
			return nil, err3
		}
	}

	return &user, nil
}
func checkError(err error) {
	if err != nil {
		log.Fatalf("An error occured - %s", err)
	}
}

//User id a user description
type User struct {
	// Login    string `json:"login"`
	// Name     string `json:"name"`
	// PublicRepos int    `json:"public_repos"`
	Login        string
	Name         string
	Public_Repos int
}
