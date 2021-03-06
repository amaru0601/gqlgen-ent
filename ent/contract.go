// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/bug/ent/contract"
	"entgo.io/bug/ent/property"
	"entgo.io/ent/dialect/sql"
)

// Contract is the model entity for the Contract schema.
type Contract struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// StartDate holds the value of the "start_date" field.
	StartDate time.Time `json:"start_date,omitempty"`
	// EndDate holds the value of the "end_date" field.
	EndDate time.Time `json:"end_date,omitempty"`
	// PayAmount holds the value of the "pay_amount" field.
	PayAmount float64 `json:"pay_amount,omitempty"`
	// PayDate holds the value of the "pay_date" field.
	PayDate time.Time `json:"pay_date,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ContractQuery when eager-loading is set.
	Edges             ContractEdges `json:"edges"`
	property_contract *int
}

// ContractEdges holds the relations/edges for other nodes in the graph.
type ContractEdges struct {
	// Users holds the value of the users edge.
	Users []*User `json:"users,omitempty"`
	// Property holds the value of the property edge.
	Property *Property `json:"property,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// UsersOrErr returns the Users value or an error if the edge
// was not loaded in eager-loading.
func (e ContractEdges) UsersOrErr() ([]*User, error) {
	if e.loadedTypes[0] {
		return e.Users, nil
	}
	return nil, &NotLoadedError{edge: "users"}
}

// PropertyOrErr returns the Property value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ContractEdges) PropertyOrErr() (*Property, error) {
	if e.loadedTypes[1] {
		if e.Property == nil {
			// The edge property was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: property.Label}
		}
		return e.Property, nil
	}
	return nil, &NotLoadedError{edge: "property"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Contract) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case contract.FieldPayAmount:
			values[i] = new(sql.NullFloat64)
		case contract.FieldID:
			values[i] = new(sql.NullInt64)
		case contract.FieldStartDate, contract.FieldEndDate, contract.FieldPayDate:
			values[i] = new(sql.NullTime)
		case contract.ForeignKeys[0]: // property_contract
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Contract", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Contract fields.
func (c *Contract) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case contract.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			c.ID = int(value.Int64)
		case contract.FieldStartDate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field start_date", values[i])
			} else if value.Valid {
				c.StartDate = value.Time
			}
		case contract.FieldEndDate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field end_date", values[i])
			} else if value.Valid {
				c.EndDate = value.Time
			}
		case contract.FieldPayAmount:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field pay_amount", values[i])
			} else if value.Valid {
				c.PayAmount = value.Float64
			}
		case contract.FieldPayDate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field pay_date", values[i])
			} else if value.Valid {
				c.PayDate = value.Time
			}
		case contract.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field property_contract", value)
			} else if value.Valid {
				c.property_contract = new(int)
				*c.property_contract = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryUsers queries the "users" edge of the Contract entity.
func (c *Contract) QueryUsers() *UserQuery {
	return (&ContractClient{config: c.config}).QueryUsers(c)
}

// QueryProperty queries the "property" edge of the Contract entity.
func (c *Contract) QueryProperty() *PropertyQuery {
	return (&ContractClient{config: c.config}).QueryProperty(c)
}

// Update returns a builder for updating this Contract.
// Note that you need to call Contract.Unwrap() before calling this method if this Contract
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Contract) Update() *ContractUpdateOne {
	return (&ContractClient{config: c.config}).UpdateOne(c)
}

// Unwrap unwraps the Contract entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Contract) Unwrap() *Contract {
	tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Contract is not a transactional entity")
	}
	c.config.driver = tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Contract) String() string {
	var builder strings.Builder
	builder.WriteString("Contract(")
	builder.WriteString(fmt.Sprintf("id=%v", c.ID))
	builder.WriteString(", start_date=")
	builder.WriteString(c.StartDate.Format(time.ANSIC))
	builder.WriteString(", end_date=")
	builder.WriteString(c.EndDate.Format(time.ANSIC))
	builder.WriteString(", pay_amount=")
	builder.WriteString(fmt.Sprintf("%v", c.PayAmount))
	builder.WriteString(", pay_date=")
	builder.WriteString(c.PayDate.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Contracts is a parsable slice of Contract.
type Contracts []*Contract

func (c Contracts) config(cfg config) {
	for _i := range c {
		c[_i].config = cfg
	}
}
