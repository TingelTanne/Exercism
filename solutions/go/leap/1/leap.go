// Package leap calculates if a year is a leap year.
package leap

// IsLeapYear takes a year as an int and returns true if it's a leap year.
func IsLeapYear(year int) bool {
	return year%4 == 0 && (year%100 != 0 || (year%100 == 0 && year%400 == 0))
}
