// Code generated by protoc-gen-gorm. DO NOT EDIT.
// versions:
// 	protoc-gen-gorm v0.0.0
// 	protoc          (unknown)
// source: tabler/tabler.proto

package tabler

import (
	context "context"
	_ "github.com/complex64/protoc-gen-gorm/gormpb"
)

// ImplementsTablerModel is the GORM model for tabler.ImplementsTabler.
type ImplementsTablerModel struct {
}

// ToProto converts a ImplementsTablerModel to its protobuf representation.
func (m *ImplementsTablerModel) ToProto() ImplementsTabler {
	panic(true)
}

// ToModel converts a ImplementsTabler to its GORM model.
func (x *ImplementsTabler) ToModel() ImplementsTablerModel {
	panic(true)
}

func (m *ImplementsTablerModel) TableName() string {
	return "name"
}
func CreateImplementsTablerModel(ctx context.Context) {}

func GetImplementsTablerModel(ctx context.Context) {}

func ListImplementsTablerModel(ctx context.Context) {}

func UpdateImplementsTablerModel(ctx context.Context) {}

func DeleteImplementsTablerModel(ctx context.Context) {}
