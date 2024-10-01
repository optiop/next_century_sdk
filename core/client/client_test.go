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

func TestMain(m *testing.M) {
	go ncSdkMock.Run()

	// wait for server to start
	time.Sleep(1 * time.Second)

	m.Run()
}

func TestClient(t *testing.T) {
	mockClient := New("test", "test", "http://localhost:1234")

	t.Run("TestGetDailyReads", func(t *testing.T) {
		dataResMock, err := mockClient.GetDailyReads("x_1234", schema.TimeRequest{Date: time.Now()})
		assert.Nil(t, err)

		// check it with sample mock data
		var dataMock []*schema.MeterData
		err = json.Unmarshal([]byte(ncMockSample.DailyReadsSample), &dataMock)
		assert.Nil(t, err)

		assert.Equal(t, dataMock, dataResMock)
	})

	t.Run("TestGetUnits", func(t *testing.T) {
		unitsResMock, err := mockClient.GetUnits("x_1234")
		assert.Nil(t, err)

		// check it with sample mock data
		var unitsMock []*schema.Unit
		err = json.Unmarshal([]byte(ncMockSample.UnitsSample), &unitsMock)
		assert.Nil(t, err)

		assert.Equal(t, unitsMock, unitsResMock)
	})

}
