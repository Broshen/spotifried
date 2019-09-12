package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"
	"sort"
)

func getTrackIntersection(user_1_tracks, user_2_tracks []Track) []Track{
	intersection := []Track{}
	user_1_set := map[string]bool{}

	for _, track := range user_1_tracks{
		user_1_set[track.Id] = false
	}

	for _, track := range user_2_tracks{
		if added, ok := user_1_set[track.Id]; ok && !added{
			user_1_set[track.Id] = true
			intersection = append(intersection, track)
		}
	}

	return intersection
}

func getArtistIntersection(user1_artists, user2_artists []Artist) []Artist{
	intersection := []Artist{}
	user1_set := map[string]bool{}

	for _, artist := range user1_artists{
		user1_set[artist.Id] = false
	}

	for _, artist := range user2_artists{
		if added, ok := user1_set[artist.Id]; ok && !added{
			user1_set[artist.Id] = true
			intersection = append(intersection, artist)
		}
	}

	return intersection
}

func getLibraryIntersections(user_1_songs, user_2_songs []Song) ([]Song, []Artist) {
	song_intersection := []Song{}
	artists := []Artist{}

	user_1_song_set := map[string]bool{}
	artist_intersection := map[string]*Artist{}


	for _, song := range user_1_songs{
		user_1_song_set[song.Track.Id] = true
	}

	for _, song := range user_2_songs{
		if _, ok := user_1_song_set[song.Track.Id]; ok {
			song_intersection = append(song_intersection, song)

			for _, artist := range song.Track.Artists {
				if _, ok := artist_intersection[artist.Name]; ok {
					artist_intersection[artist.Name].SongCount += 1
				} else {
					artist_intersection[artist.Name] = &Artist{
						Name: artist.Name,
						Id: artist.Id,
						SongCount: 1,
					}
				}
			}
		}
	}

	for _, artist := range artist_intersection{
		artists = append(artists, *artist)
	}

	return song_intersection, artists
}

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

	var user1_top_tracks, user2_top_tracks [][]Track
	var user1_top_artists, user2_top_artists [][]Artist

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
	err = json.Unmarshal([]byte(user1.TopSongs), &user1_top_tracks)
	if err != nil {
	    w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	err = json.Unmarshal([]byte(user2.TopSongs), &user2_top_tracks)
	if err != nil {
	    w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	err = json.Unmarshal([]byte(user1.TopArtists), &user1_top_artists)
	if err != nil {
	    w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	err = json.Unmarshal([]byte(user2.TopArtists), &user2_top_artists)
	if err != nil {
	    w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	songs, artists := getLibraryIntersections(user1_songs, user2_songs)

	top_tracks := getTrackIntersection(
		append(append(user1_top_tracks[0], user1_top_tracks[1]...), user1_top_tracks[2]...),
		append(append(user2_top_tracks[0], user2_top_tracks[1]...), user2_top_tracks[2]...),
	)

	top_artists := getArtistIntersection(
		append(append(user1_top_artists[0], user1_top_artists[1]...), user1_top_artists[2]...),
		append(append(user2_top_artists[0], user2_top_artists[1]...), user2_top_artists[2]...),
	)

	sort.Slice(songs, func(i, j int) bool {
		return songs[i].Track.Artists[0].Name < songs[j].Track.Artists[0].Name
	})
	sort.Slice(artists, func(i, j int) bool {
		return artists[i].SongCount > artists[j].SongCount
	})
	sort.Slice(top_tracks, func(i, j int) bool {
		return top_tracks[i].Artists[0].Name < top_tracks[j].Artists[0].Name
	})
	sort.Slice(top_artists, func(i, j int) bool {
		return top_artists[i].Name < top_artists[j].Name
	})

	response := map[string]interface{}{}
	response["songs"] = songs
	response["artists"] = artists
	response["top_tracks"] = top_tracks
	response["top_artists"] = top_artists
	response["user1"] = map[string]interface{}{
		"name": user1.DisplayName,
		"id": user1.ID,
		"songcount": len(user1_songs),
	}
	response["user2"] = map[string]interface{}{
		"name": user2.DisplayName,
		"id": user2.ID,
		"songcount": len(user2_songs),
	}

	json.NewEncoder(w).Encode(response)
}