# Monitor services that use Patan

When: 
 * multiple services are using Patan, and expose the snapshot via Rest 
 * communication between services is tracked using patan

Then this tool, given an input description, can: 
 * Draw a graph of the services and their connections
 * Show a subset of the statistics on the nodes and eges.
 
### Getting started

    go get github.com/stretchr/testify
    go get github.com/toefel18/patan-monitor
    
    run application.go with 1 input parameter: path to input-example.json
    