package igdbgo

import (
	"encoding/json"
	"errors"
	"time"
	//"fmt"
	"io/ioutil"
	//"log"
	"net/http"
	//"os"
	"strconv"
	//"time"

	"golang.org/x/net/context"
	"google.golang.org/appengine/urlfetch"
)

// GetGames returns a slice of Game types according to the passed Options argument
func GetGames(ctx context.Context, term string, amount int, sort int, order int, ID string) ([]Game, error) {
	URL, err := SetURL(term, amount, sort, order, ID)
	if err != nil {
		return nil, err
	}
	games, err := fetchData(ctx, URL)
	if err != nil {
		return nil, err
	}
	return games, nil
}

// SetURL will take a search term, an integer for the amount of items to be returned,
// and two more integers that correspond to different sorting options
// sort:  0 = omit, 1 = release dates, 2 = popularity
// order: 0 = omit, 1 = descending, 2 = ascending
func SetURL(term string, amount int, sort int, order int, ID string) (s string, err error) {
	s = rootURL
	if ID != "" {
		s = s + ID
	}
	s = s + "?fields=*&limit" + strconv.Itoa(amount) + "&offset=0"
	switch sort {
	case 0:
		break
	case 1:
		s = s + "&order=release_dates.date%3"
	case 2:
		s = s + "&order=popularity%3"
	case 3:
		s = s + "&order=rating%3"
	default:
		err = errors.New("setOptions: Invalid input in sort parameter")
		return "", err
	}
	switch order {
	case 0:
		break
	case 1:
		s = s + "Adesc"
	case 2:
		s = s + "Aasc"
	default:
		err = errors.New("setOptions: Invalid input in order parameter")
		return "", err
	}
	if term != "" {
		s = s + "&search=" + term
	}
	return s, nil
}

