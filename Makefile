# Makefile for botdemo

.PHONY: sls/init
sls/init: sls/build
	docker run --rm -v $(PWD):/app -it lambda bash -c "sls create -t aws-go -p botdemo; \
		cp assets/Dockerfile assets/Makefile assets/goinit.sh botdemo"

.PHONY: sls/build
sls/build:
	docker build -t lambda .

.PHONY: sls/dev
sls/dev:
	docker run --rm -v $(PWD):/app -it lambda

.PHONY: deploy
deploy: go/run
	docker run --rm -v $(PWD):/app -it lambda bash -c "cp -r assets/.aws ~/.aws ; cp -r assets/.aws /app/botdemo ;\
		cd botdemo ; make deploy"

.PHONY: go/build
go/build:
	cd botdemo ; docker build -t botdemo .

.PHONY: go/run
go/run: go/build
	docker run --rm -v $(PWD)/botdemo:/app/botdemo -it botdemo
