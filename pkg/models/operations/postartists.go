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

func (o *PostArtistsRequestBody) GetAlbumsRecorded() *int64 {
	if o == nil {
		return nil
	}
	return o.AlbumsRecorded
}

func (o *PostArtistsRequestBody) GetArtistGenre() *string {
	if o == nil {
		return nil
	}
	return o.ArtistGenre
}

func (o *PostArtistsRequestBody) GetArtistName() *string {
	if o == nil {
		return nil
	}
	return o.ArtistName
}

func (o *PostArtistsRequestBody) GetUsername() string {
	if o == nil {
		return ""
	}
	return o.Username
}

// PostArtists400ApplicationJSON - Invalid request
type PostArtists400ApplicationJSON struct {
	Message *string `json:"message,omitempty"`
}

func (o *PostArtists400ApplicationJSON) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

type PostArtistsResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Invalid request
	PostArtists400ApplicationJSONObject *PostArtists400ApplicationJSON
}

func (o *PostArtistsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PostArtistsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PostArtistsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *PostArtistsResponse) GetPostArtists400ApplicationJSONObject() *PostArtists400ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.PostArtists400ApplicationJSONObject
}
