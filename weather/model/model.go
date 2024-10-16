package model

import "time"

type AutoGenerated struct {
	PublicTime          time.Time `json:"publicTime"`
	PublicTimeFormatted string    `json:"publicTimeFormatted"`
	PublishingOffice    string    `json:"publishingOffice"`
	Title               string    `json:"title"`
	Link                string    `json:"link"`
	Description         struct {
		PublicTime          time.Time `json:"publicTime"`
		PublicTimeFormatted string    `json:"publicTimeFormatted"`
		HeadlineText        string    `json:"headlineText"`
		BodyText            string    `json:"bodyText"`
		Text                string    `json:"text"`
	} `json:"description"`
	Forecasts []struct {
		Date      string `json:"date"`
		DateLabel string `json:"dateLabel"`
		Telop     string `json:"telop"`
		Detail    struct {
			Weather string `json:"weather"`
			Wind    string `json:"wind"`
			Wave    string `json:"wave"`
		} `json:"detail"`
		Temperature struct {
			Min struct {
				Celsius    any `json:"celsius"`
				Fahrenheit any `json:"fahrenheit"`
			} `json:"min"`
			Max struct {
				Celsius    string `json:"celsius"`
				Fahrenheit string `json:"fahrenheit"`
			} `json:"max"`
		} `json:"temperature"`
		ChanceOfRain struct {
			T0006 string `json:"T00_06"`
			T0612 string `json:"T06_12"`
			T1218 string `json:"T12_18"`
			T1824 string `json:"T18_24"`
		} `json:"chanceOfRain"`
		Image struct {
			Title  string `json:"title"`
			URL    string `json:"url"`
			Width  int    `json:"width"`
			Height int    `json:"height"`
		} `json:"image"`
	} `json:"forecasts"`
	Location struct {
		Area       string `json:"area"`
		Prefecture string `json:"prefecture"`
		District   string `json:"district"`
		City       string `json:"city"`
	} `json:"location"`
	Copyright struct {
		Title string `json:"title"`
		Link  string `json:"link"`
		Image struct {
			Title  string `json:"title"`
			Link   string `json:"link"`
			URL    string `json:"url"`
			Width  int    `json:"width"`
			Height int    `json:"height"`
		} `json:"image"`
		Provider []struct {
			Link string `json:"link"`
			Name string `json:"name"`
			Note string `json:"note"`
		} `json:"provider"`
	} `json:"copyright"`
}