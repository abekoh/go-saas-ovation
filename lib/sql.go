package lib

import "database/sql"

func IntToNullInt32(i *int) sql.NullInt32 {
	if i == nil {
		return sql.NullInt32{
			Int32: 0,
			Valid: false,
		}
	} else {
		return sql.NullInt32{
			Int32: int32(*i),
			Valid: true,
		}
	}
}
