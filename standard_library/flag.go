package standardlibrary

import(
	"fmt"
	"flag"
)

func FlagFunctions(){
	var username *string = flag.String("username", "root", "database username")
	var password *string = flag.String("password", "", "database password")
	var address *string = flag.String("address", "localhost", "database address")
	var port *int = flag.Int("port", 0, "database port")


	fmt.Println("Username : ", *username)
	fmt.Println("Password : ", *password)
	fmt.Println("Address  : ", *address)
	fmt.Println("Port     : ", *port)
}