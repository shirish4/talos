// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package files

import (
	"io/fs"

	"github.com/cosi-project/runtime/pkg/resource"
	"github.com/cosi-project/runtime/pkg/resource/meta"
	"github.com/cosi-project/runtime/pkg/resource/typed"
)

//go:generate deep-copy -type EtcFileSpecSpec -type EtcFileStatusSpec -header-file ../../../../hack/boilerplate.txt -o deep_copy.generated.go .

// EtcFileSpecType is type of EtcFile resource.
const EtcFileSpecType = resource.Type("EtcFileSpecs.files.talos.dev")

// EtcFileSpec resource holds contents of the file which should be put to `/etc` directory.
type EtcFileSpec = typed.Resource[EtcFileSpecSpec, EtcFileSpecMD]

// EtcFileSpecSpec describes status of rendered secrets.
type EtcFileSpecSpec struct {
	Contents []byte      `yaml:"contents"`
	Mode     fs.FileMode `yaml:"mode"`
}

// NewEtcFileSpec initializes a EtcFileSpec resource.
func NewEtcFileSpec(namespace resource.Namespace, id resource.ID) *EtcFileSpec {
	return typed.NewResource[EtcFileSpecSpec, EtcFileSpecMD](
		resource.NewMetadata(namespace, EtcFileSpecType, id, resource.VersionUndefined),
		EtcFileSpecSpec{},
	)
}

// EtcFileSpecMD provides auxiliary methods for EtcFileSpec.
type EtcFileSpecMD struct{}

// ResourceDefinition implements meta.ResourceDefinitionProvider interface.
func (EtcFileSpecMD) ResourceDefinition(resource.Metadata, EtcFileSpecSpec) meta.ResourceDefinitionSpec {
	return meta.ResourceDefinitionSpec{
		Type:             EtcFileSpecType,
		Aliases:          []resource.Type{},
		DefaultNamespace: NamespaceName,
		PrintColumns:     []meta.PrintColumn{},
	}
}
