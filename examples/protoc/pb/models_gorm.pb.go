// Code generated by protoc-gen-gorm. DO NOT EDIT.
// versions:
// 	protoc-gen-gorm v0.0.0
// 	protoc          v3.19.3
// source: models.proto

package pb

import (
	_ "github.com/complex64/protoc-gen-gorm/gormpb"
)

// UserModel is the GORM model for pb.User.
type UserModel struct {
	Name string `gorm:"not null;unique;primaryKey"`
}

// AsProto converts a UserModel to its protobuf representation.
func (m *UserModel) AsProto() (*User, error) {
	x := new(User)
	x.Name = m.Name
	return x, nil
}

// AsModel converts a User to its GORM model.
func (x *User) AsModel() (*UserModel, error) {
	m := new(UserModel)
	m.Name = x.Name
	return m, nil
}