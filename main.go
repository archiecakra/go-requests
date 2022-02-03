package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	telebot "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	env "github.com/joho/godotenv"
)

// Sites to check
var site = map[string]string{
	"surabaya.go.id":         "<div class=\"teks-surabaya-2\">Gotong Royong Menuju Surabaya Kota Dunia<br>yang Maju, Humanis, Dan Berkelanjutan</div>",
	"esurat.surabaya.go.id":  "<h4>Log In eSurat</h4>",
	"sswalfa.surabaya.go.id": "<h2 class=\"text-white text-center font-weight-bold\">Ajukan Permohonan Izin Melalui SSW</h2>",
}

// Load .env configuration file
func init() {
	err := env.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

// Send telegram message
func SendMessage(text string) {
	bot, err := telebot.NewBotAPI(os.Getenv("TELEGRAM_APITOKEN"))
	if err != nil {
		panic(err)
	}
	// bot.Debug = true // Uncomment to see debug response
	msg := telebot.NewMessage(-764108168, text)
	bot.Send(msg)
}

// Get website statuses
func getStatus(site map[string]string) string {

	var response = ""

	for key, val := range site {

		res, err := http.Get("https://" + key)
		if err != nil {
			return err.Error()
		}

		defer res.Body.Close()
		body, err := io.ReadAll(res.Body)

		if strings.Contains(string(body), val) {
			response += res.Status + " | " + key + " Key found\n"
		} else {
			response += res.Status + " | " + key + " Key not found\n"
		}

	}

	return response

}

func main() {

	var results string = getStatus(site)

	fmt.Println("Checking Websites...")
	fmt.Println(results)

	fmt.Println("Sending Message...")
	SendMessage(results)

}
