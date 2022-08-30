package schema

import (
	"ddd-sample/ent/schema/def"
	"ddd-sample/ent/schema/mapping"
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
		field.Int("id").GoType(ScreeningID(0)),
		field.Bool("can_add_interview").Default(false).Comment("面接追加制御フラグ"),
		field.Time("apply_date").SchemaType(mapping.Date).Optional().Nillable().Comment("応募日"),
		field.String("applicant_email_address").SchemaType(map[string]string{dialect.MySQL: "varchar(50)"}).Comment("応募者メールアドレス"),
		field.Time("created_at").SchemaType(mapping.DateTime).Default(time.Now).Annotations(def.Created),
	}

}
