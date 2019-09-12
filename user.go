package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"
)

func getUserById(id string) (*User, error){
	user := User{}
	// TODO: make this error proof
	if err := db.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func analyzeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	vars := mux.Vars(r)
	user_id := vars["user_id"]
	user, err := getUserById(user_id)
	if err != nil {
	    w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error fetching user " + err.Error()))
		return
	}

	var artists []Artist
	var genres []Genre
	var topArtists [][]Artist
	var topSongs [][]Track

	err = json.Unmarshal([]byte(user.Artists), &artists)
	if err != nil {
	    w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error fetching user artists " + err.Error()))
		return
	}
	err = json.Unmarshal([]byte(user.Genres), &genres)
	if err != nil {
	    w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error fetching user genres " + err.Error()))
		return
	}
	err = json.Unmarshal([]byte(user.TopArtists), &topArtists)
	if err != nil {
	    w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error fetching user top artists " + err.Error()))
		return
	}
	err = json.Unmarshal([]byte(user.TopSongs), &topSongs)
	if err != nil {
	    w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error fetching user top songs " + err.Error()))
		return
	}

	response := map[string]interface{}{}
	response["username"] = user.DisplayName
	response["last_refreshed"] = user.LastRefreshed
	response["artists"] = artists
	response["genres"] = genres
	response["top_artists"] = topArtists
	response["top_songs"] = topSongs

	json.NewEncoder(w).Encode(response)
}