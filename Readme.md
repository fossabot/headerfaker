# HeaderFaker

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/aURORA-JC/headerfaker?logo=go&style=flat-square)
[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2FCustCSA%2Fheaderfaker.svg?type=shield)](https://app.fossa.com/projects/git%2Bgithub.com%2FCustCSA%2Fheaderfaker?ref=badge_shield)

HeaderFaker is a simulation training ground to Help beginners learn about http protocol with answering some question.

## Install

1. Clone the repository from Github

```shell
git clone https://github.com/aURORA-JC/headerfaker.git
```

2. Cd to project root path & build

```shell
go build
```

3. HeaderFaker will be build in the project root path

4. Install Mysql or MariaDB

## Usage

+ Files required to start the program
```shell
 + template/
 + data.json
 + config.ini
 + headerfaker
```
1. Import the database from `heakerfaker.sql`

3. Modify the configuration file `config.ini`
```ini
port = 9000
release = false

[mysql]
user = 
password = 
host = 
port = 
db = headerfaker
```

+ Simple Start

```shell
# Set runnable power (Linux need)
sudo chmod +x ./headerfaker

# Start service
./headerfaker
# HeaderFaker will start on 0.0.0.0:9000
````

+ systemd run
```shell
# Set runnable power (Linux need)
sudo chmod +x ./headerfaker

# add a new .service file in /lib/systemd/system
sudo vim /lib/systemd/system/headerfaker.service

# write & save content
[Unit]
Description=HeaderFaker Service
After=network.target

[Service]
Type=simple
User=nobody
Restart=on-failure
RestartSec=5s
ExecStart=<YOUR_HEADERFAKER_ABSOLUTE_PATH>
WorkingDirectory=<YOUR_HEADERFAKER_DIR_ABSOLUTE_PATH>
LimitNOFILE=1048576

[Install]
WantedBy=multi-user.target
# Or you can copy the .service file under the ./init dir

# Start the service
systemctl start headerfaker

# Set start at boot
systemctl enable headerfaker
```

## Contributing

PRs accepted. But new branch is developing.

## License

MIT © aURORA-JC


[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2FCustCSA%2Fheaderfaker.svg?type=large)](https://app.fossa.com/projects/git%2Bgithub.com%2FCustCSA%2Fheaderfaker?ref=badge_large)

## Author

John Chow (aURORA-JC)