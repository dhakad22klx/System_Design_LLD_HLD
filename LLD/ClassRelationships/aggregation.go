//If a class contains other classes for logical grouping only without lifecycle ownership, it is an aggregation.
package main

import (
	"fmt"
	"slices"
)

// Artist
type Artist struct {
	name string //private to this package
}

func NewArtist(name string) *Artist {
	return &Artist{name: name}
}

func (a *Artist) GetName() string {
	return a.name
}

//Song

type Song struct {
	title    string
	artist   *Artist
	duration int
}

func NewSong(title string, artist *Artist, duration int) *Song {
	return &Song{
		title:    title,
		artist:   artist,
		duration: duration,
	}
}

func (s *Song) GetTitle() string {
	return s.title
}

func (s *Song) GetArtist() *Artist {
	return s.artist
}

func (s *Song) GetDuration() int {
	return s.duration
}

//String() is a special method in Go that implements the fmt.Stringer interface,
//allowing custom string representation when printing structs.

func (s *Song) String() string {
	return fmt.Sprintf("%s by %s (%ds)", s.title, s.artist.GetName(), s.duration)
}

// Library

type Library struct {
	songs []*Song
}

func (l *Library) GetSongsCount() int {
	return len(l.songs)
}

func (l *Library) AddSong(s *Song) {
	if !slices.Contains(l.songs, s) {
		l.songs = append(l.songs, s)
	}
}

func (l *Library) GetSongs() []*Song {
	return l.songs
}

// Playlist
type Playlist struct {
	name  string
	songs []*Song
}

func NewPlaylist(name string) *Playlist {
	return &Playlist{name: name}
}

func (p *Playlist) AddSong(song *Song) {
	if !slices.Contains(p.songs, song) {
		p.songs = append(p.songs, song)
	}
}

func (p *Playlist) RemoveSong(song *Song) {
	filtered := p.songs[:0]
	for _, s := range p.songs {
		if s != song {
			filtered = append(filtered, s)
		}
	}
	p.songs = filtered
}

func (p *Playlist) GetSongsCount() int {
	return len(p.songs)
}

func (p *Playlist) GetTotalDuration() int {
	total := 0
	for _, s := range p.songs {
		total += s.GetDuration()
	}
	return total
}

func (p *Playlist) GetName() string {
	return p.name
}

func (p *Playlist) GetSongs() []*Song {
	return p.songs
}

// User : Owns playlists and can create or delete them
type User struct {
	name      string
	playlists []*Playlist
}

func NewUser(name string) *User {
	return &User{name: name}
}

func (u *User) CreatePlaylist(playlistName string) *Playlist {
	playlist := NewPlaylist(playlistName)
	u.playlists = append(u.playlists, playlist)
	return playlist
}

func (u *User) DeletePlaylist(playlist *Playlist) {
	filtered := u.playlists[:0]
	for _, p := range u.playlists {
		if p != playlist {
			filtered = append(filtered, p)
		}
	}
	u.playlists = filtered
}

func (u *User) GetName() string {
	return u.name
}
func (u *User) GetPlaylists() []*Playlist {
	return u.playlists
}

func testAggregation() {
	// Artists
	coldplay := NewArtist("Coldplay")
	adele := NewArtist("Adele")

	// Songs
	yellow := NewSong("Yellow", coldplay, 269)
	clocks := NewSong("Clocks", coldplay, 307)
	hello := NewSong("Hello", adele, 295)
	someone := NewSong("Someone Like You", adele, 285)

	// Library
	var library Library
	library.AddSong(yellow)
	library.AddSong(clocks)
	library.AddSong(hello)
	library.AddSong(someone)

	// User & playlists
	alice := NewUser("Alice")
	workout := alice.CreatePlaylist("Workout Mix")
	chill := alice.CreatePlaylist("Chill Vibes")

	workout.AddSong(yellow)
	workout.AddSong(clocks)
	workout.AddSong(hello)

	chill.AddSong(hello)
	chill.AddSong(someone)

	// Output
	fmt.Printf("Library has %d songs\n\n", library.GetSongsCount())

	fmt.Printf("%s (%d songs, %ds):\n",
		workout.GetName(), workout.GetSongsCount(), workout.GetTotalDuration())
	for _, s := range workout.GetSongs() {
		fmt.Printf("  - %s\n", s)
	}

	fmt.Println()

	fmt.Printf("%s (%d songs, %ds):\n",
		chill.GetName(), chill.GetSongsCount(), chill.GetTotalDuration())
	for _, s := range chill.GetSongs() {
		fmt.Printf("  - %s\n", s)
	}

	fmt.Println()

	workoutName := workout.GetName()
	alice.DeletePlaylist(workout)

	fmt.Printf("After deleting '%s':\n", workoutName)
	fmt.Printf("  Library still has %d songs\n", library.GetSongsCount())
	fmt.Printf("  '%s' still has %d songs\n", chill.GetName(), chill.GetSongsCount())
	fmt.Printf("  'Yellow' still exists: %s\n", yellow.GetTitle())
}
