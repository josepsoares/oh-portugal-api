package seeder

import (
	"encoding/json"
	"errors"
	"josepsoares/iberiapi/db"
	"josepsoares/iberiapi/models"
	"log"
	"os"

	"gorm.io/gorm"
)

func Seed() {
	dbConnection := db.DBConn

	// check meepley seeder

	// country seeder
	if err := InsertFromJSON(dbConnection, &models.Country{}, "portugal"); err != nil {
		log.Fatal("Error during regions bulk insert: ", err)
	}

	// regions seeder
	if err := InsertFromJSON(dbConnection, &models.Region{}, "regions"); err != nil {
		log.Fatal("Error during regions bulk insert: ", err)
	}

	// islands seeder
	if err := InsertFromJSON(dbConnection, &models.Island{}, "islands"); err != nil {
		log.Fatal("Error during islands bulk insert: ", err)
	}

	// mountains seeder
	if err := InsertFromJSON(dbConnection, &models.Mountain{}, "mountains"); err != nil {
		log.Fatal("Error during mountains bulk insert: ", err)
	}

	// rivers seeder
	if err := InsertFromJSON(dbConnection, &models.River{}, "rivers"); err != nil {
		log.Fatal("Error during rivers bulk insert: ", err)
	}

	// lagoons seeder
	if err := InsertFromJSON(dbConnection, &models.Lagoon{}, "lagoons"); err != nil {
		log.Fatal("Error during lagoons bulk insert: ", err)
	}

	// unesco world heritage sites seeder
	if err := InsertFromJSON(dbConnection, &models.UnescoWorldHeritageSite{}, "unesco_world_heritage_sites"); err != nil {
		log.Fatal("Error during unesco world heritage sites bulk insert: ", err)
	}
}

func InsertFromJSON(db *gorm.DB, model interface{}, filename string) error {
	path := "./data/"

	if db.Migrator().HasTable(model) {
		if err := db.First(model).Error; errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
	}

	content, err := os.ReadFile(path + filename + ".json")
	if err != nil {
		return err
	}

	var payload interface{}
	if err := json.Unmarshal(content, &payload); err != nil {
		return err
	}

	if err := db.Model(model).Create(payload).Error; err != nil {
		return err
	}

	return nil
}
