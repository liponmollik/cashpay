package utils

import (
	"crypto/md5"
	"encoding/csv"
	"encoding/hex"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"math"
	"math/rand"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"gopkg.in/yaml.v2"
)

// HashTo11Digits takes an integer input, hashes it using MD5, and returns
// the first 11 characters of the hexadecimal hash as a string.
func make11DigitNumber(input int) (string, error) {
	// Convert the input number to a string
	inputStr := strconv.Itoa(input)

	// Calculate the MD5 hash of the input string
	hash := md5.Sum([]byte(inputStr))
	hashStr := hex.EncodeToString(hash[:]) // Convert hash bytes to a hexadecimal string

	// Take the first 8 characters of the hash string
	substr := hashStr[:8]

	// Convert the hexadecimal substring to a decimal number
	decimal, err := strconv.ParseUint(substr, 16, 64)
	if err != nil {
		return "", err
	}

	// Ensure the decimal number is 11 digits long
	output := fmt.Sprintf("%011d", decimal)

	return output, nil
}

// HashTo11Digits takes an integer input, hashes it using MD5, and returns
// the first 11 characters of the hexadecimal hash as a string.
func HashTo11Digits(input int) string {
	inputStr := strconv.Itoa(input)
	hash := md5.Sum([]byte(inputStr))
	hashStr := hex.EncodeToString(hash[:])
	return hashStr[:11]
}

// CheckError checks if an error is not nil and panics if it is.
func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

// IsEven returns true if the given number is even, false otherwise.
func IsEven(number int) bool {
	return number%2 == 0
}

// IsOdd returns true if the given number is odd, false otherwise.
func IsOdd(number int) bool {
	return !IsEven(number)
}

// Add adds two numbers and returns the result.
func Add(a, b int) int {
	return a + b
}

// Subtract subtracts one number from another and returns the result.
func Subtract(a, b int) int {
	return a - b
}

// Multiply multiplies two numbers and returns the result.
func Multiply(a, b int) int {
	return a * b
}

// Divide divides one number by another and returns the result.
// Returns an error if the divisor is zero.
func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

// GenerateRandomNumber generates a random number within the specified range.
func GenerateRandomNumber(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min+1) + min
}

// IsPalindrome checks if a given string is a palindrome.
func IsPalindrome(s string) bool {
	s = strings.ToLower(s)
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}

// IsPrime checks if a given number is prime.
func IsPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// Fibonacci returns the nth Fibonacci number.
func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	a, b := 0, 1
	for i := 2; i <= n; i++ {
		a, b = b, a+b
	}
	return b
}

// ReverseString reverses a given string.
func ReverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < len(runes)/2; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// TitleCase converts a string to title case.
func TitleCase(s string) string {
	return strings.Title(strings.ToLower(s))
}

// ExtractNumbers extracts numbers from a given string and returns them as a slice of integers.
func ExtractNumbers(s string) []int {
	re := regexp.MustCompile(`\d+`)
	matches := re.FindAllString(s, -1)
	numbers := make([]int, len(matches))
	for i, match := range matches {
		num, _ := strconv.Atoi(match)
		numbers[i] = num
	}
	return numbers
}

// ShuffleSlice shuffles the elements of a slice.
func ShuffleSlice(slice []int) {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(slice), func(i, j int) {
		slice[i], slice[j] = slice[j], slice[i]
	})
}

// SliceContains checks if a slice contains a given value.
func SliceContains(slice []int, value int) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

// JoinStrings joins a slice of strings into a single string, separated by a delimiter.
func JoinStrings(slice []string, delimiter string) string {
	return strings.Join(slice, delimiter)
}

// SplitString splits a string into a slice of strings using a delimiter.
func SplitString(s, delimiter string) []string {
	return strings.Split(s, delimiter)
}

// FormatNumber formats a number with commas every three digits.
func FormatNumber(n int) string {
	s := strconv.Itoa(n)
	var formatted string
	for i, r := range s {
		if i > 0 && (len(s)-i)%3 == 0 {
			formatted += ","
		}
		formatted += string(r)
	}
	return formatted
}

// ConvertToBinary converts a decimal number to binary.
func ConvertToBinary(n int) string {
	return strconv.FormatInt(int64(n), 2)
}

// ConvertToDecimal converts a binary number to decimal.
func ConvertToDecimal(binary string) (int, error) {
	decimal, err := strconv.ParseInt(binary, 2, 64)
	return int(decimal), err
}

// ToLower converts a string to lowercase.
func ToLower(s string) string {
	return strings.ToLower(s)
}

