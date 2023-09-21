shutdown-monitor is a simple golang program intended to be run on a Raspberry Pi. It monitors the state of a pin, and when it is grounded the program will shutdown the Raspberry Pi.

## Installation

### Prerequisites

* [Go](https://golang.org/doc/install) - openbar-server is written in Go, so you'll need to install Go to build and run the project.
* Git - `sudo apt install git`

### Installation From Source

Clone and install the go application

```bash
git clone https://github.com/cocktailrobots/shutdown-monitor.git
cd shutdown-monitor
go install .
```

Create /etc/systemd/system/shutdown-monitor.service:

```
[Unit]
Description=shutdown monitor
After=network.target

[Service]
Type=simple
User=root
Group=root
AmbientCapabilities=CAP_NET_BIND_SERVICE

ExecStart=/home/openbar/go/bin/shutdown-monitor 26

Restart=always
RestartSec=1

MemoryAccounting=true
MemoryMax=90%

[Install]
WantedBy=multi-user.target
```

Enable and start the monitor as a systemctl service

```
sudo systemctl enable shutdown-monitor
sudo systemctl start shutdown-monitor
```
