package main

import (
	"fmt"
	"github.com/deanishe/awgo"
	"github.com/pkg/errors"
	"math"
	"net"
	"strconv"
)

// struct for /22
type ipv4Prefix struct {
	n *network
}

// struct for 192.168.33.12/21
type ipv4Network struct {
	n *network
}

// struct for 255.255.252.0
type ipv4SubnetMask struct {
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
	if self.n.prefix > maxPrefix {
		return errors.New("入力形式が不正です。")
	}

	mask := net.CIDRMask(self.n.prefix, maxPrefix)
	self.n.subnetMask = net.IPv4bcast.Mask(mask)

	buildWildCardMask(self.n, mask)
	return nil
}

func (self *ipv4Prefix) addNewItem() {
	addDefaultItem(self.n)
}

func (self *ipv4Network) calculate() error {
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

	buildWildCardMask(self.n, ipv4Net.Mask)
	return nil
}

func (self *ipv4Network) addNewItem() {
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

func (self *ipv4SubnetMask) calculate() error {
	self.n.subnetMask = net.ParseIP(self.n.query).To4()
	if self.n.subnetMask == nil {
		return errors.New("入力形式が不正です。")
	}
	var b int
	ipMask := net.IPMask(self.n.subnetMask)
	self.n.prefix, b = ipMask.Size()
	if self.n.prefix == 0 && b == 0 {
		return errors.New("入力形式が不正です。")
	}
	buildWildCardMask(self.n, ipMask)
	return nil
}

func (self *ipv4SubnetMask) addNewItem() {
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
	v := int(math.Exp2(float64(maxPrefix - n.prefix)))
	wf.NewItem(fmt.Sprintf("%d", v)).
		Subtitle(`ip count`).
		Arg(fmt.Sprintf("%d", v)).
		Valid(true).
		Icon(&aw.Icon{Value: "./img/ip_count.png", Type: ""})
}

func buildWildCardMask(n *network, mask net.IPMask) {
	n.wildcardMask = make([]byte, ipSliceLen)
	for i, v := range mask {
		n.wildcardMask[i] = v ^ 255
	}
}
