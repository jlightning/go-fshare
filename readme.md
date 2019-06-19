# Go FShare Client
## Sample code:
##### Get Download URL
```
client := fshare.NewClient(fshare.NewConfig("<your_username>", "<your_password>"))
err := client.Login()
if err != nil {
	log.Fatal(err)
}
url, err := client.GetDownloadURL("<some_fshare_url>")
if err != nil {
	log.Fatal(err)
}
```