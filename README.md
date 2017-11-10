# Midtrans Bridge
Any Matching Transaction will be forwarded into any matching URL. Helping you to
manage callback from midtrans if you have multiple application.


## HOW TO USE
Set mapping urls with your Prefix Order ID as key.

Example:
```go
urls := map[string]string{
    "VALUTAC":  "https://valutac.com/midtrans/callback",
}
```

Any transaction which has `VALUTAC` prefix on order id will be forwarded into
`https://valutac.com/midtrans/callback`. Example: `VALUTAC-01-0001-01`.


## HOW TO DEPLOY
* Build the binary with `go build`
    * Use GOOS or GOARCH to build cross platform, Ex: `GOOS=linux go build`
* Copy binary file into your deployment machine
* Create a new service file, Ex: `/etc/systemd/system/midtrans.service`:

    ```
    [Unit]
    Description=Midtrans Bridge
    After=syslog.target

    [Service]
    User=root
    ExecStart=/path/to/binary SuccessExitStatus=143

    [Install]
    WantedBy=multi-user.target
    ```
* Re-read systemd configuration

    ```sh
    $ sudo systemctl daemon-reload
    ```
* Start the service. The Application will be running on port :8080

    ```sh
    $ sudo systemctl start midtrans
    ```
* Proxy Reverse with Nginx

    ```
    upstream midtrans {
        server 127.0.0.1:8080;
    }

    server {
        ...

        location / {
            proxy_pass http://midtrans;
            proxy_next_upstream error timeout invalid_header http_500 http_502
            http_503 http_504;
            proxy_redirect off;
            proxy_buffering off;
            proxy_set_header Host $host;
            proxy_set_header X-Real-IP $remote_addr;
            proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        }
    }
    ```

* Make sure to set Notification URL into midtrans with the right path,
    "http://yourdeploymentdomain.com/"

## LICENSE

<a href="LICENSE">
<img src="https://raw.githubusercontent.com/valutac/accent/master/mit.png" width="75"></img>
</a>
