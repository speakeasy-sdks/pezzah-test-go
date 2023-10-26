// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type GetArtistsRequest struct {
	// Limits the number of items on a page
	Limit *int64 `queryParam:"style=form,explode=true,name=limit"`
	// Specifies the page number of the artists to be displayed
	Offset *int64 `queryParam:"style=form,explode=true,name=offset"`
}

func (o *GetArtistsRequest) GetLimit() *int64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *GetArtistsRequest) GetOffset() *int64 {
	if o == nil {
		return nil
	}
	return o.Offset
}

// GetArtists400ApplicationJSON - Invalid request
type GetArtists400ApplicationJSON struct {
	Message *string `json:"message,omitempty"`
}

func (o *GetArtists400ApplicationJSON) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

type GetArtists200ApplicationJSON struct {
	AlbumsRecorded *int64  `json:"albums_recorded,omitempty"`
	ArtistGenre    *string `json:"artist_genre,omitempty"`
	ArtistName     *string `json:"artist_name,omitempty"`
	Username       string  `json:"username"`
}

func (o *GetArtists200ApplicationJSON) GetAlbumsRecorded() *int64 {
	if o == nil {
		return nil
	}
	return o.AlbumsRecorded
}

func (o *GetArtists200ApplicationJSON) GetArtistGenre() *string {
	if o == nil {
		return nil
	}
	return o.ArtistGenre
}

func (o *GetArtists200ApplicationJSON) GetArtistName() *string {
	if o == nil {
		return nil
	}
	return o.ArtistName
}

func (o *GetArtists200ApplicationJSON) GetUsername() string {
	if o == nil {
		return ""
	}
	return o.Username
}

type GetArtistsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Successfully returned a list of artists
	GetArtists200ApplicationJSONObjects []GetArtists200ApplicationJSON
	// Invalid request
	GetArtists400ApplicationJSONObject *GetArtists400ApplicationJSON
}

func (o *GetArtistsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetArtistsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetArtistsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetArtistsResponse) GetGetArtists200ApplicationJSONObjects() []GetArtists200ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.GetArtists200ApplicationJSONObjects
}

func (o *GetArtistsResponse) GetGetArtists400ApplicationJSONObject() *GetArtists400ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.GetArtists400ApplicationJSONObject
}
