package queryhelper

import (
	"fmt"
	"strings"
)

func AppendOrderBy(q *string, sort string) {
	if sort == "" {
		return
	}

	var orderBy string

	for col := range strings.SplitSeq(sort, ",") {
		if strings.HasPrefix(col, "-") {
			orderBy += fmt.Sprintf(" %s DESC,", col[1:])
		} else {
			orderBy += fmt.Sprintf(" %s ASC,", col)
		}
	}

	// Remove trailing comma.
	orderBy = orderBy[:len(orderBy)-1]

	// Append ORDER BY clause to the original query string.
	*q += fmt.Sprintf(" ORDER BY %s", orderBy)
}
