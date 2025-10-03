package yandexcloud

import "github.com/go-resty/resty/v2"

const (
	// Default API endpoints
	translateAPIBaseURL = "https://translate.api.cloud.yandex.net/translate/v2"
	speechAPIBaseURL    = "https://tts.api.cloud.yandex.net/speech/v1"
)

// Client is the main client for interacting with Yandex Cloud APIs.
type Client struct {
	apiKey   string
	folderID string
	client   *resty.Client

	// Endpoints can be overridden for testing or other purposes
	TranslateEndpoint string
	DetectEndpoint    string
	LanguageEndpoint  string
	SpeechEndpoint    string
}

// NewClient creates a new client for the Yandex Cloud API.
// An API key is required. The folderID is recommended for most operations.
func NewClient(apiKey, folderID string) *Client {
	return &Client{
		apiKey:   apiKey,
		folderID: folderID,
		client:   resty.New(),

		TranslateEndpoint: translateAPIBaseURL + "/translate",
		DetectEndpoint:    translateAPIBaseURL + "/detect",
		LanguageEndpoint:  translateAPIBaseURL + "/languages",
		SpeechEndpoint:    speechAPIBaseURL + "/synthesize",
	}
}
