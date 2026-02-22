# Detailed Configuration

This document covers advanced configuration details for wmb beyond the basics needed to get the bot running locally.

## Environment Variables

| Variable | Required | Description |
|---|---|---|
| `IRC_SERVER` | Yes | IRC server address in `server:port` format |
| `IRC_CHANNEL` | Yes | The main IRC channel the bot will join and send messages to |
| `IRC_NICK` | Yes | The nickname the bot will use on IRC |
| `PASSWORD` | Yes | Password for the webhook endpoint, functioning as an API key you define |
| `NICKSERV_PASSWORD` | No | Password for NickServ authentication |
| `OTHER_IRC_CHANNELS` | No | Comma-separated list of additional channels to join on launch (besides the main channel) |
| `PORT` | No | Port for the webserver to listen on (defaults to `8080`) |
| `SELF_MODE` | No | User mode to set on self once joined to the IRC server (defaults to nothing, but should be `+B` on some servers) |
| `DBFILE` | No | File path for the database (defaults to `wmb.db` - be sure to map properly if inside Docker) |

## Docker Compose

See `docker-compose.yaml` for a ready-to-use example. The compose file includes all required and optional environment variables, with optional ones commented out. Adjust the values and run:

```
docker-compose up -d
```

Builds exist for `amd64` and `arm64` architectures in the Docker repo, so it should work on most systems including Raspberry Pis and x86_64 systems.

## nginx Reverse Proxy

If you want to put wmb behind an nginx reverse proxy, add the following to your desired nginx config:

```
location /wmb {
    rewrite /wmb/(.*) /$1  break;
    proxy_pass http://127.0.0.1:8080;
}
```

Then you can POST to `domain.tld/wmb/message` to send out a webhook.
