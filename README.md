# Gexec

Gexec is a powerful tool that allows you to convert any shell command into a web API using the Go programming language.

Clients can interact with the web API by sending HTTP GET and POST requests. Additional flags and arguments can be included in the requests and will be passed into the wrapped command or script within the API. The result of the command or script execution is then returned to the client as plain text.

## Building
To build Gexec, follow these simple steps:

```bash
go build -ldflags '-s -w' -o bin/gexec src/*
```

## Usage
To get started with Gexec, use the following command to see the available options:

```bash
bin/gexec --help
```

The output will display the available flags and their descriptions:

```
Usage of bin/gexec:
  -b string
        Path to Unix command/binary that has STDOUT (default "echo helloWorld")
  -p int
        HTTP port to listen on (default 8080)
```

## Running
If no command or port is specified, Gexec will use the default command and port settings.

As an example, let's assume you want to expose the "timedatectl" command as a web API. You can achieve this by running the following command:

```bash
./bin/gexec -p 6576 -b "timedatectl"
```

Now, clients can interact with the API by sending HTTP requests to http://0.0.0.0:6576.

You can also utilize Gexec to expose custom shell scripts and other command-line programs. For instance, if you have a Python script called `foo.py` that you wish to expose as a web API, you can do so easily:

```bash
./bin/gexec -p 6576 -b "python3 foo.py"
```

Feel free to explore the capabilities of Gexec and adapt it to your specific needs. It provides a convenient and efficient way to expose command-line functionality as a web service.