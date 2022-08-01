// Code generated by ent, DO NOT EDIT.

package instanceruntime

import (
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the instanceruntime type in the database.
	Label = "instance_runtime"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "oid"
	// FieldInput holds the string denoting the input field in the database.
	FieldInput = "input"
	// FieldData holds the string denoting the data field in the database.
	FieldData = "data"
	// FieldController holds the string denoting the controller field in the database.
	FieldController = "controller"
	// FieldMemory holds the string denoting the memory field in the database.
	FieldMemory = "memory"
	// FieldFlow holds the string denoting the flow field in the database.
	FieldFlow = "flow"
	// FieldOutput holds the string denoting the output field in the database.
	FieldOutput = "output"
	// FieldStateBeginTime holds the string denoting the statebegintime field in the database.
	FieldStateBeginTime = "state_begin_time"
	// FieldDeadline holds the string denoting the deadline field in the database.
	FieldDeadline = "deadline"
	// FieldAttempts holds the string denoting the attempts field in the database.
	FieldAttempts = "attempts"
	// FieldCallerData holds the string denoting the caller_data field in the database.
	FieldCallerData = "caller_data"
	// FieldInstanceContext holds the string denoting the instancecontext field in the database.
	FieldInstanceContext = "instance_context"
	// FieldStateContext holds the string denoting the statecontext field in the database.
	FieldStateContext = "state_context"
	// FieldMetadata holds the string denoting the metadata field in the database.
	FieldMetadata = "metadata"
	// EdgeInstance holds the string denoting the instance edge name in mutations.
	EdgeInstance = "instance"
	// EdgeCaller holds the string denoting the caller edge name in mutations.
	EdgeCaller = "caller"
	// Table holds the table name of the instanceruntime in the database.
	Table = "instance_runtimes"
	// InstanceTable is the table that holds the instance relation/edge.
	InstanceTable = "instance_runtimes"
	// InstanceInverseTable is the table name for the Instance entity.
	// It exists in this package in order to avoid circular dependency with the "instance" package.
	InstanceInverseTable = "instances"
	// InstanceColumn is the table column denoting the instance relation/edge.
	InstanceColumn = "instance_runtime"
	// CallerTable is the table that holds the caller relation/edge.
	CallerTable = "instance_runtimes"
	// CallerInverseTable is the table name for the Instance entity.
	// It exists in this package in order to avoid circular dependency with the "instance" package.
	CallerInverseTable = "instances"
	// CallerColumn is the table column denoting the caller relation/edge.
	CallerColumn = "instance_children"
)

// Columns holds all SQL columns for instanceruntime fields.
var Columns = []string{
	FieldID,
	FieldInput,
	FieldData,
	FieldController,
	FieldMemory,
	FieldFlow,
	FieldOutput,
	FieldStateBeginTime,
	FieldDeadline,
	FieldAttempts,
	FieldCallerData,
	FieldInstanceContext,
	FieldStateContext,
	FieldMetadata,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "instance_runtimes"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"instance_runtime",
	"instance_children",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)
