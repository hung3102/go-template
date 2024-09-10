package environ

import "os"

// IsLocal - local environment
func IsLocal() bool {
	return os.Getenv("IS_LOCAL") == "true"
}
