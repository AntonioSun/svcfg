////////////////////////////////////////////////////////////////////////////
// Porgram: config - Config handling
// Authors: Antonio Sun (c) 2015-2018, All rights reserved
////////////////////////////////////////////////////////////////////////////

package svcfg

import "io/ioutil"

import (
	"gopkg.in/yaml.v2"
)

////////////////////////////////////////////////////////////////////////////
// Constant and data type/structure definitions

/*
dbuser: uu
dbpassword: "pp"

pod:

 - id: v746b

   # password, dbserver & servers under pod is mandatory
   password: s3cr3ct
   dbserver: TorsvPerfDb07
   servers: TorsvPerfBje05 TorsvPerfBje06 TorsvPerfApp03 TorsvPerfApp06
   instance:

    - database: perfwhit746
      weekadj: 5
      basedate: 2016-06-19

    - database: perfwhit746b
      weekadj: 5
      basedate: 2016-06-19
      # password, dbserver & servers under instance is optional, 
      # to overwrite the pod setting
      dbserver: TorsvPerfDb07
      servers: TorsvPerfBje05 TorsvPerfBje06 TorsvPerfApp03 TorsvPerfApp06

*/

type InstanceT struct {
	Database string
	WeekAdj  int
	Dbserver string
	Servers  string
	BaseDate string
	Password string
	Serversa []string
	// data ouside yaml definition, for templating at InstanceT level
	CurDFv string // current DF ver
	Id     string // pod Id from above level
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
	Pod        []pod
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
		return err
	}
	return yaml.Unmarshal(cfgStr, &Config)
}
