package mapping

import "entgo.io/ent/dialect"

var (
	Date     = map[string]string{dialect.MySQL: "date"}
	DateTime = map[string]string{dialect.MySQL: "datetime"}
)
