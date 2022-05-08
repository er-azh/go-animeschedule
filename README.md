<h1 align="center">go-animeschedule</h1>

<h2>
⚠️ BOTH THE API AND THIS WRAPPER ARE EXPERIMENTAL. DO NOT USE THEM IN PRODUCTION. 
</h2>

<h3>How to install:</h3>
    
    go get github.com/er-azh/go-animeschedule

<h3>Usage example:</h3>

```go
package main

import (
	"fmt"
	"github.com/er-azh/go-animeschedule"
)

func main() {
    timetable, err := animeschedule.GetCurrentTimetable(nil, nil)
    if err != nil {
        panic(err)
    }
    
    fmt.Println(timetable)
}
```

<h3>Methods:</h3>
There are more endpoints but I don't know them. If you know one that isn't listed here open an issue or submit a PR.

```go
// GetCurrentTimetable gets the current timetable or the timetable in a specific week
// when timetableTime is specified it tries to get that week's timetable.
// timezone can be set to return times in a different time zone.
func GetCurrentTimetable(timetableTime *TimetableTime,timezone *time.Location) (*Timetable, error)

// GetShow gets a show by its slug
func GetShow(slug string) (*ShowDetail, error) 
```
