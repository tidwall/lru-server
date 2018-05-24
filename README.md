# `lru-server`

A convenient [LRU cache](https://en.wikipedia.org/wiki/Cache_replacement_policies#Least_recently_used_(LRU)) that's built on [hashicorp/golang-lru](https://github.com/hashicorp/golang-lru) and [tidwall/modern-server](https://github.com/tidwall/modern-server). It uses a very simple HTTP REST api and supports Let's Encrypt. 

## Building

To start using lru-server, install Go and run:

```sh
$ make
```

## Using

Start the server up.

```sh
$ ./lru-server
$ ./lru-server -s 100000    # specify an lru capacity of 100,000
```

### Set a key

```sh
$ curl -d "my value" localhost:8000/mykey
```

### Get a key

```sh
$ curl localhost:8000/mykey
my value
```

### Delete a key

```sh
$ curl -X DELETE localhost:8000/mykey
```

## Contact

Josh Baker [@tidwall](http://twitter.com/tidwall)

## License

`lru-server` source code is available under the MIT [License](/LICENSE).

