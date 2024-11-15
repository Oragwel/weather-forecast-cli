package model

type Weather struct {
    Location struct {
        Name    string `json:"name"`
        Country string `json:"country"`
    } `json:"location"`
    
    Current struct {
        TempC     float64 `json:"temp_c"`
        Condition struct {
            Text string `json:"text"`
        } `json:"condition"`
    } `json:"current"`

    Forecast struct {
        ForecastDay []struct {
            Hour []struct {
                TimeEpoch    int64   `json:"time_epoch"`
                TempC        float64 `json:"temp_c"`
                Condition    struct {
                    Text string `json:"text"`
                } `json:"condition"`
                ChanceOfRain float64 `json:"chance_of_rain"`
            } `json:"hour"`
        } `json:"forecastday"`
    } `json:"forecast"`
}