// ToUpper converts a string to uppercase.
func ToUpper(s string) string {
	return strings.ToUpper(s)
}

// RepeatString repeats a string n times.
func RepeatString(s string, n int) string {
	return strings.Repeat(s, n)
}

// Capitalize capitalizes the first letter of a string.
func Capitalize(s string) string {
	if len(s) == 0 {
		return s
	}
	return strings.ToUpper(string(s[0])) + s[1:]
}

// Truncate truncates a string to the specified length.
func Truncate(s string, length int) string {
	if len(s) <= length {
		return s
	}
	return s[:length]
}

// PadLeft pads a string with a specified character on the left to reach a certain length.
func PadLeft(s string, length int, padChar rune) string {
	if len(s) >= length {
		return s
	}
	return strings.Repeat(string(padChar), length-len(s)) + s
}

// PadRight pads a string with a specified character on the right to reach a certain length.
func PadRight(s string, length int, padChar rune) string {
	if len(s) >= length {
		return s
	}
	return s + strings.Repeat(string(padChar), length-len(s))
}

// IsEmail checks if a string is a valid email address.
func IsEmail(email string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(email)
}

// IsPhoneNumber checks if a string is a valid phone number.
func IsPhoneNumber(phone string) bool {
	re := regexp.MustCompile(`^\+\d{1,3}\d{3,}$`)
	return re.MatchString(phone)
}

// IsURL checks if a string is a valid URL.
func IsURL(url string) bool {
	re := regexp.MustCompile(`^(http|https)://[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}/?`)
	return re.MatchString(url)
}

// DateDiff calculates the difference between two dates.
func DateDiff(start, end time.Time) time.Duration {
	return end.Sub(start)
}

// DaysBetween calculates the number of days between two dates.
func DaysBetween(start, end time.Time) int {
	return int(DateDiff(start, end).Hours() / 24)
}

// IsValidDate checks if a given date is a valid date.
func IsValidDate(date string) bool {
	_, err := time.Parse("2006-01-02", date)
	return err == nil
}

// IsValidTime checks if a given time is a valid time.
func IsValidTime(timeStr string) bool {
	_, err := time.Parse("15:04:05", timeStr)
	return err == nil
}

// IsValidDateTime checks if a given datetime string is a valid datetime.
func IsValidDateTime(datetime string) bool {
	_, err := time.Parse("2006-01-02 15:04:05", datetime)
	return err == nil
}

// DaysInMonth returns the number of days in a given month and year.
func DaysInMonth(year, month int) int {
	return time.Date(year, time.Month(month)+1, 0, 0, 0, 0, 0, time.UTC).Day()
}

// IsLeapYear checks if a given year is a leap year.
func IsLeapYear(year int) bool {
	return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}

// Encrypt encrypts a string using a Caesar cipher with the specified shift.
func Encrypt(text string, shift int) string {
	result := ""
	for _, char := range text {
		if char >= 'A' && char <= 'Z' {
			result += string((char-65+int32(shift))%26 + 65)
		} else if char >= 'a' && char <= 'z' {
			result += string((char-97+int32(shift))%26 + 97)
		} else {
			result += string(char)
		}
	}
	return result
}

// Decrypt decrypts a string using a Caesar cipher with the specified shift.
func Decrypt(text string, shift int) string {
	return Encrypt(text, 26-shift)
}

// IsValidCreditCard checks if a credit card number is valid using the Luhn algorithm.
func IsValidCreditCard(number string) bool {
	sum := 0
	double := false
	for i := len(number) - 1; i >= 0; i-- {
		digit := int(number[i] - '0')
		if double {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}
		sum += digit
		double = !double
	}
	return sum%10 == 0
}

// IsIPv4 checks if a given string is a valid IPv4 address.
func IsIPv4(ip string) bool {
	re := regexp.MustCompile(`^\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}$`)
	return re.MatchString(ip)
}

// IsIPv6 checks if a given string is a valid IPv6 address.
func IsIPv6(ip string) bool {
	re := regexp.MustCompile(`^([0-9a-fA-F]{1,4}:){7}[0-9a-fA-F]{1,4}$`)
	return re.MatchString(ip)
}

// IsValidMACAddress checks if a given string is a valid MAC address.
func IsValidMACAddress(mac string) bool {
	re := regexp.MustCompile(`^([0-9A-Fa-f]{2}[:-]){5}([0-9A-Fa-f]{2})$`)
	return re.MatchString(mac)
}

// IsISBN checks if a given string is a valid ISBN-10 or ISBN-13.
func IsISBN(isbn string) bool {
	// Check if ISBN-10 or ISBN-13
	if len(isbn) != 10 && len(isbn) != 13 {
		return false
	}
	if len(isbn) == 10 {
		return IsValidISBN10(isbn)
	}
	return IsValidISBN13(isbn)
}

