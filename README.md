# news-collector
Collects news with news api and stores it in a mysql/mariadb database.

## How to Install
i) Initialize a database with the table listed in `init.sql`.
ii) Copy config.example.yaml to config.yaml and change it's values..
iii) Build with `go build ./cmd/news/main.go` or `make install`
iv) Run the executable with flag `-c <CONFIG.YAML PATH>` in order to scrap the data.

*Tip: Setup anacron to scrap news data daily*
