/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 *
 */

package raft

import (
	"go.etcd.io/etcd/raft/v3"
	"go.uber.org/zap"
)

var _ raft.Logger = &Logger{}

// Logger is a wrapper around zap.SugaredLogger to implement the raft.Logger interface.
type Logger struct {
	*zap.SugaredLogger
}

func (r Logger) Warning(v ...interface{}) {
	r.SugaredLogger.Warn(v...)
}

func (r Logger) Warningf(format string, v ...interface{}) {
	r.SugaredLogger.Warnf(format, v...)
}

func (r Logger) Debug(v ...interface{}) {
	r.SugaredLogger.Debug(v...)
}

func (r Logger) Debugf(format string, v ...interface{}) {
	r.SugaredLogger.Debugf(format, v...)
}

func (r Logger) Error(v ...interface{}) {
	r.SugaredLogger.Error(v...)
}

func (r Logger) Errorf(format string, v ...interface{}) {
	r.SugaredLogger.Errorf(format, v...)
}

func (r Logger) Info(v ...interface{}) {
	r.SugaredLogger.Info(v...)
}

func (r Logger) Infof(format string, v ...interface{}) {
	r.SugaredLogger.Infof(format, v...)
}

func (r Logger) Fatal(v ...interface{}) {
	r.SugaredLogger.Fatal(v...)
}

func (r Logger) Fatalf(format string, v ...interface{}) {
	r.SugaredLogger.Fatalf(format, v...)
}

func (r Logger) Panic(v ...interface{}) {
	r.SugaredLogger.Panic(v...)
}

func (r Logger) Panicf(format string, v ...interface{}) {
	r.SugaredLogger.Panicf(format, v...)
}
