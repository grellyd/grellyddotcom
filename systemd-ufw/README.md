# systemd Service Deployment

## systemd 

### Setup
Take `grellyddotcom.service` and put it in `/etc/systemd/system/`. 
Create the `/var/http` directory and put required resources in there.

### Run it
Load it: `systemctl daemon-reload`
Enable it: `systemctl enable grellyddotcom.service`
Start it: `systemctl start grellyddotcom.service`
Check in on it: `systemctl status grellyddotcom.service`
Restart it: `systemctl restart grellyddotcom.service`

### Troubleshooting
Debugging one-liner: `systemctl daemon-reload && systemctl restart grellyddotcom.service && journalctl -u grellyddotcom.service --no-pager -f`

## ufw 

### Setup
Take `grellydotcom-server` and put it in `/etc/ufw/applications.d/`.
Turn on ufw: `ufw enable`

### Configure it
Load it: `ufw app update GrellydDotCom`
View it: `ufw app info GrellydDotCom`
Allow it: `ufw allow GrellydDotCom`




## Resources:

- https://www.freedesktop.org/software/systemd/man/systemd.service.html
- https://www.freedesktop.org/software/systemd/man/systemd.exec.html#
- https://linuxhandbook.com/create-systemd-services/
- https://www.loggly.com/ultimate-guide/linux-logging-with-systemd/
- https://askubuntu.com/questions/409013/how-do-you-create-an-app-profile-for-ufw