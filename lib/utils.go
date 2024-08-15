package lib

import (
	"github.com/spf13/viper"
	"os"
	"regexp"
	"strconv"
)

func checkKey(key string) {
	_, envExists := os.LookupEnv(key)

	if !viper.IsSet(key) && !envExists {
		panic(key + " key is not set")
	}
}

func getString(key string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return viper.GetString(key)
}

func getStringOrPanic(key string) string {
	checkKey(key)
	return getString(key)
}

func getIntOrDefault(key string, def int) int {
	value := getString(key)

	if intValue, err := strconv.Atoi(value); err == nil {
		return intValue
	}

	return def
}

func getIntOrPanic(key string) int {
	checkKey(key)
	return getIntOrDefault(key, 0)
}

// CheckLicensePlateFormat checks if the license plate format is valid
//
//	func main() {
//		// Example usage
//		licensePlate := "B 1234 ABC" // Replace with any license plate number to test
//
//		if CheckLicensePlateFormat(licensePlate) {
//			fmt.Printf("License plate '%s' format is valid.\n", licensePlate)
//		} else {
//			fmt.Printf("License plate '%s' format is invalid.\n", licensePlate)
//		}
//	}
func CheckLicensePlateFormat(licensePlate string) bool {
	// Regular expression for Indonesian license plate format
	// Example: B 1234 AB or B 1234 XYZ
	// Adjust the regular expression according to specific formats if needed.
	regex := `^[A-Z]{1,2}\s\d{1,4}\s[A-Z]{1,3}$`

	match, _ := regexp.MatchString(regex, licensePlate)
	return match
}
