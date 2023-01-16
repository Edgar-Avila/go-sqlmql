package parser

// Comparison (=|<|>|<=|>=) between a column and a literal
type Comparison struct {
	Col string  `parser:"@Ident"`
	Op  string  `parser:"@('='|'<'|'>'|'<='|'>=')"`
	Lit Literal `parser:"@@"`
}

// Factor of a boolean expression
type BoolFactor struct {
	Not bool       `parser:"@Not?"`
	Cmp Comparison `parser:"@@"`
}

// Term of a boolean expression
type BoolTerm struct {
	Factors []BoolFactor `parser:"@@ (And @@)*"`
}

// Where clause
type Where struct {
	Terms []BoolTerm `parser:"Where @@ (Or @@)*"`
}
