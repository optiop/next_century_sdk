package client

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/optiop/next_century_sdk/core/schema"
	ncSdkMock "github.com/optiop/next_century_sdk/mock"
	ncMockSample "github.com/optiop/next_century_sdk/mock/sample"
	"github.com/stretchr/testify/assert"
)

func TestClient(t *testing.T) {
	go ncSdkMock.Run()

	// wait for server to start
	time.Sleep(1 * time.Second)

	mockClient := New("test", "test", "http://localhost:1234")

	// Test GetDailyReads
	dataResMock, err := mockClient.GetDailyReads("x_1234", schema.TimeRequest{Date: time.Now()})
	assert.Nil(t, err)

	// check it with sample mock data
	var dataMock []*schema.MeterData
	err = json.Unmarshal([]byte(ncMockSample.DailyReadsSample), &dataMock)
	assert.Nil(t, err)

	assert.Equal(t, dataMock, dataResMock)
}
