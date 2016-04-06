all:
	go build --ldflags "-X main.Builddate=`date -u '+%Y-%m-%d_%I:%M:%S%p'` -X main.Githash=`git rev-parse HEAD`"
