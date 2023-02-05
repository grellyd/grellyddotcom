prd:
	hugo
	go run grellyddotcom.go

dev:
	hugo --buildDrafts --buildFuture
	go run grellyddotcom.go


# ssh root@dev.grellyd.com
# cd grellyddotcom
deploy-dev:
	git pull origin dev -f
	git submodule update --init --recursive
	hugo --buildDrafts --buildFuture
	cp -rf public /var/http
	go install
	systemctl restart grellyddotcom.service
	systemctl status grellyddotcom.service

# ssh root@grellyd.com
# cd grellyddotcom
deploy-prod:
	git pull origin prod -f
	git submodule update --init --recursive
	hugo
	cp -rf public /var/http
	go install
	systemctl restart grellyddotcom.service
	systemctl status grellyddotcom.service


logs-dev:
	ssh root@dev.grellyd.com
	journalctl -u grellyddotcom.service --no-pager -f

logs-prod:
	ssh root@grellyd.com
	journalctl -u grellyddotcom.service --no-pager -f
