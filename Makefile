prd:
	hugo
	go run grellyddotcom.go

dev:
	hugo --buildDrafts --buildFuture
	go run grellyddotcom.go


deploy-dev:
	ssh root@dev.grellyd.com
	cd grellyddotcom
	git pull origin dev -f
	hugo --buildDrafts --buildFuture
	go install
	systemctl restart grellyddotcom.service

deploy-prod:
	ssh root@grellyd.com
	cd grellyddotcom
	git pull origin prod -f
	hugo
	go install
	systemctl restart grellyddotcom.service


logs-dev:
	ssh root@dev.grellyd.com
	journalctl -u grellyddotcom.service --no-pager -f

logs-prod:
	ssh root@grellyd.com
	journalctl -u grellyddotcom.service --no-pager -f
