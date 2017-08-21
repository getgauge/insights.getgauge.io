# insights.getgauge.io

steps to generate insights:

* clone this repo.
* copy `analytics private key` in `ga.pem` file.
* set `GA_TOKEN_URL`, `GA_SERVICE_ACCOUNT_EMAIL` and `GA_VIEW_ID` env variables.
* run `go get ./...`
* run `go run *.go`

it will generate html pages in `_insights` dir.
