package igdbgo

// Game is a struct that holds all relevant information on a particular game
type Game struct {
	ID           int        `json:"id"`
	Name         string     `json:"name"`
	Slug         string     `json:"slug,omitempty"`
	URL          string     `json:"url"`
	CreatedAt    int        `json:"created_at,omitempty"`
	UpdatedAt    int        `json:"updated_at,omitempty"`
	Summary      string     `json:"summary,omitempty"`
	Storyline    string     `json:"storyline,omitempty"`
	Collection   int        `json:"colletion,omitempty"`
	Hypes        int        `json:"hypes,omitempty"`
	Rating       float64    `json:"rating,omitempty"`
	Popularity   float64    `json:"popularity,omitempty"`
	AggRating    float64    `json:"aggregated_rating,omitempty"`
	RatingCount  int        `json:"rating_count,omitempty"`
	Developers   []int      `json:"developers,omitempty"`
	Publishers   []int      `json:"publishers,omitempty"`
	Engines      []int      `json:"game_engines,omitempty"`
	Category     int        `json:"category,omitempty"`
	TimeToBeat   Completion `json:"time_to_beat,omitempty"`
	GameModes    []int      `json:"game_modes,omitempty"`
	Keywords     []int      `json:"keywords,omitempty"`
	Themes       []int      `json:"themes,omitempty"`
	Genres       []int      `json:"genres,omitempty"`
	FirstRelease int        `json:"first_release_date,omitempty"`
	ReleaseDates []Release  `json:"release_dates,omitempty"`
	AltNames     []AltName  `json:"alternative_names,omitempty"`
	Screenshots  []Image    `json:"screenshots,omitempty"`
	Videos       []Video    `json:"videos,omitempty"`
	Cover        Image      `json:"cover,omitempty"`
	ESRB         Rating     `json:"esrb,omitempty"`
	PEGI         Rating     `json:"pegi,omitempty"`
}

// Completion is a struct that holds the time to beat the game under different constraints
type Completion struct {
	Completely int `json:"completely,omitempty"`
	Hastly     int `json:"hastly,omitempty"`
	Normally   int `json:"normally,omitempty"`
}

// Release is a struct that holds relevant release date information for a Game type
type Release struct {
	Category int `json:"category,omitempty"`
	Platform int `json:"platform,omitempty"`
	Date     int `json:"date,omitempty"`
	Region   int `json:"region,omitempty"`
}

// AltName is a struct that holds alternate names for a Game type
type AltName struct {
	Name    string `json:"name"`
	Comment string `json:"comment,omitempty"`
}

// Image is a struct that holds the ID to reach the image along with its dimensions
type Image struct {
	ID     string `json:"cloudinary_id"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

// Video is a struct that holds the name of the referenced video along with its ID.
type Video struct {
	Name string `json:"name,omitempty"`
	ID   string `json:"video_id"`
}

// Rating is a struct that holds age and content rating information from the ESRB or PEGI
type Rating struct {
	Synopsis string `json:"synopsis,omitempty"`
	Rating   int    `json:"rating,omitempty"`
}

/* Deprecated, just like everything else having to do with options yo
// Options is a struct that holds the different criteria for a database query or search
type Options struct {
	Term   string // term to be searched for
	Amount string // number of results to be returned
	Sort   string // 1 = release dates, 2 = popularity
	Order  string // 1 = descending, 2 = ascending
}
*/
