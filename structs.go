package main

type User struct{
  ID        string `gorm:"primary_key"`
  DisplayName   string
  AccessToken   string
  RefreshToken  string
  LastRefreshed string
  Songs         string
  Artists       string
  Genres        string
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

type ArtistResponse struct {
  Artists [] Artist
}

type Artist struct {
  Id string
  Name string
  Genres []string
  SongCount int
}

type Genre struct {
  Name string
  SongCount int
  ArtistCount int
  SubGenres StringSet
}

type StringSet struct {
  set map[string]bool
}

func NewStringSet() StringSet {
  return StringSet{
    set: map[string]bool{},
  }
}

func (s *StringSet) Add(str string){
  if _, ok := s.set[str]; !ok {
    s.set[str] = true
  }
}
