package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
)

const baseURL = "https://api.green-api.com/waInstance"

type ApiResponse struct {
	Data interface{} `json:"data"`
}

type SettingResponce struct {
	Wid                               string `json:"wid"`
	CountryInstance                   string `json:"countryInstance"`
	TypeAccount                       string `json:"typeAccount"`
	WebhookUrl                        string `json:"webhookUrl"`
	WebhookUrlToken                   string `json:"webhookUrlToken"`
	DelaySendMessagesMilliseconds     int    `json:"delaySendMessagesMilliseconds"`
	MarkIncomingMessagesReaded        string `json:"markIncomingMessagesReaded"`
	MarkIncomingMessagesReadedOnReply string `json:"markIncomingMessagesReadedOnReply"`
	SharedSession                     string `json:"sharedSession"`
	ProxyInstance                     string `json:"proxyInstance"`
	OutgoingWebhook                   string `json:"outgoingWebhook"`
	OutgoingMessageWebhook            string `json:"outgoingMessageWebhook"`
	OutgoingAPIMessageWebhook         string `json:"outgoingAPIMessageWebhook"`
	IncomingWebhook                   string `json:"incomingWebhook"`
	DeviceWebhook                     string `json:"deviceWebhook"`
	StatusInstanceWebhook             string `json:"statusInstanceWebhook"`
	StateWebhook                      string `json:"stateWebhook"`
	EnableMessagesHistory             string `json:"enableMessagesHistory"`
	KeepOnlineStatus                  string `json:"keepOnlineStatus"`
	PollMessageWebhook                string `json:"pollMessageWebhook"`
	IncomingBlockWebhook              string `json:"incomingBlockWebhook"`
	IncomingCallWebhook               string `json:"incomingCallWebhook"`
}

type StateInstanceResponce struct {
	StateInstance string `json:"stateInstance"`
}
type SendMessageResponce struct {
	IdMessage string `json:"idMessage"`
}
type SendFileByUrlResponce struct {
	IdMessage string `json:"idMessage"`
}

func main() {
	http.HandleFunc("/", serveHTML)
	http.HandleFunc("/getSettings", handleGetSettings)
	http.HandleFunc("/getStateInstance", handleGetStateInstance)
	http.HandleFunc("/sendMessage", handleSendMessage)
	http.HandleFunc("/sendFileByUrl", handleSendFileByUrl)

	fmt.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func serveHTML(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func handleGetSettings(w http.ResponseWriter, r *http.Request) {
	idInstance := r.URL.Query().Get("idInstance")
	apiToken := r.URL.Query().Get("apiToken")
	url := fmt.Sprintf("%s%s/getSettings/%s", baseURL, idInstance, apiToken)
	log.Println(url)
	response, err := http.Get(url)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var apiResponse SettingResponce
	json.Unmarshal(body, &apiResponse)
	fmt.Println(apiResponse)
	json.NewEncoder(w).Encode(apiResponse)
}

func handleGetStateInstance(w http.ResponseWriter, r *http.Request) {
	idInstance := r.URL.Query().Get("idInstance")
	apiToken := r.URL.Query().Get("apiToken")
	url := fmt.Sprintf("%s%s/getStateInstance/%s", baseURL, idInstance, apiToken)
	log.Println(url)
	response, err := http.Get(url)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var apiResponse StateInstanceResponce
	json.Unmarshal(body, &apiResponse)
	fmt.Println(apiResponse)
	json.NewEncoder(w).Encode(apiResponse)
}

func handleSendMessage(w http.ResponseWriter, r *http.Request) {
	idInstance := r.URL.Query().Get("idInstance")
	apiToken := r.URL.Query().Get("apiToken")
	chatId := r.URL.Query().Get("chatId")
	message := r.URL.Query().Get("message")
	apiURL := fmt.Sprintf("%s%s/sendMessage/%s", baseURL, idInstance, apiToken)
	log.Println(apiURL)
	data := url.Values{}
	data.Set("chatId", chatId)
	data.Set("message", message)

	response, err := http.PostForm(apiURL, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var apiResponse SendMessageResponce
	json.Unmarshal(body, &apiResponse)
	fmt.Println(apiResponse)
	json.NewEncoder(w).Encode(apiResponse)
}

func handleSendFileByUrl(w http.ResponseWriter, r *http.Request) {
	idInstance := r.URL.Query().Get("idInstance")
	apiToken := r.URL.Query().Get("apiToken")
	chatId := r.URL.Query().Get("chatId")
	urlFile := r.URL.Query().Get("urlFile")
	fileName := r.URL.Query().Get("fileName")
	apiURL := fmt.Sprintf("%s%s/sendFileByUrl/%s", baseURL, idInstance, apiToken)
	log.Println(apiURL)
	data := url.Values{}
	data.Set("chatId", chatId)
	data.Set("urlFile", urlFile)
	data.Set("fileName", fileName)

	response, err := http.PostForm(apiURL, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var apiResponse SendFileByUrlResponce
	json.Unmarshal(body, &apiResponse)
	fmt.Println(apiResponse)
	json.NewEncoder(w).Encode(apiResponse)
}
