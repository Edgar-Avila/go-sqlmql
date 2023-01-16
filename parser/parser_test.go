package parser

import (
	"encoding/json"
	"testing"
)

func PrettyPrint(t *testing.T, obj interface{}) {
    t.Helper()
    b, _ := json.MarshalIndent(obj, "", "  ")
    t.Log(string(b))
}

func HelpParse(t *testing.T, text string) (*SqlFile) {
    t.Helper()
    parser, err := NewParser()
    if err != nil {
        t.Fatal(err)
    }
    parsed, err := parser.ParseString("", text)
    if err != nil {
        t.Fatal(err)
    }
    return parsed
}

func TestCreateStatement(t *testing.T) {
    text := "CREATE TABLE students (code VARCHAR(100) PRIMARY KEY);"
    parsed := HelpParse(t, text)
    PrettyPrint(t, parsed)
}

func TestDropStatement(t *testing.T) {
    text := "DROP TABLE IF EXISTS students, users, classes;"
    parsed := HelpParse(t, text)
    PrettyPrint(t, parsed)
}
