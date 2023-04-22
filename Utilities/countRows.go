package utilities

import "database/sql"

func CountRows(rows *sql.Rows) int {
	count := 0

	for rows.Next() {
		count += 1
	}

	return count;
}