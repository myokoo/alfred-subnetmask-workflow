package main

import (
	"github.com/deanishe/awgo"
	"github.com/pkg/errors"
	"net"
	"strconv"
)

type ipv4Prefix struct {
	n *network
}
type ipv4SubnetMask struct {
	n *network
}
type ipv4BroadcastAddr struct {
	n *network
}

func (self *ipv4Prefix) calculate() error {
	var err error
	self.n.prefix, err = strconv.Atoi(self.n.query[1:])
	if err != nil {
		return err
	}

	mask := net.CIDRMask(self.n.prefix, 32)
	self.n.subnetMask = net.IPv4bcast.Mask(mask)
	return nil
}

func (self *ipv4Prefix) addNewItem() {
	addDefaultItem(self.n)
}

func (self *ipv4SubnetMask) calculate() error {
	var err error
	var ipv4Net *net.IPNet

	self.n.ip, ipv4Net, err = net.ParseCIDR(self.n.query)
	if err != nil {
		return err
	}

	self.n.networkAddr = ipv4Net.IP
	self.n.subnetMask = net.IPv4bcast.Mask(ipv4Net.Mask)
	self.n.prefix, _ = ipv4Net.Mask.Size()

	self.n.broadcastAddr = make([]byte, len(self.n.ip.To4()))
	copy(self.n.broadcastAddr, self.n.ip.To4())
	for i, v := range ipv4Net.Mask {
		self.n.broadcastAddr[i] |= v ^ 255
	}
	return nil
}

func (self *ipv4SubnetMask) addNewItem() {
	wf.NewItem(self.n.ip.String()).
		Subtitle(`ip`).
		Icon(&aw.Icon{Value: "./img/ip.png", Type: ""})
	wf.NewItem(self.n.networkAddr.String() + " ~ " + self.n.broadcastAddr.String()).
		Subtitle(`network range`).
		Icon(&aw.Icon{Value: "./img/range.png", Type: ""})
	addDefaultItem(self.n)
}

func (self *ipv4BroadcastAddr) calculate() error {
	ip := net.ParseIP(self.n.query).To4()
	if ip == nil {
		return errors.New("入力形式が不正です。")
	}
	mask := net.IPv4Mask(ip[0], ip[1], ip[2], ip[3])
	self.n.prefix, _ = mask.Size()
	self.n.subnetMask = ip
	if self.n.prefix == 0 && !ip.Equal(net.ParseIP(`0.0.0.0`)) {
		return errors.New("入力形式が不正です。")
	}
	return nil
}

func (self *ipv4BroadcastAddr) addNewItem() {
	addDefaultItem(self.n)
}

func addDefaultItem(n *network) {
	wf.NewItem(`/` + strconv.Itoa(n.prefix)).
		Subtitle(`network prefix`).
		Icon(&aw.Icon{Value: "./img/prefix.png", Type: ""})
	wf.NewItem(n.subnetMask.String()).
		Subtitle(`subnet mask`).
		Icon(&aw.Icon{Value: "./img/subnet_mask.png", Type: ""})
}
