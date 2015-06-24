package goma

import (
	"database/sql"
)

// Goma is sql.DB access wrapper.
type Goma struct {
	*sql.DB
	options Options
}

// Open is create goma client.
// - database open
func Open(configPath string) (*sql.DB, error) {
	opts, err := NewOptions(configPath)
	if err != nil {
		return nil, err
	}
	return OpenOptions(opts)
}

// OpenOptions is create goma client.
// - database open
func OpenOptions(options Options) (*sql.DB, error) {
	return sql.Open(options.Driver, options.Source())
}

// Close sql.DB close.
func (d *Goma) Close() error {
	err := d.DB.Close()
	return err
}
