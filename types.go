package yandexcloud

// =================================================================================
// Translation Types
// =================================================================================

// TranslateOptions provides optional parameters for a translation request.
type TranslateOptions struct {
	SourceLanguage string
	FolderID       string
	Model          string
	GlossaryConfig *GlossaryConfig
	EnableSpeller  bool
	Format         string // "PLAIN_TEXT" or "HTML"
}

// TranslationRequest is the request body for the Translate API.
type TranslationRequest struct {
	TargetLanguageCode string          `json:"targetLanguageCode"`
	Texts              []string        `json:"texts"`
	FolderID           string          `json:"folderId,omitempty"`
	SourceLanguageCode string          `json:"sourceLanguageCode,omitempty"`
	Model              string          `json:"model,omitempty"`
	GlossaryConfig     *GlossaryConfig `json:"glossaryConfig,omitempty"`
	Speller            *bool           `json:"speller,omitempty"`
	Format             string          `json:"format,omitempty"`
}

// GlossaryConfig specifies the glossary to use for translation.
type GlossaryConfig struct {
	GlossaryData *GlossaryData `json:"glossaryData"`
}

// GlossaryData contains pairs of source and translated text.
type GlossaryData struct {
	GlossaryPairs []GlossaryPair `json:"glossaryPairs"`
}

// GlossaryPair is a single source-text/translated-text pair.
type GlossaryPair struct {
	SourceText     string `json:"sourceText"`
	TranslatedText string `json:"translatedText"`
}

// TranslationResponse is the response from the Translate API.
type TranslationResponse struct {
	Translations []TranslatedText `json:"translations"`
}

// TranslatedText contains the translated text and the detected source language.
type TranslatedText struct {
	Text                 string `json:"text"`
	DetectedLanguageCode string `json:"detectedLanguageCode"`
}

// =================================================================================
// Language Detection Types
// =================================================================================

// DetectLanguageOptions provides optional parameters for language detection.
type DetectLanguageOptions struct {
	FolderID          string
	LanguageCodeHints []string
}

// DetectLanguageRequest is the request body for the Detect Language API.
type DetectLanguageRequest struct {
	Text              string   `json:"text"`
	LanguageCodeHints []string `json:"languageCodeHints,omitempty"`
	FolderID          string   `json:"folderId,omitempty"`
}

// DetectLanguageResponse is the response from the Detect Language API.
type DetectLanguageResponse struct {
	LanguageCode string `json:"languageCode"`
}

// ListLanguagesRequest is the request body for the List Languages API.
type ListLanguagesRequest struct {
	FolderID string `json:"folderId,omitempty"`
}

// ListLanguagesResponse is the response from the List Languages API.
type ListLanguagesResponse struct {
	Languages []Language `json:"languages"`
}

// Language represents a language supported by the API.
type Language struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

// =================================================================================
// Text-to-Speech (Speech Synthesis) Types
// =================================================================================

// TTSRequest is the request body for the Text-to-Speech API.
type TTSRequest struct {
	Text            string          `json:"text"`
	Hints           []Hint          `json:"hints,omitempty"`
	OutputAudioSpec OutputAudioSpec `json:"outputAudioSpec"`
	LoudnessDb      float64         `json:"loudnessDb,omitempty"`
}

// Hint provides additional synthesis parameters.
type Hint struct {
	Voice string `json:"voice"`
}

// OutputAudioSpec specifies the desired audio format.
type OutputAudioSpec struct {
	ContainerAudio ContainerAudio `json:"containerAudio"`
}

// ContainerAudio specifies the audio container format.
type ContainerAudio struct {
	ContainerAudioType string `json:"containerAudioType"` // "MP3", "WAV", etc.
}

// TTSResponse is the response from the Text-to-Speech API.
type TTSResponse struct {
	Result struct {
		AudioChunk struct {
			Data string `json:"data"`
		} `json:"audioChunk"`
	} `json:"result"`
}
