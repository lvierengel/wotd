default:

deploy:
	GOOS=linux GOARCH=amd64 go build -o /tmp/wotd-linux-amd64 ./cmd/wotd
	ssh -T root@describe.today "service wotd stop"
	scp /tmp/wotd-linux-amd64 root@describe.today:/usr/local/bin/wotd
	ssh -T root@describe.today "service wotd start"
