package database

import (
	"app/models"
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/mmcloughlin/geohash"
	"gorm.io/gorm"
)

func SeedDatabase() {
	// reading json file

	cwd, _ := os.Getwd()
	fmt.Println("Current working directory:", cwd)

	// Reading file

	file, fileErr := os.Open("./database/yelp.json")

	if fileErr != nil {
		log.Fatal("Error opening the file: ", fileErr.Error())
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {

		line := scanner.Text()

		var payload = models.Businesses{UserID: 1}

		if err := json.Unmarshal([]byte(line), &payload); err != nil {
			log.Println("Error parsing JSON:", err)
			continue
		}

		// geoHash

		hash := geohash.EncodeWithPrecision(payload.Latitude, payload.Longitude, 12)

		payload.GeoHash = hash

		dbInstance := Database.DB

		if err := dbInstance.Create(&payload).Error; err != nil {
			if errors.Is(err, gorm.ErrDuplicatedKey) {
				fmt.Printf(err.Error())
			}
			continue

		}

	}

	if err := scanner.Err(); err != nil {
		log.Fatal("Error reading file:", err)
	}

}
