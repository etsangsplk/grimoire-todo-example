package changeset

import (
	"html"
)

// EscapeString escapes special characters like "<" to become "&lt;". this is helper for html.EscapeString
func EscapeString(ch *Changeset, field string) {
	ApplyString(ch, field, html.EscapeString)
}
