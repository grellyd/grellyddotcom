build/prd:
	hugo

build/dev:
	hugo --buildDrafts --buildFuture

run/prd:
	hugo
	go run grellyddotcom.go

run/dev:
	hugo --buildDrafts --buildFuture
	go run grellyddotcom.go

define deploy-dev
	git pull origin dev -f && \
	git submodule update --init --recursive && \
	hugo --buildDrafts --buildFuture && \
	cp -rf public /var/http && \
	go install && \
	systemctl restart grellyddotcom.service && \
	systemctl status grellyddotcom.service
endef

deploy/dev:
	$(deploy-dev)

remote/deploy/dev:
	ssh root@dev.grellyd.com "cd grellyddotcom && $(deploy-dev)"

define deploy-prd
	git pull origin prd -f && \
	git submodule update --init --recursive && \
	hugo && \
	cp -rf public /var/http && \
	go install && \
	systemctl restart grellyddotcom.service && \
	systemctl status grellyddotcom.service
endef

deploy/prd:
	$(deploy-dev)

remote/deploy/prd:
	ssh root@grellyd.com "cd grellyddotcom && $(deploy-prd)"

define logs
	journalctl -u grellyddotcom.service --no-pager -f
endef

logs:
	$(logs)

remote/logs/dev:
	ssh root@dev.grellyd.com "$(logs)"

remote/logs/prd:
	ssh root@grellyd.com "$(logs)"
