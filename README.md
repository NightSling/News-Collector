# news-collector
Collects news with news api and stores it in a mysql/mariadb database.

## How to Install
1. Initialize a database with the table listed in `init.sql`.
2. Copy config.example.yaml to config.yaml and change it's values..
3. Build with `go build ./cmd/news/main.go` or `make install`
4. Run the executable with flag `-c <CONFIG.YAML PATH>` in order to scrap the data.

*Tip: Setup anacron to scrap news data daily*
