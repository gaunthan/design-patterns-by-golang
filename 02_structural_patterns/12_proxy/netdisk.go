package proxy

import "fmt"

type NetDisk interface {
	Get(user, filename string) string
}

type NetDiskImpl struct{}

func (NetDiskImpl) Get(user, filename string) string {
	return fmt.Sprintf("data of %s", filename)
}

type NetDiskProxy struct {
	diskImpl NetDiskImpl
}

func NewNetDiskProxy() *NetDiskProxy {
	return &NetDiskProxy{}
}

func (p NetDiskProxy) Get(user, filename string) string {
	// check access authority
	if !hasFileAuthority(user, filename) {
		return ""
	}

	data := p.diskImpl.Get(user, filename)

	// update access statistics
	updateAccessStatistics(user, filename)

	return data
}

func hasFileAuthority(user, filename string) bool {
	return user != ""
}

func updateAccessStatistics(user, filename string) {}
