package cfgmysql

import "github.com/bitwormhole/starter/application"

// ExportConfigForMySQL ...
func ExportConfigForMySQL(cb application.ConfigBuilder) error {
	// return nil
	return autoGenConfig(cb)
}
