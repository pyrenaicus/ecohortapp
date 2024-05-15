package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type PreUrl struct {
	Url    string `json:"datos"` //Definim el camp url i la vinculació amb l'objecte json
	Client *http.Client
}

func main() {
	url, _ := GetPreUrl()
	fmt.Println(url)
}

func GetPreUrl() (string, error) {
	url := "https://opendata.aemet.es/opendata/api/prediccion/especifica/municipio/diaria/08001/?api_key=eyJhbGciOiJIUzI1NiJ9.eyJzdWIiOiJwYXVndWVycmVyb0BnbWFpbC5jb20iLCJqdGkiOiI1ZGViMWYxMC04ODQyLTRlZWMtYmNmZS03YmFhNGQ4ZmQ1YWMiLCJpc3MiOiJBRU1FVCIsImlhdCI6MTcxNTc5MDU1MiwidXNlcklkIjoiNWRlYjFmMTAtODg0Mi00ZWVjLWJjZmUtN2JhYTRkOGZkNWFjIiwicm9sZSI6IiJ9.k_lImURCiLP4l6a4PqgqILNu2KClIP58cjaSPz2LFPI"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("cache-control", "no-cache")

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		log.Println("La conexió amb aemet ha tingut l'error", err)
		return "", err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println("Error recuperant el json", err)
		return "", err
	}

	preUrl := PreUrl{}
	err = json.Unmarshal(body, &preUrl)
	if err != nil {
		log.Println("Error decodificant el Json", err)
		return "", err
	}

	return preUrl.Url, err

}
