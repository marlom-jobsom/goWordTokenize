# CIn/UFPE: Master Degree Class [IN1118](https://sites.google.com/a/cin.ufpe.br/in1118/home)

## Exercise #1 Middleware Based on TCP/UDP
The middleware consists of a tokenizer, which receives a text and produces the unique word tokens.

```shell
$ go run client/client.go --help
Usage of client:
  -protocol string
    	The client network prototol (tcp or udp)
  -text string
    	Text to be tokenize
```

Next sections there are instructions of usage for both protocols.

### TCP and UDP servers
```shell
$ cd [WORD-TOKENIZE-MIDDLEWARE-SOCKET]

# Bring up both TCP and UDP servers
$ go run server/tcp/server.go
$ go run server/udp/server.go
```

### Request Servers With Client
```shell
$ cd [WORD-TOKENIZE-MIDDLEWARE-SOCKET]

# Request the server
$ go run client/client.go --protocol "PROTOCOL" --text "TEXT TO BE TOKENIZED"
```
