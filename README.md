# go_test_basic_module
Testing Go Capabilities

15.01.2016:
-----------

- A Basic module which does:
    - Consume messages from MCP and publish, based on messages data, to relevant workers exchange
    - Consume messages from workers and publish, based on messages data, to MCP exchange
    - Perform whatever logic is needed based on incoming messages
    Pending:
        - Publish job responses to more exchanges? (like APIs, etc ...)

Easiest way to test it is running the process and inspecting the logs

    go run src/github.com/ugarcia/go_test_basic_module/main.go

Dependencies:
-------------
- RabbitMq >= 3.5.1
- Golang >= 1.5.x (and proper system config for $GOPATH and $GOROOT)
- Package github.com/ugarcia/go_test_common
