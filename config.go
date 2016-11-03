////////////////////////////////////////////////////////////////////////////
// Porgram: config - Config handling
// Authors: Antonio Sun (c) 2015, All rights reserved
////////////////////////////////////////////////////////////////////////////

package main

import (
	"io/ioutil"
	"strings"
)

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

var config struct {
	DbUser     string
	DbPassword string
	Pod        []pod
	ENV        map[string]string
	Today      string
}

////////////////////////////////////////////////////////////////////////////
// Function definitions

func configGet() {
	cfgStr, err := ioutil.ReadFile(rootArgv.SvConfig)
	err = yaml.Unmarshal(cfgStr, &config)
	check(err)
	//fmt.Printf("] %+v\r\n", config)
}

func Basename(s string) string {
	n := strings.LastIndexByte(s, '.')
	if n > 0 {
		return s[:n]
	}
	return s
}
