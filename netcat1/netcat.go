// Netcat1 is a read-only TCP client.
package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	mustCopy(os.Stdout, conn)
	//netcat1

	// defer conn.Close()
	// go mustCopy(os.Stdout, conn)
	// mustCopy(conn, os.Stdin)
	//netcat2

	// done := make(chan struct{})
	// go func() {
	//     io.Copy(os.Stdout, conn) // NOTE: ignoring errors
	//     log.Println("done")
	//     done <- struct{}{} // signal the main goroutine
	// }()
	// mustCopy(conn, os.Stdin)
	// conn.Close()
	// <-done // wait for background goroutine to finish
	//netcat3
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
