prd:
	hugo
	go build

dev:
	hugo --buildDrats --buildFuture
	go run grellyddotcom.go

deploy:
	echo "Not implemented yet"