build:
	go build -o smart-wallet ./...

run:
	./smart-wallet

clean:
	rm -f smart-wallet
