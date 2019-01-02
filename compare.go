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

	redirect_uri := "http://" + r.Host+"/authenticated"

	user1 := getUserById(user1_id)
	user2 := getUserById(user2_id)

	user1.AccessToken, user1.RefreshToken, _ = getTokens(user1.RefreshToken, redirect_uri, true)
	user2.AccessToken, user2.RefreshToken, _ = getTokens(user2.RefreshToken, redirect_uri, true)
	defer db.Save(&user1)
	defer db.Save(&user2)

	user1_songs := getAllUserSongs(user1.AccessToken, user1.RefreshToken)
	user2_songs := getAllUserSongs(user2.AccessToken, user2.RefreshToken)

	songs := getSongIntersection(user1_songs, user2_songs)

	json.NewEncoder(w).Encode(songs)
}