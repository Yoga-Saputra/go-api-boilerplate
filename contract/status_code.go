// This package hold enum values or custom type of wallet service response code and status
package contract

import (
	"fmt"
	"strings"
)

type (
	// StatusCode custom type to hold value for standard status code.
	StatusCode uint
)

// Constant standar int.
const ssc = 2410

// Status Code & Message enum.
const (
	OK StatusCode = (iota + ssc)

	INTERNALERROR

	COMINGSOON
	MAINTENANCE

	FAILEDXFRNOTFOUND
	FAILEDXFRALREADYSETTLED
	ERRORPROCESSING
	STILLPROCESSING

	enumStatusLimit
)

// String give StatusCode type string value
func (s StatusCode) String(custom ...string) string {
	if len(custom) > 0 {
		return custom[0]
	}

	return [...]string{
		"OK",

		"Internal Error {0}",

		"This Feature is Cooming Soon",
		"This Feature Under Maintenance",

		"Failed Transfer Not Found",
		"Failed Transfer Already Settled",
		"Error occurred while processing",
		"Still processing by other user",
	}[s-ssc]
}

// String give StatusCode type string value
func (s StatusCode) FormatedString(appended ...string) (status string) {
	for k, v := range appended {
		status = strings.ReplaceAll(s.String(), fmt.Sprintf("{%d}", k), v)
	}

	return
}

// StatusCodeLists return all status code enum into map[uint]string list.
func StatusCodeLists() (l map[uint]string) {
	l = make(map[uint]string, enumStatusLimit)
	for i := StatusCode(0 + ssc); i < enumStatusLimit; i++ {
		l[uint(i)] = i.String()
	}

	return
}
