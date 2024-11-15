
# Weather Forecast CLI Application

This Weather Forecast CLI is a simple command-line tool that fetches and displays current weather and hourly forecast information for a given location. The app provides temperature, rain probability, and weather conditions for each hour of the day, with color-coded output to indicate favorable or unfavorable conditions.

## Features
- **Current Weather**: Shows the current temperature, condition, and location.
- **Hourly Forecast**: Provides a breakdown of hourly weather conditions, including temperature and chance of rain.
- **Color Coding**: Displays green for favorable conditions and red for high rain chances.

## Installation

To use this application, you need to have:
- **Go** installed (version 1.16 or later)
- Access to the `weatherapi.com` API with a valid API key.

### Clone the repository
```bash
git clone https://github.com/Oragwel/weather-forecast-cli.git
cd weather-forecast-cli
```

### Dependencies
Install the required dependencies by running:
```bash
go mod tidy
```

### Model Setup
The application uses a `Weather` struct defined in `model/model.go`. Ensure you have the correct structure to match the API response.

## Project Structure

```
weather-forecast-cli/
│
├── main.go                  # Main application code
├── model/
│   └── model.go             # Contains the Weather struct to parse the JSON response
├── main_test.go             # Test cases for the application
├── go.mod                   # Go module file
└── README.md                # Project documentation
```


## Usage

Run the application with a location argument:
```bash
go run main.go "City, Country"
```
If no location is specified, it defaults to **Kenya**.

Example:
```bash
go run main.go "Paris, France"
```

## Configuration
Replace the API key in the URL with your own from `weatherapi.com`:
```go
http.Get("http://api.weatherapi.com/v1/forecast.json?key=YOUR_API_KEY&q=" + q + "&days=1&aqi=no&alerts=no")
```

## Error Handling

The application will terminate and display an error message if:
- The API is unreachable.
- The API key is invalid or expired.
- The location is incorrect or data is unavailable.

## Sample Output

```
Paris, France: 15°C, Clear
14:00 - 14°C, 30% chance of rain, Clear
15:00 - 16°C, 10% chance of rain, Sunny
16:00 - 17°C, 50% chance of rain, Cloudy
```
*The above output is color-coded where favorable weather is in green and high rain chance is in red.*

## License
This project is open-source and available under the [MIT License](LICENSE).

## Acknowledgements
- This application uses the [WeatherAPI](https://www.weatherapi.com/) service.
- [Fatih Color](https://github.com/fatih/color) is used for color-coded CLI outputs.
