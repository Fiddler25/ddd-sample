package def

import "entgo.io/ent/dialect/entsql"

var (
	Created = &entsql.Annotation{Default: "CURRENT_TIMESTAMP"}
	Updated = &entsql.Annotation{Default: "CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"}
)
