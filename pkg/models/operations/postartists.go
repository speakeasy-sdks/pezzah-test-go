// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type PostArtistsRequestBody struct {
	AlbumsRecorded *int64  `json:"albums_recorded,omitempty"`
	ArtistGenre    *string `json:"artist_genre,omitempty"`
	ArtistName     *string `json:"artist_name,omitempty"`
	Username       string  `json:"username"`
}

// PostArtists400ApplicationJSON - Invalid request
type PostArtists400ApplicationJSON struct {
	Message *string `json:"message,omitempty"`
}

type PostArtistsResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Invalid request
	PostArtists400ApplicationJSONObject *PostArtists400ApplicationJSON
}
