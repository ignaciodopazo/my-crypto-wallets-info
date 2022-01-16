package domain

import (
)

// Interface that aims to be implemeted by API-specific URL builders
// upon its requeriments for their query parameters.
//
// Minimum recomended fields:
//
//  baseURL string
//  queryParams QueryParams
type URLBuilder interface {
	// Creates a new url builder object
	NewURLBuilder() *URLBuilder
	// Define the base URL for the request.
	SetBaseURL(string)
	// Add needed query parameter (a key, value pair) to the builder.
	AddQueryParam(string, string)
	// Generate the URL with the current base URL and query parameters present in the object.
	GenerateURL() string
}

type defaultURLBuilder struct {
	baseURL     string
	queryParams QueryParams
}

// When an API doesn't need a specific URL construction for request to be sent to it,
// this default URL builder can be used.

var DefaultURLBuilder = NewURLBuilder()

func NewURLBuilder() *defaultURLBuilder {
	return &defaultURLBuilder{
		baseURL:     "",
		queryParams: NewQueryParams(),
	}
}

func (b *defaultURLBuilder) SetBaseURL(url string) {
	b.baseURL = url
}

func (b *defaultURLBuilder) AddQueryParam(key, value string) {
	b.queryParams.AddQueryParam(key, value)
}

func (b *defaultURLBuilder) GenerateURL() string {
	if b.baseURL == "" {
		return ""
	}
	result := b.baseURL
	firstPass := true
	for key, val := range b.queryParams {
		if !firstPass {
			result = result + "&" + key + "=" + val
		} else {
			result = result + "?" + key + "=" + val
			firstPass = false
		}
	}
	return result
}
