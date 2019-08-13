package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"
)

func shareHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	user_id := vars["user_id"]

	redirect_uri := "http://" + r.Host+"/authenticate?state=" + user_id
	http.Redirect(w,r, redirect_uri, 301)
}

func compareHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	user1_id := vars["user1_id"]
	user2_id := vars["user2_id"]

	user1 := getUserById(user1_id)
	user2 := getUserById(user2_id)

	var user1_songs, user2_songs []Song

	err := json.Unmarshal([]byte(user1.Songs), &user1_songs)
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

	songs := getSongIntersection(user1_songs, user2_songs)

	json.NewEncoder(w).Encode(songs)
}