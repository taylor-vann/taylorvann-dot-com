package storex

import (
	"fmt"
	"testing"
)

const createTable = `
CREATE TABLE IF NOT EXISTS %s (
  id BIGSERIAL PRIMARY KEY,
	message VARCHAR(512) NOT NULL
);
`
const insertRow = `
INSERT INTO %s (
	message
)
VALUES 
  ($1)
RETURNING
	*;
`
const readRow = `SELECT * FROM %s;`
const dropTable = `DROP TABLE %s;`
const storexTest = "storex_unit_tests"
const message = "yo dawg!"

func TestExec(t *testing.T) {
	result, err := Exec(
		fmt.Sprintf(
			createTable,
			storexTest,
		),
	)

	if err != nil {
		t.Error("error in storex exec")
	}
	if result == nil {
		t.Error("storex result nil")
	}

	dropResult, errDropResult := Exec(
		fmt.Sprintf(
			dropTable,
			storexTest,
		),
	)

	if errDropResult != nil {
		t.Error("error in storex exec drop")
	}
	if dropResult == nil {
		t.Error("storex drop result nil")
	}
}

func TestQuery(t *testing.T) {
	result, err := Exec(
		fmt.Sprintf(
			createTable,
			storexTest,
		),
	)

	if err != nil {
		t.Error("error in storex exec")
	}
	if result == nil {
		t.Error("storex result nil")
	}

	insertResult, errInsertResult := Query(
		fmt.Sprintf(
			insertRow,
			storexTest,
		),
		message,
	)

	if errInsertResult != nil {
		t.Error("error in storex exec")
	}
	if insertResult == nil {
		t.Error("storex result nil")
	}

	readResult, errReadResult := Query(
		fmt.Sprintf(
			readRow,
			storexTest,
		),
	)

	readResult.Next()
	var id int
	var returnedMessage string
	errReturn := readResult.Scan(&id, &returnedMessage)
	if errReturn != nil {
		t.Error("error reading row")
	}

	if returnedMessage != message {
		t.Error("returned message is not equal to original message")
	}

	if errReadResult != nil {
		fmt.Println(errReadResult)
		t.Error("error in storex exec")
	}
	if readResult == nil {
		t.Error("storex result nil")
	}

	dropResult, errDropResult := Exec(
		fmt.Sprintf(
			dropTable,
			storexTest,
		),
	)

	if errDropResult != nil {
		t.Error("error in storex exec drop")
	}
	if dropResult == nil {
		t.Error("storex drop result nil")
	}
}
