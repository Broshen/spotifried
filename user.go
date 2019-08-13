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
	w.Header().Set("Access-Control-Allow-Origin", "*")
	vars := mux.Vars(r)
	user_id := vars["user_id"]
	user := getUserById(user_id)

	var artists []Artist
	var genres []Genre

	err := json.Unmarshal([]byte(user.Artists), &artists)
	if err != nil {
	    w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	err = json.Unmarshal([]byte(user.Genres), &genres)
	if err != nil {
	    w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	response := map[string]interface{}{}
	response["artists"] = artists
	response["genres"] = genres

	json.NewEncoder(w).Encode(response)
}