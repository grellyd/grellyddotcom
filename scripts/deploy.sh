#! /bin/bash

set -e

ENV=unknwon

while [[ $# -gt 0 ]]; do
	case "$1" in 
		-d|--dev)
			ENV="dev"
			;;
        -p|--prd)
            ENV="prd"
            ;;
		*)
			echo "Unknown option: $1"
			exit 1
			;;
	esac
done

if [[ $ENV = "unknown" ]]; then
    echo "Unknown env"
fi


git pull origin prd -f
git submodule update --init --recursive
case $ENV in 
    "dev")
     hugo --buildDrafts --buildFuture
     ;;
    "prd")
     hugo
     ;;
esac
rm -f /var/http/grellyddotcom
rm -rf /var/http/public
rm -rf /var/http/templates
cp -r ./public /var/http
CGO_ENABLED=1 go build grellyddotcom.go
cp grellyddotcom /var/http/
cp -r ./templates /var/http/
ls /var/http/
go version
systemctl restart grellyddotcom.service
systemctl status grellyddotcom.service
