# golang-websocket-with-redis-sample

## How to run app
### using redis-cli
```bash
make up
open assets/index.html
redis-cli --pass password publish sample "hello"
```

### using websocat
```bash
make up
websocat ws://127.0.0.1:8000/ws
>>> hello
<<< pong
```
