package cfgmysqldemo

import "github.com/bitwormhole/starter/application"

// ExportConfigForMySQLDemo ...
func ExportConfigForMySQLDemo(cb application.ConfigBuilder) error {
	// return nil
	return autoGenConfig(cb)
}
