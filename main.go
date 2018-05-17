package main

import (
	"net"
	"os"
	"net/http"
	"log"
	"github.com/robfig/cron"
	"encoding/json"
	"bytes"
	"html/template"
	"path/filepath"
)

type NodeInformation struct {
	Hostname    string
	IpAddresses []string
}

func main() {
	// only one argument means no argument passed while app was executed
	isServer := len(os.Args) == 1

	nodes := make(map[string]NodeInformation, 0)

	if isServer {
		runServerMode(nodes)
	} else {
		runClientMode()
	}

}
// Returns the local hostname and ip addresses
func getNodeInformation() *NodeInformation {
	return &NodeInformation{
		Hostname:    getHostname(),
		IpAddresses: getIpAddresses(),
	}
}
// The function runServerMode start ipNotifier in server mode.
// The server provides two endpoints.
//
// '/' will show a list of all clients, that sent their ip addresse to the server.
//
// '/update' endpoint handles JSON post calls from ipNotifier clients.
func runServerMode(nodes map[string]NodeInformation) {
	dir := filepath.Dir(os.Args[0])
	t, _ := template.ParseFiles(dir + "/index.html")
	nodeInformation := getNodeInformation();
	nodes[nodeInformation.Hostname] = *nodeInformation

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t.Execute(w, nodes)
	})

	http.HandleFunc("/update", func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		var node NodeInformation
		err := decoder.Decode(&node)
		if err != nil {
			panic(err)
		}
		defer r.Body.Close()
		log.Println(node)

		nodes[node.Hostname] = node
		log.Println(nodes)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}

// The function runClientMode start ipNotifier in client mode
// where the local ip addresses will be send via a rest POST call
// to the serverAddress every minute. The server needs to start
// ipNotifier in server mode.
func runClientMode() {
	serverAddress := os.Args[1]
	c := cron.New()
	c.AddFunc("0 * * * * *", func() {
		b := new(bytes.Buffer)
		json.NewEncoder(b).Encode(*getNodeInformation())
		res, _ := http.Post(serverAddress, "application/json; charset=utf-8", b)
		log.Println(res)
	})
	c.Start()
	// run forever
	select {}
}

// Returns all local ip addresses
func getIpAddresses() []string {
	var addresses []string
	ifaces, _ := net.Interfaces()
	for _, i := range ifaces {
		addrs, _ := i.Addrs()
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			addresses = append(addresses, ip.String())
		}
	}
	return addresses
}

// Returns the local hostname
func getHostname() string {
	hn, _ := os.Hostname()
	return hn
}
