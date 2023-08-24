package utils

func GetGeoHashLength(coverRange float64) int {
	switch {
	case coverRange <= 0.5:
		return 7
	case coverRange > 0.5 && coverRange <= 2:
		return 6
	case coverRange > 2 && coverRange <= 5:
		return 5
	case coverRange > 5 && coverRange <= 40:
		return 4
	case coverRange > 40 && coverRange <= 156:
		return 3
	case coverRange > 156 && coverRange <= 1250:
		return 2
	case coverRange > 1250 && coverRange < 5000:
		return 1
	default:
		return 6
	}
}
