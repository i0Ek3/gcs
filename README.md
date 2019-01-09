# gcs

> Go Client-Server model.

Simple Client-Server model, and to know how this work.   


## Run

You can download [release](https://github.com/i0Ek3/gcs/releases) file to run it directly, or clone this repo then build it:

```Shell
$ git clone git@github.com:i0Ek3/gcs.git
$ cd gcs/src/
$ go build client.go ; go build server.go
$ ./server
$ ./client
```


## Log

2018-01-09: v1.0, client send data to server only.





## Structure

```Shell

____src      // go files
      |____client.go
      |____server.go
      |____dial.go
      |____socket.go
      ...
      

____bin      // excutable files, un-upload
      |____client
      |____server
      ...


```


## Note

Credit for [Unknown](https://github.com/Unknwon).

Reference: [https://github.com/Unknwon/the-way-to-go_ZH_CN](https://github.com/Unknwon/the-way-to-go_ZH_CN)




