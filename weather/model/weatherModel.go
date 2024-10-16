package model

type Temperature struct {
	Min struct {
		Celsius	any `json:"celsius"`
	} `json:"min"`
	Max struct {
		Celsius any `json:"celsius"`
	} `json:"max"`
}

type ChanceOfRain struct {
	T0612 string 	`json:"T06_12"`
	T1218 string 	`json:"T12_18"`
}

type Image struct {
	Title				 string `json:"title"`
	Url			    	 string `json:"previewImageUrl"`
	Width  				 int `json:"width"`
	Height 				 int `json:"height"`
}

type Forecasts []struct {
	Telop	string 	`json:"telop"`
	Temperature 	`json:"temperature"`
	ChanceOfRain 	`json:"chanceOfRain"`
	Image					`json:"image"`
}