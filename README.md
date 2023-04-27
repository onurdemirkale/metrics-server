- Before running the commands below, update the `ExecStart` field with the absolute path to the `startup-script.sh`.

- Build the Go binary if necessary.

```
go build main.go
```

- Update generate-metrics and startup script with executable permissions.

```
chmod +x generate-metrics.sh
chmod +x startup-script.sh
```

- Generate metrics.
```
sudo ./generate-metrics.sh
```

- Copy the metrics service.
```
sudo cp systemd-config/metrics.service /etc/systemd/system/
```

- Start metrics service.
```
sudo systemctl daemon-reload
sudo systemctl enable metrics.service
sudo systemctl start metrics.service
```
