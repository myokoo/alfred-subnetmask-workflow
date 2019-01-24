package main

import (
	"github.com/deanishe/awgo"
	"github.com/docopt/docopt-go"
	"net"
	"os"
	"strings"
)

type networkManagement interface {
	calculate() error
	addNewItem()
}

type network struct {
	query         string
	qType         networkManagement
	ip            net.IP
	prefix        int
	subnetMask    net.IP
	networkAddr   net.IP
	broadcastAddr net.IP
	wildcardMask  net.IP
}

var wf *aw.Workflow

func init() {
	wf = aw.New()
}

func main() {

	usage := `alfred-subnetmask-workflow

Usage:
  alfred-subnetmask-workflow <query>
  alfred-subnetmask-workflow -h | --help
  alfred-subnetmask-workflow -v | --version

Options:
  -h            Show this screen.
  --version     Show version.
`
	opts, _ := docopt.ParseArgs(usage, os.Args[1:], "1.0.0")

	network := new(network)
	network.query, _ = opts.String("<query>")
	network.applyProcessType()

	if err := network.qType.calculate(); err != nil {
		wf.NewItem("入力内容が不正です。")
		wf.SendFeedback()
		return
	}
	network.qType.addNewItem()
	wf.SendFeedback()
}

func (n *network) applyProcessType() {
	if strings.HasPrefix(n.query, `/`) {
		// ex: /24
		n.qType = &ipv4Prefix{n}
	} else if strings.Contains(n.query, `/`) {
		// ex: 192.168.24.11/24 or 192.168.24.11/255.255.255.0
		n.qType = &ipv4Network{n}
	} else {
		// ex: 255.255.255.0
		n.qType = &ipv4SubnetMask{n}
	}
}
