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

func TestCreateStatement(t *testing.T) {
    parser, err := NewParser()
    if err != nil {
        t.Fatal(err)
    }

    text := "CREATE TABLE students (Id VARCHAR(100) NOT NULL);"
    parsed, err := parser.ParseString("", text)
    if err != nil {
        t.Fatal(err)
    }
    PrettyPrint(t, &parsed)
}
