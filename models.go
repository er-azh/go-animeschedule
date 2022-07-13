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
	EpisodeDate             time.Time         `json:"EpisodeDate"`
	DelayedFrom             time.Time         `json:"DelayedFrom"`
	DelayedUntil            time.Time         `json:"DelayedUntil"`
	Romaji                  *string           `json:"Romaji,omitempty"`
	Streams                 map[string]string `json:"Streams"`
	English                 *string           `json:"English,omitempty"`
	Japanese                *string           `json:"Japanese,omitempty"`
	LengthMin               *int              `json:"LengthMin,omitempty"`
	Episodes                *int              `json:"Episodes,omitempty"`
	AirType                 AirType           `json:"AirType"`
	ImageVersionRoute       string            `json:"ImageVersionRoute"`
	Status                  string            `json:"Status"`
	AiringStatus            string            `json:"AiringStatus"`
	Route                   string            `json:"Route"`
	Title                   string            `json:"Title"`
	EpisodeNumber           int               `json:"EpisodeNumber"`
	SubtractedEpisodeNumber int               `json:"SubtractedEpisodeNumber"`
	Chinese                 bool              `json:"Chinese"`
}

// ShowDetail contains details about a show
type ShowDetail struct {
	JpnTime         time.Time         `json:"JpnTime"`
	DelayedFrom     time.Time         `json:"DelayedFrom"`
	DelayedUntil    time.Time         `json:"DelayedUntil"`
	SubDelayedUntil time.Time         `json:"SubDelayedUntil"`
	DubDelayedFrom  time.Time         `json:"DubDelayedFrom"`
	DubDelayedUntil time.Time         `json:"DubDelayedUntil"`
	DubPremier      time.Time         `json:"DubPremier"`
	DubTime         time.Time         `json:"DubTime"`
	SubDelayedFrom  time.Time         `json:"SubDelayedFrom"`
	Premier         time.Time         `json:"Premier"`
	SubTime         time.Time         `json:"SubTime"`
	SubPremier      time.Time         `json:"SubPremier"`
	Websites        map[string]string `json:"Websites"`
	Season          struct {
		Route  string `json:"Route"`
		Season string `json:"Season"`
		Title  string `json:"Title"`
		Year   string `json:"Year"`
	} `json:"Season"`
	Names struct {
		Abbreviation string `json:"Abbreviation"`
		English      string `json:"English"`
		Japanese     string `json:"Japanese"`
		Romaji       string `json:"Romaji"`
	} `json:"Names"`
	Status            string `json:"Status"`
	Month             string `json:"Month"`
	Description       string `json:"Description"`
	Route             string `json:"Route"`
	ImageVersionRoute string `json:"ImageVersionRoute"`
	Title             string `json:"Title"`
	Relations         struct {
		Alternatives []string `json:"Alternatives"`
		Other        []string `json:"Other"`
		Parents      []string `json:"Parents"`
		Prequels     []string `json:"Prequels"`
		Sequels      []string `json:"Sequels"`
		SideStories  []string `json:"SideStories"`
		Spinoffs     []string `json:"Spinoffs"`
	} `json:"Relations"`
	Sources    []Keyword `json:"Sources"`
	Studios    []Keyword `json:"Studios"`
	Genres     []Keyword `json:"Genres"`
	MediaTypes []Keyword `json:"MediaTypes"`
	Year       int       `json:"Year"`
	LengthMin  int       `json:"LengthMin"`
	Episodes   int       `json:"Episodes"`
	Days       Weekdays  `json:"Days"`
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
