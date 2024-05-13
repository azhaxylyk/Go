package pkg

import (
	"encoding/json"
	"net/http"
)

var (
	bands     []Band
	relations Relations
)

type Band struct {
	ID           int                 `json:"id"`
	Image        string              `json:"image"`
	Name         string              `json:"name"`
	Members      []string            `json:"members"`
	CreationDate int                 `json:"creationDate"`
	FirstAlbum   string              `json:"firstAlbum"`
	Locations    map[string][]string `json:"-"`
	ConcertDates string              `json:"concertDates"`
	Relations    string              `json:"relations"`
}

type Relations struct {
	Index []struct {
		ID             int                 `json:"id"`
		DatesLocations map[string][]string `json:"datesLocations"`
	} `json:"index"`
}

func GetBandInfo(artistAPI, locationsAPI string) ([]Band, Relations, error) {
	respArtist, err := http.Get(artistAPI)
	if err != nil {
		return nil, Relations{}, err
	}
	defer respArtist.Body.Close()
	err = json.NewDecoder(respArtist.Body).Decode(&bands)
	if err != nil {
		return nil, Relations{}, err
	}
	respRelations, err := http.Get(locationsAPI)
	if err != nil {
		return nil, Relations{}, err
	}
	defer respRelations.Body.Close()
	err = json.NewDecoder(respRelations.Body).Decode(&relations)
	if err != nil {
		return nil, Relations{}, err
	}
	return bands, relations, nil
}
