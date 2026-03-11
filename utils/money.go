package utils

import "math"

const (
	TaxRate8  = 0.08
	TaxRate10 = 0.10
)

// CalcTax8 calculates 8% tax for an amount (returns tax portion)
func CalcTax8(amount float64) float64 {
	return math.Floor(amount * TaxRate8)
}

// CalcTax10 calculates 10% tax for an amount (returns tax portion)
func CalcTax10(amount float64) float64 {
	return math.Floor(amount * TaxRate10)
}

// AmountWithTax8 returns amount + 8% tax
func AmountWithTax8(amount float64) float64 {
	return amount + CalcTax8(amount)
}

// AmountWithTax10 returns amount + 10% tax
func AmountWithTax10(amount float64) float64 {
	return amount + CalcTax10(amount)
}

// RoundFloor rounds down to nearest integer
func RoundFloor(amount float64) float64 {
	return math.Floor(amount)
}

// RoundCeil rounds up to nearest integer
func RoundCeil(amount float64) float64 {
	return math.Ceil(amount)
}

// RoundNearest rounds to nearest integer
func RoundNearest(amount float64) float64 {
	return math.Round(amount)
}

// CalcSubtotal calculates subtotal from unit price and quantity
func CalcSubtotal(unitPrice float64, qty float64) float64 {
	return unitPrice * qty
}

// CalcTotal sums a slice of amounts
func CalcTotal(amounts []float64) float64 {
	var total float64
	for _, a := range amounts {
		total += a
	}
	return total
}
