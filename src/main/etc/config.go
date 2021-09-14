package etc

import "github.com/bitwormhole/starter/application"

func ExportMySQLDriverConfig(cb application.ConfigBuilder) error {

	return autoGenConfig(cb)
}
