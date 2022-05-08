package animeschedule

import "time"

const (
	WebsiteOfficial    = "Official"
	WebsiteMAL         = "MAL"
	WebsiteAniList     = "AniList"
	WebsiteKitsu       = "Kitsu"
	WebsiteAnimePlanet = "AnimePlanet"
	WebsiteAniDB       = "AniDB"
	WebsiteCrunchyroll = "Crunchyroll"
	WebsiteVRV         = "VRV"
	WebsiteYouTube     = "YouTube"
	WebsiteHiDive      = "HiDive"
)

type AirType string

const (
	AirTypeRaw AirType = "raw"
	AirTypeSub AirType = "sub"
)

// Timetable is a week's timetable
type Timetable []TimetableShow

// TimetableShow is a single entry in a Timetable
type TimetableShow struct {
	AirType                 AirType           `json:"AirType"`
	AiringStatus            string            `json:"AiringStatus"`
	Chinese                 bool              `json:"Chinese"`
	DelayedFrom             time.Time         `json:"DelayedFrom"`
	DelayedUntil            time.Time         `json:"DelayedUntil"`
	English                 *string           `json:"English,omitempty"`
	EpisodeDate             time.Time         `json:"EpisodeDate"`
	EpisodeNumber           int               `json:"EpisodeNumber"`
	Episodes                *int              `json:"Episodes,omitempty"`
	ImageVersionRoute       string            `json:"ImageVersionRoute"`
	Japanese                *string           `json:"Japanese,omitempty"`
	LengthMin               *int              `json:"LengthMin,omitempty"`
	Romaji                  *string           `json:"Romaji,omitempty"`
	Route                   string            `json:"Route"`
	Status                  string            `json:"Status"`
	Streams                 map[string]string `json:"Streams"`
	SubtractedEpisodeNumber int               `json:"SubtractedEpisodeNumber"`
	Title                   string            `json:"Title"`
}

// ShowDetail contains details about a show
type ShowDetail struct {
	Days              Weekdays  `json:"Days"`
	DelayedFrom       time.Time `json:"DelayedFrom"`
	DelayedUntil      time.Time `json:"DelayedUntil"`
	Description       string    `json:"Description"`
	DubDelayedFrom    time.Time `json:"DubDelayedFrom"`
	DubDelayedUntil   time.Time `json:"DubDelayedUntil"`
	DubPremier        time.Time `json:"DubPremier"`
	DubTime           time.Time `json:"DubTime"`
	Episodes          int       `json:"Episodes"`
	Genres            []Keyword `json:"Genres"`
	ImageVersionRoute string    `json:"ImageVersionRoute"`
	JpnTime           time.Time `json:"JpnTime"`
	LengthMin         int       `json:"LengthMin"`
	MediaTypes        []Keyword `json:"MediaTypes"`
	Month             string    `json:"Month"`
	Names             struct {
		Abbreviation string `json:"Abbreviation"`
		English      string `json:"English"`
		Japanese     string `json:"Japanese"`
		Romaji       string `json:"Romaji"`
	} `json:"Names"`
	Premier   time.Time `json:"Premier"`
	Relations struct {
		Alternatives []string `json:"Alternatives"`
		Other        []string `json:"Other"`
		Parents      []string `json:"Parents"`
		Prequels     []string `json:"Prequels"`
		Sequels      []string `json:"Sequels"`
		SideStories  []string `json:"SideStories"`
		Spinoffs     []string `json:"Spinoffs"`
	} `json:"Relations"`
	Route  string `json:"Route"`
	Season struct {
		Route  string `json:"Route"`
		Season string `json:"Season"`
		Title  string `json:"Title"`
		Year   string `json:"Year"`
	} `json:"Season"`
	Sources         []Keyword         `json:"Sources"`
	Status          string            `json:"Status"`
	Studios         []Keyword         `json:"Studios"`
	SubDelayedFrom  time.Time         `json:"SubDelayedFrom"`
	SubDelayedUntil time.Time         `json:"SubDelayedUntil"`
	SubPremier      time.Time         `json:"SubPremier"`
	SubTime         time.Time         `json:"SubTime"`
	Title           string            `json:"Title"`
	Websites        map[string]string `json:"Websites"`
	Year            int               `json:"Year"`
}

type Weekdays struct {
	Sunday    bool `json:"sunday"`
	Monday    bool `json:"monday"`
	Tuesday   bool `json:"tuesday"`
	Wednesday bool `json:"wednesday"`
	Thursday  bool `json:"thursday"`
	Friday    bool `json:"friday"`
	Saturday  bool `json:"saturday"`
}

type Keyword struct {
	Name   string `json:"Name"`
	Route  string `json:"Route"`
	Weight int    `json:"Weight"`
}
