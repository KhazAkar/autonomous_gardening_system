package backend

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"os"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

func connectToMqttServer() MQTT.Client {
	broker := os.Getenv("MQTT_BROKER")
	port := os.Getenv("MQTT_PORT")
	cafile := os.Getenv("MQTT_CA_CERT")
	clientCert := os.Getenv("MQTT_CLIENT_CERT")
	clientKey := os.Getenv("MQTT_CLIENT_KEY")

	caCert, err := os.ReadFile(cafile)
	if err != nil {
		fmt.Println("Error loading CA cert:", err)
		return nil
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	clientCertPEM, err := os.ReadFile(clientCert)
	if err != nil {
		fmt.Println("Error loading client cert:", err)
		return nil
	}
	clientKeyPEM, err := os.ReadFile(clientKey)
	if err != nil {
		fmt.Println("Error loading client key:", err)
		return nil
	}
	cert, err := tls.X509KeyPair(clientCertPEM, clientKeyPEM)
	if err != nil {
		fmt.Println("Error creating X509 key pair:", err)
		return nil
	}

	tlsConfig := &tls.Config{
		RootCAs:      caCertPool,
		Certificates: []tls.Certificate{cert},
	}

	opts := MQTT.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tls://%s:%s", broker, port))
	opts.SetTLSConfig(tlsConfig)

	client := MQTT.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
	}
	return client
}

func MqttClient() {
	client := connectToMqttServer()
}
