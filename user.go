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

func getArtistIntersection(user1_artists, user2_artists []Artist) []Artist{
	intersection := []Artist{}
	user1_set := map[string]bool{}

	for _, artist := range user1_artists{
		user1_set[artist.Id] = true
	}

	for _, artist := range user2_artists{
		_, ok := user1_set[artist.Id]
		if ok{
			intersection = append(intersection, artist)
		}
	}

	return intersection
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