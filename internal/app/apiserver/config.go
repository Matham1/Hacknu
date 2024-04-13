package apiserver

import (
	"log"
	"net/http"
	"encoding/json"
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
				Text: "Sample text",
				Questions: []struct {
					Question string   `json:"question"`
					Options  []string `json:"options"`
					Correct  int      `json:"correct"`
				}{
					{
						Question: "Sample question",
						Options:  []string{"Option 1", "Option 2", "Option 3"},
						Correct:  1, // Index of correct option
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

// Define your routes
func SetupRoutes() {
	http.HandleFunc("/test", TestHandler)
}

// Start the server
func StartServer() {
	SetupRoutes()
	addr := GetConfig().BindAddr
	log.Printf("Starting server on %s", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}