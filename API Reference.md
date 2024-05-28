# Table of Contents
- [Table of Contents](#table-of-contents)
- [`/message`](#message)
  - [POST](#post)
    - [Request Body](#request-body)
  - [GET](#get)
    - [URL Schema](#url-schema)
- [`/directedMessage`](#directedmessage)
  - [POST](#post-1)
    - [Request Body](#request-body-1)

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
## GET

Sends a message.

### URL Schema

The URL should include the following query parameters:

- `Message` (string): The content of the message.
- `Password` (string): The password for authentication.
- `ColourCode` (integer, optional): The color code for the message. See [IRC Formatting](https://modern.ircdocs.horse/formatting.html) for more details.
- `Broadcast` (boolean, optional): Whether the message should be broadcasted to all channels setup in the enviornment variables.

Example:
```http
http://localhost:8080/message?Message=Hello,%20World!&Password=password&ColourCode=1&Broadcast=true
```

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
    - `Broadcast` (boolean, optional): Whether the message should be broadcasted.

Example:

```json
{
    "Target": "target1",
    "IncomingMessage": {
        "Message": "Hello, world!",
        "Password": "correct_password",
        "ColourCode": 3,
        "Broadcast": true
    }
}
