package schema

import (
	"ddd-sample/ent/schema/def"
	"ddd-sample/ent/schema/mapping"
	"ddd-sample/ent/schema/property"
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
	"time"
)

type ScreeningID int

type Screening struct {
	ent.Schema
}

func (Screening) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").MaxLen(50).NotEmpty().Unique().Immutable(),
		field.Enum("screening_status").GoType(property.ScreeningStatus("")).Default(string(property.NotApplied)).Comment("採用選考ステータス"),
		field.Time("apply_date").SchemaType(mapping.Date).Optional().Nillable().Comment("応募日"),
		field.String("applicant_email_address").SchemaType(map[string]string{dialect.MySQL: "varchar(50)"}).Comment("応募者メールアドレス"),
		field.Time("created_at").SchemaType(mapping.DateTime).Default(time.Now).Annotations(def.Created),
	}

}
