package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

func main() {
	binary := flag.String("b", "echo HelloWorld", "Path to the Unix command/binary that outputs to STDOUT")
	port := flag.Int("p", 8080, "HTTP port to listen on")
	help := flag.Bool("h", false, "Show command-line help")
	flag.Parse()

	if *help {
		flag.Usage()
		return
	}

	if *binary == "" {
		log.Println("Path to Unix command/binary not specified.")
		return
	}

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var argString string
		if r.Body != nil {
			defer r.Body.Close()
			data, err := io.ReadAll(r.Body)
			if err != nil {
				logger.Print(err)
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			argString = string(data)
		}

		fields := strings.Fields(*binary)
		args := append(fields, strings.Fields(argString)...)
		logger.Printf("Command: [%s %s]", args[0], strings.Join(args[1:], " "))

		output, err := exec.Command(args[0], args[1:]...).Output()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "text/plain")
		_, err = w.Write(output)
		if err != nil {
			logger.Print(err)
		}
	})

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [options]\n", os.Args[0])
		fmt.Fprintln(os.Stderr, "Options:")
		flag.PrintDefaults()
	}

	logger.Printf("Listening on port %d...", *port)
	logger.Printf("Exposed binary: %s", *binary)
	err := http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", *port), nil)
	if err != nil {
		logger.Fatal(err)
	}
}
