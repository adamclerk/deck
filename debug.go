package deck

import "fmt"

// Debugf statement for use to debug games
func Debugf(debug bool, format string, a ...interface{}) (int, error) {
	if debug {
		return fmt.Printf(format, a...)
	}
	return 0, nil
}
