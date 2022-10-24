package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

func main() {
	binary := flag.String("b", "", "Path to the executable binary")
	port := flag.Int("p", 63011, "HTTP port to listen on")
	flag.Parse()

	if *binary == "" {
		fmt.Println("Path to binary not specified.")
		return
	}

	l := log.New(os.Stdout, "", log.Ldate|log.Ltime)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var argString string
		if r.Body != nil {
			data, err := ioutil.ReadAll(r.Body)
			if err != nil {
				l.Print(err)
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			argString = string(data)
		}

		fields := strings.Fields(*binary)
		args := append(fields[1:], strings.Fields(argString)...)
		l.Printf("Command: [%s %s]", fields[0], strings.Join(args, " "))

		output, err := exec.Command(fields[0], args...).Output()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "text/plain")
		w.Write(output)
	})

	l.Printf("Listening on port %d...", *port)
	l.Printf("Exposed binary: %s", *binary)
	http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", *port), nil)
}
