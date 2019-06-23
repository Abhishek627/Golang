package httproute

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/spf13/viper"
)

// ServerHTTP will serve GET request on port 3000
func ServerHTTP() {
	LoadConfig()
	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello there"))
	})

	port := viper.GetString("PORT")
	http.ListenAndServe(port, r)
	log.Println("Started listening on port ", port)
}

// LoadConfig uses viper to initialize config
func LoadConfig() {
	viper.SetConfigFile("httproute/config.json")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Can't read the config file")
	}
	log.Printf("Port to listen is %s", viper.Get("PORT"))

}
