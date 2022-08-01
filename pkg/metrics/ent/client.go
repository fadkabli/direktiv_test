// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"log"

	"github.com/direktiv/direktiv/pkg/metrics/ent/migrate"

	"github.com/direktiv/direktiv/pkg/metrics/ent/metrics"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Metrics is the client for interacting with the Metrics builders.
	Metrics *MetricsClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Metrics = NewMetricsClient(c.config)
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:     ctx,
		config:  cfg,
		Metrics: NewMetricsClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		ctx:     ctx,
		config:  cfg,
		Metrics: NewMetricsClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Metrics.
//		Query().
//		Count(ctx)
//
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.Metrics.Use(hooks...)
}

// MetricsClient is a client for the Metrics schema.
type MetricsClient struct {
	config
}

// NewMetricsClient returns a client for the Metrics from the given config.
func NewMetricsClient(c config) *MetricsClient {
	return &MetricsClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `metrics.Hooks(f(g(h())))`.
func (c *MetricsClient) Use(hooks ...Hook) {
	c.hooks.Metrics = append(c.hooks.Metrics, hooks...)
}

// Create returns a builder for creating a Metrics entity.
func (c *MetricsClient) Create() *MetricsCreate {
	mutation := newMetricsMutation(c.config, OpCreate)
	return &MetricsCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Metrics entities.
func (c *MetricsClient) CreateBulk(builders ...*MetricsCreate) *MetricsCreateBulk {
	return &MetricsCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Metrics.
func (c *MetricsClient) Update() *MetricsUpdate {
	mutation := newMetricsMutation(c.config, OpUpdate)
	return &MetricsUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *MetricsClient) UpdateOne(m *Metrics) *MetricsUpdateOne {
	mutation := newMetricsMutation(c.config, OpUpdateOne, withMetrics(m))
	return &MetricsUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *MetricsClient) UpdateOneID(id int) *MetricsUpdateOne {
	mutation := newMetricsMutation(c.config, OpUpdateOne, withMetricsID(id))
	return &MetricsUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Metrics.
func (c *MetricsClient) Delete() *MetricsDelete {
	mutation := newMetricsMutation(c.config, OpDelete)
	return &MetricsDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *MetricsClient) DeleteOne(m *Metrics) *MetricsDeleteOne {
	return c.DeleteOneID(m.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *MetricsClient) DeleteOneID(id int) *MetricsDeleteOne {
	builder := c.Delete().Where(metrics.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &MetricsDeleteOne{builder}
}

// Query returns a query builder for Metrics.
func (c *MetricsClient) Query() *MetricsQuery {
	return &MetricsQuery{
		config: c.config,
	}
}

// Get returns a Metrics entity by its id.
func (c *MetricsClient) Get(ctx context.Context, id int) (*Metrics, error) {
	return c.Query().Where(metrics.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *MetricsClient) GetX(ctx context.Context, id int) *Metrics {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *MetricsClient) Hooks() []Hook {
	return c.hooks.Metrics
}
