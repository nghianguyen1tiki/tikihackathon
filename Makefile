run: build web

local:
	heroku local web

build:
	go build -o bin/app -v cmd/app/main.go

deploy:
	git push heroku main

remote:
	heroku open