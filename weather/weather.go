package weather

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Weather struct {
	Area		string `json:"targetArea"`
	HeadLine	string `json:"headlineText"`
	Body		string `json:"text"`
}

func GetWeather() (str string, err error) {
	body, err := HttpGetBody("https://weather.tsukumijima.net/api/forecast/city/270000")
	if err != nil {
		return str, err
	}
	weather, err := FormatWeather(body)
	if err != nil {
		return str, err
	}

	result := weather.ToS()

	return result, nil
}

func HttpGetBody(url string) ([]byte, error) {
	res, err := http.Get(url)
	if err != nil {
		err = fmt.Errorf("Get Http Error: %s", err)
		return nil, err
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		err = fmt.Errorf("IO Read Error:: %s", err)
		return nil, err
	}

	defer res.Body.Close()

	return body, nil
}

func FormatWeather(body []byte) (*Weather, error) {
	weather := new(Weather)
	if err := json.Unmarshal(body, weather); err != nil {
		err = fmt.Errorf("JSON Unmarshal error: %s", err)
		return nil, err
	}
	return weather, nil
}

func (w *Weather) ToS() string {
	area := fmt.Sprintf("%sの天気です。\n", w.Area)
	head := fmt.Sprintf("%s\n", w.HeadLine)
	body := fmt.Sprintf("%s\n", w.Body)
	result := area + head + body

	return result
}