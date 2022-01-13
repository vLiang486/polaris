/**
 * Tencent is pleased to support the open source community by making Polaris available.
 *
 * Copyright (C) 2019 THL A29 Limited, a Tencent company. All rights reserved.
 *
 * Licensed under the BSD 3-Clause License (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * https://opensource.org/licenses/BSD-3-Clause
 *
 * Unless required by applicable law or agreed to in writing, software distributed
 * under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR
 * CONDITIONS OF ANY KIND, either express or implied. See the License for the
 * specific language governing permissions and limitations under the License.
 */

package v1

import wrappers "github.com/golang/protobuf/ptypes/wrappers"

// NewUserResponse 创建回复带用户信息
func NewUserResponse(code uint32, user *User) *Response {
	return &Response{
		Code: &wrappers.UInt32Value{Value: code},
		Info: &wrappers.StringValue{Value: code2info[code]},
		User: user,
	}
}

// NewUserResponse 创建回复带用户信息
func NewUserResponseWithMsg(code uint32, info string, user *User) *Response {
	return &Response{
		Code: &wrappers.UInt32Value{Value: code},
		Info: &wrappers.StringValue{Value: info},
		User: user,
	}
}

// NewUserGroupResponse 创建回复带用户组信息
func NewUserGroupResponse(code uint32, user *UserGroup) *Response {
	return &Response{
		Code:      &wrappers.UInt32Value{Value: code},
		Info:      &wrappers.StringValue{Value: code2info[code]},
		UserGroup: user,
	}
}

func NewModifyUserGroupResponse(code uint32, group *ModifyUserGroup) *Response {
	return &Response{
		Code:            &wrappers.UInt32Value{Value: code},
		Info:            &wrappers.StringValue{Value: code2info[code]},
		ModifyUserGroup: group,
	}
}

// NewUserGroupRelationResponse
func NewUserGroupRelationResponse(code uint32, relation *UserGroupRelation) *Response {
	return &Response{
		Code:     &wrappers.UInt32Value{Value: code},
		Info:     &wrappers.StringValue{Value: code2info[code]},
		Relation: relation,
	}
}

// NewAuthStrategyResponse
func NewAuthStrategyResponse(code uint32, req *AuthStrategy) *Response {
	return &Response{
		Code:         &wrappers.UInt32Value{Value: code},
		Info:         &wrappers.StringValue{Value: code2info[code]},
		AuthStrategy: req,
	}
}

// NewAuthStrategyResponseWithMsg 
//  @param code 
//  @param msg 
//  @param req 
//  @return *Response 
func NewAuthStrategyResponseWithMsg(code uint32, msg string, req *AuthStrategy) *Response {
	return &Response{
		Code:         &wrappers.UInt32Value{Value: code},
		Info:         &wrappers.StringValue{Value: msg},
		AuthStrategy: req,
	}
}

// NewModifyAuthStrategyResponse
//  @param code
//  @param req
//  @return *Response
func NewModifyAuthStrategyResponse(code uint32, req *ModifyAuthStrategy) *Response {
	return &Response{
		Code:               &wrappers.UInt32Value{Value: code},
		Info:               &wrappers.StringValue{Value: code2info[code]},
		ModifyAuthStrategy: req,
	}
}

// NewLoginResponse
//  @param code
//  @param loginResponse
//  @return *Response
func NewLoginResponse(code uint32, loginResponse *LoginResponse) *Response {
	return &Response{
		Code:          &wrappers.UInt32Value{Value: code},
		Info:          &wrappers.StringValue{Value: code2info[code]},
		LoginResponse: loginResponse,
	}
}
