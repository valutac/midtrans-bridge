# Midtrans Bridge
Any Matching Transaction will be forwarded into any matching URL. Helping you to
manage callback from midtrans if you have more than one application.


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
Deployment can be done with supervisord or systemd.
* Make sure to set Notification URL into midtrans with the right path,
    "http://yourdeploymentdomain.com/".
