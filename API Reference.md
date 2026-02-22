# Table of Contents
- [Table of Contents](#table-of-contents)
- [`/message`](#message)
  - [POST](#post)
    - [Request Body](#request-body)
    - [Response](#response)
  - [GET](#get)
    - [URL Schema](#url-schema)
    - [Response](#response-1)
- [`/directedMessage`](#directedmessage)
  - [POST](#post-1)
    - [Request Body](#request-body-1)
    - [Response](#response-2)
- [`/subscribe/message`](#subscribemessage)
  - [POST](#post-2)
    - [Request Body](#request-body-2)
    - [Response](#response-3)
- [`/unsubscribe/message`](#unsubscribemessage)
  - [POST](#post-3)
    - [Request Body](#request-body-3)
    - [Response](#response-4)


# `/message`

This endpoint is used to send a message. It supports both POST and GET methods.

## POST

Sends a message.

### Request Body

The request body should be a JSON object with the following properties:

- `Message` (string): The content of the message.
- `Password` (string): The password for authentication.
- `ColourCode` (integer, optional): The color code for the message. See [IRC Formatting](https://modern.ircdocs.horse/formatting.html) for more details.
- `Broadcast` (boolean, optional): Whether the message should be broadcasted.

Example:

```json
{
    "Message": "Hello, world!",
    "Password": "correct_password",
    "ColourCode": 3,
    "Broadcast": true
}
```

### Response

| Status Code | Description |
|---|---|
| `200 OK` | Message sent successfully |
| `400 Bad Request` | Invalid JSON in request body, or message contains newline characters |
| `401 Unauthorized` | Invalid password |
| `500 Internal Server Error` | Failed to send message to IRC server |

## GET

Sends a message.

### URL Schema

The URL should include the following query parameters:

- `Message` (string): The content of the message.
- `Password` (string): The password for authentication.
- `ColourCode` (integer, optional): The color code for the message. See [IRC Formatting](https://modern.ircdocs.horse/formatting.html) for more details.
- `Broadcast` (boolean, optional): Whether the message should be broadcasted to all channels setup in the enviornment variables.

Example:
```
http://localhost:8080/message?Message=Hello,%20World!&Password=password&ColourCode=1&Broadcast=true
```

### Response

| Status Code | Description |
|---|---|
| `200 OK` | Message sent successfully |
| `400 Bad Request` | Invalid query parameters, or message contains newline characters |
| `401 Unauthorized` | Invalid password |
| `500 Internal Server Error` | Failed to send message to IRC server |

# `/directedMessage`

This endpoint is used to send a directed message. It only supports the POST method.

## POST

Sends a directed message to a specific target.

### Request Body

The request body should be a JSON object with the following properties:

- `Target` (string): The target to which the message should be sent.
- `IncomingMessage` (object): The message to be sent. It should have the following properties:
    - `Message` (string): The content of the message.
    - `Password` (string): The password for authentication.
    - `ColourCode` (integer, optional): The color code for the message. See [IRC Formatting](https://modern.ircdocs.horse/formatting.html) for more details.

Example:

```json
{
    "Target": "target1",
    "IncomingMessage": {
        "Message": "Hello, world!",
        "Password": "correct_password",
        "ColourCode": 3,
    }
}
```

### Response

| Status Code | Description |
|---|---|
| `200 OK` | Message sent successfully |
| `400 Bad Request` | Invalid JSON in request body, or message contains newline characters |
| `401 Unauthorized` | Invalid password |
| `500 Internal Server Error` | Failed to send message to IRC server |

# `/subscribe/message`

This endpoint is used to subscribe to message notifications. It only supports the POST method.

## POST

Subscribes to message notifications. They will be sent (via POST) to the provided URL on each message recieved in `Target` (ie, a channel)

Message payload sent to the URL will be like this:
```json
{
    "Target": "target_channel",
    "Message": "This is a test message.",
    "IRCUser": {
        "Nick": "TestNick",
        "User": "TestUser",
        "Host": "TestHost"
    },
    "Timestamp": "2024-05-28T12:34:56.789123456Z"
}
```

- **Target** (string): The target where the message was sent to. This will likely be the channel name.
- **Message** (string): The content of the message that was sent.
- **IRCUser** (IrcUser): An object representing the IRC user who is sending the message. It includes the following fields:
  - **Nick** (string): The nickname of the IRC user.
  - **User** (string): The username of the IRC user.
  - **Host** (string): The host of the IRC user.
- **Timestamp** (time.Time): The time when the message was created, in ISO 8601 format. This is generated using `time.Now()` in Go, which provides the current time.

You can process this accoridngly in whatever you are using to recieve the message, then respond accordingly. At the time of writing this, the message will NOT be re-sent in the event you fail to recieve it for any reason.

Respond with 200 OK if processed okay for future-proofing, as error handling/retries will eventualy be built-in.

### Request Body

The request body should be a JSON object with the following properties:

- `Target` (string): The target to which the subscription should be applied.
- `URL` (string): The URL to which notifications should be sent.
- `Password` (string): The password for authentication.

Example:

```json
{
    "Target": "target1",
    "URL": "http://example.com/notify",
    "Password": "correct_password"
}
```

### Response

| Status Code | Description |
|---|---|
| `200 OK` | Subscription successful |
| `400 Bad Request` | Invalid JSON in request body, or subscription failed |
| `401 Unauthorized` | Invalid password |

- `status` (string): Indicates success or failure.
- `message` (string): A message providing additional details.

Example Success Response (`200 OK`):

```json
{
    "status": "success",
    "message": "Subscription successful"
}
```

Example Failure Response (`400 Bad Request`):

```json
{
    "status": "failure",
    "message": "Subscription failed"
}
```

# `/unsubscribe/message`

This endpoint is used to unsubscribe from message notifications. It only supports the POST method.

## POST

Unsubscribes from message notifications.

### Request Body

The request body should be a JSON object with the following properties:

- `Target` (string): The target from which the subscription should be removed.
- `URL` (string): The URL to which notifications were being sent.
- `Password` (string): The password for authentication.

Example:

```json
{
    "Target": "target1",
    "URL": "http://example.com/notify",
    "Password": "correct_password"
}
```

### Response

| Status Code | Description |
|---|---|
| `200 OK` | Unsubscription successful |
| `400 Bad Request` | Invalid JSON in request body, or unsubscription failed |
| `401 Unauthorized` | Invalid password |

- `status` (string): Indicates success or failure.
- `message` (string): A message providing additional details.

Example Success Response (`200 OK`):
```json
{
    "status": "success",
    "message": "Unsubscription successful"
}
```

Example Failure Response (`400 Bad Request`):

```json
{
    "status": "failure",
    "message": "Unsubscription failed"
}
```