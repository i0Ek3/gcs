# gcs

> Go Client-Server model.

Simple Client-Server model, and to know how this work.   


## Run

You can download [release]() to run it directly, or clone this repo then build it:

```Shell
$ git clone git@github.com:i0Ek3/gcs.git
$ cd gcs/src/$v  // $v is version dir
$ go build client.go ; go build server.go
$ ./server
$ ./client
```



## Log

2018-01-09: v1, client send data to server only.



## Structure

```Shell

____src      // go files
      |____v1
      |____v2
      ...
      |____vx

____bin      // excutable files
      |____v1
      |____v2
      ...
      |____vx

```


## Note

Credit for [Unknown](https://github.com/Unknwon).
Reference: [https://github.com/Unknwon/the-way-to-go_ZH_CN](https://github.com/Unknwon/the-way-to-go_ZH_CN)




