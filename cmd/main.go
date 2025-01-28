package main

import (
	"log"

	"github.com/gbrn7/music_catalog/internal/configs"
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

	r := gin.Default

	r.Run(cfg.Service.Port)
}
