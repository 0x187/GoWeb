# Go Web
Make Any Shell Command as a Web API with Golang

Clients invoke the web API by sending HTTP GET and POST requests. Clients can also send in additional flags and arguments to be passed into the command/script wrapped within the web API. Result of the command/script execution is sent back to the client as a plain text payload.

## Running 

As an example, assume you need to expose the "date" command as a web API. You can simply run the tool as follows:

```bash
./GoWeb -p 8080 -b "curl -L google.com"
```
Now, the clients can invoke the API by sending an HTTP request to http://0.0.0.0:8080


You can also use this tool to expose custom shell scripts and other command-line programs. For example, if you have a Python script foo.py which you wish to expose as a web API, all you have to do is:
```bash
./GoWeb -p 8080 -b "python3 Mio.py"
```
## Last words

<img src="https://raw.githubusercontent.com/0x187/ClearText/main/68747470733a2f2f692e696d6775722e636f6d2f774d34553835682e6a7067.jpg">
