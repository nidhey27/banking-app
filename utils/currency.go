package utils

// Constants for all supported currencies
const (
	USD = "USD"
	INR = "INR"
	CAD = "CAD"
)

// IsSupportedCurrency returns true if the currecny is supported
func IsSupportedCurrency(currency string) bool {

	switch currency {
	case USD, CAD, INR:
		return true

	}

	return false
}
