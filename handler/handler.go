/*
Copyright © 2019 Portworx

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
package handler

import (
	_ "github.com/portworx/pxc/handler/alerts"
	_ "github.com/portworx/pxc/handler/cloudmigration"
	_ "github.com/portworx/pxc/handler/cluster"
	_ "github.com/portworx/pxc/handler/clusterpair"
	_ "github.com/portworx/pxc/handler/context"
	_ "github.com/portworx/pxc/handler/logs"
	_ "github.com/portworx/pxc/handler/node"
	_ "github.com/portworx/pxc/handler/pvc"
	_ "github.com/portworx/pxc/handler/volume"
	_ "github.com/portworx/pxc/handler/volumeclone"
	_ "github.com/portworx/pxc/handler/volumesnapshot"
	_ "github.com/portworx/pxc/handler/volumestats"
)
