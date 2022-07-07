package stat

import (
	"fmt"
	"runtime"
	"strings"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/net"
)

type NetStat struct {
	IP     string
	MAC    string
	Driver string
	// Network
	TotalBytesSent    int
	TotalBytesSentAvg float64
	TotalBytesRecv    int
	TotalBytesRecvAvg float64
	InBytesRecv       int
	InBytesRecvOld    int
	InBytesRecvAvg    float64
	InBytesRecvDiff   int
	OutBytesRecv      int
	OutBytesRecvOld   int
	OutBytesRecvAvg   float64
	OutBytesRecvDiff  int
	InBytesSent       int
	InBytesSentOld    int
	InBytesSentAvg    float64
	InBytesSentDiff   int
}

var (
	snapshot_len int32         = 1024
	promiscuous  bool          = false
	timeout      time.Duration = 30 * time.Second
	handle       *pcap.Handle
)

func NewNetStat(timing int) *NetStat {
	ip, mac, _, _ := getIP()
	netStat := &NetStat{
		IP:     ip,
		MAC:    mac,
		Driver: "",
		// Network
		TotalBytesSent:    0,
		TotalBytesSentAvg: 0,
		TotalBytesRecv:    0,
		TotalBytesRecvAvg: 0,
		InBytesRecv:       0,
		InBytesRecvAvg:    0,
		InBytesRecvDiff:   0,
		OutBytesRecv:      0,
		OutBytesRecvAvg:   0,
		OutBytesRecvDiff:  0,
		InBytesSent:       0,
		InBytesSentAvg:    0,
		InBytesSentDiff:   0,
	}
	go netStat.TimingStat(timing)
	return netStat
}

func (s *NetStat) TimingStat(timing int) {
	go s.InOutIODiffStat() //windows 下gopacket 抓包部分设备下无法使用，获取不到网卡信息
	for {
		bootTime, _ := host.BootTime()
		now := uint64(time.Now().Unix())
		info, _ := net.IOCounters(true)
		for _, v := range info {
			s.TotalBytesSent += int(v.BytesSent)
			s.TotalBytesRecv += int(v.BytesRecv)
		}
		s.TotalBytesSentAvg = float64(s.TotalBytesSent) / float64(now-bootTime)
		s.TotalBytesRecvAvg = float64(s.TotalBytesRecv) / float64(now-bootTime)
		s.InBytesRecvAvg = float64(s.InBytesRecv) / float64(now-bootTime)
		s.OutBytesRecvAvg = float64(s.OutBytesRecv) / float64(now-bootTime)
		s.InBytesSentAvg = float64(s.InBytesSent) / float64(now-bootTime)
		s.InBytesRecvDiff = s.InBytesRecv - s.InBytesRecvOld
		s.OutBytesRecvDiff = s.OutBytesRecv - s.OutBytesRecvOld
		s.InBytesSentDiff = s.InBytesSent - s.InBytesSentOld
		s.InBytesRecvOld = s.InBytesRecv
		s.OutBytesRecvOld = s.OutBytesRecv
		s.InBytesSentOld = s.InBytesSent
		time.Sleep(time.Second * time.Duration(timing))
	}
}

func (s *NetStat) InOutIODiffStat() {
	devices, err := pcap.FindAllDevs()
	if err != nil {
		fmt.Println("err", err.Error())
	}
	for _, device := range devices {
		for _, address := range device.Addresses {
			if strings.HasPrefix(address.IP.String(), "192.168") {
				s.Driver = device.Name
			}
		}
	}
	if s.Driver == "" {
		runtime.Goexit() // 终止当前 goroutine
	}
	handle, err := pcap.OpenLive(s.Driver, snapshot_len, promiscuous, timeout)
	if err != nil {
		fmt.Println("err", err.Error())
	}
	defer handle.Close()
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {
		// Process packet here
		ipLayer := packet.Layer(layers.LayerTypeIPv4)
		if ipLayer != nil {
			ip, _ := ipLayer.(*layers.IPv4)
			if ip.DstIP.String() == s.IP {
				// in
				if strings.HasPrefix(ip.SrcIP.String(), string([]byte(s.IP)[:9])) {
					s.InBytesRecv += int(ip.Length)
				} else {
					s.OutBytesRecv += int(ip.Length)
				}
			} else {
				// out
				if strings.HasPrefix(ip.DstIP.String(), string([]byte(s.IP)[:9])) {
					s.InBytesSent += int(ip.Length)
				}
			}

		}
	}
}

func getIP() (string, string, string, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		fmt.Println(err)
	}
	for _, iface := range ifaces {
		addrs := iface.Addrs
		mac := iface.HardwareAddr
		if err != nil {
			return "", "", "", err
		}
		if strings.HasPrefix(addrs[1].Addr, "192.168.2") {
			return string(addrs[1].Addr[:len(addrs[1].Addr)-3]), mac, iface.Name, nil
		}
	}
	return "", "", "", nil
}

// func NetStatTest() {
// 	ip, mac, _ := getIP()
// 	fmt.Println(ip, mac)
// 	info, _ := net.IOCounters(true)
// 	for index, v := range info {
// 		fmt.Printf("%v:%v send:%v recv:%v\n", index, v, v.BytesSent, v.BytesRecv)
// 	}
// }

