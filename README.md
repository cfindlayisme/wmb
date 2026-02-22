Very simple bot to send IRC messages from webhooks.

## Launching
See `docker-compose.yaml` for an example of how to launch the bot. Just have to change environment variables and run `docker-compose up -d`

## Configuration
See [Detailed Configuration](Detailed%20Configuration.md) for a comprehensive reference of all environment variables, Docker Compose usage, and nginx reverse proxy setup.

## API documentation
See [API Reference](API%20Reference.md) for more information on how to use the API's endpoints in your projects. This details all the schema, endpoints, and variables available.

## Examples
See [Examples](Examples.md) for some examples of how to use the bot in your projects.

## Building
See [Building](Building.md) for information on how to build the bot if not using a published build.

## Resource usage
For those of you running this on tiny VMs, the resource usage is pretty low - barely any CPU usage, and about 12MB of RAM usage (which is pretty good considering it's running an IRC client and a web server.)

## But why?
I wanted to get notifications in IRC from my various systems, and particularly my GitHub action pipelines. This does that.
