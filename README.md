# Minimal IRC client

To use it create add a config.toml in the same directory as the executable with the same fields that are in `example-config.toml`

The syntax for sending a message is:  
`PRIVMSG #global :this is a message :p`

Make sure you respond to pings or you will be forecebly disconnected. To respond do this:
```sh
PING 123456
PONG 123456
```
