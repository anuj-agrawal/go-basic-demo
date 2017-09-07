package numbers

func Sign(a int) string {
	switch {
	case a > 0:
		return "+"
	default:
		return "-"
	}
}
