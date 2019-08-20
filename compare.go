package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"
	"sort"
)

func shareHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	user_id := vars["user_id"]

	redirect_uri := api_url + "/authenticate?state=" + user_id
	http.Redirect(w,r, redirect_uri, 301)
}

func compareHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	vars := mux.Vars(r)
	user1_id := vars["user1_id"]
	user2_id := vars["user2_id"]

	user1, err := getUserById(user1_id)
	if err != nil {
	    w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	user2, err := getUserById(user2_id)
	if err != nil {
	    w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	var user1_songs, user2_songs []Song
	var user1_artists, user2_artists []Artist

	err = json.Unmarshal([]byte(user1.Songs), &user1_songs)
	if err != nil {
	    w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	err = json.Unmarshal([]byte(user2.Songs), &user2_songs)
	if err != nil {
	    w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	err = json.Unmarshal([]byte(user1.Artists), &user1_artists)
	if err != nil {
	    w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	err = json.Unmarshal([]byte(user2.Artists), &user2_artists)
	if err != nil {
	    w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	songs := getSongIntersection(user1_songs, user2_songs)
	artists := getArtistIntersection(user1_artists, user2_artists)

	sort.Slice(songs, func(i, j int) bool {
		return songs[i].Track.Name < songs[j].Track.Name
	})
	sort.Slice(artists, func(i, j int) bool {
		return artists[i].Name < artists[j].Name
	})

	response := map[string]interface{}{}
	response["songs"] = songs
	response["artists"] = artists
	response["user1"] = map[string]interface{}{
		"name": user1.DisplayName,
		"id": user1.ID,
		"songcount": len(user1_songs),
	}
	response["user2"] = map[string]interface{}{
		"name": user2.DisplayName,
		"id": user1.ID,
		"songcount": len(user2_songs),
	}

	json.NewEncoder(w).Encode(response)
}