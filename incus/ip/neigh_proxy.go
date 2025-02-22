package ip

import (
	"net"
	"strings"

	"github.com/cyphar/incus/shared"
)

// NeighProxy represents arguments for neighbour proxy manipulation.
type NeighProxy struct {
	DevName string
	Addr    net.IP
}

// Show list neighbour proxy entries.
func (n *NeighProxy) Show() ([]NeighProxy, error) {
	out, err := shared.RunCommand("ip", "neigh", "show", "proxy", "dev", n.DevName)
	if err != nil {
		return nil, err
	}

	lines := shared.SplitNTrimSpace(out, "\n", -1, true)
	entries := make([]NeighProxy, 0, len(lines))

	for _, line := range lines {
		fields := strings.Fields(line)
		if len(fields) <= 0 {
			continue
		}

		ip := net.ParseIP(fields[0])
		if ip == nil {
			continue
		}

		entries = append(entries, NeighProxy{
			DevName: n.DevName,
			Addr:    ip,
		})
	}

	return entries, nil
}

// Add a neighbour proxy entry.
func (n *NeighProxy) Add() error {
	_, err := shared.RunCommand("ip", "neigh", "add", "proxy", n.Addr.String(), "dev", n.DevName)
	if err != nil {
		return err
	}

	return nil
}

// Delete a neighbour proxy entry.
func (n *NeighProxy) Delete() error {
	_, err := shared.RunCommand("ip", "neigh", "delete", "proxy", n.Addr.String(), "dev", n.DevName)
	if err != nil {
		return err
	}

	return nil
}
