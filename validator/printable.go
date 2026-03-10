package validator

// IsPrintableString returns true when s contains only characters allowed
// by ITU-T X.680 PrintableString (subset of ASCII printable characters).
func IsPrintableString(s string) bool {
	for i := 0; i < len(s); i++ {
		b := s[i]
		switch {
		case b >= 'A' && b <= 'Z':
		case b >= 'a' && b <= 'z':
		case b >= '0' && b <= '9':
		case b == ' ': // SPACE
		case b == '\'': // APOSTROPHE
		case b == '(':
		case b == ')':
		case b == '+':
		case b == ',':
		case b == '-':
		case b == '.':
		case b == '/':
		case b == ':':
		case b == '=':
		case b == '?':
		default:
			return false
		}
	}
	return true
}
