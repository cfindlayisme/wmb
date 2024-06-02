Very simple bot to send IRC messages from webhooks.

## Launching
See `docker-compose.yaml` for an example of how to launch the bot. Just have to change envionment variables and run `docker-compose up -d`

Builds exist for `amd64` amd `arm64` architectures in the docker repo, so it should work on most systems incl raspberry pis and x86_64 systems.

## Envionment Variables
- `IRC_SERVER` in server:port format
- `IRC_CHANNEL` 
- `IRC_NICK`
- `PASSWORD` (Not IRC password, but password for the webhook. Sort of like an API key you define)
- `NICKSERV_PASSWORD` (Optional)
- `OTHER_IRC_CHANNELS` (Optional, comma separated list of channels to join on launch that are not the main channel)
- `PORT` (Optional, what port to have the webserver listen on, defaults to 8080)
- `SELF_MODE` (Optional, what mode to set on self once joined to IRC server. Defaults nothing, but should be +B on some servers)
- `DBFILE` (Optional, what file to use for the database, defaults to `wmb.db` - be sure to map properly if inside docker)


## API documentation
See [API Reference](API%20Reference.md) for more information on how to use the API's endpoints in your projects. This details all the schema, endpoints, and variables available.

## Examples
See [Examples](Examples.md) for some examples of how to use the bot in your projects.

## Building
See [Building](Building.md) for information on how to build the bot if not using a published build.

## nginx reverse proxy
```
location /wmb {
    rewrite /wmb/(.*) /$1  break;
    proxy_pass http://127.0.0.1:8080;
}
```

In your desired config. Then you can POST to domain.tld/wmb/message to send out a webhook.

## Resource usage
For those of you running this on tiny VMs, the resource usage is pretty low - barely any CPU usage, and about 12MB of RAM usage (which is pretty good considering it's running an IRC client and a web server.)

## But why?
I wanted to get notifications in IRC from my various systems, and particularly my GitHub action pipelines. This does that.