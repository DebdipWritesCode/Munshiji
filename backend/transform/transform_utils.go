package transform

func derefString(s *string) string {
	if s != nil {
		return *s
	}
	return ""
}

func derefFloat64(f *float64) float64 {
	if f != nil {
		return *f
	}
	return 0
}
