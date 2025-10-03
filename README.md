# 🌐 Go Yandex Cloud Translate

[![Go Report Card](https://goreportcard.com/badge/github.com/ahrorbekOlimjonov/go-yandex-translate)](https://goreportcard.com/report/github.com/ahrorbekOlimjonov/go-yandex-translate)
[![Go Reference](https://pkg.go.dev/badge/github.com/ahrorbekOlimjonov/go-yandex-translate.svg)](https://pkg.go.dev/github.com/ahrorbekOlimjonov/go-yandex-translate)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Release](https://img.shields.io/github/v/release/ahrorbekOlimjonov/go-yandex-translate.svg)](https://github.com/ahrorbekOlimjonov/go-yandex-translate/releases)

An idiomatic and lightweight Go client for the [Yandex Cloud Translate](https://cloud.yandex.com/en/services/translate) and [SpeechKit (Text-to-Speech)](https://cloud.yandex.com/en/services/speechkit) APIs.

This library provides a clean, fluent interface to seamlessly integrate Yandex's powerful machine translation and speech synthesis capabilities into your Go applications.

---

## ✨ Features

-   ✅ **Translate Text**: Simple and batch translation with language detection.
-   ✅ **Detect Language**: Reliably identify the language of any text.
-   ✅ **List Languages**: Get an up-to-date list of all supported languages.
-   ✅ **Synthesize Speech**: Convert text into natural-sounding speech (MP3).
-   ✅ **Lightweight**: Built on `resty` for efficient HTTP requests with minimal overhead.
-   ✅ **Well-Tested**: Robust unit tests using mock servers to ensure reliability.

---

## 🚀 Getting Started

### 1. Prerequisites

Before you begin, you'll need:
*   A [Yandex Cloud Account](https://cloud.yandex.com/en/).
*   Your **Folder ID**.
*   An **API Key** for your service account. You can find instructions on how to get one [here](https://cloud.yandex.com/en/docs/iam/operations/api-key/create).

### 2. Installation

Install the package using `go get`:
```sh
go get -u github.com/ahrorbekOlimjonov/go-yandex-translate