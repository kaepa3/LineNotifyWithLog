build:
	go build main.go

pi:
	GOOS=linux GOARCH=arm GOARM=6 go build -o linenotify main.go
