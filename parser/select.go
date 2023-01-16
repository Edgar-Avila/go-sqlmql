package parser

import "strings"

// Direction of an ORDER BY (ASC, DESC)
type Direction string

// Capture captures a DIRECTION (ASC, DESC) (Default ASC)
func (d *Direction) Capture(values []string) error {
	*d = "ASC"
	if len(values) > 0 && strings.EqualFold(values[0], "DESC") {
		*d = "DESC"
	}
	return nil
}

// Sort specification for an ORDER BY
type SortSpec struct {
	Col string    `parser:"@Ident"`
	Dir Direction `parser:"@((Asc|Desc)?)"`
}

// Select statement
type SelectStatement struct {
	Distinct bool       `parser:"Select Distinct?"`
	Cols     []string   `parser:"((@'*')|(@Ident (',' @Ident)*))"`
	From     string     `parser:"From @Ident"`
	Where    *Where     `parser:"@@?"`
	GroupBy  []string   `parser:"(Group By @Ident (',' @Ident)*)?"`
	OrderBy  []SortSpec `parser:"(Order By @@ (',' @@)*)?"`
	Limit    int64      `parser:"(Limit @Int)? ';'?"`
}

func (s *SelectStatement) stmt() {}
