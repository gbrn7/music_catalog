package tracks

import (
	"context"

	trackactivities "github.com/gbrn7/music_catalog/internal/models/trackActivities"
	"github.com/gbrn7/music_catalog/internal/repository/spotify"
)

//go:generate mockgen -source=service.go -destination=service_mock_test.go -package=tracks
type spotifyOutbound interface {
	Search(ctx context.Context, query string, limit, offset int) (*spotify.SpotifySearchResponse, error)
}

type trackActivitiesRepository interface {
	Create(ctx context.Context, model trackactivities.TrackActivity) error
	Update(ctx context.Context, model trackactivities.TrackActivity) error
	Get(ctx context.Context, userID uint, spotifyID string) (*trackactivities.TrackActivity, error)
	GetBulkBySpotifyIDs(ctx context.Context, userID uint, spotifyIDs []string) (map[string]trackactivities.TrackActivity, error)
}

type service struct {
	spotifyOutbound     spotifyOutbound
	trackActivitiesRepo trackActivitiesRepository
}

func NewService(spotifyOutbound spotifyOutbound, trackActivitiesRepo trackActivitiesRepository) *service {
	return &service{
		spotifyOutbound:     spotifyOutbound,
		trackActivitiesRepo: trackActivitiesRepo,
	}
}
