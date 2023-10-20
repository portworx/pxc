/*
Copyright © 2020 Portworx

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
package main

import (
	_ "github.com/portworx/pxc/component/examples/golang/handler"
	pxc "github.com/portworx/pxc/pkg/component"
)

var (
	// ComponentName is set by the Makefile
	ComponentName = "dev"

	// ComponentVersion is set by the Makefile
	ComponentVersion = "(dev)"
)

func main() {
	c := pxc.NewComponent(&pxc.ComponentConfig{
		Name:    ComponentName,
		Short:   "This is a short message from cm",
		Version: ComponentVersion,
	})
	c.Execute()
}
