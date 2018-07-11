/*
Copyright 2018 codestation

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

package services

import (
	"fmt"
	"path"
	"strings"
	"time"

	"github.com/mholt/archiver"
)

// Tarball has the config options for the Tarball service
type Tarball struct {
	Name     string
	Path     string
	Compress bool
}

// Backup creates a tarball of the specified directory
func (f *Tarball) Backup() (string, error) {
	var name string

	if f.Name != "" {
		name = f.Name + "-backup"
	} else {
		name = path.Base(f.Path) + "-backup"
	}

	filepath := fmt.Sprintf("%s/%s-%s.tar", SaveDir, name, time.Now().Format("20060102150405"))
	filepath = path.Clean(filepath)

	var err error

	if f.Compress {
		filepath += ".gz"
		err = archiver.TarGz.Make(filepath, []string{f.Path})
	} else {
		err = archiver.Tar.Make(filepath, []string{f.Path})
	}

	if err != nil {
		return "", fmt.Errorf("cannot create tarball on %s, %v", filepath, err)
	}

	return filepath, nil
}

// Restore extracts a tarball to the specified directory
func (f *Tarball) Restore(filepath string) error {
	var err error

	if strings.HasSuffix(filepath, ".gz") {
		err = archiver.TarGz.Open(filepath, f.Path)
	} else {
		err = archiver.Tar.Open(filepath, f.Path)
	}

	if err != nil {
		return fmt.Errorf("cannot unpack tarball on %s, %v", filepath, err)
	}

	return nil
}