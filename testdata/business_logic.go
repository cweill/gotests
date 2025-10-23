package testdata

import (
	"errors"
	"fmt"
)

// CalculateDiscount applies a percentage discount to a price
func CalculateDiscount(price float64, percentage int) (float64, error) {
	if price < 0 {
		return 0, errors.New("price cannot be negative")
	}

	if percentage < 0 || percentage > 100 {
		return 0, errors.New("percentage must be between 0 and 100")
	}

	discount := price * float64(percentage) / 100.0
	finalPrice := price - discount

	// Round to 2 decimal places
	return float64(int(finalPrice*100+0.5)) / 100, nil
}

// IsEligible checks if someone is eligible based on age and license
func IsEligible(age int, hasLicense bool) bool {
	if age < 0 {
		return false
	}

	// Must be at least 18 and have a license
	return age >= 18 && hasLicense
}

// FormatCurrency formats an amount as currency with the given code
func FormatCurrency(amount float64, code string) (string, error) {
	if code == "" {
		return "", errors.New("currency code cannot be empty")
	}

	// Simple currency codes
	supportedCodes := map[string]string{
		"USD": "$",
		"EUR": "€",
		"GBP": "£",
		"JPY": "¥",
	}

	symbol, ok := supportedCodes[code]
	if !ok {
		return "", fmt.Errorf("unsupported currency code: %s", code)
	}

	// Format with 2 decimal places (except JPY which uses 0)
	if code == "JPY" {
		return fmt.Sprintf("%s%.0f", symbol, amount), nil
	}

	return fmt.Sprintf("%s%.2f", symbol, amount), nil
}

// CalculateShippingCost calculates shipping based on weight and distance
func CalculateShippingCost(weightKg float64, distanceKm int) (float64, error) {
	if weightKg <= 0 {
		return 0, errors.New("weight must be positive")
	}

	if distanceKm <= 0 {
		return 0, errors.New("distance must be positive")
	}

	// Base rate: $5
	cost := 5.0

	// Add per kg charge: $2/kg
	cost += weightKg * 2.0

	// Add distance charge
	if distanceKm <= 50 {
		// Local: no extra charge
	} else if distanceKm <= 200 {
		// Regional: $10
		cost += 10.0
	} else {
		// Long distance: $10 + $0.05/km over 200
		cost += 10.0 + float64(distanceKm-200)*0.05
	}

	return float64(int(cost*100+0.5)) / 100, nil
}

// ApplyLoyaltyPoints applies loyalty points as a discount
func ApplyLoyaltyPoints(totalAmount float64, points int) float64 {
	if totalAmount <= 0 || points <= 0 {
		return totalAmount
	}

	// Each point is worth $0.01
	discount := float64(points) * 0.01

	// Discount cannot exceed total amount
	if discount > totalAmount {
		return 0
	}

	return totalAmount - discount
}

// ValidateOrderQuantity checks if an order quantity is valid
func ValidateOrderQuantity(quantity, minOrder, maxStock int) error {
	if quantity <= 0 {
		return errors.New("quantity must be positive")
	}

	if minOrder > 0 && quantity < minOrder {
		return fmt.Errorf("minimum order quantity is %d", minOrder)
	}

	if maxStock > 0 && quantity > maxStock {
		return fmt.Errorf("only %d items in stock", maxStock)
	}

	return nil
}
