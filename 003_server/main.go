package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panic(err)
	}

	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Fatalln(err)
		}

		go handle(conn)
	}
}

// func handle(conn net.Conn) {
// 	err := conn.SetDeadline(time.Now().Add(10 * time.Second))
// 	if err != nil {
// 		log.Println("CONN TIMEOUT")
// 	}
// 	scanner := bufio.NewScanner(conn)
// 	for scanner.Scan() {
// 		ln := scanner.Text()
// 		fmt.Println(ln)
// 		fmt.Fprintf(conn, "I heard you say: %s\n", ln)
// 	}
// 	defer conn.Close()

// 	fmt.Println("Code got Here.")
// }

func handle(conn net.Conn) {
	defer conn.Close()
	request(conn)
	response(conn)
}

func request(conn net.Conn) {
	i := 0
	scanner := bufio.NewScanner(conn)
	hdr := make(map[string]string)
	for scanner.Scan() {
		ln := scanner.Text()
		// fmt.Println(ln)
		if i == 0 {
			m := strings.Fields(ln)[0]
			fmt.Println("***METHOD", m)
		}
		if i > 0 && ln != "" {
			hdrSplit := strings.Split(ln, ":")
			hdr[hdrSplit[0]] = hdrSplit[1]
		}
		if ln == "" {
			//headers are done break
			for key, val := range hdr {
				fmt.Println(key, " : ", val)
			}
			break
		}
		i++
	}
}

func response(conn net.Conn) {
	body := `<!DOCTYPE html><html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Hello World!</title>
</head>
<body>
<h1>Hello World</h1>
</body>
</html>`
	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}
