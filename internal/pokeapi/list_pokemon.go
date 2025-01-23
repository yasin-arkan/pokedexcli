package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListPokemon(name string) (Pokemon, error) {

	url := baseURL + "/pokemon/" + name

	var pokemonResp Pokemon

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return Pokemon{}, err
	}

	resp, err := c.httpClient.Do(req)

	if err != nil {
		return Pokemon{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)

	if err != nil {
		return Pokemon{}, err
	}

	err = json.Unmarshal(dat, &pokemonResp)

	if err != nil {
		return Pokemon{}, err
	}

	return pokemonResp, nil

}
