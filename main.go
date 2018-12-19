package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

type options struct {
	address string
	dir     string
}

func main() {
	opt := options{
		address: "127.0.0.1:8000",
		dir:     "./",
	}
	flag.Usage = func() {
		fmt.Println(os.Args[0], " [address] [dir] [option]")
		flag.PrintDefaults()
	}

	flag.StringVar(&opt.dir, "dir", "./", "root serve path")
	flag.Parse()

	if flag.NArg() > 0 {
		opt.address = flag.Arg(0)
	}
	if flag.NArg() > 1 {
		opt.dir = flag.Arg(1)
	}

	log.Println(opt)
	if err := http.ListenAndServe(opt.address, http.FileServer(http.Dir(opt.dir))); err != nil {
		log.Fatalln(err)
	}
}
