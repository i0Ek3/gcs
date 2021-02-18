# gcs

A simple Go Client-Server model demo.   

## Process

- Parse host and port
- Initial server
- Listen port and accept connection
- Error checking

## How-To

You can download [release](https://github.com/i0Ek3/gcs/releases) file to run it directly, or clone this repo then build it:

```Shell
$ git clone git@github.com:i0Ek3/gcs.git
$ cd gcs/
$ go build client.go ; go build server.go
$ ./server 127.0.0.1 50000
$ ./client 50000
```

## Screenshot

![](https://github.com/i0Ek3/gcs/blob/master/pic/example.jpg)

## To-Do

- [x] Expand C/S model
- [x] More commands support
- [x] ...


## Reference

- [https://github.com/Unknwon/the-way-to-go_ZH_CN](https://github.com/Unknwon/the-way-to-go_ZH_CN)




