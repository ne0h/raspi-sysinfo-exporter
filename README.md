# RasPi System Information Exporter

Shows

* System load (load1, load5, load15)
* Memory usage (total, free, available)
* Temperature
* Platform information

Example

```json
{
	"platform": {
		"model": "Raspberry Pi 3 Model B Rev 1.2",
		"serial": "000000000056e779"
	},
	"load": {
		"load1": 0.11,
		"load5": 0.28,
		"load15": 0.37
	},
	"memory": {
		"total": 948300,
		"free": 37560,
		"available": 102160
	},
	"temperature": 55.306
}
```

## Build and Lint

```bash
# build
make build

# lint
make inst-golangci-lint
make lint
```

## Use

```bash
/usr/local/bin/raspi-sysinfo-exporter -h
Usage of /usr/local/bin/raspi-sysinfo-exporter:
  -addr string
        Listen to <interface>:<port>, e.g. 'localhost:8082' (default ":8082")
```

### Setup

```bash
# copy raspi-sysinfo-exporter to /usr/local/bin
# ensure thar /usr/local/bin is part of $PATH
cp raspi-sysinfo-exporter /usr/local/bin

# install systemd service to start the service automatically at boot time
cp raspi-sysinfo-exporter.service /etc/systemd/system
systemctl start raspi-sysinfo-exporter.service
systemctl enable raspi-sysinfo-exporter.service
```

#### Setup on LibreELEC

The minimal system of LibreELEC requires setting up the exporter in a slightly different way:

* The systemd service file has to be copied to `/storage/.config/system.d`
