package yandexcloud

import (
	"context"
	"fmt"
)

// DetectLanguage detects the language of the provided text.
func (c *Client) DetectLanguage(ctx context.Context, text string, opts *DetectLanguageOptions) (string, error) {
	if text == "" {
		return "", fmt.Errorf("text is empty")
	}

	request := DetectLanguageRequest{
		Text:     text,
		FolderID: c.folderID,
	}

	if opts != nil {
		if opts.FolderID != "" {
			request.FolderID = opts.FolderID
		}
		if len(opts.LanguageCodeHints) > 0 {
			request.LanguageCodeHints = opts.LanguageCodeHints
		}
	}

	var response DetectLanguageResponse

	resp, err := c.client.R().
		SetContext(ctx).
		SetHeader("Authorization", "Api-Key "+c.apiKey).
		SetBody(request).
		SetResult(&response).
		Post(c.DetectEndpoint)

	if err != nil {
		return "", fmt.Errorf("language detection request failed: %w", err)
	}

	if resp.IsError() {
		return "", fmt.Errorf("language detection API returned error: status code %d, body: %s", resp.StatusCode(), resp.String())
	}

	return response.LanguageCode, nil
}

// ListLanguages retrieves the list of supported languages for translation.
func (c *Client) ListLanguages(ctx context.Context, folderID string) ([]Language, error) {
	request := ListLanguagesRequest{
		FolderID: c.folderID,
	}

	if folderID != "" {
		request.FolderID = folderID
	}

	var response ListLanguagesResponse

	resp, err := c.client.R().
		SetContext(ctx).
		SetHeader("Authorization", "Api-Key "+c.apiKey).
		SetBody(request).
		SetResult(&response).
		Post(c.LanguageEndpoint)

	if err != nil {
		return nil, fmt.Errorf("list languages request failed: %w", err)
	}

	if resp.IsError() {
		return nil, fmt.Errorf("list languages API returned error: status code %d, body: %s", resp.StatusCode(), resp.String())
	}

	return response.Languages, nil
}