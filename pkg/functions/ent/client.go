// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/direktiv/direktiv/pkg/functions/ent/migrate"

	"github.com/direktiv/direktiv/pkg/functions/ent/services"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Services is the client for interacting with the Services builders.
	Services *ServicesClient
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
	c.Services = NewServicesClient(c.config)
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
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:      ctx,
		config:   cfg,
		Services: NewServicesClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
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
		ctx:      ctx,
		config:   cfg,
		Services: NewServicesClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Services.
//		Query().
//		Count(ctx)
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
	c.Services.Use(hooks...)
}

// ServicesClient is a client for the Services schema.
type ServicesClient struct {
	config
}

// NewServicesClient returns a client for the Services from the given config.
func NewServicesClient(c config) *ServicesClient {
	return &ServicesClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `services.Hooks(f(g(h())))`.
func (c *ServicesClient) Use(hooks ...Hook) {
	c.hooks.Services = append(c.hooks.Services, hooks...)
}

// Create returns a builder for creating a Services entity.
func (c *ServicesClient) Create() *ServicesCreate {
	mutation := newServicesMutation(c.config, OpCreate)
	return &ServicesCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Services entities.
func (c *ServicesClient) CreateBulk(builders ...*ServicesCreate) *ServicesCreateBulk {
	return &ServicesCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Services.
func (c *ServicesClient) Update() *ServicesUpdate {
	mutation := newServicesMutation(c.config, OpUpdate)
	return &ServicesUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ServicesClient) UpdateOne(s *Services) *ServicesUpdateOne {
	mutation := newServicesMutation(c.config, OpUpdateOne, withServices(s))
	return &ServicesUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ServicesClient) UpdateOneID(id int) *ServicesUpdateOne {
	mutation := newServicesMutation(c.config, OpUpdateOne, withServicesID(id))
	return &ServicesUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Services.
func (c *ServicesClient) Delete() *ServicesDelete {
	mutation := newServicesMutation(c.config, OpDelete)
	return &ServicesDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *ServicesClient) DeleteOne(s *Services) *ServicesDeleteOne {
	return c.DeleteOneID(s.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *ServicesClient) DeleteOneID(id int) *ServicesDeleteOne {
	builder := c.Delete().Where(services.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ServicesDeleteOne{builder}
}

// Query returns a query builder for Services.
func (c *ServicesClient) Query() *ServicesQuery {
	return &ServicesQuery{
		config: c.config,
	}
}

// Get returns a Services entity by its id.
func (c *ServicesClient) Get(ctx context.Context, id int) (*Services, error) {
	return c.Query().Where(services.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ServicesClient) GetX(ctx context.Context, id int) *Services {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *ServicesClient) Hooks() []Hook {
	return c.hooks.Services
}
