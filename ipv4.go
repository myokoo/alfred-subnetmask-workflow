package main

import (
	"fmt"
	"github.com/deanishe/awgo"
	"github.com/pkg/errors"
	"math"
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

const maxPrefix = 32
const ipSliceLen = 4

func (self *ipv4Prefix) calculate() error {
	var err error
	self.n.prefix, err = strconv.Atoi(self.n.query[1:])
	if err != nil {
		return err
	}

	mask := net.CIDRMask(self.n.prefix, 32)
	self.n.subnetMask = net.IPv4bcast.Mask(mask)

	self.n.wildcardMask = make([]byte, ipSliceLen)
	for i, v := range mask {
		self.n.wildcardMask[i] = v ^ 255
	}
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

	self.n.broadcastAddr = make([]byte, ipSliceLen)
	copy(self.n.broadcastAddr, self.n.ip.To4())
	for i, v := range ipv4Net.Mask {
		self.n.broadcastAddr[i] |= v ^ 255
	}

	self.n.wildcardMask = make([]byte, ipSliceLen)
	for i, v := range ipv4Net.Mask {
		self.n.wildcardMask[i] = v ^ 255

	}
	return nil
}

func (self *ipv4SubnetMask) addNewItem() {
	wf.NewItem(self.n.ip.String()).
		Subtitle(`ip`).
		Arg(self.n.ip.String()).
		Valid(true).
		Icon(&aw.Icon{Value: "./img/ip.png", Type: ""})
	wf.NewItem(self.n.networkAddr.String() + " ~ " + self.n.broadcastAddr.String()).
		Subtitle(`network range`).
		Arg(self.n.networkAddr.String() + " ~ " + self.n.broadcastAddr.String()).
		Valid(true).
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
	if self.n.prefix == 0 && !ip.Equal(net.IPv4zero) {
		return errors.New("入力形式が不正です。")
	}

	self.n.wildcardMask = make([]byte, ipSliceLen)
	for i, v := range mask {
		self.n.wildcardMask[i] = v ^ 255
	}
	return nil
}

func (self *ipv4BroadcastAddr) addNewItem() {
	addDefaultItem(self.n)
}

func addDefaultItem(n *network) {
	wf.NewItem(`/` + strconv.Itoa(n.prefix)).
		Subtitle(`network prefix`).
		Arg(`/` + strconv.Itoa(n.prefix)).
		Valid(true).
		Icon(&aw.Icon{Value: "./img/prefix.png", Type: ""})
	wf.NewItem(n.subnetMask.String()).
		Subtitle(`subnet mask`).
		Arg(n.subnetMask.String()).
		Valid(true).
		Icon(&aw.Icon{Value: "./img/subnet_mask.png", Type: ""})
	wf.NewItem(n.wildcardMask.String()).
		Subtitle(`wildcard mask`).
		Arg(n.wildcardMask.String()).
		Valid(true).
		Icon(&aw.Icon{Value: "./img/wildcard_mask.png", Type: ""})
	v := int(math.Exp2(float64(maxPrefix-n.prefix)))
	wf.NewItem(fmt.Sprintf("%d", v)).
		Subtitle(`ip count`).
		Arg(fmt.Sprintf("%d", v)).
		Valid(true).
		Icon(&aw.Icon{Value: "./img/ip_count.png", Type: ""})
}
