package weather

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"weatherLineBot/weather/model"
)


type Weather struct {
	Title	string  			`json:"title"`
	model.Forecasts 		`json:"forecasts"`
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
	// fmt.Println(weather)

	result := weather.ToS()
	fmt.Println(result)
	// result := weather.ExampleToS()

	return result, nil
}

func HttpGetBody(url string) ([]byte, error) {
	res, err := http.Get(url)
	if err != nil {
		err = fmt.Errorf("get http error: %s", err)
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

func TempCheck(temp any) string {
	var checkedTemp string
	if temp != nil {
		checkedTemp = fmt.Sprintf("%s", temp)
	} else {
		checkedTemp = "不明"
	}
	return checkedTemp
}

func (w *Weather) ToS() string {
	var minTemp string
	var maxTemp string

	area := fmt.Sprintf("%sです。\n", w.Title)
	body := fmt.Sprintf("今日の大阪の天気：%s\n", w.Forecasts[0].Telop)
	
	minTemp = TempCheck(w.Forecasts[0].Temperature.Min.Celsius)
	min_cel := fmt.Sprintf("今日の最低気温：%s\n", minTemp)

	maxTemp = TempCheck(w.Forecasts[0].Temperature.Max.Celsius)
	max_cel := fmt.Sprintf("今日の最高気温：%s\n", maxTemp)

	T06_12 := fmt.Sprintf("今日の朝から昼の降水確率：%s\n", w.Forecasts[0].ChanceOfRain.T0612)
	T12_18 := fmt.Sprintf("今日の昼から夕方の降水確率：%s\n", w.Forecasts[0].ChanceOfRain.T1218)

	image := fmt.Sprintf("%s\n", w.Forecasts[0].Image.URL)
	result := area + body + max_cel + min_cel + T06_12 + T12_18 + image
	return result
}