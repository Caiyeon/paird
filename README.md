### Starting the application
Make sure your vault token is set as env var: `export VAULT_TOKEN={secret_token}`

Then: `go run server.go`

And then test the server by pinging it:
`curl http://127.0.0.1:8000/v1/ping`



### Uploading static frontend files
Upload to docs folder (replacing index.html and other files/folders) and go to https://caiyeon.github.io/paird/



### Accessing the webserver via ssh
`ssh root@tonycai.me`

Contact me if your ssh key has not been added
