package problem_2

// gcdOfStrings returns the largest string x such that
// x divides both str1 and str2.
func gcdOfStrings(str1, str2 string) string {
	// Quick rejection:
	// If str1+str2 != str2+str1, there is no common base pattern.
	if str1+str2 != str2+str1 {
		return ""
	}

	// Otherwise, the answer is the prefix of length gcd(len(str1), len(str2))
	gcdLen := gcd(len(str1), len(str2))
	return str1[:gcdLen]
}

// gcd returns the greatest common divisor of two ints.
func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func main() {
	gcdOfStrings("ABABAB", "ABAB")
}
