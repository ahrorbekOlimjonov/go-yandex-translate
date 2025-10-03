
package yandexcloud

import (
	"encoding/base64"
	"fmt"
)

// SynthesizeSpeech converts text to speech and returns the audio data as bytes (e.g., MP3).
// The `voice` parameter specifies the voice to use, e.g., "filipp", "jane".
func (c *Client) SynthesizeSpeech(text, voice string) ([]byte, error) {
	request := TTSRequest{
		Text: text,
		Hints: []Hint{
			{Voice: voice},
		},
		OutputAudioSpec: OutputAudioSpec{
			ContainerAudio: ContainerAudio{
				ContainerAudioType: "MP3",
			},
		},
	}

	var response TTSResponse

	resp, err := c.client.R().
		SetHeader("Authorization", "Api-Key "+c.apiKey).
		SetBody(request).
		SetResult(&response).
		Post(c.SpeechEndpoint)

	if err != nil {
		return nil, fmt.Errorf("error making speech synthesis request: %w", err)
	}
	if resp.IsError() {
		return nil, fmt.Errorf("yandex tts API error: status code %d, body: %s", resp.StatusCode(), resp.String())
	}

	audioBytes, err := base64.StdEncoding.DecodeString(response.Result.AudioChunk.Data)
	if err != nil {
		return nil, fmt.Errorf("error decoding base64 audio data: %w", err)
	}

	return audioBytes, nil
}