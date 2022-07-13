package file

import (
	"bufio"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func Read1(path string) {
	fi, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer fi.Close()
	buf := make([]byte, 1024)
	for {
		n, err := fi.Read(buf)
		if err != nil && err != io.EOF {
			log.Fatal(err)
		}
		if 0 == n {
			break
		}
	}
}

func Read2(path string) {
	fi, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer fi.Close()
	r := bufio.NewReader(fi)
	buf := make([]byte, 1024)
	for {
		n, err := r.Read(buf)
		if err != nil && err != io.EOF {
			log.Fatal(err)
		}
		if 0 == n {
			break
		}
	}
}

func Read3(path string) {
	fi, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer fi.Close()
	_, err = ioutil.ReadAll(fi)
}
