package main

import (
	"bufio"
	"fmt"
	"github.com/fisher310/goder/fib"
	"net"
	"os"
)

func writeFile(filename string) {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	f := fib.Fibonacci()
	for i := 0; i < 20; i++ {
		fmt.Fprintln(writer, f())
	}
}

func main() {
	conn, err := net.Dial("tcp", "www.baidu.com:80")
	if err != nil {
		panic(err)
	}
	n, err := fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
	if err != nil {
		panic(err)
	}
	fmt.Println(n)
	status, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		panic(err)
	}
	fmt.Println(status)

}
