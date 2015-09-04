package wls

import (
	"encoding/json"
	"fmt"
	"testing"
)

var servers_json = `{
  "body": {
    "items": [
      {
        "name": "adminserver",
        "state": "RUNNING",
        "health": " HEALTH_OK "
      },
      {
        "name": "ms1",
        "state": "SHUTDOWN",
        "health": ""
      }
     ]
   },
   "messages": [
  ]
 }`

func TestUnmarshalServersJson(t *testing.T) {
	wrapper := &ServerWrapper{}
	if err := json.Unmarshal([]byte(servers_json), wrapper); err != nil {
		t.Error(err)
	}
	t.Log(wrapper)
	var servers_json_tests = []struct {
		in  string
		out string
	}{
		{wrapper.Body.Items[0].Health, " HEALTH_OK "},
		{wrapper.Body.Items[0].Name, "adminserver"},
		{wrapper.Body.Items[0].State, "RUNNING"},
		{wrapper.Body.Items[1].Name, "ms1"},
		{wrapper.Body.Items[1].State, "SHUTDOWN"},
		{wrapper.Body.Items[1].Health, ""},
	}

	for _, tt := range servers_json_tests {
		if tt.in != tt.out {
			t.Errorf("want %q, got %q", tt.out, tt.in)
		}
	}

}

var single_server = `{
  "body": {
    "item": {
      "name": "adminserver",
      "clusterName": null,

      "state": "RUNNING",

      "currentMachine": "machine-0",
      "weblogicVersion": "WebLogic Server 12.1.1.0.0 Thu May 5 01:17:16 2011 PDT",
      "openSocketsCurrentCount": 2,
      "health": "HEALTH_OK",

      "heapSizeCurrent": 536870912,
      "heapFreeCurrent": 39651944,
      "heapSizeMax": 1073741824,
      "javaVersion": "1.6.0_20",
      "osName": "Linux",
      "osVersion": "2.6.18-238.0.0.0.1.el5xen",

      "jvmProcessorLoad": 0.25
     }
    },
     "messages": [
    ]
  }`

func TestUnmarshalSingleServer(t *testing.T) {
	wrapper := &ServerWrapper{}
	if err := json.Unmarshal([]byte(single_server), wrapper); err != nil {
		t.Error(err)
	}
	//	t.Log(wrapper)
	var servers_json_tests = []struct {
		in  string
		out string
	}{
		{wrapper.Body.Item.Name, "adminserver"},
		{wrapper.Body.Item.ClusterName, ""},
		{wrapper.Body.Item.State, "RUNNING"},
		{wrapper.Body.Item.CurrentMachine, "machine-0"},
		{wrapper.Body.Item.WeblogicVersion, "WebLogic Server 12.1.1.0.0 Thu May 5 01:17:16 2011 PDT"},
		{fmt.Sprint(wrapper.Body.Item.OpenSocketsCurrentCount), "2"},
		{wrapper.Body.Item.Health, "HEALTH_OK"},
		{fmt.Sprint(wrapper.Body.Item.HeapSizeCurrent), "536870912"},
		{fmt.Sprint(wrapper.Body.Item.HeapFreeCurrent), "39651944"},
		{wrapper.Body.Item.JavaVersion, "1.6.0_20"},
		{wrapper.Body.Item.OsName, "Linux"},
		{wrapper.Body.Item.OsVersion, "2.6.18-238.0.0.0.1.el5xen"},
		{fmt.Sprint(wrapper.Body.Item.JvmProcessorLoad), "0.25"},
	}

	for _, tt := range servers_json_tests {
		if tt.in != tt.out {
			t.Errorf("want %q, got %q", tt.out, tt.in)
		}
	}
}

var single_cluster = `{
    "body": {
        "item": {
            "name": "mycluster1",
            "servers": [
                {
                    "name": "ms1",
                    "state": "RUNNING",
                    "health": "OK",
                    "clusterMaster": false,
                    "dropOutFrequency": "Never",
                    "resendRequestsCount": 0,
                    "fragmentsSentCount": 3708,
                    "fragmentsReceivedCount": 3631
                },
                {
                    "name": "ms2",
                    "state": "RUNNING",
                    "health": "OK"
                }
            ]
        }
    },
    "messages": []
}`

