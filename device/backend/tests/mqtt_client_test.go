package backend_test

import (
	"os"
	"testing"
)

func TestConnectToMqttServer(t *testing.T) {
	// Set up environment variables for testing
	os.Setenv("MQTT_BROKER", "test_broker")
	os.Setenv("MQTT_PORT", "1883")
	os.Setenv("MQTT_USERNAME", "test_user")
	os.Setenv("MQTT_PASSWORD", "test_password")

	// Call the function to test
	backend.connectToMqttServer()

	// Add assertions to verify the behavior
	// For example, you can check if the client is connected
	// assert.True(t, client.IsConnected())
}
