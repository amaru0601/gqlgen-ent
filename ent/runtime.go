// Code generated by entc, DO NOT EDIT.

package ent

import (
	"entgo.io/bug/ent/schema"
	"entgo.io/bug/ent/user"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	propertyFields := schema.Property{}.Fields()
	_ = propertyFields
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescNames is the schema descriptor for names field.
	userDescNames := userFields[0].Descriptor()
	// user.NamesValidator is a validator for the "names" field. It is called by the builders before save.
	user.NamesValidator = userDescNames.Validators[0].(func(string) error)
	// userDescLastnames is the schema descriptor for lastnames field.
	userDescLastnames := userFields[1].Descriptor()
	// user.LastnamesValidator is a validator for the "lastnames" field. It is called by the builders before save.
	user.LastnamesValidator = userDescLastnames.Validators[0].(func(string) error)
	// userDescEmail is the schema descriptor for email field.
	userDescEmail := userFields[3].Descriptor()
	// user.EmailValidator is a validator for the "email" field. It is called by the builders before save.
	user.EmailValidator = userDescEmail.Validators[0].(func(string) error)
}
