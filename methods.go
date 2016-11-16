package igdbgo

import (
	"errors"
	"time"
)

// GetDate takes the FirstRelease variable from a Game type and returns the year, month, and day
func (g *Game) GetDate() (year int, month int, day int) {
	milli := int64(g.FirstRelease)
	t := time.Unix(0, milli*int64(time.Millisecond))
	year, m, day := t.Date()
	month = int(m)
	return
}

// CheckFuture returns true if the Game's release is in the future and false otherwise
func (g *Game) CheckFuture() bool {
	y, m, d := g.GetDate()
	if y >= time.Now().Year() && m >= int(time.Now().Month()) && d > time.Now().Day() {
		return true
	}
	return false
}

// GetGenres returns a slice of strings of the game's genres based on its genre IDs
func (g *Game) GetGenres() (s []string) {
	for _, val := range g.Genres {
		switch val {
		case 2:
			s = append(s, "Point-and-Click")
		case 4:
			s = append(s, "Fighting")
		case 5:
			s = append(s, "Shooter")
		case 7:
			s = append(s, "Music")
		case 8:
			s = append(s, "Platform")
		case 9:
			s = append(s, "Puzzle")
		case 10:
			s = append(s, "Racing")
		case 11:
			s = append(s, "Real Time Strategy")
		case 12:
			s = append(s, "RPG")
		case 13:
			s = append(s, "Simulator")
		case 14:
			s = append(s, "Sport")
		case 15:
			s = append(s, "Strategy")
		case 16:
			s = append(s, "Turn-Based-Strategy")
		case 24:
			s = append(s, "Tactical")
		case 25:
			s = append(s, "Hack and Slash")
		case 26:
			s = append(s, "Quiz/Trivia")
		case 30:
			s = append(s, "Pinball")
		case 31:
			s = append(s, "Adventure")
		case 32:
			s = append(s, "Indie")
		case 33:
			s = append(s, "Arcade")
		default:
			s = append(s, "Unknown")
		}
	}
	return
}

// GetImageURL returns the full URL to the cloudinary image of the cover
func (g *Game) GetImageURL() string {
	return imageRoot + "cover_big/" + g.Cover.ID + ".jpg"
}

// GetVideoURL returns the full URL to the trailer hosted on Youtube
func (g *Game) GetVideoURL() (string, error) {
	if g.Videos[0].ID == "" {
		return "", errors.New("Video ID not found.")
	}
	return videoRoot + g.Videos[0].ID, nil
}

// timeToMilli takes Go's time type and returns a millisecond value since Epoch
func timeToMilli(t time.Time) int64 {
	return (t.UnixNano() / (int64(time.Millisecond) / int64(time.Nanosecond)))
}