// func TestGoPacket() {
// 	devices, err := pcap.FindAllDevs()
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	for _, device := range devices {
// 		fmt.Println("\nName: ", device.Name)
// 		fmt.Println("Description: ", device.Description)
// 		fmt.Println("Devices addresses:", device.Description)
// 		for _, address := range device.Addresses {
// 			fmt.Println("- IP address: ", address.IP)
// 			fmt.Println("- Subnet mask: ", address.Netmask)
// 			if strings.HasPrefix(address.IP.String(), "192.168.1") {
// 				deviceName = device.Name
// 				break
// 			}
// 		}
// 	}
// 	handle, err = pcap.OpenLive(devicea, snapshot_len, promiscuous, timeout)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer handle.Close()

// 	// var filter string = "tcp and port 80"
// 	// err = handle.SetBPFFilter(filter)
// 	// if err != nil {
// 	// 	log.Fatal(err)
// 	// }
// 	// Use the handle as a packet source to process all packets
// 	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
// 	for packet := range packetSource.Packets() {
// 		// Process packet here
// 		// var payload gopacket.Payload
// 		// payload = packet.Layers()[1].LayerPayload()
// 		// fmt.Println(payload.DecodeFromBytes(payload))
// 		// fmt.Println(packet.Layers()[1].LayerType().String())
// 		// fmt.Println(packet.String())
// 		ipLayer := packet.Layer(layers.LayerTypeIPv4)
// 		if ipLayer != nil {
// 			ip, _ := ipLayer.(*layers.IPv4)
// 			if ip.SrcIP.String() == "192.168.1.103"{

// 				// in
// 				if strings.HasPrefix(ip.SrcIP.String(), "192.168.1"){

// 				}else{

// 				}

// 			}else{

// 		}
// 		// 	if packet.Layer(layers.LayerTypeTCP) != nil {

// 		// 	}
// 	}
// }

// func printPacketInfo(packet gopacket.Packet) {
// 	// Let's see if the packet is an ethernet packet
// 	ethernetLayer := packet.Layer(layers.LayerTypeEthernet)
// 	if ethernetLayer != nil {
// 		fmt.Println("Ethernet layer detected.")
// 		ethernetPacket, _ := ethernetLayer.(*layers.Ethernet)
// 		fmt.Println("Source MAC: ", ethernetPacket.SrcMAC)
// 		fmt.Println("Destination MAC: ", ethernetPacket.DstMAC)
// 		// Ethernet type is typically IPv4 but could be ARP or other
// 		fmt.Println("Ethernet type: ", ethernetPacket.EthernetType)
// 		fmt.Println()
// 	}

// 	// Let's see if the packet is IP (even though the ether type told us)
// 	ipLayer := packet.Layer(layers.LayerTypeIPv4)
// 	if ipLayer != nil {
// 		fmt.Println("IPv4 layer detected.")
// 		ip, _ := ipLayer.(*layers.IPv4)

// 		// IP layer variables:
// 		// Version (Either 4 or 6)
// 		// IHL (IP Header Length in 32-bit words)
// 		// TOS, Length, Id, Flags, FragOffset, TTL, Protocol (TCP?),
// 		// Checksum, SrcIP, DstIP
// 		fmt.Printf("From %s to %s\n", ip.SrcIP, ip.DstIP)
// 		fmt.Println("Protocol: ", ip.Protocol)
// 		fmt.Println()
// 	}

// 	// Let's see if the packet is TCP
// 	tcpLayer := packet.Layer(layers.LayerTypeTCP)
// 	if tcpLayer != nil {
// 		fmt.Println("TCP layer detected.")
// 		tcp, some := tcpLayer.(*layers.TCP)
// 		// tcp := tcpLayer
// 		fmt.Println(reflect.TypeOf(tcp), reflect.TypeOf(tcpLayer))
// 		// fmt.Println(tcpLayer)
// 		fmt.Println("some=", some)

// 		// TCP layer variables:
// 		// SrcPort, DstPort, Seq, Ack, DataOffset, Window, Checksum, Urgent
// 		// Bool flags: FIN, SYN, RST, PSH, ACK, URG, ECE, CWR, NS
// 		// fmt.Printf("From port %d to %d\n", tcpLayer.SrcPort, tcpLayer.DstPort)
// 		fmt.Printf("From port %d to %d\n", tcp.SrcPort, tcp.DstPort)
// 		fmt.Println("Sequence number: ", tcp.Seq)
// 		fmt.Println()
// 	}

// 	// Iterate over all layers, printing out each layer type
// 	fmt.Println("All packet layers:")
// 	for _, layer := range packet.Layers() {
// 		fmt.Println("- ", layer.LayerType())
// 	}

// 	// When iterating through packet.Layers() above,
// 	// if it lists Payload layer then that is the same as
// 	// this applicationLayer. applicationLayer contains the payload
// 	applicationLayer := packet.ApplicationLayer()
// 	if applicationLayer != nil {
// 		fmt.Println("Application layer/Payload found.")
// 		fmt.Printf("%s\n", applicationLayer.Payload())

// 		// Search for a string inside the payload
// 		if strings.Contains(string(applicationLayer.Payload()), "HTTP") {
// 			fmt.Println("HTTP found!")
// 		}
// 	}

// 	// Check for errors
// 	if err := packet.ErrorLayer(); err != nil {
// 		fmt.Println("Error decoding some part of the packet:", err)
// 	}
// }
