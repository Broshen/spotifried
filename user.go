package main

import (
	"fmt"
	"strconv"
	"time"
	"net/http"
	"io/ioutil"
	"encoding/json"
)

type User struct{
	ID 				string `gorm:"primary_key"`
	DisplayName 	string
	AccessToken		string
	RefreshToken	string
}

type TrackResponse struct {
  Items [] Song
  Next string
}

type Song struct {
    Added_at string
    Track struct {
      Artists [] struct {
        Id string
        Name string
      }
      Id string
      Name string
    }
}


func getOrCreateUser(access_token, refresh_token string) User{
	r, _ := http.NewRequest("GET", "https://api.spotify.com/v1/me", nil)
	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("Accept", "application/json")
	r.Header.Add("Authorization", fmt.Sprintf("Bearer %s", access_token))

	client := &http.Client{}
	resp, _ := client.Do(r)
	var respVal map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&respVal)

	id := getStringFromJSON(respVal, "id")
	user := User{
		ID: id,
		DisplayName: getStringFromJSON(respVal, "display_name"),
		AccessToken: access_token,
		RefreshToken: refresh_token,
	}
	db.Save(&user)
	return user
}

func getUserById(id string) User{
	user := User{}
	// TODO: make this error proof
	db.Where("id = ?", id).First(&user)

	return user
}

func getAllUserSongs(access_token, refresh_token string) []Song{

	next := "https://api.spotify.com/v1/me/tracks?offset=0&limit=50"
	client := &http.Client{}
	var songs []Song

	for next != "" {

		r, _ := http.NewRequest("GET", next, nil) // URL-encoded payload
		r.Header.Add("Content-Type", "application/json")
		r.Header.Add("Accept", "application/json")
		r.Header.Add("Authorization", fmt.Sprintf("Bearer %s", access_token))

		resp, err := client.Do(r)
		if resp.StatusCode == 429{ // rate limited
			fmt.Println("Request was rate limited, retry after " + resp.Header["Retry-After"][0] + " seconds")
			waittime, _ := strconv.Atoi(resp.Header["Retry-After"][0])
			time.Sleep(time.Duration(waittime) * time.Second)
			continue
		} else if resp.StatusCode != 200{
			bodyBytes, _ := ioutil.ReadAll(resp.Body)
			bodyString := string(bodyBytes)
			panic(fmt.Sprintf("Error while getting user songs: %s", bodyString))
			break

		} else if err != nil{
			panic(fmt.Sprintf("Error while getting user songs: %s", err))
		}

		var respVal TrackResponse
		err = json.NewDecoder(resp.Body).Decode(&respVal)

		// TODO: handle JSON decoding errors
		if err != nil {
			panic(fmt.Sprintf("%s", err))
		}
		next = respVal.Next
		songs = append(songs, respVal.Items...)
	}

	return songs
}

func getSongIntersection(user_1_songs, user_2_songs []Song) []Song{
	intersection := []Song{}
	user_1_set := map[string]bool{}

	for _, song := range user_1_songs{
		user_1_set[song.Track.Id] = true
	}

	for _, song := range user_2_songs{
		_, ok := user_1_set[song.Track.Id]
		if ok{
			intersection = append(intersection, song)
		}
	}

	return intersection
}
