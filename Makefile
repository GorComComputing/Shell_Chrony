.PHONY: main
main: *.go deps
	GOOS=linux GOARCH=arm go build -o ChShell .


.PHONY:deps
deps:
#	go get github.com/gorilla/sessions



