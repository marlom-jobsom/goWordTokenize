# CIn/UFPE: Master Degree Class [IN1118](https://sites.google.com/a/cin.ufpe.br/in1118/home)

## Middleware Based on TCP, UDP and RPC
The middleware consists of a tokenizer, which receives a text and produces the unique word tokens.

### CLI Client
```shell
$ go run cmd/client/main.go --help
Usage of /tmp/go-build217990265/b001/exe/main:
  -protocol string
    	The client_request_handler network prototol (tcp or udp) (default "tcp")
  -rpc
    	Uses RPC client_request_handler (over TCP only)
  -text string
    	Text to be tokenize (default "Silent is is golden")
```

### CLI Server
```shell
$ go run cmd/server/main.go --help
Usage of /tmp/go-build010583863/b001/exe/main:
  -protocol string
    	The server_request_handler network prototol (tcp or udp) (default "tcp")
  -rpc
    	Uses RPC server_request_handler (over TCP only)
```

## User case example
1. Open a terminal and starts the server over the desired communication method. (let's use RPC over TCP as example)
2. Open another terminal and execute a request for the RPC over TCP as well given a text to be tokenize

```shell
# Bring up the server
$ go run cmd/server/main.go -protocol tcp -rpc
2018/10/09 20:19:12 Bring up RPC server over TCP
2018/10/09 20:19:12 Address [::]:5000
```

```shell
# Client request
$ go run cmd/client/main.go -protocol tcp -rpc -text "The text to be tokenized"
2018/10/09 20:20:29 Sending request (rpc): The text to be tokenized
2018/10/09 20:20:29 Receiving response (rpc): {[to be tokenized The text] 407.443µs}
```
