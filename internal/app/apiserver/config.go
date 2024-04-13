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
		ID        int `json:"id"`
		Text      string `json:"text"`
		Questions []struct {
			ID int `json:"id"`
			Question string   `json:"question"`
			Options  []string `json:"options"`
			Correct  int      `json:"correct"`
		} `json:"questions"`
	} `json:"texts"`
}

// Sample Data 
var quizData QuizProps

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

func LoadData() {
	quizData = QuizProps{
		Texts: []struct {
			ID        int `json:"id"`
			Text      string `json:"text"`
			Questions []struct {
				ID       int      `json:"id"`
				Question string   `json:"question"`
				Options  []string `json:"options"`
				Correct  int      `json:"correct"`
			} `json:"questions"`
		}{
			{
				ID: 1,
				Text: "Қайырлы, Тарбағатай! Сен тау болып жаралғалсаң, құтты қоныс қойнауың мал мен жанға талай толып, талай солды ғой. Сенің еңбек оқиғалары шыққан сөздеріңде, адам өміріне көрсетіп берген нәтижелерді сен қалай білесің? Біз білмейміз деп ойлайсың, қай бола аламыз? Қазақ дауысындағы көңілді қара түнек Қаратау алдында тарих қойылды, қарым-қатынасымен шамырады тау - сен оларды шығарасың, Тарбағатай! Бір халқтың екі түрлі кездесуін көрсетіп алып, онда жерге түсіріп жатырсың, қылықтығым. Бақыт пен сортың арасын салықпен бөліп, туған еліңнің шығысындағы бір жетекшілік болып, осы түрде менің жүріп алған жолдарымды түзу қайырымсыз емес, қайырлы алтын адам!",
				Questions: []struct {
					ID       int      `json:"id"`
					Question string   `json:"question"`
					Options  []string `json:"options"`
					Correct  int      `json:"correct"`
				}{
					{
						ID: 1,
						Question: "Қандай сипаттама Қаратауға қатысты айтылған?",
						Options: []string{
							"Қазақтың қара шаңырағы.",
							"Қтты қоныс қойнауы.",
							"Туған елдің шығыстағы қамалы.",
							"Бір халықг екі тҥрлі тағдыры бар.",
						},
						Correct: 0,
					},
					{
						ID: 2,
						Question: "Автор Тарбағатай тауын қалай қабылдайды?",
						Options: []string{
							"Табиғаттың ерекше сыйы ретінде.",
							"Жастық шағының куҽсі ретінде.",
							"Манғаздықтан, асқақтықтың символы ретінде.",
							"Балалық шағының куҽсі ретінде.",
						},
						Correct: 1,
					},
				},
			},
		},
	}
}
