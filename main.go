package main

import (
	"app-bioskop/cmd"
	"app-bioskop/internal/data/repository"
	"app-bioskop/internal/wire"
	"app-bioskop/pkg/database"
	"app-bioskop/pkg/utils"
	"log"
)

func main() {
	config, err := utils.ReadConfiguration()
	if err != nil {
		log.Fatalf("Gagal membaca konfigurasi: %v", err)
	}
	logger, err := utils.InitLogger("./logs/app-", true)
	db, err := database.InitDB(config.DB)
	if err != nil {
		log.Fatalf("Gagal konek ke database: %v", err)
	}
	defer db.Close()

	repo := repository.AllRepository(db, logger)
	router := wire.Wiring(repo, logger, config.SMTP)
	cmd.APiserver(router)
}
