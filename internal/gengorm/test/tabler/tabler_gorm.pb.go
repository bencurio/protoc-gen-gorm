// Code generated by protoc-gen-gorm. DO NOT EDIT.
// versions:
// 	protoc-gen-gorm v0.0.0
// 	protoc          (unknown)
// source: tabler/tabler.proto

package tabler

import (
	_ "github.com/complex64/protoc-gen-gorm/gormpb"
)

// ImplementsTablerModel is the GORM model for tabler.ImplementsTabler.
type ImplementsTablerModel struct {
}

// AsProto converts a ImplementsTablerModel to its protobuf representation.
func (m *ImplementsTablerModel) AsProto() (*ImplementsTabler, error) {
	x := new(ImplementsTabler)
	return x, nil
}

// AsModel converts a ImplementsTabler to its GORM model.
func (x *ImplementsTabler) AsModel() (*ImplementsTablerModel, error) {
	m := new(ImplementsTablerModel)
	return m, nil
}

func (m *ImplementsTablerModel) TableName() string {
	return "name"
}