package foo // import "webridge.dev/foo"

import "fmt"

// Bar returns a welcome message to name.
func Bar(name string) string {
	return fmt.Sprintf("Hello %s!", name)
}
