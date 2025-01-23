package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

// ListLocations -
func (c *Client) ListLocationDetails(area_name string) (RespDetailedLocations, error) {
	url := baseURL + "/location-area/" + area_name

	var locationsResp RespDetailedLocations

	if data, ok := c.cache.Get(url); ok {
		err := json.Unmarshal(data, &locationsResp)
		if err != nil {
			return RespDetailedLocations{}, err
		}
		return locationsResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return RespDetailedLocations{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespDetailedLocations{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespDetailedLocations{}, err
	}

	err = json.Unmarshal(dat, &locationsResp)
	if err != nil {
		return RespDetailedLocations{}, err
	}

	c.cache.Add(url, dat)

	return locationsResp, nil
}
