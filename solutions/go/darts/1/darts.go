package darts

func Score(x, y float64) int {
	switch {
	case x <= 1:
		switch {
		case y <= 1:
			return 10
		case y <= 5:
			return 5
		case y <= 10:
			return 1
		default:
			return 0
		}

	case x <= 5:
		switch {
		case y <= 5:
			return 5
		case y <= 10:
			return 1
		default:
			return 0
		}

	case x <= 10:
		if y <= 10 {
			return 1
		} else {
			return 0
		}

	default:
		return 0
	}
}