func TestUnmarshalSingleCluster(t *testing.T) {
	wrapper := &ClusterWrapper{}
	if err := json.Unmarshal([]byte(single_cluster), wrapper); err != nil {
		t.Error(err)
	}
	if len(wrapper.Body.Item.Servers) == 0 {
		t.Errorf("Servers in wrapper.Body.Item is 0, should be 2")
	}
	var servers_json_tests = []struct {
		in  string
		out string
	}{
		{wrapper.Body.Item.Name, "mycluster1"},
		{wrapper.Body.Item.Servers[0].Name, "ms1"},
		{wrapper.Body.Item.Servers[0].State, "RUNNING"},
		{wrapper.Body.Item.Servers[0].Health, "OK"},
		{fmt.Sprint(wrapper.Body.Item.Servers[0].IsClusterMaster), "false"},
		{wrapper.Body.Item.Servers[0].DropOutFrequency, "Never"},
		{fmt.Sprint(wrapper.Body.Item.Servers[0].ResendRequestsCount), "0"},
		{fmt.Sprint(wrapper.Body.Item.Servers[0].FragmentsSentCount), "3708"},
		{fmt.Sprint(wrapper.Body.Item.Servers[0].FragmentsReceivedCount), "3631"},
		{wrapper.Body.Item.Servers[1].Name, "ms2"},
		{wrapper.Body.Item.Servers[1].State, "RUNNING"},
		{wrapper.Body.Item.Servers[1].Health, "OK"},
		{fmt.Sprint(wrapper.Body.Item.Servers[1].DropOutFrequency), ""},
	}

	for _, tt := range servers_json_tests {
		if tt.in != tt.out {
			t.Errorf("want %q, got %q", tt.out, tt.in)
		}
	}
}

var clusters = `{
    "body": {
        "items": [
            {
                "name": "mycluster1",
                "servers": [
                    {
                        "name": "ms1",
                        "state": "RUNNING",
                        "health": "HEALTH_OK"
                    },
                    {
                        "name": "ms2",
                        "state": "RUNNING",
                        "health": "HEALTH_OVERLOADED"
                    }
                ]
            }
        ]
    },
    "messages": []
}`

func TestUnmarshalMultipleClusters(t *testing.T) {
	wrapper := &ClusterWrapper{}
	if err := json.Unmarshal([]byte(clusters), wrapper); err != nil {
		t.Error(err)
	}
	if len(wrapper.Body.Items) == 0 {
		t.Errorf("Clusters count should be 2, was 0")
	}
	var servers_json_tests = []struct {
		in  string
		out string
	}{
		{wrapper.Body.Items[0].Name, "mycluster1"},
		{wrapper.Body.Items[0].Servers[0].Name, "ms1"},
		{wrapper.Body.Items[0].Servers[0].State, "RUNNING"},
		{wrapper.Body.Items[0].Servers[0].Health, "HEALTH_OK"},
		{wrapper.Body.Items[0].Servers[1].Name, "ms2"},
		{wrapper.Body.Items[0].Servers[1].State, "RUNNING"},
		{wrapper.Body.Items[0].Servers[1].Health, "HEALTH_OVERLOADED"},
	}

	for _, tt := range servers_json_tests {
		if tt.in != tt.out {
			t.Errorf("want %q, got %q", tt.out, tt.in)
		}
	}
}

var applications = `{
    "body": {
        "items": [
            {
                "name": "appscopedejbs",
                "type": "ear",
                "state": "STATE_ACTIVE",
                "health": " HEALTH_OK"
            },
            {
                "name": "MyWebApp",
                "type": "war",
                "state": "STATE_NEW"
            }
        ]
    },
    "messages": []
}`

func TestUnmarshalMultipleApplications(t *testing.T) {
	wrapper := &ApplicationWrapper{}
	if err := json.Unmarshal([]byte(applications), wrapper); err != nil {
		t.Error(err)
	}
	if len(wrapper.Body.Items) == 0 {
		t.Errorf("Applications count should be 2, was 0")
	}
	var applications_json_tests = []struct {
		in  string
		out string
	}{
		{wrapper.Body.Items[0].Name, "appscopedejbs"},
		{wrapper.Body.Items[0].AppType, "ear"},
		{wrapper.Body.Items[0].State, "STATE_ACTIVE"},
		{wrapper.Body.Items[0].Health, " HEALTH_OK"},
		{wrapper.Body.Items[1].Name, "MyWebApp"},
		{wrapper.Body.Items[1].AppType, "war"},
		{wrapper.Body.Items[1].State, "STATE_NEW"},
	}

	for _, tt := range applications_json_tests {
		if tt.in != tt.out {
			t.Errorf("want %q, got %q", tt.out, tt.in)
		}
	}
}

var application = `{
    "body": {
        "item": {
            "name": "appscopedejbs",
            "type": "ear",
            "health": " HEALTH_OK ",
            "state": "STATE_ACTIVE",
            "targetStates": [
                {
                    "target": "ms1",
                    "state": "STATE_ACTIVE"
                },
                {
                    "target": "ms2",
                    "state": "STATE_ACTIVE"
                }
            ],
            "dataSources": [],
            "entities": [],
            "workManagers": [
                {
                    "name": "default",
                    "server": "ms1",
                    "pendingRequests": 0,
                    "completedRequests": 0
                }
            ],
            "minThreadsConstraints": [
                {
                    "name": "minThreadsConstraints-0",
                    "server": "ms1",
                    "completedRequests": 0,
                    "pendingRequests": 0,
                    "executingRequests": 0,
                    "outOfOrderExecutionCount": 0,
                    "mustRunCount": 0,
                    "maxWaitTime": 0,
                    "currentWaitTime": 0
                }
            ],
            "maxThreadsConstraints": [
                {
                    "name": "maxThreadsConstraints-0",
                    "server": "ms1",
                    "executingRequests": 0,
                    "deferredRequests": 0
                }
            ],
            "requestClasses": [
                {
                    "name": "requestClasses-0",
                    "server": "ms1",
                    "requestClassType": "fairshare",
                    "completedCount": 0,
                    "totalThreadUse": 0,
                    "pendingRequestCount": 0,
                    "virtualTimeIncrement": 0
                }
            ]
        }
    },
    "messages": []
}`

