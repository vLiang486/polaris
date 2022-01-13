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

package auth

import (
	"context"
	"errors"
	"fmt"
	"os"
	"sync"

	"github.com/polarismesh/polaris-server/cache"
)

// Config 鉴权能力的相关配置参数
type Config struct {
	Name   string
	Option map[string]interface{}
}

var (
	// Slots store slots
	Slots      = make(map[string]AuthManager)
	once       = &sync.Once{}
	authMgn    AuthManager
	finishInit bool = false
)

/**
 * RegisterAuthManager 注册一个新的Store
 */
func RegisterAuthManager(s AuthManager) error {
	name := s.Name()
	if _, ok := Slots[name]; ok {
		return errors.New("auth manager name is exist")
	}

	Slots[name] = s
	return nil
}

/**
 * GetStore 获取Store
 */
func GetAuthManager() (AuthManager, error) {
	if !finishInit {
		return nil, errors.New("AuthManager has not done Initialize")
	}
	return authMgn, nil
}

// Initialize 初始化
func Initialize(ctx context.Context, authOpt *Config, cacheMgn *cache.NamingCache) error {
	var err error
	once.Do(func() {
		err = initialize(ctx, authOpt, cacheMgn)
	})

	if err != nil {
		return err
	}

	finishInit = true
	return nil
}

/**
 * @brief 包裹了初始化函数，在GetStore的时候会在自动调用，全局初始化一次
 */
func initialize(ctx context.Context, authOpt *Config, cacheMgn *cache.NamingCache) error {
	name := authOpt.Name
	if name == "" {
		return errors.New("auth manager Name is empty")
	}

	mgn, ok := Slots[name]
	if !ok {
		return errors.New("no such name AuthManager")
	}

	authMgn = mgn

	if err := authMgn.Initialize(authOpt, cacheMgn); err != nil {
		fmt.Printf("auth manager do initialize err: %s", err.Error())
		os.Exit(-1)
	}
	return nil
}