// fetchData returns a raw slice of Game types based on the URL parameter.
// All options and specifications must be in the URL before executing this function.
func fetchData(ctx context.Context, URL string) ([]Game, error) {
	client := urlfetch.Client(ctx)
	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("X-Mashape-Key", key)
	req.Header.Add("Accept", "application/json")

	resp, err := client.Do(req)
	if err != nil { // No response from host
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()

	data := make([]Game, 10, 50)
	if err = json.Unmarshal(body, &data); err != nil {
		return nil, err
	}

	return data, nil
}

//GetTop returns the top 10 rated games with certain filters in place to prevent outliers
func GetTop(ctx context.Context) ([]Game, error) {
	/* Deprecated code when SetOptions was necessary; revert if needed
	opt, err := SetOptions("", 10, 3, 1) //any title, 10 listings, by rating, descending
	if err != nil {
		return nil, err
	}
	URL := setURL(opt)
	URL = URL + "&filter[rating_count][gte]=10"
	*/
	URL, err := SetURL("", 10, 3, 1, "") //any title, 10 listings, by rating, descending
	if err != nil {
		return nil, err
	}
	URL = URL + "&filter[rating_count][gte]=10"

	games, err := fetchData(ctx, URL)
	if err != nil {
		return nil, err
	}
	return games, nil
}

//GetPop returns the 10 most popular games
func GetPop(ctx context.Context) ([]Game, error) {
	URL, err := SetURL("", 10, 2, 1, "")
	if err != nil {
		return nil, err
	}
	URL = URL + "&filter[rating_count][gte]=10"

	games, err := fetchData(ctx, URL)
	if err != nil {
		return nil, err
	}
	return games, nil
}

//GetUpcoming returns the 10 latest soon to be released games
func GetUpcoming(ctx context.Context) ([]Game, error) {
	/* Deprecated code when SetOptions was necessary; revert if needed
	opt, err := SetOptions("", 10, 1, 1) //any title, 10 listings, by release, descending
	if err != nil {
		return nil, err
	}
	URL := setURL(opt)
	*/
	URL, err := SetURL("", 10, 1, 2, "") //any title, 10 listings, by release, descending
	if err != nil {
		return nil, err
	}

	nowInt := timeToMilli(time.Now())
	nowStr := strconv.FormatInt(nowInt, 10)
	URL = URL + "&filter[first_release_date][gte]=" + nowStr

	/*
		t := time.Now()
		year, month, day := t.Date()

		URL = URL + "&filter[release_dates.date][lt]=" + string(year) + "-" + string(month) + "-" + string(day)
	*/
	games, err := fetchData(ctx, URL)
	if err != nil {
		return nil, err
	}
	return games, nil
}

/* OLD VERSION
// SetURL will take a search term, an integer for the amount of items to be returned,
// and two more integers that correspond to different sorting options
// sort:  0 = omit, 1 = release dates, 2 = popularity
// order: 0 = omit, 1 = descending, 2 = ascending
func SetURL(term string, amount int, sort int, order int, ID int) (s string, err error) {
	s = rootURL + "&limit" + strconv.Itoa(amount) + "&offset=0"
	switch sort {
	case 0:
		break
	case 1:
		s = s + "&order=release_dates.date%3"
	case 2:
		s = s + "&order=popularity%3"
	case 3:
		s = s + "&order=rating%3"
	default:
		err = errors.New("setOptions: Invalid input in sort parameter")
		return "", err
	}
	switch order {
	case 0:
		break
	case 1:
		s = s + "Adesc"
	case 2:
		s = s + "Aasc"
	default:
		err = errors.New("setOptions: Invalid input in order parameter")
		return "", err
	}
	if term != "" {
		s = s + "&search=" + term
	}
	if ID != 0 {
		s = s + "&filter[id][eq]=" + strconv.Itoa(ID)
	}
	return s, nil
}
*/

/*
//SetOptions is now deprecated. Revert if necessary.

// SetOptions is a helper function that will take a search term, an integer for the amount of items to be returned,
// and two more integers that correspond to different sorting options
// amount: limited from 0 to 50
// sort: 1 = release dates, 2 = popularity
// order: 1 = descending, 2 = ascending
func SetOptions(term string, amount int, sort int, order int) (o Options, err error) {
	o.Term = term
	o.Amount = strconv.Itoa(amount)
	switch sort {
	case 0:
		o.Sort = ""
	case 1:
		o.Sort = "release_dates.date"
	case 2:
		o.Sort = "popularity"
	case 3:
		o.Sort = "rating"
	default:
		err = errors.New("setOptions: Invalid input in sort parameter")
		return
	}
	switch order {
	case 0:
		o.Order = ""
	case 1:
		o.Order = "Adesc"
	case 2:
		o.Order = "Aasc"
	default:
		err = errors.New("setOptions: Invalid input in order parameter")
		return
	}
	return
}
*/

/*
//Part of setOptions deprecation. Revert if setOptions is reverted
func setURL(o Options) string {
	if o.Term == "" {
		return rootURL + "&limit=" + o.Amount + "&offset=0&order=" + o.Sort + "%3" + o.Order
	}
	return rootURL + "&limit=" + o.Amount + "&offset=0&order=" + o.Sort + "%3" + o.Order + "&search=" + o.Term

	// *&limit=5&offset=0&order=release_dates.date%3Adesc&search=
	// *&limit=10&offset=0&order=popularity%3Adesc&search=
}
*/

/* OLD VERSION
// GetGames returns a slice of Game types according to the passed Options argument
func GetGames(ctx context.Context, term string, amount int, sort int, order int, ID int) ([]Game, error) {
	URL, err := SetURL(term, amount, sort, order, ID)
	if err != nil {
		return nil, err
	}
	games, err := fetchData(ctx, URL)
	if err != nil {
		return nil, err
	}
	return games, nil
}
*/

/*
// GetGame returns a Game type according to its ID
func GetGame(ctx context.Context, ID int) (Game, error) {
	URL := SetURL()
}
*/
