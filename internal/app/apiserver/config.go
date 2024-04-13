package apiserver

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/spf13/viper"
)

type AppConfig struct {
	BindAddr string
	LogLevel string
}

// Define the interface QuizProps
type QuizProps struct {
	Texts []struct {
		Text      string `json:"text"`
		Questions []struct {
			Question string   `json:"question"`
			Options  []string `json:"options"`
			Correct  int      `json:"correct"`
		} `json:"questions"`
	} `json:"texts"`
}

var config AppConfig

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("configs/")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %s\n", err)
	}

	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalf("Error unmarshaling config: %s\n", err)
	}
}

func GetConfig() AppConfig {
	return config
}

// Handler function for /test endpoint
func TestHandler(w http.ResponseWriter, r *http.Request) {
	// Define the sample data conforming to QuizProps
	quizData := QuizProps{
		Texts: []struct {
			Text      string `json:"text"`
			Questions []struct {
				Question string   `json:"question"`
				Options  []string `json:"options"`
				Correct  int      `json:"correct"`
			} `json:"questions"`
		}{
			{
				Text: "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.",
				Questions: []struct {
					Question string   `json:"question"`
					Options  []string `json:"options"`
					Correct  int      `json:"correct"`
				}{
					{
						Question: "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur.",
						Options: []string{
							"Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.",
							"Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.",
							"Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur.",
							"Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.",
						},
						Correct: 0,
					},
					{
						Question: "Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.",
						Options: []string{
							"Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.",
							"Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.",
							"Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur.",
							"Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.",
						},
						Correct: 1,
					},
				},
			},
		},
	}

	// Convert quizData to JSON
	jsonData, err := json.Marshal(quizData)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Set Content-Type header
	w.Header().Set("Content-Type", "application/json")

	// Write JSON response
	w.Write(jsonData)
}
