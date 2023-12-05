package pkg

var data Data

type Coor struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}

type Data struct {
	List []struct {
		DataTime int64 `json:"dt"`
		Main     struct {
			Temp        float64 `json:"temp"`
			FeelsLike   float64 `json:"feels_like"`
			TempMin     float64 `json:"temp_min"`
			TempMax     float64 `json:"temp_max"`
			Pressure    float64 `json:"pressure"`
			SeaLevel    float64 `json:"sea_level"`
			GroundLevel float64 `json:"grnd_level"`
			Humidity    int     `json:"humidity"`
		} `json:"main"`
		Weather []struct {
			Id          int    `json:"id"`
			Main        string `json:"main"`
			Description string `json:"description"`
		} `json:"weather"`
		Wind struct {
			Speed float64 `json:"speed"`
			Gust  float64 `json:"gust"`
			Deg   float64 `json:"deg"`
		} `json:"wind"`
		Visibility int64 `json:"visibility"`
		Pop        int   `json:"pop"`
	} `json:"list"`
	City struct {
		Name    string `json:"name"`
		Country string `json:"country"`
		Sunrise int64  `json:"surise"`
		Sunset  int64  `json:"sunset"`
	}
}
