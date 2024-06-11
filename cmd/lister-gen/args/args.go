/*
Copyright 2017 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package args

import (
	"fmt"

	"github.com/spf13/pflag"
	clientgenargs "k8s.io/code-generator/cmd/client-gen/args"
	"k8s.io/code-generator/cmd/client-gen/types"
)

// Args is used by the gengo framework to pass args specific to this generator.
type Args struct {
	OutputDir    string // must be a directory path
	OutputPkg    string // must be a Go import-path
	GoHeaderFile string

	// PluralExceptions specify list of exceptions used when pluralizing certain types.
	// For example 'Endpoints:Endpoints', otherwise the pluralizer will generate 'Endpointes'.
	PluralExceptions []string

	// Overrides for which types should be included in the informer.
	IncludedTypesOverrides map[types.GroupVersion][]string
}

// New returns default arguments for the generator.
func New() *Args {
	return &Args{}
}

// AddFlags add the generator flags to the flag set.
func (args *Args) AddFlags(fs *pflag.FlagSet) {
	fs.StringVar(&args.OutputDir, "output-dir", "",
		"the base directory under which to generate results")
	fs.StringVar(&args.OutputPkg, "output-pkg", "",
		"the base Go import-path under which to generate results")
	fs.StringSliceVar(&args.PluralExceptions, "plural-exceptions", args.PluralExceptions,
		"list of comma separated plural exception definitions in Type:PluralizedType format")
	fs.StringVar(&args.GoHeaderFile, "go-header-file", "",
		"the path to a file containing boilerplate header text; the string \"YEAR\" will be replaced with the current 4-digit year")
	pflag.Var(clientgenargs.NewGVTypesValue(&args.IncludedTypesOverrides, []string{}), "included-types-overrides", "list of group/version/type for which lister should be generated. By default, lister is generated for all types which have genclient in types.go. This overrides that. For each groupVersion in this list, only the types mentioned here will be included. The default check of genclient will be used for other group versions.")
}

// Validate checks the given arguments.
func (args *Args) Validate() error {
	if len(args.OutputDir) == 0 {
		return fmt.Errorf("--output-dir must be specified")
	}
	if len(args.OutputPkg) == 0 {
		return fmt.Errorf("--output-pkg must be specified")
	}
	return nil
}
