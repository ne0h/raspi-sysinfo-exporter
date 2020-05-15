# RasPi System Information Exporter

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
