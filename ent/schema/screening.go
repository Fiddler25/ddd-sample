package schema

import (
	"ddd-sample/ent/schema/property"
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
)

type ScreeningID int

type Screening struct {
	ent.Schema
}

func (Screening) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

func (Screening) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").MaxLen(50).NotEmpty().Unique().Immutable(),
		field.Enum("screening_status").GoType(property.ScreeningStatus("")).Default(string(property.NotApplied)).Comment("採用選考ステータス"),
		field.Time("apply_date").SchemaType(map[string]string{dialect.MySQL: "date"}).Optional().Nillable().Comment("応募日"),
		field.String("applicant_email_address").MaxLen(50).NotEmpty().Comment("応募者メールアドレス"),
	}

}
