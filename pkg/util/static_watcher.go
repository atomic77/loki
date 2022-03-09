package util

import (
	"context"

	"github.com/grafana/dskit/services"
)

type staticNotifications interface {
	AddressAdded(address string)
	AddressRemoved(address string)
}

type staticWatcher struct {
	notifications staticNotifications
	address       string
}

// NewStaticWatcher returns a trivial service that does a passthrough of a statically defined Address
// without going through DNS resolution
func NewStaticWatcher(address string, notifications staticNotifications) (services.Service, error) {

	w := &staticWatcher{
		notifications: notifications,
		address:       address,
	}
	return services.NewIdleService(w.fakeDNSResolve, nil), nil
}

func (w *staticWatcher) fakeDNSResolve(servCtx context.Context) error {
	// A "resolve" function that passes the address as is
	w.notifications.AddressAdded(w.address)
	return nil
}