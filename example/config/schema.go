// Code generated by genfig (schema built from 'default.yml') on 2019-07-17T21:56:11+02:00; DO NOT EDIT.

package config

type Config struct {
	Secrets    []string
	Apis       ConfigApis
	Wip        bool
	Version    string
	Project    string
	Server     ConfigServer
	Db         ConfigDb
	Randomizer ConfigRandomizer
	LongDesc   ConfigLongDesc
}

type ConfigApis struct {
	Google ConfigApisGoogle
}

type ConfigApisGoogle struct {
	Uri string
}

type ConfigDb struct {
	Pass string
	Uri  string
	User string
}

type ConfigLongDesc struct {
	En string
	De string
}

type ConfigRandomizer struct {
	Threshold float64
}

type ConfigServer struct {
	Port int64
	Host string
}
