deploy:
	@git push heroku main

open:
	@heroku open

set-go-version:
	heroku config:set GOVERSION=go1.16