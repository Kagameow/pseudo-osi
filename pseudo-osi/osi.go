/*
Write a program that implements message flow from the top layer to the bottom layer
of the 7-layer protocol model. Your program should include a separate protocol function
for each layer. Protocol headers are sequence up to 64 characters.
Each protocol function has two parameters:
a message passed from the higher layer protocol (a char buffer)
and the size of the message.
This function attaches its header in front of the message,
prints the new message on the standard output,
and then invokes the protocol function of the lower-layer protocol.
Program input is an application message.
*/
package main

import (
	"fmt"
)

type Layer struct {
	name   string
	header string
}

var OSILayers = []Layer{
	{"Application", "APPLICATION_HEADER: "},
	{"Presentation", "PRESENTATION_HEADER: "},
	{"Session", "SESSION_HEADER: "},
	{"Transport", "TRANSPORT_HEADER: "},
	{"Network", "NETWORK_HEADER: "},
	{"Data Link", "DATA_LINK_HEADER: "},
	{"Physical", "PHYSICAL_HEADER: "},
}

func indexOf[T comparable](collection []T, el T) int {
	for i, x := range collection {
		if x == el {
			return i
		}
	}
	return -1
}

func (l Layer) addProtocolHeaderAndPassLower(message string, size int) {
	message = l.header + message
	fmt.Printf("%s layer: %s lenght: %d\n", l.name, message, size)
	index := indexOf[Layer](OSILayers, l)

	if index+2 > len(OSILayers) {
		return
	}

	OSILayers[index+1].addProtocolHeaderAndPassLower(message, len(message))
}

func main() {
	message := "Hello Network!"
	OSILayers[0].addProtocolHeaderAndPassLower(message, len(message))
}
