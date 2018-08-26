# CIn/UFPE: Master Degree Class [IN1118](https://sites.google.com/a/cin.ufpe.br/in1118/home)

## Exercise #1 Middleware Based on TCP/UDP

The middleware consists of a tokenizer, which receives a text and produces the unique word tokens.

```shell
$ go run [PROTOCOL]/client/client.go --help
Usage of client:
  -text string
        Text to be tokenize
```

Next sections there are instructions of usage for both protocols.

### TCP
1. Open a terminal with the server
2. Use a second terminal to request the TCP server

```shell
$ cd [WORD-TOKENIZE-MIDDLEWARE-SOCKET]

# Bring up the TCP server
$ go run tcp/server/server.go

# Request the server
$ go run tcp/client/client.go --text "TEXT TO BE TOKENIZED"
```

### UDP
1. Open a terminal with the server
2. Use a second terminal to request the UDP server

```shell
$ cd [WORD-TOKENIZE-MIDDLEWARE-SOCKET]

# Bring up the TCP server
$ go run udp/server/server.go

# Request the server
$ go run udp/client/client.go --text "TEXT TO BE TOKENIZED"
```