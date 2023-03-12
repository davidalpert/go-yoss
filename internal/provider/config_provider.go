package provider

import (
	"github.com/spf13/pflag"
)

type Interface interface {
	GetValue(key string) (string, error)
	GetValueTree(prefix string) (map[string]string, error)
	SetValue(key string, value string) error
}

type Options struct {
	ProviderName string
	Provider     Interface
	Region       string // a.k.a. Datacenter
	Namespace    string // a.k.a. Environment
	OutDir       string
	Debug        bool
}

// AddProviderOptions adds flags to a pflag.FlagSet
func (o *Options) AddProviderOptions(c *pflag.FlagSet) {
	c.StringVar(&o.Region, "region", "default", "region")
	c.StringVarP(&o.Namespace, "namespace", "n", "default", "namespace")
	c.StringVar(&o.OutDir, "out-dir", "out", "output dir for result or any diagnostics")
}

type NewProviderFn = func(o *Options) (Interface, error)
