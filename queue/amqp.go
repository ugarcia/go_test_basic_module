package queue

import (
    "fmt"
    "strings"

    "github.com/streadway/amqp"

    "github.com/ugarcia/go_test_common/models"
    "github.com/ugarcia/go_test_common/mq"
)


// Global variables for queue/channel etc.
var q *mq.AMQP

// Constants for this module
const MQ_URL = "amqp://guest:guest@mq.gamewheel.local:5672/"
const QUEUE = "basic_module_q"
const ROUTE = "modules.basic"
const ID = "modules.basic"
const EXCHANGE = "modules"
var EXCHANGES = []string{"mcp", "modules", "workers"}

/**
 * Initialize Queues, Channels and Consumer Loops
 */
func Init() {

    // Create struct
    q = new(mq.AMQP)

    // Init it
    q.Init(MQ_URL)

    // Defer closing
    defer q.Close()

    // Register exchanges
    q.RegisterExchanges(EXCHANGES)

    // Register queues
    q.RegisterQueues([]string{QUEUE})

    // Bind queues
    q.BindQueuesToExchange([]string{QUEUE}, EXCHANGE, ROUTE)

    // Start consuming
    q.Consume(QUEUE, receiveQueueMessage)
}

/**
 * Receives a message from basic modules queue and adds it to queries channel
 */
func receiveQueueMessage(msg models.QueueMessage, d amqp.Delivery) {

    // Get source exchange channel
    exchange := strings.Split(msg.Sender, ".")[0]

    // Lookup original exchange channel from message, then call logic and send response
    switch exchange {

        // Coming from MCP
        case "mcp":
            handleMcpMessage(msg)

        // Coming from a Worker
        case "workers":
            handleWorkerMessage(msg)

        // Unknown source
        default:
            fmt.Println("Unknown message exchange source for Basic Module")
    }

    // TODO: Pass this delivery object along ans send ACK only after finishing everything???
    d.Ack(false)
}

/**
 * Handles logic for a message coming from MCP
 */
func handleMcpMessage(msg models.QueueMessage) {

    // Lookup Target
    switch msg.Target {

        // Data request target
        case "data":

            // Lookup Action
            switch msg.Action {

                // DB action
                case "index", "post", "get", "update", "delete":

                    // TODO: Logic here for needed steps. for now considering one-off tasks

                    msg.Receiver = "workers.db"

                // TODO: More Worker receivers here

            // Unknown action
            default:
                fmt.Printf("Invalid action '%s' for Basic Module", msg.Action)
                return
            }

        // Unknown target
        default:
            fmt.Printf("Invalid message target '%s' for Basic Module: ", msg.Target)
            return
    }

    // Set the sender to current module
    msg.Sender = ID

    // Send the message to queue
    q.SendMessage(msg)
}

/**
 * Handles logic for a message coming from a Worker
 */
func handleWorkerMessage(msg models.QueueMessage) {

    // Lookup worker sender
    switch msg.Sender {

        // Coming from DB Worker
        case "workers.db":

            // Lookup Target
            switch msg.Target {

                // Data request target
                case "data":

                    // Lookup Action
                    switch msg.Action {

                        // DB action
                        case "index", "post", "get", "update", "delete":

                            // TODO: Logic here for needed steps. for now considering one-off tasks

                            msg.Receiver = msg.Source

                        // TODO: More Worker receivers here

                        // Unknown action
                        default:
                            fmt.Printf("Invalid action '%s' for Basic Module", msg.Action)
                            return
                    }

                // Unknown target
                default:
                    fmt.Printf("Invalid message target '%s' for Basic Module: ", msg.Target)
                    return
            }

        // Unknown target
        default:
            fmt.Printf("Invalid message source '%s' for Basic Module: ", msg.Sender)
            return
    }

    // Set the sender to current module
    msg.Sender = ID

    // Send the message to queue
    q.SendMessage(msg)
}
