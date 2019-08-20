package main

import (
	"fmt"
	// "log"
	"net/http"
	"net/url"
	"strings"
	"strconv"
	"encoding/json"
	"io/ioutil"
	// "errors"
)

func getStringFromJSON(obj map[string]interface{}, key string) string {
	if val, ok := obj[key].(string); ok {
	    return val
	} else {
	    panic(fmt.Sprintf("Key %s does not exist", key))
	}
}

func getTokens(code, redirect_uri string, refresh bool) (string, string, error){
	data := url.Values{}
	data.Add("client_id", client_id)
	data.Add("client_secret", client_secret)

	if refresh == false{
		data.Add("grant_type", "authorization_code")
		data.Add("code", code)
		data.Add("redirect_uri", redirect_uri)
	} else{
		data.Add("grant_type", "refresh_token")
		data.Add("refresh_token", code)
	}

	r, err := http.NewRequest("POST", "https://accounts.spotify.com/api/token", strings.NewReader(data.Encode())) // URL-encoded payload
	if err != nil{
		panic(fmt.Sprintf("Error while creating getTokens request: %s", err))
	}
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	client := &http.Client{}
	resp, err := client.Do(r)

	// if an error occurred, redirect to error page
	if resp.StatusCode != 200{
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		bodyString := string(bodyBytes)
		panic(fmt.Sprintf("Error while fetching tokens: %s", bodyString))
	} else if err != nil{
		panic(fmt.Sprintf("Error while fetching tokens: %s", err))
	} else{
	// otherwise, parse the body
		respVal := map[string]interface{}{}

		err := json.NewDecoder(resp.Body).Decode(&respVal)

		if err != nil{
			panic(fmt.Sprintf("Error while decoding response JSON: %s", err))
		} else{
			access_token := getStringFromJSON(respVal, "access_token")
			refresh_token := code
			if refresh == false{
				refresh_token = getStringFromJSON(respVal, "refresh_token")
			}
			// unused for now
			// token_type :=respVal["token_type"]
			// scope :=respVal["scope"]
			// expires_in :=respVal["expires_in"]
			
			return access_token, refresh_token, nil
		}
	}
}


func authHandler(w http.ResponseWriter, r *http.Request) {
	keys := r.URL.Query()
	other_user_id := keys.Get("state")
	
	auth_url := fmt.Sprintf(
		"https://accounts.spotify.com/authorize?client_id=%s&redirect_uri=%s&response_type=code&scope=%s&state=%s",
		client_id,
		api_url + "authenticated",
		scopes,
		other_user_id)
	http.Redirect(w, r, auth_url, 301)
}

func authenticatedHandler(w http.ResponseWriter, r *http.Request) {
	keys := r.URL.Query()
	code := keys.Get("code")
	other_user_id := keys.Get("state")

	// use code to get access tokens
	access_token, refresh_token, err := getTokens(code, api_url+"authenticated", false)
	if err != nil{
		panic(fmt.Sprintf("Error getting access tokens:", err))
	}

	user, err := getAllUserData(access_token, refresh_token)
	if err != nil{
		panic(fmt.Sprintf("Error getting user data:", err))
	}

	// if it's the initial user request - i.e. there is no "other user",
	// log the user into the db and return it as a JSON object
	if other_user_id == "" {
		profile_url := fmt.Sprintf("%sprofile/%s", frontend_url, user.ID)
		http.Redirect(w,r, profile_url, 301)
	}else{
	// otherwise, we have two users to compare, redirect to the compare page
		compare_url := fmt.Sprintf("%scompare/%s/%s", frontend_url, other_user_id, user.ID)
		http.Redirect(w,r, compare_url, 301)
	}
}