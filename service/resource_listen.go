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

package service

import (
	"context"

	"github.com/polarismesh/polaris-server/common/model"
	"github.com/polarismesh/polaris-server/common/utils"
)

// ResourceHook The listener is placed before and after the resource operation, only normal flow
type ResourceHook interface {

	// Before
	//  @param ctx
	//  @param resourceType
	Before(ctx context.Context, resourceType model.Resource)

	// After
	//  @param ctx
	//  @param resourceType
	//  @param res
	After(ctx context.Context, resourceType model.Resource, res *ResourceEvent)
}

type ResourceEvent struct {
	Origin   interface{}
	IsRemove bool
}

// Before
//  @param ctx
//  @param resourceType
func (svr *serverAuthAbility) Before(ctx context.Context, resourceType model.Resource) {
	// do nothing
}

// After
//  @param ctx
//  @param resourceType
//  @param res
func (svr *serverAuthAbility) After(ctx context.Context, resourceType model.Resource, res *ResourceEvent) {
	switch resourceType {
	case model.RNamespace:
		svr.onNamespaceResource(ctx, res)
	case model.RService:
		svr.onServiceResource(ctx, res)
	}
}

// onNamespaceResource 
//  @receiver svr 
//  @param ctx 
//  @param res 
func (svr *serverAuthAbility) onNamespaceResource(ctx context.Context, res *ResourceEvent) {
	authCtx := ctx.Value(utils.ContextAuthContextKey).(*model.AcquireContext)

	svr.authMgn.AfterResourceOperation(authCtx)
}

// onServiceResource 
//  @receiver svr 
//  @param ctx 
//  @param res 
func (svr *serverAuthAbility) onServiceResource(ctx context.Context, res *ResourceEvent) {
	authCtx := ctx.Value(utils.ContextAuthContextKey).(*model.AcquireContext)

	svr.authMgn.AfterResourceOperation(authCtx)
}

// onConfigGroupResource 
//  @receiver svr 
//  @param ctx 
//  @param res 
func (svr *serverAuthAbility) onConfigGroupResource(ctx context.Context, res *ResourceEvent) {
	authCtx := ctx.Value(utils.ContextAuthContextKey).(*model.AcquireContext)

	svr.authMgn.AfterResourceOperation(authCtx)
}
