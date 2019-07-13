package main

import (
	"bufio"
	"flag"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"
)

func handleConn(conn net.Conn, dest string) {
	defer conn.Close()
	destConn, err := net.DialTimeout("tcp", dest, time.Second)
	if err != nil {
		log.Println("connect to dest failed", err)
		return
	}
	id := strconv.FormatInt(time.Now().UnixNano(), 36)
	log.Printf("%v %v begin [%v <-> %v]: %v", id, strings.Repeat("=", 20),
		conn.RemoteAddr(), destConn.RemoteAddr(), strings.Repeat("=", 20))
	defer destConn.Close()
	p := func(src, dest *net.TCPConn, wg *sync.WaitGroup, prefix string) {
		defer wg.Done()
		defer dest.CloseWrite()
		defer src.CloseRead()
		pr, pw := io.Pipe()
		defer pw.Close()
		mw := io.MultiWriter(dest, pw)
		go func() {
			br := bufio.NewReader(pr)
			for {
				line, err := br.ReadBytes('\n')
				if err != nil {
					if err == io.EOF {
						if len(line) != 0 {
							log.Printf("%v %v %v", id, prefix, string(line))
						}
						return
					}
					log.Printf("%v ERROR %v", id, err)
					return
				}
				log.Printf("%v %v %v", id, prefix, string(line))
			}
		}()
		_, err := io.Copy(mw, src)
		if err != nil {
			log.Println(err)
		}

	}
	wg := &sync.WaitGroup{}
	wg.Add(2)
	go p(conn.(*net.TCPConn), destConn.(*net.TCPConn), wg, "-->")
	go p(destConn.(*net.TCPConn), conn.(*net.TCPConn), wg, "<--")
	wg.Wait()
	log.Printf("%v %v end   [%v <-> %v]: %v", id, strings.Repeat("=", 20),
		conn.RemoteAddr(), destConn.RemoteAddr(), strings.Repeat("=", 20))
}

func startFakeHttpServer() (addr string) {
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		log.Panicln(err)
	}
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		io.Copy(ioutil.Discard, r.Body)
		rw.WriteHeader(200)
	})
	go http.Serve(ln, nil)
	return ln.Addr().String()
}

func main() {
	log.SetFlags(log.Lmicroseconds | log.Ldate)
	listen := flag.String("l", ":9999", "listen addr")
	dest := flag.String("d", "localhost:8080", "dest addr")
	flag.Parse()
	if *dest == "" {
		*dest = startFakeHttpServer()
		log.Println("use fake http server", *dest)
	}
	conn, err := net.DialTimeout("tcp", *dest, time.Second)
	if err != nil {
		log.Panicln("cannot connect to dest: ", *dest, err)
	}
	conn.Close()
	log.Printf("http dump server running at %v and proxy to %v", *listen, *dest)
	nl, err := net.Listen("tcp", *listen)
	if err != nil {
		log.Panicln(err)
	}
	for {
		conn, err := nl.Accept()
		if err != nil {
			log.Panicln(err)
		}
		go handleConn(conn, *dest)
	}
}
