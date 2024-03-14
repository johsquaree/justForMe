# Data Transmission Between MQTT and HTTP Channel

This project entails a Go application that utilizes an MQTT client to transmit received messages to an HTTP channel.

MQTT is a lightweight messaging protocol commonly employed for data communication among IoT devices.

This application processes messages received via MQTT and forwards them to an HTTP channel.

## Getting Started

To run this application, you need to have Go language installed on your computer. Additionally, a few steps are required to resolve the project's dependencies.

### Requirements

-Go language installed
-Eclipse Paho MQTT library installed

### Installing Dependencies

We utilize the Eclipse Paho MQTT library to run the project. You can install this library with the following command:

```bash
go get github.com/eclipse/paho.mqtt.golang
