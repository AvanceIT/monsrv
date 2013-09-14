/*
Package secure provides encryption methods.

These methods are used to handle encrypted streams such as the login process for
the client.
*/
package secure

import (
	"fmt"
	"net"
)

// Function ClientAuth handles the server side of the client authentication
// process. It is a challenge/response process similar to the method used
// in OpenSSH
func ClientAuth(c net.Conn) bool {
	var banner string = "Monsrv\n\n"

	fmt.Fprintf(c, banner)

	return true
}
