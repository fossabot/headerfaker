# HeaderFaker

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/aURORA-JC/headerfaker?logo=go&style=flat-square)

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

## Usage
+ Simple Start

```shell
# Move data file to program root dir
mv ./database/data.json .

# Set runnable power (Linux need)
sudo chmod +x ./headerfaker

# Start service
./headerfaker
# HeaderFaker will start on 0.0.0.0:11451
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
LimitNOFILE=1048576

[Install]
WantedBy=multi-user.target

# Start the service
systemctl start headerfaker

# Set start at boot
systemctl enable headerfaker
```

## Contributing

PRs accepted.

## License

MIT © aURORA-JC

## Author

John Chow (aURORA-JC)