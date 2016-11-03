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
      dbserver: TorsvPerfDb07
      servers: TorsvPerfBje05 TorsvPerfBje06 TorsvPerfApp03 TorsvPerfApp06

    - database: perfwhit746b
      dbserver: TorsvPerfDb07
      servers: TorsvPerfBje05 TorsvPerfBje06 TorsvPerfApp03 TorsvPerfApp06

*/

type instance struct {
	Database string
	Dbserver string
	Servers  string
	BaseDate string
	Password string
	Serversa []string
}

type pod struct {
	Id       string
	Instance []instance
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
