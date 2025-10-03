package yandexcloud

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestClient_Translate(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var reqBody TranslationRequest
		err := json.NewDecoder(r.Body).Decode(&reqBody)
		require.NoError(t, err)

		assert.Equal(t, "en", reqBody.TargetLanguageCode)
		assert.Equal(t, []string{"Привет мир"}, reqBody.Texts)

		w.Header().Set("Content-Type", "application/json")
		response := TranslationResponse{
			Translations: []TranslatedText{
				{Text: "Hello world", DetectedLanguageCode: "ru"},
			},
		}
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	client := NewClient("fake-api-key", "fake-folder-id")
	client.TranslateEndpoint = server.URL

	text, detectedLang, err := client.Translate(context.Background(), "Привет мир", "en", nil)

	require.NoError(t, err)
	assert.Equal(t, "Hello world", text)
	assert.Equal(t, "ru", detectedLang)
}