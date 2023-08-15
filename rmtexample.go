package rmtexample

import "fmt"

func GenerateHacktivMessage(name string) string {
	return fmt.Sprintf("Welcome to hacktiv8 %s", name)
}

func PrintWelcomeMessage(name string) {
	fmt.Println(GenerateHacktivMessage(name))
}
