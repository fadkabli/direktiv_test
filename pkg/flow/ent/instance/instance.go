// Code generated by ent, DO NOT EDIT.

package instance

import (
	"time"

	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the instance type in the database.
	Label = "instance"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "oid"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldEndAt holds the string denoting the end_at field in the database.
	FieldEndAt = "end_at"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldAs holds the string denoting the as field in the database.
	FieldAs = "as"
	// FieldErrorCode holds the string denoting the errorcode field in the database.
	FieldErrorCode = "error_code"
	// FieldErrorMessage holds the string denoting the errormessage field in the database.
	FieldErrorMessage = "error_message"
	// FieldInvoker holds the string denoting the invoker field in the database.
	FieldInvoker = "invoker"
	// EdgeNamespace holds the string denoting the namespace edge name in mutations.
	EdgeNamespace = "namespace"
	// EdgeWorkflow holds the string denoting the workflow edge name in mutations.
	EdgeWorkflow = "workflow"
	// EdgeRevision holds the string denoting the revision edge name in mutations.
	EdgeRevision = "revision"
	// EdgeLogs holds the string denoting the logs edge name in mutations.
	EdgeLogs = "logs"
	// EdgeVars holds the string denoting the vars edge name in mutations.
	EdgeVars = "vars"
	// EdgeRuntime holds the string denoting the runtime edge name in mutations.
	EdgeRuntime = "runtime"
	// EdgeChildren holds the string denoting the children edge name in mutations.
	EdgeChildren = "children"
	// EdgeEventlisteners holds the string denoting the eventlisteners edge name in mutations.
	EdgeEventlisteners = "eventlisteners"
	// Table holds the table name of the instance in the database.
	Table = "instances"
	// NamespaceTable is the table that holds the namespace relation/edge.
	NamespaceTable = "instances"
	// NamespaceInverseTable is the table name for the Namespace entity.
	// It exists in this package in order to avoid circular dependency with the "namespace" package.
	NamespaceInverseTable = "namespaces"
	// NamespaceColumn is the table column denoting the namespace relation/edge.
	NamespaceColumn = "namespace_instances"
	// WorkflowTable is the table that holds the workflow relation/edge.
	WorkflowTable = "instances"
	// WorkflowInverseTable is the table name for the Workflow entity.
	// It exists in this package in order to avoid circular dependency with the "workflow" package.
	WorkflowInverseTable = "workflows"
	// WorkflowColumn is the table column denoting the workflow relation/edge.
	WorkflowColumn = "workflow_instances"
	// RevisionTable is the table that holds the revision relation/edge.
	RevisionTable = "instances"
	// RevisionInverseTable is the table name for the Revision entity.
	// It exists in this package in order to avoid circular dependency with the "revision" package.
	RevisionInverseTable = "revisions"
	// RevisionColumn is the table column denoting the revision relation/edge.
	RevisionColumn = "revision_instances"
	// LogsTable is the table that holds the logs relation/edge.
	LogsTable = "log_msgs"
	// LogsInverseTable is the table name for the LogMsg entity.
	// It exists in this package in order to avoid circular dependency with the "logmsg" package.
	LogsInverseTable = "log_msgs"
	// LogsColumn is the table column denoting the logs relation/edge.
	LogsColumn = "instance_logs"
	// VarsTable is the table that holds the vars relation/edge.
	VarsTable = "var_refs"
	// VarsInverseTable is the table name for the VarRef entity.
	// It exists in this package in order to avoid circular dependency with the "varref" package.
	VarsInverseTable = "var_refs"
	// VarsColumn is the table column denoting the vars relation/edge.
	VarsColumn = "instance_vars"
	// RuntimeTable is the table that holds the runtime relation/edge.
	RuntimeTable = "instance_runtimes"
	// RuntimeInverseTable is the table name for the InstanceRuntime entity.
	// It exists in this package in order to avoid circular dependency with the "instanceruntime" package.
	RuntimeInverseTable = "instance_runtimes"
	// RuntimeColumn is the table column denoting the runtime relation/edge.
	RuntimeColumn = "instance_runtime"
	// ChildrenTable is the table that holds the children relation/edge.
	ChildrenTable = "instance_runtimes"
	// ChildrenInverseTable is the table name for the InstanceRuntime entity.
	// It exists in this package in order to avoid circular dependency with the "instanceruntime" package.
	ChildrenInverseTable = "instance_runtimes"
	// ChildrenColumn is the table column denoting the children relation/edge.
	ChildrenColumn = "instance_children"
	// EventlistenersTable is the table that holds the eventlisteners relation/edge.
	EventlistenersTable = "events"
	// EventlistenersInverseTable is the table name for the Events entity.
	// It exists in this package in order to avoid circular dependency with the "events" package.
	EventlistenersInverseTable = "events"
	// EventlistenersColumn is the table column denoting the eventlisteners relation/edge.
	EventlistenersColumn = "instance_eventlisteners"
)

// Columns holds all SQL columns for instance fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldEndAt,
	FieldStatus,
	FieldAs,
	FieldErrorCode,
	FieldErrorMessage,
	FieldInvoker,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "instances"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"namespace_instances",
	"revision_instances",
	"workflow_instances",
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
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)
