build:
	go build ./cmd/server-app/
	go build ./cmd/client-app/

clean:
	rm server-app
	rm client-app