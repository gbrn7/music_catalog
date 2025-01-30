package main

import (
	"log"
	"net/http"

	"github.com/gbrn7/music_catalog/internal/configs"
	membershipHandler "github.com/gbrn7/music_catalog/internal/handler/memberships"
	tracksHandler "github.com/gbrn7/music_catalog/internal/handler/tracks"
	"github.com/gbrn7/music_catalog/internal/models/memberships"
	trackactivities "github.com/gbrn7/music_catalog/internal/models/trackactivities"
	membershipsRepo "github.com/gbrn7/music_catalog/internal/repository/memberships"
	"github.com/gbrn7/music_catalog/internal/repository/spotify"
	trackactivitiesRepo "github.com/gbrn7/music_catalog/internal/repository/trackactivities"
	membershipsSvc "github.com/gbrn7/music_catalog/internal/service/memberships"
	"github.com/gbrn7/music_catalog/internal/service/tracks"
	"github.com/gbrn7/music_catalog/pkg/httpclient"
	"github.com/gbrn7/music_catalog/pkg/internalsql"
	"github.com/gin-gonic/gin"
)

func main() {
	var (
		cfg *configs.Config
	)

	err := configs.Init(
		configs.WithConfigFolder(
			[]string{"./internal/configs/"},
		),
		configs.WithConfigFile("config"),
		configs.WithConfigType("yaml"),
	)
	if err != nil {
		log.Fatal("Gagal inisiasi config", err)
	}

	cfg = configs.Get()

	db, err := internalsql.Connect(cfg.Database.DataSourceName)
	if err != nil {
		log.Fatalf("failed to connext to database, err: %+v", err)
	}

	db.AutoMigrate(&memberships.User{})
	db.AutoMigrate(&trackactivities.TrackActivity{})

	r := gin.Default()

	httpClient := httpclient.NewClient(&http.Client{})
	spotifyOutbound := spotify.NewSpotifyOutbound(cfg, httpClient)

	membershipsRepo := membershipsRepo.NewRepository(db)
	trackActivitiesRepo := trackactivitiesRepo.NewRepository(db)

	membershipsSvc := membershipsSvc.NewService(cfg, membershipsRepo)
	trackSvc := tracks.NewService(spotifyOutbound, trackActivitiesRepo)

	membershipHandler := membershipHandler.NewHandler(r, membershipsSvc)
	membershipHandler.RegisterRoute()

	tracksHandler := tracksHandler.NewHandler(r, trackSvc)
	tracksHandler.RegisterRoute()

	r.Run(cfg.Service.Port)
}
