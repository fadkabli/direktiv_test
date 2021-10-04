// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// CloudEventsColumns holds the columns for the "cloud_events" table.
	CloudEventsColumns = []*schema.Column{
		{Name: "oid", Type: field.TypeUUID},
		{Name: "event_id", Type: field.TypeString, Unique: true},
		{Name: "event", Type: field.TypeJSON},
		{Name: "fire", Type: field.TypeTime},
		{Name: "created", Type: field.TypeTime},
		{Name: "processed", Type: field.TypeBool},
		{Name: "namespace_cloudevents", Type: field.TypeUUID, Nullable: true},
	}
	// CloudEventsTable holds the schema information for the "cloud_events" table.
	CloudEventsTable = &schema.Table{
		Name:       "cloud_events",
		Columns:    CloudEventsColumns,
		PrimaryKey: []*schema.Column{CloudEventsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "cloud_events_namespaces_cloudevents",
				Columns:    []*schema.Column{CloudEventsColumns[6]},
				RefColumns: []*schema.Column{NamespacesColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// EventsColumns holds the columns for the "events" table.
	EventsColumns = []*schema.Column{
		{Name: "oid", Type: field.TypeUUID},
		{Name: "events", Type: field.TypeJSON},
		{Name: "correlations", Type: field.TypeJSON},
		{Name: "signature", Type: field.TypeBytes, Nullable: true},
		{Name: "count", Type: field.TypeInt},
		{Name: "instance_eventlisteners", Type: field.TypeUUID, Nullable: true},
		{Name: "workflow_wfevents", Type: field.TypeUUID, Nullable: true},
	}
	// EventsTable holds the schema information for the "events" table.
	EventsTable = &schema.Table{
		Name:       "events",
		Columns:    EventsColumns,
		PrimaryKey: []*schema.Column{EventsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "events_instances_eventlisteners",
				Columns:    []*schema.Column{EventsColumns[5]},
				RefColumns: []*schema.Column{InstancesColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "events_workflows_wfevents",
				Columns:    []*schema.Column{EventsColumns[6]},
				RefColumns: []*schema.Column{WorkflowsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// EventsWaitsColumns holds the columns for the "events_waits" table.
	EventsWaitsColumns = []*schema.Column{
		{Name: "oid", Type: field.TypeUUID},
		{Name: "events", Type: field.TypeJSON},
		{Name: "events_wfeventswait", Type: field.TypeUUID, Nullable: true},
	}
	// EventsWaitsTable holds the schema information for the "events_waits" table.
	EventsWaitsTable = &schema.Table{
		Name:       "events_waits",
		Columns:    EventsWaitsColumns,
		PrimaryKey: []*schema.Column{EventsWaitsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "events_waits_events_wfeventswait",
				Columns:    []*schema.Column{EventsWaitsColumns[2]},
				RefColumns: []*schema.Column{EventsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// InodesColumns holds the columns for the "inodes" table.
	InodesColumns = []*schema.Column{
		{Name: "oid", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "name", Type: field.TypeString, Nullable: true},
		{Name: "type", Type: field.TypeString},
		{Name: "attributes", Type: field.TypeJSON, Nullable: true},
		{Name: "inode_children", Type: field.TypeUUID, Nullable: true},
		{Name: "namespace_inodes", Type: field.TypeUUID, Nullable: true},
	}
	// InodesTable holds the schema information for the "inodes" table.
	InodesTable = &schema.Table{
		Name:       "inodes",
		Columns:    InodesColumns,
		PrimaryKey: []*schema.Column{InodesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "inodes_inodes_children",
				Columns:    []*schema.Column{InodesColumns[6]},
				RefColumns: []*schema.Column{InodesColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "inodes_namespaces_inodes",
				Columns:    []*schema.Column{InodesColumns[7]},
				RefColumns: []*schema.Column{NamespacesColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "inode_name_inode_children",
				Unique:  true,
				Columns: []*schema.Column{InodesColumns[3], InodesColumns[6]},
			},
		},
	}
	// InstancesColumns holds the columns for the "instances" table.
	InstancesColumns = []*schema.Column{
		{Name: "oid", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "end_at", Type: field.TypeTime, Nullable: true},
		{Name: "status", Type: field.TypeString},
		{Name: "as", Type: field.TypeString},
		{Name: "error_code", Type: field.TypeString, Nullable: true},
		{Name: "error_message", Type: field.TypeString, Nullable: true},
		{Name: "namespace_instances", Type: field.TypeUUID, Nullable: true},
		{Name: "revision_instances", Type: field.TypeUUID, Nullable: true},
		{Name: "workflow_instances", Type: field.TypeUUID, Nullable: true},
	}
	// InstancesTable holds the schema information for the "instances" table.
	InstancesTable = &schema.Table{
		Name:       "instances",
		Columns:    InstancesColumns,
		PrimaryKey: []*schema.Column{InstancesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "instances_namespaces_instances",
				Columns:    []*schema.Column{InstancesColumns[8]},
				RefColumns: []*schema.Column{NamespacesColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "instances_revisions_instances",
				Columns:    []*schema.Column{InstancesColumns[9]},
				RefColumns: []*schema.Column{RevisionsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "instances_workflows_instances",
				Columns:    []*schema.Column{InstancesColumns[10]},
				RefColumns: []*schema.Column{WorkflowsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// InstanceRuntimesColumns holds the columns for the "instance_runtimes" table.
	InstanceRuntimesColumns = []*schema.Column{
		{Name: "oid", Type: field.TypeUUID},
		{Name: "input", Type: field.TypeBytes},
		{Name: "data", Type: field.TypeString},
		{Name: "controller", Type: field.TypeString, Nullable: true},
		{Name: "memory", Type: field.TypeString, Nullable: true},
		{Name: "flow", Type: field.TypeJSON, Nullable: true},
		{Name: "output", Type: field.TypeString, Nullable: true},
		{Name: "state_begin_time", Type: field.TypeTime, Nullable: true},
		{Name: "deadline", Type: field.TypeTime, Nullable: true},
		{Name: "attempts", Type: field.TypeInt, Nullable: true},
		{Name: "caller_data", Type: field.TypeString, Nullable: true},
		{Name: "instance_context", Type: field.TypeString, Nullable: true},
		{Name: "state_context", Type: field.TypeString, Nullable: true},
		{Name: "instance_runtime", Type: field.TypeUUID, Unique: true, Nullable: true},
		{Name: "instance_children", Type: field.TypeUUID, Nullable: true},
	}
	// InstanceRuntimesTable holds the schema information for the "instance_runtimes" table.
	InstanceRuntimesTable = &schema.Table{
		Name:       "instance_runtimes",
		Columns:    InstanceRuntimesColumns,
		PrimaryKey: []*schema.Column{InstanceRuntimesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "instance_runtimes_instances_runtime",
				Columns:    []*schema.Column{InstanceRuntimesColumns[13]},
				RefColumns: []*schema.Column{InstancesColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "instance_runtimes_instances_children",
				Columns:    []*schema.Column{InstanceRuntimesColumns[14]},
				RefColumns: []*schema.Column{InstancesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// LogMsgsColumns holds the columns for the "log_msgs" table.
	LogMsgsColumns = []*schema.Column{
		{Name: "oid", Type: field.TypeUUID},
		{Name: "t", Type: field.TypeTime},
		{Name: "msg", Type: field.TypeString},
		{Name: "instance_logs", Type: field.TypeUUID, Nullable: true},
		{Name: "namespace_logs", Type: field.TypeUUID, Nullable: true},
		{Name: "workflow_logs", Type: field.TypeUUID, Nullable: true},
	}
	// LogMsgsTable holds the schema information for the "log_msgs" table.
	LogMsgsTable = &schema.Table{
		Name:       "log_msgs",
		Columns:    LogMsgsColumns,
		PrimaryKey: []*schema.Column{LogMsgsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "log_msgs_instances_logs",
				Columns:    []*schema.Column{LogMsgsColumns[3]},
				RefColumns: []*schema.Column{InstancesColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "log_msgs_namespaces_logs",
				Columns:    []*schema.Column{LogMsgsColumns[4]},
				RefColumns: []*schema.Column{NamespacesColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "log_msgs_workflows_logs",
				Columns:    []*schema.Column{LogMsgsColumns[5]},
				RefColumns: []*schema.Column{WorkflowsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// NamespacesColumns holds the columns for the "namespaces" table.
	NamespacesColumns = []*schema.Column{
		{Name: "oid", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "name", Type: field.TypeString, Unique: true},
	}
	// NamespacesTable holds the schema information for the "namespaces" table.
	NamespacesTable = &schema.Table{
		Name:       "namespaces",
		Columns:    NamespacesColumns,
		PrimaryKey: []*schema.Column{NamespacesColumns[0]},
	}
	// RefsColumns holds the columns for the "refs" table.
	RefsColumns = []*schema.Column{
		{Name: "oid", Type: field.TypeUUID},
		{Name: "immutable", Type: field.TypeBool, Default: true},
		{Name: "name", Type: field.TypeString},
		{Name: "revision_refs", Type: field.TypeUUID, Nullable: true},
		{Name: "workflow_refs", Type: field.TypeUUID, Nullable: true},
	}
	// RefsTable holds the schema information for the "refs" table.
	RefsTable = &schema.Table{
		Name:       "refs",
		Columns:    RefsColumns,
		PrimaryKey: []*schema.Column{RefsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "refs_revisions_refs",
				Columns:    []*schema.Column{RefsColumns[3]},
				RefColumns: []*schema.Column{RevisionsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "refs_workflows_refs",
				Columns:    []*schema.Column{RefsColumns[4]},
				RefColumns: []*schema.Column{WorkflowsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "ref_name_workflow_refs",
				Unique:  true,
				Columns: []*schema.Column{RefsColumns[2], RefsColumns[4]},
			},
		},
	}
	// RevisionsColumns holds the columns for the "revisions" table.
	RevisionsColumns = []*schema.Column{
		{Name: "oid", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "hash", Type: field.TypeString},
		{Name: "source", Type: field.TypeBytes},
		{Name: "workflow_revisions", Type: field.TypeUUID, Nullable: true},
	}
	// RevisionsTable holds the schema information for the "revisions" table.
	RevisionsTable = &schema.Table{
		Name:       "revisions",
		Columns:    RevisionsColumns,
		PrimaryKey: []*schema.Column{RevisionsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "revisions_workflows_revisions",
				Columns:    []*schema.Column{RevisionsColumns[4]},
				RefColumns: []*schema.Column{WorkflowsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// RoutesColumns holds the columns for the "routes" table.
	RoutesColumns = []*schema.Column{
		{Name: "oid", Type: field.TypeUUID},
		{Name: "weight", Type: field.TypeInt},
		{Name: "ref_routes", Type: field.TypeUUID, Nullable: true},
		{Name: "workflow_routes", Type: field.TypeUUID, Nullable: true},
	}
	// RoutesTable holds the schema information for the "routes" table.
	RoutesTable = &schema.Table{
		Name:       "routes",
		Columns:    RoutesColumns,
		PrimaryKey: []*schema.Column{RoutesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "routes_refs_routes",
				Columns:    []*schema.Column{RoutesColumns[2]},
				RefColumns: []*schema.Column{RefsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "routes_workflows_routes",
				Columns:    []*schema.Column{RoutesColumns[3]},
				RefColumns: []*schema.Column{WorkflowsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// VarDataColumns holds the columns for the "var_data" table.
	VarDataColumns = []*schema.Column{
		{Name: "oid", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "size", Type: field.TypeInt},
		{Name: "hash", Type: field.TypeString},
		{Name: "data", Type: field.TypeBytes},
	}
	// VarDataTable holds the schema information for the "var_data" table.
	VarDataTable = &schema.Table{
		Name:       "var_data",
		Columns:    VarDataColumns,
		PrimaryKey: []*schema.Column{VarDataColumns[0]},
	}
	// VarRefsColumns holds the columns for the "var_refs" table.
	VarRefsColumns = []*schema.Column{
		{Name: "oid", Type: field.TypeUUID},
		{Name: "name", Type: field.TypeString, Nullable: true},
		{Name: "instance_vars", Type: field.TypeUUID, Nullable: true},
		{Name: "namespace_vars", Type: field.TypeUUID, Nullable: true},
		{Name: "var_data_varrefs", Type: field.TypeUUID, Nullable: true},
		{Name: "workflow_vars", Type: field.TypeUUID, Nullable: true},
	}
	// VarRefsTable holds the schema information for the "var_refs" table.
	VarRefsTable = &schema.Table{
		Name:       "var_refs",
		Columns:    VarRefsColumns,
		PrimaryKey: []*schema.Column{VarRefsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "var_refs_instances_vars",
				Columns:    []*schema.Column{VarRefsColumns[2]},
				RefColumns: []*schema.Column{InstancesColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "var_refs_namespaces_vars",
				Columns:    []*schema.Column{VarRefsColumns[3]},
				RefColumns: []*schema.Column{NamespacesColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "var_refs_var_data_varrefs",
				Columns:    []*schema.Column{VarRefsColumns[4]},
				RefColumns: []*schema.Column{VarDataColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "var_refs_workflows_vars",
				Columns:    []*schema.Column{VarRefsColumns[5]},
				RefColumns: []*schema.Column{WorkflowsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// WorkflowsColumns holds the columns for the "workflows" table.
	WorkflowsColumns = []*schema.Column{
		{Name: "oid", Type: field.TypeUUID},
		{Name: "live", Type: field.TypeBool, Default: true},
		{Name: "log_to_events", Type: field.TypeString, Nullable: true},
		{Name: "inode_workflow", Type: field.TypeUUID, Unique: true, Nullable: true},
		{Name: "namespace_workflows", Type: field.TypeUUID, Nullable: true},
	}
	// WorkflowsTable holds the schema information for the "workflows" table.
	WorkflowsTable = &schema.Table{
		Name:       "workflows",
		Columns:    WorkflowsColumns,
		PrimaryKey: []*schema.Column{WorkflowsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "workflows_inodes_workflow",
				Columns:    []*schema.Column{WorkflowsColumns[3]},
				RefColumns: []*schema.Column{InodesColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "workflows_namespaces_workflows",
				Columns:    []*schema.Column{WorkflowsColumns[4]},
				RefColumns: []*schema.Column{NamespacesColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		CloudEventsTable,
		EventsTable,
		EventsWaitsTable,
		InodesTable,
		InstancesTable,
		InstanceRuntimesTable,
		LogMsgsTable,
		NamespacesTable,
		RefsTable,
		RevisionsTable,
		RoutesTable,
		VarDataTable,
		VarRefsTable,
		WorkflowsTable,
	}
)

func init() {
	CloudEventsTable.ForeignKeys[0].RefTable = NamespacesTable
	EventsTable.ForeignKeys[0].RefTable = InstancesTable
	EventsTable.ForeignKeys[1].RefTable = WorkflowsTable
	EventsWaitsTable.ForeignKeys[0].RefTable = EventsTable
	InodesTable.ForeignKeys[0].RefTable = InodesTable
	InodesTable.ForeignKeys[1].RefTable = NamespacesTable
	InstancesTable.ForeignKeys[0].RefTable = NamespacesTable
	InstancesTable.ForeignKeys[1].RefTable = RevisionsTable
	InstancesTable.ForeignKeys[2].RefTable = WorkflowsTable
	InstanceRuntimesTable.ForeignKeys[0].RefTable = InstancesTable
	InstanceRuntimesTable.ForeignKeys[1].RefTable = InstancesTable
	LogMsgsTable.ForeignKeys[0].RefTable = InstancesTable
	LogMsgsTable.ForeignKeys[1].RefTable = NamespacesTable
	LogMsgsTable.ForeignKeys[2].RefTable = WorkflowsTable
	RefsTable.ForeignKeys[0].RefTable = RevisionsTable
	RefsTable.ForeignKeys[1].RefTable = WorkflowsTable
	RevisionsTable.ForeignKeys[0].RefTable = WorkflowsTable
	RoutesTable.ForeignKeys[0].RefTable = RefsTable
	RoutesTable.ForeignKeys[1].RefTable = WorkflowsTable
	VarRefsTable.ForeignKeys[0].RefTable = InstancesTable
	VarRefsTable.ForeignKeys[1].RefTable = NamespacesTable
	VarRefsTable.ForeignKeys[2].RefTable = VarDataTable
	VarRefsTable.ForeignKeys[3].RefTable = WorkflowsTable
	WorkflowsTable.ForeignKeys[0].RefTable = InodesTable
	WorkflowsTable.ForeignKeys[1].RefTable = NamespacesTable
}