Very simple bot to turn JSON webhooks into IRC messages.

Work in progress - not ready for usage

# Envionment Variables
- IRC_SERVER in server:port format
- IRC_CHANNEL 
- IRC_NICK

# Send a message to IRC
`curl -X POST -H "Content-Type: application/json" -d '{"message":"Hello, World! stuff", "
password":"password"}' http://localhost:8080/webhook`