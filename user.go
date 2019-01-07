package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"
)

func getUserById(id string) User{
	user := User{}
	// TODO: make this error proof
	db.Where("id = ?", id).First(&user)

	return user
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

func analyzeHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	user_id := vars["user_id"]

	redirect_uri := "http://" + r.Host+"/authenticated"

	user := getUserById(user_id)

	user.AccessToken, user.RefreshToken, _ = getTokens(user.RefreshToken, redirect_uri, true)
	defer db.Save(&user)

	user_songs := getAllUserSongs(user.AccessToken, user.RefreshToken)
	artists, genres := getAllUserArtistsAndGenres(user.AccessToken, user.RefreshToken, user_songs)

	response := map[string]interface{}{}
	response["artists"] = artists
	response["genres"] = genres
	json.NewEncoder(w).Encode(response)
}