// IsValidISBN10 checks if a given string is a valid ISBN-10.
func IsValidISBN10(isbn string) bool {
	if len(isbn) != 10 {
		return false
	}
	sum := 0
	for i := 0; i < 9; i++ {
		digit := int(isbn[i] - '0')
		if digit < 0 || digit > 9 {
			return false
		}
		sum += (i + 1) * digit
	}
	checksum := isbn[9]
	if checksum == 'X' {
		sum += 10 * 10
	} else {
		digit := int(checksum - '0')
		if digit < 0 || digit > 9 {
			return false
		}
		sum += 10 * digit
	}
	return sum%11 == 0
}

// IsValidISBN13 checks if a given string is a valid ISBN-13.
func IsValidISBN13(isbn string) bool {
	if len(isbn) != 13 {
		return false
	}
	sum := 0
	for i := 0; i < 12; i++ {
		digit := int(isbn[i] - '0')
		if digit < 0 || digit > 9 {
			return false
		}
		if i%2 == 0 {
			sum += digit
		} else {
			sum += 3 * digit
		}
	}
	checksum := int(isbn[12] - '0')
	if checksum < 0 || checksum > 9 {
		return false
	}
	return (10-sum%10)%10 == checksum
}

// IsValidPostalCode checks if a given string is a valid postal code.
func IsValidPostalCode(postalCode string) bool {
	re := regexp.MustCompile(`^[A-Z0-9]{3,10}$`)
	return re.MatchString(postalCode)
}

// IsValidPassword checks if a given string is a valid password.
func IsValidPassword(password string) bool {
	// Password must be at least 8 characters long and contain at least one uppercase letter, one lowercase letter, one digit, and one special character.
	re := regexp.MustCompile(`^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[@$!%*?&])[A-Za-z\d@$!%*?&]{8,}$`)
	return re.MatchString(password)
}

// IsValidUsername checks if a given string is a valid username.
func IsValidUsername(username string) bool {
	// Username must contain only alphanumeric characters and underscores, and be between 3 and 20 characters long.
	re := regexp.MustCompile(`^[a-zA-Z0-9_]{3,20}$`)
	return re.MatchString(username)
}

// IsValidURL checks if a given string is a valid URL.
func IsValidURL(url string) bool {
	re := regexp.MustCompile(`^(http|https):\/\/[^\s\/$.?#].[^\s]*$`)
	return re.MatchString(url)
}

// StripHTML removes HTML tags from a string.
func StripHTML(html string) string {
	re := regexp.MustCompile(`<[^>]*>`)
	return re.ReplaceAllString(html, "")
}

// IsValidJSON checks if a given string is a valid JSON.
func IsValidJSON(jsonStr string) bool {
	var js map[string]interface{}
	return json.Unmarshal([]byte(jsonStr), &js) == nil
}

// IsValidXML checks if a given string is a valid XML.
func IsValidXML(xmlStr string) bool {
	return xml.Unmarshal([]byte(xmlStr), new(interface{})) == nil
}

// IsValidCSV checks if a given string is a valid CSV.
func IsValidCSV(csvStr string) bool {
	r := csv.NewReader(strings.NewReader(csvStr))
	_, err := r.ReadAll()
	return err == nil
}

// IsValidYAML checks if a given string is a valid YAML.
func IsValidYAML(yamlStr string) bool {
	var out map[string]interface{}
	return yaml.Unmarshal([]byte(yamlStr), &out) == nil
}

// IsValidMarkdown checks if a given string is a valid Markdown.
func IsValidMarkdown(markdown string) bool {
	md := goldmark.New(
		goldmark.WithExtensions(extension.GFM),
	)
	return md.Convert([]byte(markdown), os.Stdout) == nil
}

// IsValidHTML checks if a given string is a valid HTML.
func IsValidHTML(html string) bool {
	_, err := html.Parse(strings.NewReader(html))
	return err == nil
}

// IsValidCSS checks if a given string is valid CSS.
func IsValidCSS(css string) bool {
	_, err := css.Parse(css, nil)
	return err == nil
}

// IsValidSQL checks if a given string is a valid SQL statement.
func IsValidSQL(sql string) bool {
	_, err := sqlparser.Parse(sql)
	return err == nil
}

// IsValidRegExp checks if a given string is a valid regular expression.
func IsValidRegExp(regexp string) bool {
	_, err := regexp.Compile(regexp)
	return err == nil
}
