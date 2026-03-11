package utils

import (
	"kms_golang/database"
	"kms_golang/models"
)

// GetAddressByZip returns prefecture, city, town for a given zip code
func GetAddressByZip(zip string) (*models.Prefecture, *models.City, *models.Town, error) {
	var town models.Town
	if err := database.DB.Where("post_code = ?", zip).First(&town).Error; err != nil {
		return nil, nil, nil, err
	}

	var city models.City
	if town.CityID != nil {
		if err := database.DB.Where("city_id = ?", *town.CityID).First(&city).Error; err != nil {
			return nil, nil, &town, err
		}
	}

	var pref models.Prefecture
	if city.PrefectureID != nil {
		if err := database.DB.Where("prefecture_id = ?", *city.PrefectureID).First(&pref).Error; err != nil {
			return nil, &city, &town, err
		}
	}

	return &pref, &city, &town, nil
}

// GetPrefectures returns all prefectures
func GetPrefectures() ([]models.Prefecture, error) {
	var prefs []models.Prefecture
	err := database.DB.Order("prefecture_id ASC").Find(&prefs).Error
	return prefs, err
}

// GetCitiesByPref returns all cities for a prefecture (by string prefecture_id)
func GetCitiesByPref(prefID string) ([]models.City, error) {
	var cities []models.City
	err := database.DB.Where("prefecture_id = ?", prefID).Order("city_id ASC").Find(&cities).Error
	return cities, err
}

// GetTownsByCity returns all towns for a city (by string city_id)
func GetTownsByCity(cityID string) ([]models.Town, error) {
	var towns []models.Town
	err := database.DB.Where("city_id = ?", cityID).Order("town_id ASC").Find(&towns).Error
	return towns, err
}
