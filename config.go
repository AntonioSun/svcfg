////////////////////////////////////////////////////////////////////////////
// Porgram: config - Config handling
// Authors: Antonio Sun (c) 2015-2018, All rights reserved
////////////////////////////////////////////////////////////////////////////

package svcfg

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v2"
)

////////////////////////////////////////////////////////////////////////////
// Constant and data type/structure definitions

/*
dbuser: uu
dbpassword: "pp"

# Password that can trickle down to Pod Instances
instpw: xyz
# whether to append PodId to Instance Database name
#appendid: true

pod:

 - id: v746b

   # password, dbserver & servers under pod is mandatory
   password: s3cr3ct
   dbserver: TorsvPerfDb07
   servers: TorsvPerfBje05 TorsvPerfBje06 TorsvPerfApp03 TorsvPerfApp06
   instance:

    - database: perfcust1
      weekadj: 5
      basedate: 2016-06-19

    - database: perfwhit
      weekadj: 5
      basedate: 2016-06-19
      # password, dbserver & servers under instance is optional,
      # to overwrite the pod setting
      password: n3ws3cr3ct
      dbserver: TorsvPerfDb07
      servers: TorsvPerfBje05 TorsvPerfBje06 TorsvPerfApp03 TorsvPerfApp06

*/

type InstanceT struct {
	Database string
	WeekAdj  int
	Dbserver string
	Servers  string `json:"-"` // Servers str, ignored in json output
	BaseDate string
	Password string
	Servera  []string // Servers array
	// data ouside yaml definition, for templating at InstanceT level
	CurDFv  string `json:"-"` // current DF ver, ignored in json output
	Version string // pod Id (Version) from above level
}

type pod struct {
	Id       string
	Password string
	Dbserver string
	Servers  string
	Instance []InstanceT
}

var Config struct {
	DbUser     string
	DbPassword string
	InstPW     string // Password that can trickle down to Pod Instances
	Pod        []pod
	AppendId   bool // whether to append PodId to Instance Database name
	ENV        map[string]string
	Today      string
}

////////////////////////////////////////////////////////////////////////////
// Function definitions

// ConfigGet will get the servers config from the servers definition file
// svConfig and save in the module variable Config
func ConfigGet(svConfig string) error {
	cfgStr, err := ioutil.ReadFile(svConfig)
	if err != nil {
		ex, e := os.Executable()
		if e != nil {
			return e
		}
		svConfig = filepath.Dir(ex) + string(filepath.Separator) + svConfig
		cfgStr, err = ioutil.ReadFile(svConfig)
		if err != nil {
			return err
		}
	}

	err = yaml.Unmarshal(cfgStr, &Config)
	if err != nil {
		return err
	}

	// fill optional instance level setting from pod, if empty
	for pp, pods := range Config.Pod {
		if len(pods.Password) == 0 {
			Config.Pod[pp].Password = Config.InstPW
		}
		for ii, inst := range pods.Instance {
			pods.Instance[ii].Version = pods.Id
			if Config.AppendId {
				pods.Instance[ii].Database += pods.Id
			}
			if len(inst.Password) == 0 {
				pods.Instance[ii].Password = Config.Pod[pp].Password
			}
			if len(inst.Dbserver) == 0 {
				pods.Instance[ii].Dbserver = pods.Dbserver
			}
			if len(inst.Servers) == 0 {
				pods.Instance[ii].Servers = pods.Servers
			}
			// expand from "Servers" to "Servera"
			pods.Instance[ii].Servera = strings.Fields(pods.Instance[ii].Servers)
		}
	}

	return nil
}

// GetInst will get the *Instance, given the PotID and Instance prefix
func GetInst(pid, sid string) *InstanceT {
	for _, pods := range Config.Pod {
		if pods.Id != pid {
			continue
		}
		for _, inst := range pods.Instance {
			if sid != inst.Database {
				continue
			}
			return &inst
		}
	}
	return nil
}
