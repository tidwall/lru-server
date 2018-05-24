lru-server: main.go
	go build -o lru-server
clean:
	rm -f lru-server
install: lru-server
	cp -f lru-server /usr/local/bin/lru-server
uninstall:
	rm -f /usr/local/bin/lru-server
