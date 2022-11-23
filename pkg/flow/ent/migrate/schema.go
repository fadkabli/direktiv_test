// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// CloudEventsColumns holds the columns for the "cloud_events" table.
	CloudEventsColumns = []*schema.Column{
		{Name: "oid", Type: field.TypeUUID},
		{Name: "event_id", Type: field.TypeString},
		{Name: "event", Type: field.TypeJSON},
		{Name: "fire", Type: field.TypeTime},
		{Name: "created", Type: field.TypeTime},
		{Name: "processed", Type: field.TypeBool},
		{Name: "namespace_cloudevents", Type: field.TypeUUID},
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
		Indexes: []*schema.Index{
			{
				Name:    "cloudevents_event_id_namespace_cloudevents",
				Unique:  true,
				Columns: []*schema.Column{CloudEventsColumns[1], CloudEventsColumns[6]},
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
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "instance_eventlisteners", Type: field.TypeUUID, Nullable: true},
		{Name: "namespace_namespacelisteners", Type: field.TypeUUID},
		{Name: "workflow_wfevents", Type: field.TypeUUID},
	}
	// EventsTable holds the schema information for the "events" table.
	EventsTable = &schema.Table{
		Name:       "events",
		Columns:    EventsColumns,
		PrimaryKey: []*schema.Column{EventsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "events_instances_eventlisteners",
				Columns:    []*schema.Column{EventsColumns[7]},
				RefColumns: []*schema.Column{InstancesColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "events_namespaces_namespacelisteners",
				Columns:    []*schema.Column{EventsColumns[8]},
				RefColumns: []*schema.Column{NamespacesColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "events_workflows_wfevents",
				Columns:    []*schema.Column{EventsColumns[9]},
				RefColumns: []*schema.Column{WorkflowsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// EventsWaitsColumns holds the columns for the "events_waits" table.
	EventsWaitsColumns = []*schema.Column{
		{Name: "oid", Type: field.TypeUUID},
		{Name: "events", Type: field.TypeJSON},
		{Name: "events_wfeventswait", Type: field.TypeUUID},
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
		{Name: "expandedType", Type: field.TypeString, Nullable: true},
		{Name: "read_only", Type: field.TypeBool, Nullable: true, Default: false},
		{Name: "inode_children", Type: field.TypeUUID, Nullable: true},
		{Name: "namespace_inodes", Type: field.TypeUUID},
	}
	// InodesTable holds the schema information for the "inodes" table.
	InodesTable = &schema.Table{
		Name:       "inodes",
		Columns:    InodesColumns,
		PrimaryKey: []*schema.Column{InodesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "inodes_inodes_children",
				Columns:    []*schema.Column{InodesColumns[8]},
				RefColumns: []*schema.Column{InodesColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "inodes_namespaces_inodes",
				Columns:    []*schema.Column{InodesColumns[9]},
				RefColumns: []*schema.Column{NamespacesColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "inode_name_inode_children",
				Unique:  true,
				Columns: []*schema.Column{InodesColumns[3], InodesColumns[8]},
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
		{Name: "invoker", Type: field.TypeString, Nullable: true},
		{Name: "namespace_instances", Type: field.TypeUUID},
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
				Columns:    []*schema.Column{InstancesColumns[9]},
				RefColumns: []*schema.Column{NamespacesColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "instances_revisions_instances",
				Columns:    []*schema.Column{InstancesColumns[10]},
				RefColumns: []*schema.Column{RevisionsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "instances_workflows_instances",
				Columns:    []*schema.Column{InstancesColumns[11]},
				RefColumns: []*schema.Column{WorkflowsColumns[0]},
				OnDelete:   schema.SetNull,
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
		{Name: "metadata", Type: field.TypeString, Nullable: true},
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
				Columns:    []*schema.Column{InstanceRuntimesColumns[14]},
				RefColumns: []*schema.Column{InstancesColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "instance_runtimes_instances_children",
				Columns:    []*schema.Column{InstanceRuntimesColumns[15]},
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
		{Name: "mirror_activity_logs", Type: field.TypeUUID, Nullable: true},
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
				Symbol:     "log_msgs_mirror_activities_logs",
				Columns:    []*schema.Column{LogMsgsColumns[4]},
				RefColumns: []*schema.Column{MirrorActivitiesColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "log_msgs_namespaces_logs",
				Columns:    []*schema.Column{LogMsgsColumns[5]},
				RefColumns: []*schema.Column{NamespacesColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "log_msgs_workflows_logs",
				Columns:    []*schema.Column{LogMsgsColumns[6]},
				RefColumns: []*schema.Column{WorkflowsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// MirrorsColumns holds the columns for the "mirrors" table.
	MirrorsColumns = []*schema.Column{
		{Name: "oid", Type: field.TypeUUID},
		{Name: "url", Type: field.TypeString},
		{Name: "ref", Type: field.TypeString},
		{Name: "cron", Type: field.TypeString},
		{Name: "public_key", Type: field.TypeString},
		{Name: "private_key", Type: field.TypeString},
		{Name: "passphrase", Type: field.TypeString},
		{Name: "commit", Type: field.TypeString},
		{Name: "last_sync", Type: field.TypeTime, Nullable: true},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
		{Name: "inode_mirror", Type: field.TypeUUID, Unique: true, Nullable: true},
		{Name: "namespace_mirrors", Type: field.TypeUUID},
	}
	// MirrorsTable holds the schema information for the "mirrors" table.
	MirrorsTable = &schema.Table{
		Name:       "mirrors",
		Columns:    MirrorsColumns,
		PrimaryKey: []*schema.Column{MirrorsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "mirrors_inodes_mirror",
				Columns:    []*schema.Column{MirrorsColumns[10]},
				RefColumns: []*schema.Column{InodesColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "mirrors_namespaces_mirrors",
				Columns:    []*schema.Column{MirrorsColumns[11]},
				RefColumns: []*schema.Column{NamespacesColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// MirrorActivitiesColumns holds the columns for the "mirror_activities" table.
	MirrorActivitiesColumns = []*schema.Column{
		{Name: "oid", Type: field.TypeUUID},
		{Name: "type", Type: field.TypeString},
		{Name: "status", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "end_at", Type: field.TypeTime, Nullable: true},
		{Name: "controller", Type: field.TypeString, Nullable: true},
		{Name: "deadline", Type: field.TypeTime, Nullable: true},
		{Name: "mirror_activities", Type: field.TypeUUID, Nullable: true},
		{Name: "namespace_mirror_activities", Type: field.TypeUUID},
	}
	// MirrorActivitiesTable holds the schema information for the "mirror_activities" table.
	MirrorActivitiesTable = &schema.Table{
		Name:       "mirror_activities",
		Columns:    MirrorActivitiesColumns,
		PrimaryKey: []*schema.Column{MirrorActivitiesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "mirror_activities_mirrors_activities",
				Columns:    []*schema.Column{MirrorActivitiesColumns[8]},
				RefColumns: []*schema.Column{MirrorsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "mirror_activities_namespaces_mirror_activities",
				Columns:    []*schema.Column{MirrorActivitiesColumns[9]},
				RefColumns: []*schema.Column{NamespacesColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// NamespacesColumns holds the columns for the "namespaces" table.
	NamespacesColumns = []*schema.Column{
		{Name: "oid", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "config", Type: field.TypeString, Default: "\n{\n\t\"broadcast\": {\n\t  \"workflow.create\": false,\n\t  \"workflow.update\": false,\n\t  \"workflow.delete\": false,\n\t  \"directory.create\": false,\n\t  \"directory.delete\": false,\n\t  \"workflow.variable.create\": false,\n\t  \"workflow.variable.update\": false,\n\t  \"workflow.variable.delete\": false,\n\t  \"namespace.variable.create\": false,\n\t  \"namespace.variable.update\": false,\n\t  \"namespace.variable.delete\": false,\n\t  \"instance.variable.create\": false,\n\t  \"instance.variable.update\": false,\n\t  \"instance.variable.delete\": false,\n\t  \"instance.started\": false,\n\t  \"instance.success\": false,\n\t  \"instance.failed\": false\n\t}\n  }\n"},
		{Name: "name", Type: field.TypeString, Unique: true, Size: 64},
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
		{Name: "created_at", Type: field.TypeTime},
		{Name: "revision_refs", Type: field.TypeUUID},
		{Name: "workflow_refs", Type: field.TypeUUID},
	}
	// RefsTable holds the schema information for the "refs" table.
	RefsTable = &schema.Table{
		Name:       "refs",
		Columns:    RefsColumns,
		PrimaryKey: []*schema.Column{RefsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "refs_revisions_refs",
				Columns:    []*schema.Column{RefsColumns[4]},
				RefColumns: []*schema.Column{RevisionsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "refs_workflows_refs",
				Columns:    []*schema.Column{RefsColumns[5]},
				RefColumns: []*schema.Column{WorkflowsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "ref_name_workflow_refs",
				Unique:  true,
				Columns: []*schema.Column{RefsColumns[2], RefsColumns[5]},
			},
		},
	}
	// RevisionsColumns holds the columns for the "revisions" table.
	RevisionsColumns = []*schema.Column{
		{Name: "oid", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "hash", Type: field.TypeString},
		{Name: "source", Type: field.TypeBytes},
		{Name: "metadata", Type: field.TypeJSON},
		{Name: "workflow_revisions", Type: field.TypeUUID},
	}
	// RevisionsTable holds the schema information for the "revisions" table.
	RevisionsTable = &schema.Table{
		Name:       "revisions",
		Columns:    RevisionsColumns,
		PrimaryKey: []*schema.Column{RevisionsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "revisions_workflows_revisions",
				Columns:    []*schema.Column{RevisionsColumns[5]},
				RefColumns: []*schema.Column{WorkflowsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// RoutesColumns holds the columns for the "routes" table.
	RoutesColumns = []*schema.Column{
		{Name: "oid", Type: field.TypeUUID},
		{Name: "weight", Type: field.TypeInt},
		{Name: "ref_routes", Type: field.TypeUUID},
		{Name: "workflow_routes", Type: field.TypeUUID},
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
	// ServicesColumns holds the columns for the "services" table.
	ServicesColumns = []*schema.Column{
		{Name: "oid", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "url", Type: field.TypeString},
		{Name: "name", Type: field.TypeString, Unique: true},
		{Name: "data", Type: field.TypeString},
		{Name: "namespace_services", Type: field.TypeUUID},
	}
	// ServicesTable holds the schema information for the "services" table.
	ServicesTable = &schema.Table{
		Name:       "services",
		Columns:    ServicesColumns,
		PrimaryKey: []*schema.Column{ServicesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "services_namespaces_services",
				Columns:    []*schema.Column{ServicesColumns[6]},
				RefColumns: []*schema.Column{NamespacesColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "services_name_namespace_services",
				Unique:  true,
				Columns: []*schema.Column{ServicesColumns[4], ServicesColumns[6]},
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
		{Name: "mime_type", Type: field.TypeString, Default: "application/json"},
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
		{Name: "behaviour", Type: field.TypeString, Nullable: true},
		{Name: "instance_vars", Type: field.TypeUUID, Nullable: true},
		{Name: "namespace_vars", Type: field.TypeUUID, Nullable: true},
		{Name: "var_data_varrefs", Type: field.TypeUUID},
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
				Columns:    []*schema.Column{VarRefsColumns[3]},
				RefColumns: []*schema.Column{InstancesColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "var_refs_namespaces_vars",
				Columns:    []*schema.Column{VarRefsColumns[4]},
				RefColumns: []*schema.Column{NamespacesColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "var_refs_var_data_varrefs",
				Columns:    []*schema.Column{VarRefsColumns[5]},
				RefColumns: []*schema.Column{VarDataColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "var_refs_workflows_vars",
				Columns:    []*schema.Column{VarRefsColumns[6]},
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
		{Name: "read_only", Type: field.TypeBool, Nullable: true},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
		{Name: "inode_workflow", Type: field.TypeUUID, Unique: true, Nullable: true},
		{Name: "namespace_workflows", Type: field.TypeUUID},
	}
	// WorkflowsTable holds the schema information for the "workflows" table.
	WorkflowsTable = &schema.Table{
		Name:       "workflows",
		Columns:    WorkflowsColumns,
		PrimaryKey: []*schema.Column{WorkflowsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "workflows_inodes_workflow",
				Columns:    []*schema.Column{WorkflowsColumns[5]},
				RefColumns: []*schema.Column{InodesColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "workflows_namespaces_workflows",
				Columns:    []*schema.Column{WorkflowsColumns[6]},
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
		MirrorsTable,
		MirrorActivitiesTable,
		NamespacesTable,
		RefsTable,
		RevisionsTable,
		RoutesTable,
		ServicesTable,
		VarDataTable,
		VarRefsTable,
		WorkflowsTable,
	}
)

func init() {
	CloudEventsTable.ForeignKeys[0].RefTable = NamespacesTable
	EventsTable.ForeignKeys[0].RefTable = InstancesTable
	EventsTable.ForeignKeys[1].RefTable = NamespacesTable
	EventsTable.ForeignKeys[2].RefTable = WorkflowsTable
	EventsWaitsTable.ForeignKeys[0].RefTable = EventsTable
	InodesTable.ForeignKeys[0].RefTable = InodesTable
	InodesTable.ForeignKeys[1].RefTable = NamespacesTable
	InstancesTable.ForeignKeys[0].RefTable = NamespacesTable
	InstancesTable.ForeignKeys[1].RefTable = RevisionsTable
	InstancesTable.ForeignKeys[2].RefTable = WorkflowsTable
	InstanceRuntimesTable.ForeignKeys[0].RefTable = InstancesTable
	InstanceRuntimesTable.ForeignKeys[1].RefTable = InstancesTable
	LogMsgsTable.ForeignKeys[0].RefTable = InstancesTable
	LogMsgsTable.ForeignKeys[1].RefTable = MirrorActivitiesTable
	LogMsgsTable.ForeignKeys[2].RefTable = NamespacesTable
	LogMsgsTable.ForeignKeys[3].RefTable = WorkflowsTable
	MirrorsTable.ForeignKeys[0].RefTable = InodesTable
	MirrorsTable.ForeignKeys[1].RefTable = NamespacesTable
	MirrorActivitiesTable.ForeignKeys[0].RefTable = MirrorsTable
	MirrorActivitiesTable.ForeignKeys[1].RefTable = NamespacesTable
	RefsTable.ForeignKeys[0].RefTable = RevisionsTable
	RefsTable.ForeignKeys[1].RefTable = WorkflowsTable
	RevisionsTable.ForeignKeys[0].RefTable = WorkflowsTable
	RoutesTable.ForeignKeys[0].RefTable = RefsTable
	RoutesTable.ForeignKeys[1].RefTable = WorkflowsTable
	ServicesTable.ForeignKeys[0].RefTable = NamespacesTable
	VarRefsTable.ForeignKeys[0].RefTable = InstancesTable
	VarRefsTable.ForeignKeys[1].RefTable = NamespacesTable
	VarRefsTable.ForeignKeys[2].RefTable = VarDataTable
	VarRefsTable.ForeignKeys[3].RefTable = WorkflowsTable
	WorkflowsTable.ForeignKeys[0].RefTable = InodesTable
	WorkflowsTable.ForeignKeys[1].RefTable = NamespacesTable
}
