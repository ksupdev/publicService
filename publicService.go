package publicService

import "fmt"

func SayHi(name string) string {
    // Return a greeting that embeds the name in a message.
    message := fmt.Sprintf("Hi, %v. My name's publicService", name)
    return message
}