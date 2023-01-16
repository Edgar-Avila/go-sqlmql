package parser

import "strings"

type Literal interface{ GetVal() interface{} }

type boolCapture bool

func (b *boolCapture) Capture(values []string) error {
	*b = boolCapture(strings.EqualFold(values[0], "TRUE"))
	return nil
}

type stringCapture string

func (s *stringCapture) Capture(values []string) error {
	*s = stringCapture(strings.Trim(values[0], "'\""))
	return nil
}

type Int struct {
	Value int `parser:"@Int"`
}

func (i Int) GetVal() interface{} { return i.Value }

type Decimal struct {
	Value float64 `parser:"@Decimal"`
}

func (d Decimal) GetVal() interface{} { return d.Value }

type String struct {
	Value stringCapture `parser:"@String"`
}

func (s String) GetVal() interface{} { return s.Value }

type Bool struct {
	Value boolCapture `parser:"@('TRUE'|'FALSE')"`
}

func (b Bool) GetVal() interface{} { return b.Value }
