package tailscalesd

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"tailscale.com/client/tailscale/v2"
)

var errFailedAPIRequest = errors.New("failed API request")

type apiDiscoverer struct {
	client *tailscale.Client
}

func PublicAPIDiscoverer(client *tailscale.Client) Discoverer {
	return apiDiscoverer{
		client: client,
	}
}

func (a apiDiscoverer) Devices(ctx context.Context) ([]Device, error) {
	if a.client == nil {
		return nil, fmt.Errorf("no Tailscale client")
	}
	start := time.Now()
	lv := prometheus.Labels{
		"api": "public",
	}
	defer func() {
		apiRequestLatencyHistogram.With(lv).Observe(float64(time.Since(start).Milliseconds()))
	}()

	devices, err := a.client.Devices().List(ctx)

	if err != nil {
		apiRequestErrorCounter.With(lv).Inc()
		return nil, err
	}

	apiRequestCounter.With(prometheus.Labels{
		"api": "public",
	}).Inc()
	returnDevices := make([]Device, len(devices))
	for _, d := range devices {
		returnDevices = append(returnDevices, translateAPIDeviceToDevice(d, a.client.Tailnet))
	}
	return returnDevices, nil
}

func translateAPIDeviceToDevice(apiDevice tailscale.Device, tailnet string) Device {
	var d Device
	d.ID = apiDevice.ID
	d.Hostname = apiDevice.Name
	d.OS = apiDevice.OS
	d.Tags = apiDevice.Tags
	d.API = "public"
	d.ClientVersion = apiDevice.ClientVersion
	d.Authorized = apiDevice.Authorized
	d.Addresses = apiDevice.Addresses
	d.Authorized = apiDevice.Authorized
	d.Tailnet = tailnet
	return d
}
