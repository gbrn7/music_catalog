package spotify

type SearchResponse struct {
	Limit  int                  `json:"limit"`
	Offset int                  `json:"offset"`
	Items  []SpotifyTrackObject `json:"item"`
	Total  int                  `json:"total"`
}

type SpotifyTracks struct {
	HREF     string               `json:"href"`
	Limit    int                  `json:"limit"`
	Next     *string              `json:"next"`
	Offset   int                  `json:"offset"`
	Previous *string              `json:"previous"`
	Total    int                  `json:"total"`
	Items    []SpotifyTrackObject `json:"item"`
}

type SpotifyTrackObject struct {
	// album type related fields
	AlbumType        string   `json:"albumType"`
	AlbumTotalTracks int      `json:"totalTracks"`
	AlbumImagesURL   []string `json:"albumImagesURL"`
	AlbumName        string   `json:"albumName"`

	// track related fields
	ArtistsName []string `json:"artistName"`

	// track re;ated fields
	Explicit bool   `json:"explicit"`
	ID       string `json:"id"`
	Name     string `json:"name"`
	IsLiked  *bool  `json:"isLiked"`
}
