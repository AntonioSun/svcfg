////////////////////////////////////////////////////////////////////////////
// Porgram: config - Config handling
// Authors: Antonio Sun (c) 2015-2016, All rights reserved
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

   instance:

    - database: perfwhit746
      weekadj: 5
      basedate: 2016-06-19
      password: s3cr3t
      dbserver: TorsvPerfDb07
      servers: TorsvPerfBje05 TorsvPerfBje06 TorsvPerfApp03 TorsvPerfApp06

    - database: perfwhit746b
      weekadj: 5
      basedate: 2016-06-19
      password: s3cr3t
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
	CurDFv   string
}

type pod struct {
	Id       string
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
