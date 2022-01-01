package parse

import "flag"

func Flags() (username, password *string, port *int) {
	username = flag.String("user", "", "database client username:password")
	password = flag.String("pass", "", "database client password")
	port = flag.Int("port", 8000, "server port number")

	flag.Parse()
	return
}
