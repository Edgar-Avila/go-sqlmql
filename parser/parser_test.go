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

func TestInsertStatement(t *testing.T) {
    text := 
    "INSERT INTO students\n" +
    "(id, name, lastname, graduated, score)\n" +
    "VALUES\n" + 
    "(1, 'John', 'Doe', TRUE, 10.0),\n" +
    "(2, 'Jane', 'Doe', FALSE, 10.0);"

    parsed := HelpParse(t, text)
    PrettyPrint(t, parsed)
}

func TestSelectStatement(t *testing.T) {
    text := 
    "SELECT * FROM students\n" +
    "WHERE score < 5.0 AND graduated = TRUE OR lastname = \"Doe\"\n" +
    "GROUP BY graduated, lastname\n" +
    "ORDER BY score DESC\n" +
    "LIMIT 10;"

    parsed := HelpParse(t, text)
    PrettyPrint(t, parsed)
}
