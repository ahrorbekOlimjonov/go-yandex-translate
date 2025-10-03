package yandexcloud

import (
	"context"
	"fmt"
)

const (
	FormatPlainText = "PLAIN_TEXT"
	FormatHTML      = "HTML"
)

// Translate translates a single text string to the target language.
// It returns the translated text and the detected source language code.
func (c *Client) Translate(ctx context.Context, text string, targetLanguage string, opts *TranslateOptions) (string, string, error) {
	texts := []string{text}
	translations, err := c.TranslateMultiple(ctx, texts, targetLanguage, opts)
	if err != nil {
		return "", "", err
	}

	if len(translations) == 0 {
		return "", "", fmt.Errorf("no translation received")
	}

	return translations[0].Text, translations[0].DetectedLanguageCode, nil
}

// TranslateMultiple translates a slice of texts to the target language.
func (c *Client) TranslateMultiple(ctx context.Context, texts []string, targetLanguage string, opts *TranslateOptions) ([]TranslatedText, error) {
	if len(texts) == 0 {
		return nil, fmt.Errorf("texts slice is empty")
	}

	request := TranslationRequest{
		TargetLanguageCode: targetLanguage,
		Texts:              texts,
		FolderID:           c.folderID, // Use folderID from client by default
	}

	// Apply options if provided
	if opts != nil {
		if opts.SourceLanguage != "" {
			request.SourceLanguageCode = opts.SourceLanguage
		}
		if opts.FolderID != "" {
			request.FolderID = opts.FolderID // Option overrides client default
		}
		if opts.Model != "" {
			request.Model = opts.Model
		}
		if opts.GlossaryConfig != nil {
			request.GlossaryConfig = opts.GlossaryConfig
		}
		if opts.EnableSpeller {
			speller := true
			request.Speller = &speller
		}
		if opts.Format != "" {
			request.Format = opts.Format
		} else {
			request.Format = FormatPlainText
		}
	}

	var response TranslationResponse

	resp, err := c.client.R().
		SetContext(ctx).
		SetHeader("Authorization", "Api-Key "+c.apiKey).
		SetBody(request).
		SetResult(&response).
		Post(c.TranslateEndpoint)

	if err != nil {
		return nil, fmt.Errorf("translation request failed: %w", err)
	}

	if resp.IsError() {
		return nil, fmt.Errorf("translation API returned error: status code %d, body: %s", resp.StatusCode(), resp.String())
	}

	return response.Translations, nil
}