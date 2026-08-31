package advanced

import (
    "fmt"
    "strconv"
)

// DescribeNumber should return a string describing the number.
func DescribeNumber(f float64) string {
	return fmt.Sprintf("This is the number %v", strconv.FormatFloat(f, 'f', 1, 64))
}

type NumberBox interface {
	Number() int
}

// DescribeNumberBox should return a string describing the NumberBox.
func DescribeNumberBox(nb NumberBox) string {
	return fmt.Sprintf("This is a box containing the number %v", strconv.FormatFloat(float64(nb.Number()), 'f', 1, 64))
}

type FancyNumber struct {
	n string
}

func (i FancyNumber) Value() string {
	return i.n
}

type FancyNumberBox interface {
	Value() string
}

// ExtractFancyNumber should return the integer value for a FancyNumber
// and 0 if any other FancyNumberBox is supplied.
func ExtractFancyNumber(fnb FancyNumberBox) int {
	value, ok := fnb.(FancyNumber)
	if (ok) {
		integ, _ := strconv.Atoi(value.n)
        return integ
	}
	return 0
}

// DescribeFancyNumberBox should return a string describing the FancyNumberBox.
func DescribeFancyNumberBox(fnb FancyNumberBox) string {
	value, ok := fnb.(FancyNumber)
	finalValue := "0.0"
	if (ok) {
		integ, _ := strconv.Atoi(value.n)
		finalValue = strconv.FormatFloat(float64(integ), 'f', 1, 64)
	}
	return fmt.Sprintf("This is a fancy box containing the number %v", finalValue)
}

// DescribeAnything should return a string describing whatever it contains.
func DescribeAnything(i any) string {
	switch i.(type) {
		case float64:
			return DescribeNumber(i.(float64))
		case int:
			i = i.(int)
			return DescribeNumber(float64(i.(int)))
		case NumberBox:
			return DescribeNumberBox(i.(NumberBox))
		case FancyNumberBox:
			return DescribeFancyNumberBox(i.(FancyNumberBox))
		default:
			return "Return to sender"
	}
}