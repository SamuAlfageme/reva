// Copyright 2018-2021 CERN
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// In applying this license, CERN does not waive the privileges and immunities
// granted to it by virtue of its status as an Intergovernmental Organization
// or submit itself to any jurisdiction.

package scope

import (
	authpb "github.com/cs3org/go-cs3apis/cs3/auth/provider/v1beta1"
	provider "github.com/cs3org/go-cs3apis/cs3/storage/provider/v1beta1"
	types "github.com/cs3org/go-cs3apis/cs3/types/v1beta1"
	"github.com/cs3org/reva/pkg/utils"
)

func userScope(scope *authpb.Scope, resource interface{}) (bool, error) {
	// Always return true. Registered users can access all paths.
	// TODO(ishank011): Add checks for read/write permissions.
	return true, nil
}

// GetOwnerScope returns the default owner scope with access to all resources.
func GetOwnerScope() (map[string]*authpb.Scope, error) {
	ref := &provider.Reference{Path: "/"}
	val, err := utils.MarshalProtoV1ToJSON(ref)
	if err != nil {
		return nil, err
	}
	return map[string]*authpb.Scope{
		"user": &authpb.Scope{
			Resource: &types.OpaqueEntry{
				Decoder: "json",
				Value:   val,
			},
			Role: authpb.Role_ROLE_OWNER,
		},
	}, nil
}
