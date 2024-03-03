package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
)

// DeviceGroup holds the schema definition for the DeviceGroup entity.
type DeviceGroup struct {
	ent.Schema
}

// Fields of the DeviceGroup.
func (DeviceGroup) Fields() []ent.Field {
	return []ent.Field{}
}

// Edges of the DeviceGroup.
func (DeviceGroup) Edges() []ent.Edge {
	return nil
}

// Mixin of the DeviceGroup.
func (DeviceGroup) Mixin() []ent.Mixin {
	return nil
}

// Indexes of the DeviceGroup.
func (DeviceGroup) Indexes() []ent.Index {
	return nil
}

// Annotations of the DeviceGroup
func (DeviceGroup) Annotations() []schema.Annotation {
	return nil
}
