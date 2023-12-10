all/dev: edits push/dev remote/deploy/dev remote/logs/dev

push/dev:
	git push

edits:
	git add -u
	git commit -m "Edits"

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
	rm -f /var/http/grellyddotcom && \
	rm -rf /var/http/public && \
	rm -rf /var/http/templates && \
	cp -r ./public /var/http && \
	go build grellyddotcom.go && \
	cp grellyddotcom /var/http/ && \
	cp -r ./templates /var/http/ && \
	ls /var/http/ && \
	go version && \
	systemctl restart grellyddotcom.service && \
	systemctl status grellyddotcom.service
endef

deploy/dev:
	$(deploy-dev)

remote/dev:
	ssh root@dev.grellyd.com

remote/deploy/dev:
	ssh root@dev.grellyd.com "cd grellyddotcom && $(deploy-dev)"

define deploy-prd
	git pull origin prd -f && \
	git submodule update --init --recursive && \
	hugo && \
	rm -f /var/http/grellyddotcom && \
	rm -rf /var/http/public && \
	rm -rf /var/http/templates && \
	cp -r ./public /var/http && \
	go build grellyddotcom.go && \
	cp grellyddotcom /var/http/ && \
	cp -r ./templates /var/http/ && \
	ls /var/http/ && \
	go version && \
	systemctl restart grellyddotcom.service && \
	systemctl status grellyddotcom.service
endef

deploy/prd:
	$(deploy-dev)

remote/prd:
	ssh root@grellyd.com

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

new/writing:
	NAME ?= $(shell bash -c 'read -s -p "Name: " name; echo $$name'')
	hugo new content writing/$(NAME)