func TestUnmarshalApplication(t *testing.T) {
	wrapper := &ApplicationWrapper{}
	if err := json.Unmarshal([]byte(application), wrapper); err != nil {
		t.Error(err)
	}
	//	t.Log(wrapper.Body.Item)
	//	t.Log(wrapper.Body.Item.TargetStates[0])
	//	t.Log(wrapper.Body.Item.TargetStates[1])
	//	t.Log(wrapper.Body.Item.WorkManagers[0])
	t.Log(wrapper.Body.Item.MinThreadsConstraints[0])
	//	t.Log(wrapper.Body.Item.MaxThreadsConstraints[0])
	t.Log(wrapper.Body.Item.RequestClasses[0])
	var application_json_tests = []struct {
		in  string
		out string
	}{
		{wrapper.Body.Item.Name, "appscopedejbs"},
		{wrapper.Body.Item.AppType, "ear"},
		{wrapper.Body.Item.State, "STATE_ACTIVE"},
		{wrapper.Body.Item.Health, " HEALTH_OK "},
		{wrapper.Body.Item.TargetStates[0].State, "STATE_ACTIVE"},
		{wrapper.Body.Item.TargetStates[0].Target, "ms1"},
		{wrapper.Body.Item.TargetStates[1].State, "STATE_ACTIVE"},
		{wrapper.Body.Item.TargetStates[1].Target, "ms2"},
		{wrapper.Body.Item.WorkManagers[0].Name, "default"},
		{wrapper.Body.Item.WorkManagers[0].Server, "ms1"},
		{fmt.Sprint(wrapper.Body.Item.WorkManagers[0].PendingRequests), "0"},
		{fmt.Sprint(wrapper.Body.Item.WorkManagers[0].CompletedRequests), "0"},
		{wrapper.Body.Item.MinThreadsConstraints[0].Name, "minThreadsConstraints-0"},
		{wrapper.Body.Item.MinThreadsConstraints[0].Server, "ms1"},
		{fmt.Sprint(wrapper.Body.Item.MinThreadsConstraints[0].CompletedRequests), "0"},
		{fmt.Sprint(wrapper.Body.Item.MinThreadsConstraints[0].PendingRequests), "0"},
		{fmt.Sprint(wrapper.Body.Item.MinThreadsConstraints[0].ExecutingRequests), "0"},
		{fmt.Sprint(wrapper.Body.Item.MinThreadsConstraints[0].OutOfOrderExecutionCount), "0"},
		{fmt.Sprint(wrapper.Body.Item.MinThreadsConstraints[0].MustRunCount), "0"},
		{fmt.Sprint(wrapper.Body.Item.MinThreadsConstraints[0].MaxWaitTime), "0"},
		{fmt.Sprint(wrapper.Body.Item.MinThreadsConstraints[0].CurrentWaitTime), "0"},
		{wrapper.Body.Item.MaxThreadsConstraints[0].Name, "maxThreadsConstraints-0"},
		{wrapper.Body.Item.MaxThreadsConstraints[0].Server, "ms1"},
		{fmt.Sprint(wrapper.Body.Item.MaxThreadsConstraints[0].DeferredRequests), "0"},
		{fmt.Sprint(wrapper.Body.Item.MaxThreadsConstraints[0].ExecutingRequests), "0"},
		{wrapper.Body.Item.RequestClasses[0].Name, "requestClasses-0"},
		{wrapper.Body.Item.RequestClasses[0].Server, "ms1"},
		{wrapper.Body.Item.RequestClasses[0].RequestClassType, "fairshare"},
		{fmt.Sprint(wrapper.Body.Item.RequestClasses[0].CompletedCount), "0"},
		{fmt.Sprint(wrapper.Body.Item.RequestClasses[0].TotalThreadUse), "0"},
		{fmt.Sprint(wrapper.Body.Item.RequestClasses[0].PendingRequestCount), "0"},
		{fmt.Sprint(wrapper.Body.Item.RequestClasses[0].VirtualTimeIncrement), "0"},
	}

	for _, tt := range application_json_tests {
		if tt.in != tt.out {
			t.Errorf("want %q, got %q", tt.out, tt.in)
		}
	}
}
