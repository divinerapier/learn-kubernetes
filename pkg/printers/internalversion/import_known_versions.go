/*
Copyright 2016 The Kubernetes Authors.

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

package internalversion

import (
	// These imports are the API groups the client will support.
	// TODO: Remove these manual install once we don't need legacy scheme in get comman
	_ "github.com/divinerapier/learn-kubernetes/pkg/apis/apps/install"
	_ "github.com/divinerapier/learn-kubernetes/pkg/apis/authentication/install"
	_ "github.com/divinerapier/learn-kubernetes/pkg/apis/authorization/install"
	_ "github.com/divinerapier/learn-kubernetes/pkg/apis/autoscaling/install"
	_ "github.com/divinerapier/learn-kubernetes/pkg/apis/batch/install"
	_ "github.com/divinerapier/learn-kubernetes/pkg/apis/certificates/install"
	_ "github.com/divinerapier/learn-kubernetes/pkg/apis/coordination/install"
	_ "github.com/divinerapier/learn-kubernetes/pkg/apis/core/install"
	_ "github.com/divinerapier/learn-kubernetes/pkg/apis/events/install"
	_ "github.com/divinerapier/learn-kubernetes/pkg/apis/extensions/install"
	_ "github.com/divinerapier/learn-kubernetes/pkg/apis/policy/install"
	_ "github.com/divinerapier/learn-kubernetes/pkg/apis/rbac/install"
	_ "github.com/divinerapier/learn-kubernetes/pkg/apis/scheduling/install"
	_ "github.com/divinerapier/learn-kubernetes/pkg/apis/settings/install"
	_ "github.com/divinerapier/learn-kubernetes/pkg/apis/storage/install"
)
