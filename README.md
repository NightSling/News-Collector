# news-collector
Collects news with news api and stores it in a mysql/mariadb database.

## How to Install
1. Initialize a database with the table listed in `init.sql`.
2. Copy config.example.yaml to config.yaml and change it's values..
3. Build with `go build ./cmd/news/main.go` or `make install`
4. Run the executable with flag `-c <CONFIG.YAML PATH>` in order to scrap the data.

## How to setup schedular
1. Install a crontab implementation such as `cronie`.
2. `make install` or make sure you know the path to your exectuable.
3. Make sure you know the path of your `config.yaml`.
4. Run `crontab -e` 
5. Paste this `0 */2 * * * /usr/bin/news-scrap -c /home/daysling/.news.yaml` (Change the config file path and you are good to go!)
