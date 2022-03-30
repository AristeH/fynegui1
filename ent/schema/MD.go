package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// MDTabel holds the schema definition for the MDTabel entity.
type MDTabel struct {
	ent.Schema
}

// Fields of the MDTabel .
func (MDTabel) Fields() []ent.Field {
	return []ent.Field{
		field.String("namerus").MaxLen(300).StructTag(`json:"ИмяРус,omitempty"`).StorageKey("Namerus"),
		field.String("nameeng").MaxLen(300).StructTag(`json:"ИмяАнгл,omitempty"`).StorageKey("Nameeng"),
		field.String("synonym").MaxLen(300).StructTag(`json:"Синоним,omitempty"`).StorageKey("Synonym"),
		field.String("file").MaxLen(300).StructTag(`json:"ИмяФайла,omitempty"`).StorageKey("File"),
		field.String("type").MaxLen(300).StructTag(`json:"Тип,omitempty"`).StorageKey("Type"),
		field.String("id").MaxLen(36).NotEmpty().Unique().Immutable().StructTag(`json:"ссылка,omitempty"`).StorageKey("ID"),
	}
}

// Edges of the ДанныеФормыЭлементКоллекции.
func (MDTabel) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("mdsubsystems", MDSubSystems.Type).Ref("mdtables"),
		edge.To("mdrekvizits", MDRekvizit.Type),
		
	}
}

// MDRekvizit holds the schema definition for the MDRekvizit entity.
type MDRekvizit struct {
	ent.Schema
}

// Fields of the MDRekvizit .
func (MDRekvizit) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").MaxLen(36).NotEmpty().Unique().Immutable().StructTag(`json:"ссылка,omitempty"`).StorageKey("ID"),
		field.String("namerus").MaxLen(300).StructTag(`json:"ИмяРус,omitempty"`).StorageKey("Namerus"),
		field.String("nameeng").MaxLen(300).StructTag(`json:"ИмяАнгл,omitempty"`).StorageKey("Nameeng"),
		field.String("synonym").MaxLen(300).StructTag(`json:"Синоним,omitempty"`).StorageKey("Synonym"),
		field.String("por").MaxLen(300).StructTag(`json:"ПорядокВывода,omitempty"`).StorageKey("por"),
		field.Float("widthElem").StructTag(`json:"ШиринаЭлемента,omitempty"`),
		field.Float("widthSpisok").StructTag(`json:"ШиринаКолонки,omitempty"`),
		field.String("type").MaxLen(300).StructTag(`json:"Тип,omitempty"`).StorageKey("Type"),
		field.String("owner_id").MaxLen(36).NotEmpty().Optional(),
	}
}

// Edges of the ДанныеФормыЭлементКоллекции.
func (MDRekvizit) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", MDTabel.Type).Unique().Ref("mdrekvizits").Field("owner_id"),
	}
}

// MDSubSystems holds the schema definition for the MDSubSystems entity.
type MDSubSystems struct {
	ent.Schema
}

// Fields of the MDSubSystems .
func (MDSubSystems) Fields() []ent.Field {
	return []ent.Field{
		field.String("namerus").MaxLen(300).StructTag(`json:"ИмяРус,omitempty"`).StorageKey("Namerus"),
		field.String("nameeng").MaxLen(300).StructTag(`json:"ИмяАнгл,omitempty"`).StorageKey("Nameeng"),
		field.String("synonym").MaxLen(150).StructTag(`json:"Синоним,omitempty"`).StorageKey("Synonym"),
		field.String("id").MaxLen(36).NotEmpty().Unique().Immutable().StructTag(`json:"ссылка,omitempty"`).StorageKey("ID"),
		field.String("parent").Optional().StructTag(`json:"Родитель,omitempty"`).StorageKey("Parent"),
	}
}

// Edges of the ДанныеФормыЭлементКоллекции.
func (MDSubSystems) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("child_mdsubsystems", MDSubSystems.Type),
		edge.From("parent_mdsubsystems", MDSubSystems.Type).Ref("child_mdsubsystems").Unique().Field("parent").StructTag(`json:"родитель,omitempty"`),
		edge.To("mdtables", MDTabel.Type),
	}
}
