package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// MDTabel holds the schema definition for the MDTabel entity.
type MDTypeTabel struct {
	ent.Schema
}

// Fields of the MDTypeTabel .
func (MDTypeTabel) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").MaxLen(36).NotEmpty().Unique().Immutable().StructTag(`json:"ссылка,omitempty"`).StorageKey("ID"),
		field.String("nameeng").MaxLen(300).StructTag(`json:"ИмяАнгл,omitempty"`).StorageKey("Nameeng"),
		field.String("synonym").MaxLen(300).StructTag(`json:"ИмяРус,omitempty"`).StorageKey("Synonym"),
		field.String("por").StructTag(`json:"Порядок,omitempty"`).StorageKey("Por"),
		field.String("parent").Optional().StructTag(`json:"Родитель,omitempty"`).StorageKey("Parent"),
	}
}

// Edges of the ДанныеФормыЭлементКоллекции.
func (MDTypeTabel) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("child_mdtypetabels", MDTypeTabel.Type),
		edge.From("parent_mdtypetabels", MDTypeTabel.Type).Ref("child_mdtypetabels").Unique().Field("parent").StructTag(`json:"родитель,omitempty"`),

		edge.To("mdtypetabels", MDTabel.Type),	
	}
}

// MDTabel holds the schema definition for the MDTabel entity.
type MDTabel struct {
	ent.Schema
}


// Fields of the MDTabel .
func (MDTabel) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").MaxLen(36).NotEmpty().Unique().Immutable().StructTag(`json:"ссылка,omitempty"`).StorageKey("ID"),
		field.String("nameeng").MaxLen(300).StructTag(`json:"ИмяАнгл,omitempty"`).StorageKey("Nameeng"),
		field.String("synonym").MaxLen(300).StructTag(`json:"Синоним,omitempty"`).StorageKey("Synonym"),
		field.String("por").StructTag(`json:"Порядок,omitempty"`).StorageKey("Por"),
		field.String("parent").Optional().StructTag(`json:"Родитель,omitempty"`).StorageKey("Parent"),
		field.String("types_id").MaxLen(36).Optional(),

		field.String("file").MaxLen(300).StructTag(`json:"ИмяФайла,omitempty"`).StorageKey("File"),

	}
}

// Edges of the ДанныеФормыЭлементКоллекции.
func (MDTabel) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("child_mdtabel", MDTabel.Type),
		edge.From("parent_mdtabel", MDTabel.Type).Ref("child_mdtabel").Unique().Field("parent").StructTag(`json:"родитель,omitempty"`),

		edge.From("mdsubsystems", MDSubSystems.Type).Ref("mdtables"),
		edge.To("mdrekvizits", MDRekvizit.Type),
		edge.From("mdtypetabel", MDTypeTabel.Type).Unique().Ref("mdtypetabels").Field("types_id"),
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
		field.String("nameeng").MaxLen(300).StructTag(`json:"ИмяАнгл,omitempty"`).StorageKey("Nameeng"),
		field.String("synonym").MaxLen(300).StructTag(`json:"Синоним,omitempty"`).StorageKey("Synonym"),
		field.String("por").MaxLen(300).StructTag(`json:"ПорядокВывода,omitempty"`).StorageKey("por"),
		field.String("type").MaxLen(300).StructTag(`json:"Тип,omitempty"`).StorageKey("Type"),
		field.String("owner_id").MaxLen(36).NotEmpty().Optional(),

		field.Float("widthSpisok").StructTag(`json:"ШиринаКолонки,omitempty"`),
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
		field.String("id").MaxLen(36).NotEmpty().Unique().Immutable().StructTag(`json:"ссылка,omitempty"`).StorageKey("ID"),
		field.String("nameeng").MaxLen(300).StructTag(`json:"ИмяАнгл,omitempty"`).StorageKey("Nameeng"),
		field.String("synonym").MaxLen(150).StructTag(`json:"Синоним,omitempty"`).StorageKey("Synonym"),
		field.String("por").StructTag(`json:"Порядок,omitempty"`).StorageKey("Por"),		
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

// MDTabel holds the schema definition for the MDTabel entity.
type MDForms struct {
	ent.Schema
}

// Fields of the MDTypeTabel .
func (MDForms) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").MaxLen(36).NotEmpty().Unique().Immutable().StructTag(`json:"ссылка,omitempty"`).StorageKey("ID"),
		field.String("idform").MaxLen(36).StructTag(`json:"Родитель,omitempty"`).StorageKey("idform"),		
		field.String("conteiner"),
		field.String("parent").MaxLen(36).Optional().StructTag(`json:"Родитель,omitempty"`).StorageKey("Parent"),
	}
}

// Edges of the ДанныеФормыЭлементКоллекции.
func (MDForms) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("child_mdforms", MDForms.Type),
		edge.From("parent_mdforms", MDForms.Type).Ref("child_mdforms").Unique().Field("parent").StructTag(`json:"родитель,omitempty"`),
	}
}