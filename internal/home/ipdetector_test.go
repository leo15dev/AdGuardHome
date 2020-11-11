package home

import (
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIPDetector_detectSpecialNetwork(t *testing.T) {
	var ipd *ipDetector

	t.Run("newIPDetector", func(t *testing.T) {
		var err error
		ipd, err = newIPDetector()
		assert.Nil(t, err)
	})

	testCases := []struct {
		name string
		ip   net.IP
		want bool
	}{{
		name: "Not specific",
		ip:   net.ParseIP("8.8.8.8"),
		want: false,
	}, {
		name: "This host on this network",
		ip:   net.ParseIP("0.0.0.0"),
		want: true,
	}, {
		name: "Private-Use",
		ip:   net.ParseIP("10.0.0.0"),
		want: true,
	}, {
		name: "Shared Address Space",
		ip:   net.ParseIP("100.64.0.0"),
		want: true,
	}, {
		name: "Loopback",
		ip:   net.ParseIP("127.0.0.0"),
		want: true,
	}, {
		name: "Link Local",
		ip:   net.ParseIP("169.254.0.0"),
		want: true,
	}, {
		name: "Private-Use",
		ip:   net.ParseIP("172.16.0.0"),
		want: true,
	}, {
		name: "IETF Protocol Assignments",
		ip:   net.ParseIP("192.0.0.0"),
		want: true,
	}, {
		name: "DS-Lite",
		ip:   net.ParseIP("192.0.0.0"),
		want: true,
	}, {
		name: "Documentation (TEST-NET-1)",
		ip:   net.ParseIP("192.0.2.0"),
		want: true,
	}, {
		name: "6to4 Relay Anycast",
		ip:   net.ParseIP("192.88.99.0"),
		want: true,
	}, {
		name: "Private-Use",
		ip:   net.ParseIP("192.168.0.0"),
		want: true,
	}, {
		name: "Benchmarking",
		ip:   net.ParseIP("198.18.0.0"),
		want: true,
	}, {
		name: "Documentation (TEST-NET-2)",
		ip:   net.ParseIP("198.51.100.0"),
		want: true,
	}, {
		name: "Documentation (TEST-NET-3)",
		ip:   net.ParseIP("203.0.113.0"),
		want: true,
	}, {
		name: "Reserved",
		ip:   net.ParseIP("240.0.0.0"),
		want: true,
	}, {
		name: "Limited Broadcast",
		ip:   net.ParseIP("255.255.255.255"),
		want: true,
	}, {
		name: "Loopback Address",
		ip:   net.ParseIP("::1"),
		want: true,
	}, {
		name: "Unspecified Address",
		ip:   net.ParseIP("::"),
		want: true,
	}, {
		name: "IPv4-IPv6 Translat.",
		ip:   net.ParseIP("64:ff9b::"),
		want: true,
	}, {
		name: "Discard-Only Address Block",
		ip:   net.ParseIP("100::"),
		want: true,
	}, {
		name: "IETF Protocol Assignments",
		ip:   net.ParseIP("2001::"),
		want: true,
	}, {
		name: "TEREDO",
		ip:   net.ParseIP("2001::"),
		want: true,
	}, {
		name: "Benchmarking",
		ip:   net.ParseIP("2001:2::"),
		want: true,
	}, {
		name: "Documentation",
		ip:   net.ParseIP("2001:db8::"),
		want: true,
	}, {
		name: "ORCHID",
		ip:   net.ParseIP("2001:10::"),
		want: true,
	}, {
		name: "6to4",
		ip:   net.ParseIP("2002::"),
		want: true,
	}, {
		name: "Unique-Local",
		ip:   net.ParseIP("fc00::"),
		want: true,
	}, {
		name: "Linked-Scoped Unicast",
		ip:   net.ParseIP("fe80::"),
		want: true,
	}}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, ipd.detectSpecialNetwork(tc.ip))
		})
	}
}
