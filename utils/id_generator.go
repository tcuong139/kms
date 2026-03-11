package utils

import (
	"fmt"
	"kms_golang/database"
)

// GenerateID generates a prefixed sequential ID from the database
// e.g. C000001 for customers, P000001 for properties
func GenerateID(prefix string, table string, column string) (string, error) {
	var maxID string
	query := fmt.Sprintf("SELECT COALESCE(MAX(%s), '') FROM %s WHERE %s LIKE ?", column, table, column)
	if err := database.DB.Raw(query, prefix+"%").Scan(&maxID).Error; err != nil {
		return "", err
	}

	var seq int
	if maxID == "" {
		seq = 1
	} else {
		// Extract the numeric part after the prefix
		numStr := maxID[len(prefix):]
		fmt.Sscanf(numStr, "%d", &seq)
		seq++
	}

	return fmt.Sprintf("%s%06d", prefix, seq), nil
}

// GenerateCustomerID generates a Customer CD like C000001
func GenerateCustomerID() (string, error) {
	return GenerateID("C", "customers", "customer_cd")
}

// GeneratePropertyID generates a Property CD like P000001
func GeneratePropertyID() (string, error) {
	return GenerateID("P", "prop_basics", "prop_cd")
}

// GenerateSekosakiID generates a Sekosaki CD like S000001
func GenerateSekosakiID() (string, error) {
	return GenerateID("S", "sekosaki", "sekosaki_cd")
}

// GenerateReceptionNumber generates a reception accept number like R000001
func GenerateReceptionNumber() (string, error) {
	return GenerateID("R", "reception", "accept_number")
}
