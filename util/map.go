package util

import "net"

type HostMap struct {
	hostMap map[string]net.IP
	myHost  string
	myIP    net.IP
}

func NewHostMap(host string, ip net.IP) *HostMap {
	initMap := make(map[string]net.IP)
	return &HostMap{
		hostMap: initMap,
		myHost:  host,
		myIP:    ip}
}

func (m *HostMap) GetIP(host string) net.IP {
	//TODO check
	return m.hostMap[host]
}

func (m *HostMap) GetHost(ip net.IP) string {
	//TODO implement
	return ""
}

func (m *HostMap) SetMine(host string, ip net.IP) {
	m.myHost = host
	m.myIP = ip
}

func (m *HostMap) Add(host string, ip net.IP) {
	if host != m.myHost {
		m.hostMap[host] = ip
	}
}

func (m *HostMap) Remove(host string) {
	//TODO unset(m.hostMap[host])
}

func (m *HostMap) Empty(host string) {
	//TODO unset(m.hostMap[host])
}

func (m *HostMap) String() string {
	s := m.myHost + "\t\t" + m.myIP.String()
	for host, ip := range m.hostMap {
		if host != m.myHost {
			s = s + "\n" + host + "\t\t" + ip.String()
		}
	}
	//TODO String Append Performance Improve!
	return s
}

func (m *HostMap) Dump() string {
	//TODO implement: dump map to files
	return ""
}
