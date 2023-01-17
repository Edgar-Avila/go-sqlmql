package parser

type SqlFile struct {
	Statements []Statement `parser:"(@@ ';'?)+"`
}
