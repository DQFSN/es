package mongo

import (
	"testing"
)

func Test_Connect(t *testing.T)  {

	var client MyClient

	client.Connect()
	t.Log(client.Query("user"))
	client.Disconnect()
}
