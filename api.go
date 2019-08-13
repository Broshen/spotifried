package main

import (
	"errors"
	"fmt"
	"sort"
	"strconv"
	"time"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"github.com/gorilla/mux"
)

var client = &http.Client{}

func getAPI(access_token, url string) (*http.Response, error){
	r, _ := http.NewRequest("GET", url, nil)
	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("Accept", "application/json")
	r.Header.Add("Authorization", fmt.Sprintf("Bearer %s", access_token))

	resp, err := client.Do(r)

	if resp.StatusCode != 200{
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		bodyString := string(bodyBytes)
		return resp, errors.New(fmt.Sprintf("Error while accessing spotify API at %s: %s", url, bodyString))
	} else if err != nil{
		return resp, errors.New(fmt.Sprintf("Error while accessing spotify API at %s: %s", url, err))
	} else{
		return resp, nil
	}
}


func getUser(access_token, refresh_token string) User{

	resp, err := getAPI(access_token, "https://api.spotify.com/v1/me")
	if err != nil{
		panic(err)
	}

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

func getAllUserSongs(access_token, refresh_token string) []Song{

	next := "https://api.spotify.com/v1/me/tracks?offset=0&limit=50"
	var songs []Song

	for next != "" {

		resp, err := getAPI(access_token, next) 
		if resp.StatusCode == 429{ // rate limited
			fmt.Println("Request to get user songs was rate limited, retry after " + resp.Header["Retry-After"][0] + " seconds")
			waittime, _ := strconv.Atoi(resp.Header["Retry-After"][0])
			time.Sleep(time.Duration(waittime) * time.Second)
			continue
		} else if err != nil {
			panic(err)
		}

		var respVal TrackResponse
		err = json.NewDecoder(resp.Body).Decode(&respVal)
		if err != nil {
			panic(fmt.Sprintf("%s", err))
		}
		next = respVal.Next
		songs = append(songs, respVal.Items...)
	}

	return songs
}

func getArtists(access_token, artist_ids string) ([]Artist, error) {
	// wrap request sending in a for loop in case of rate limiting
	for {
		resp, err := getAPI(access_token, "https://api.spotify.com/v1/artists?ids="+artist_ids) 

		if resp.StatusCode == 429{ // rate limited
			fmt.Println("Request to get artists was rate limited, retry after " + resp.Header["Retry-After"][0] + " seconds")
			waittime, _ := strconv.Atoi(resp.Header["Retry-After"][0])
			time.Sleep(time.Duration(waittime) * time.Second)
			continue
		} else if err != nil{
			return nil, err
		}

		var respVal ArtistResponse
		err = json.NewDecoder(resp.Body).Decode(&respVal)
		if err != nil {
			return nil, err
		}

		return respVal.Artists, nil
	}
}

func getAllUserArtistsAndGenres(access_token, refresh_token string, songs []Song) ([]Artist, []Genre) {
	// first get all the artist ids from the user's list of songs
	artistsSongCount := map[string]int{}

	for _, song := range songs{
		for _, artist := range song.Track.Artists{
			if _, ok := artistsSongCount[artist.Id]; ok{
				artistsSongCount[artist.Id] += 1
			} else{
				artistsSongCount[artist.Id] = 1
			}
		}
	}

	// iterate over all artist id's in batches of 50, and get artist genres from the spotify API
	i := 0
	ids := ""
	artists := []Artist{}

	for id := range artistsSongCount{
		// add the next id to the list
		if ids == ""{
			ids += id
		} else{
			ids += "," + id
		}
		i += 1

		// if we have 50, reset the list, send the request, and process the response
		if i == 50{

			a, err := getArtists(access_token, ids)
			if err != nil {
				panic(err)
			}
			for i, _ := range a{
				a[i].SongCount = artistsSongCount[a[i].Id]
			}
			artists = append(artists, a...)

			i = 0
			ids = ""
		}
	}
	// process any leftovers
	if i != 0 {
		a, err := getArtists(access_token, ids)
		if err != nil {
			panic(err)
		}
		for i, _ := range a{
			a[i].SongCount = artistsSongCount[a[i].Id]
		}
		artists = append(artists, a...)
	}


	// count the number of songs/genre and artists/genre
	genresMap := map[string]*Genre{}
	genres := []Genre{}

	for _, artist := range artists{
		for _, genre := range artist.Genres {
			if _, ok := genresMap[genre]; ok{
				genresMap[genre].ArtistCount += 1
				genresMap[genre].SongCount += artistsSongCount[artist.Id]
			} else{
				genresMap[genre] = &Genre{
					Name: genre,
					ArtistCount: 1,
					SongCount: artistsSongCount[artist.Id],
				}
			}
		}
	}
	//convert map into a list
	for  _, genre := range genresMap {
	   genres = append(genres, *genre)
	}


	// sort results
	sort.Slice(artists, func(i, j int) bool{
		return artists[i].SongCount > artists[j].SongCount
	})
	sort.Slice(genres, func(i, j int) bool{
		if genres[i].SongCount == genres[j].SongCount{
			return genres[i].ArtistCount > genres[j].ArtistCount
		}
		return genres[i].SongCount > genres[j].SongCount
	})
	return artists, genres
}

func fetchHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	user_id := vars["user_id"]
	user := getUserById(user_id)
	defer db.Save(&user)

	redirect_uri := "http://" + r.Host+"/authenticated"
	user.AccessToken, user.RefreshToken, _ = getTokens(user.RefreshToken, redirect_uri, true)
	user_songs := getAllUserSongs(user.AccessToken, user.RefreshToken)
	artists, genres := getAllUserArtistsAndGenres(user.AccessToken, user.RefreshToken, user_songs)

	songsByteArr, err := json.Marshal(user_songs)
	if err != nil {
	    w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	artistsByteArr, err := json.Marshal(artists)
	if err != nil {
	    w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	genresByteArr, err := json.Marshal(genres)
	if err != nil {
	    w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	user.Songs = string(songsByteArr)
	user.Artists = string(artistsByteArr)
	user.Genres = string(genresByteArr)
	user.LastRefreshed = string(time.Now().Format("01-02-2006 15:04:05"))

	response := map[string]interface{}{}
	response["songs"] = user_songs
	response["artists"] = artists
	response["genres"] = genres
	json.NewEncoder(w).Encode(response)
}
