# Systemd Service Deployment

## Setup
Take the `grellyddotcom.service` and toss it in `/etc/systemd/system/`. 
Create the `/var/http` directory and put required resources in there.

## Run it
Load it: `systemctl daemon-reload`
Enable it: `systemctl enable grellyddotcom.service`
Start it: `systemctl start grellyddotcom.service`
Check in on it: `systemctl status grellyddotcom.service`
Restart it: `systemctl restart grellyddotcom.service`

## Troubleshooting
Debugging one-liner: `systemctl daemon-reload && systemctl restart grellyddotcom.service && journalctl -u grellyddotcom.service --no-pager -f`


## Resources:

https://www.freedesktop.org/software/systemd/man/systemd.service.html
https://www.freedesktop.org/software/systemd/man/systemd.exec.html#
https://linuxhandbook.com/create-systemd-services/
https://www.loggly.com/ultimate-guide/linux-logging-with-systemd/