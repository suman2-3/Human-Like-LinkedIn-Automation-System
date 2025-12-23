package messaging

import "fmt"

func RenderTemplate(name string) string {
	return fmt.Sprintf(
		"Hi %s, thanks for connecting!",
		name,
	)
}
