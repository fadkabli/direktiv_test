// Code generated by ent, DO NOT EDIT.

package namespace

import (
	"time"

	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the namespace type in the database.
	Label = "namespace"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "oid"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldConfig holds the string denoting the config field in the database.
	FieldConfig = "config"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// EdgeInodes holds the string denoting the inodes edge name in mutations.
	EdgeInodes = "inodes"
	// EdgeWorkflows holds the string denoting the workflows edge name in mutations.
	EdgeWorkflows = "workflows"
	// EdgeMirrors holds the string denoting the mirrors edge name in mutations.
	EdgeMirrors = "mirrors"
	// EdgeMirrorActivities holds the string denoting the mirror_activities edge name in mutations.
	EdgeMirrorActivities = "mirror_activities"
	// EdgeInstances holds the string denoting the instances edge name in mutations.
	EdgeInstances = "instances"
	// EdgeLogs holds the string denoting the logs edge name in mutations.
	EdgeLogs = "logs"
	// EdgeVars holds the string denoting the vars edge name in mutations.
	EdgeVars = "vars"
	// EdgeCloudevents holds the string denoting the cloudevents edge name in mutations.
	EdgeCloudevents = "cloudevents"
	// EdgeNamespacelisteners holds the string denoting the namespacelisteners edge name in mutations.
	EdgeNamespacelisteners = "namespacelisteners"
	// EdgeCloudeventfilters holds the string denoting the cloudeventfilters edge name in mutations.
	EdgeCloudeventfilters = "cloudeventfilters"
	// CloudEventFiltersFieldID holds the string denoting the ID field of the CloudEventFilters.
	CloudEventFiltersFieldID = "id"
	// Table holds the table name of the namespace in the database.
	Table = "namespaces"
	// InodesTable is the table that holds the inodes relation/edge.
	InodesTable = "inodes"
	// InodesInverseTable is the table name for the Inode entity.
	// It exists in this package in order to avoid circular dependency with the "inode" package.
	InodesInverseTable = "inodes"
	// InodesColumn is the table column denoting the inodes relation/edge.
	InodesColumn = "namespace_inodes"
	// WorkflowsTable is the table that holds the workflows relation/edge.
	WorkflowsTable = "workflows"
	// WorkflowsInverseTable is the table name for the Workflow entity.
	// It exists in this package in order to avoid circular dependency with the "workflow" package.
	WorkflowsInverseTable = "workflows"
	// WorkflowsColumn is the table column denoting the workflows relation/edge.
	WorkflowsColumn = "namespace_workflows"
	// MirrorsTable is the table that holds the mirrors relation/edge.
	MirrorsTable = "mirrors"
	// MirrorsInverseTable is the table name for the Mirror entity.
	// It exists in this package in order to avoid circular dependency with the "mirror" package.
	MirrorsInverseTable = "mirrors"
	// MirrorsColumn is the table column denoting the mirrors relation/edge.
	MirrorsColumn = "namespace_mirrors"
	// MirrorActivitiesTable is the table that holds the mirror_activities relation/edge.
	MirrorActivitiesTable = "mirror_activities"
	// MirrorActivitiesInverseTable is the table name for the MirrorActivity entity.
	// It exists in this package in order to avoid circular dependency with the "mirroractivity" package.
	MirrorActivitiesInverseTable = "mirror_activities"
	// MirrorActivitiesColumn is the table column denoting the mirror_activities relation/edge.
	MirrorActivitiesColumn = "namespace_mirror_activities"
	// InstancesTable is the table that holds the instances relation/edge.
	InstancesTable = "instances"
	// InstancesInverseTable is the table name for the Instance entity.
	// It exists in this package in order to avoid circular dependency with the "instance" package.
	InstancesInverseTable = "instances"
	// InstancesColumn is the table column denoting the instances relation/edge.
	InstancesColumn = "namespace_instances"
	// LogsTable is the table that holds the logs relation/edge.
	LogsTable = "log_msgs"
	// LogsInverseTable is the table name for the LogMsg entity.
	// It exists in this package in order to avoid circular dependency with the "logmsg" package.
	LogsInverseTable = "log_msgs"
	// LogsColumn is the table column denoting the logs relation/edge.
	LogsColumn = "namespace_logs"
	// VarsTable is the table that holds the vars relation/edge.
	VarsTable = "var_refs"
	// VarsInverseTable is the table name for the VarRef entity.
	// It exists in this package in order to avoid circular dependency with the "varref" package.
	VarsInverseTable = "var_refs"
	// VarsColumn is the table column denoting the vars relation/edge.
	VarsColumn = "namespace_vars"
	// CloudeventsTable is the table that holds the cloudevents relation/edge.
	CloudeventsTable = "cloud_events"
	// CloudeventsInverseTable is the table name for the CloudEvents entity.
	// It exists in this package in order to avoid circular dependency with the "cloudevents" package.
	CloudeventsInverseTable = "cloud_events"
	// CloudeventsColumn is the table column denoting the cloudevents relation/edge.
	CloudeventsColumn = "namespace_cloudevents"
	// NamespacelistenersTable is the table that holds the namespacelisteners relation/edge.
	NamespacelistenersTable = "events"
	// NamespacelistenersInverseTable is the table name for the Events entity.
	// It exists in this package in order to avoid circular dependency with the "events" package.
	NamespacelistenersInverseTable = "events"
	// NamespacelistenersColumn is the table column denoting the namespacelisteners relation/edge.
	NamespacelistenersColumn = "namespace_namespacelisteners"
	// CloudeventfiltersTable is the table that holds the cloudeventfilters relation/edge.
	CloudeventfiltersTable = "cloud_event_filters"
	// CloudeventfiltersInverseTable is the table name for the CloudEventFilters entity.
	// It exists in this package in order to avoid circular dependency with the "cloudeventfilters" package.
	CloudeventfiltersInverseTable = "cloud_event_filters"
	// CloudeventfiltersColumn is the table column denoting the cloudeventfilters relation/edge.
	CloudeventfiltersColumn = "namespace_cloudeventfilters"
)

// Columns holds all SQL columns for namespace fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldConfig,
	FieldName,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
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
	// DefaultConfig holds the default value on creation for the "config" field.
	DefaultConfig string
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)
