package wls

import (
	"bytes"
	"encoding/json"
	"fmt"

	ui "github.com/gizak/termui"
)

// Server is a specific Server instance deployed to the domain under the given AdminServer
type Server struct {
	Name                    string  `json:"name"`
	State                   string  `json:"state"`
	Health                  string  `json:"health"`
	ClusterName             string  `json:"clusterName,omitempty"`
	CurrentMachine          string  `json:",omitempty"`
	WebLogicVersion         string  `json:",omitempty"`
	OpenSocketsCurrentCount float64 `json:",omitempty"`
	HeapSizeCurrent         int     `json:",omitempty"`
	HeapFreeCurrent         int     `json:",omitempty"`
	JavaVersion             string  `json:",omitempty"`
	OsName                  string  `json:",omitempty"`
	OsVersion               string  `json:",omitempty"`
	JvmProcessorLoad        float64 `json:",omitempty"`
}

// NewWidget will create a termui widget
func (s *Server) NewWidget() ui.Bufferer {
	uis := ui.NewPar(s.Name)
	// TODO: add rest of stuff
	ui.NewRow()
	return uis
}

// GoString produces a GoString formatted for the console and a CLI interface.
func (s *Server) GoString() string {
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("Name:        %-14v| State:           %-14v| Health:        %-55v\n", s.Name, s.State, s.Health))
	buffer.WriteString(fmt.Sprintf("Cluster:     %-14v| CurrentMachine:  %-14v| JVM Load:      %-55v\n", s.ClusterName, s.CurrentMachine, s.JvmProcessorLoad))
	buffer.WriteString(fmt.Sprintf("Sockets #:   %-14v| Heap Sz Cur:     %-14v| Heap Free Cur: %-55v\n", s.OpenSocketsCurrentCount, s.HeapSizeCurrent, s.HeapFreeCurrent))
	buffer.WriteString(fmt.Sprintf("Java Ver:    %-14v| OS Name:         %-14v| OS Version:    %-55v\n", s.JavaVersion, s.OsName, s.OsVersion))
	buffer.WriteString(fmt.Sprintf("WLS Version: %-14v\n", s.WebLogicVersion))
	return buffer.String()
}

// Servers returns all servers configured in a domain and provides run-time information for each server, including the server state and health.
// isFullFormat determines whether to return a fully-filled out list of Servers, or only a shortened version of the Servers list.
func (a *AdminServer) Servers(isFullFormat bool) ([]Server, error) {
	url := fmt.Sprintf("%v%v/servers", a.AdminURL, MonitorPath)
	if isFullFormat {
		url = url + "?format=full"
	}
	w, err := requestAndUnmarshal(url, a)
	if err != nil {
		return nil, err
	}
	var servers []Server
	if err := json.Unmarshal(w.Body.Items, &servers); err != nil {
		return nil, err
	}
	return servers, nil
}

// Server returns information for a specified server in a domain, including the server state, health, and JVM heap availability.
func (a *AdminServer) Server(serverName string) (*Server, error) {
	url := fmt.Sprintf("%v%v/servers/%v", a.AdminURL, MonitorPath, serverName)
	w, err := requestAndUnmarshal(url, a)
	if err != nil {
		return nil, err
	}
	var server Server
	if err := json.Unmarshal(w.Body.Item, &server); err != nil {
		return nil, err
	}
	return &server, nil
}
