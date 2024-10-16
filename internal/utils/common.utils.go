package utils

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func HexStr(n int64) string {
  return strconv.FormatInt(int64(n), 16)
}


func ParseSizeToBytes(sizeStr string) (uint64, error) {
	// Define the unit multipliers in bytes
	unitMultipliers := map[string]uint64{
		"B":   1,
		"KB":  1024,
		"Kb":  1024 / 8,
		"MB":  1024 * 1024,
		"Mb":  (1024 * 1024) / 8,
		"GB":  1024 * 1024 * 1024,
		"Gb":  (1024 * 1024 * 1024) / 8,
	}

	// Use regex to extract the numeric part and the unit part
	re := regexp.MustCompile(`(?i)^(\d+)([A-Za-z]+)$`)
	matches := re.FindStringSubmatch(sizeStr)
	if len(matches) != 3 {
		return 0, fmt.Errorf("invalid format: %s", sizeStr)
	}

	// Parse the numeric part
	value, err := strconv.ParseUint(matches[1], 10, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid number: %s", matches[1])
	}

	// Get the unit (case-insensitive)
	unit := strings.ToUpper(matches[2])
	if strings.HasSuffix(matches[2], "b") && !strings.HasSuffix(matches[2], "B") {
		unit = strings.Title(matches[2]) // Ensure lowercase 'b' for bits
	}

	// Multiply the value by the appropriate unit multiplier
	multiplier, ok := unitMultipliers[unit]
	if !ok {
		return 0, fmt.Errorf("unknown unit: %s", matches[2])
	}

	return value * multiplier, nil
}
