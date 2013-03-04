//
//  Weather proxy device.
//
package main

import (
	zmq "github.com/pebbe/zmq3"
)

func main () {
    //  This is where the weather server sits
    frontend, _ := zmq.NewSocket(zmq.XSUB)
	defer frontend.Close()
    frontend.Connect("tcp://192.168.55.210:5556")

    //  This is our public endpoint for subscribers
    backend, _ := zmq.NewSocket(zmq.XPUB)
	defer backend.Close()
    backend.Bind("tcp://10.1.1.0:8100")

    //  Run the proxy until the user interrupts us
    zmq.Proxy(frontend, backend, nil)
}
