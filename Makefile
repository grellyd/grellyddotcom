prd:
	hugo
	go build

dev:
	hugo --buildDrafts --buildFuture
	go run grellyddotcom.go

deploy:
	echo "Not implemented yet"
