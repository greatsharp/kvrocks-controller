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
package logger

import (
	"fmt"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var zapLogger *zap.Logger

func Get() *zap.Logger {
	return zapLogger
}

func init() {
	zapConfig := zap.NewProductionConfig()
	zapConfig.EncoderConfig.TimeKey = "timestamp"
	zapConfig.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	zapLogger, _ = zapConfig.Build()
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.TimeKey = "time"

	return zapcore.NewJSONEncoder(encoderConfig)
}

func getWriteSyncer(filename string, maxBackups, maxAge, maxSize int, compress bool) (zapcore.WriteSyncer, error) {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   filename,
		MaxBackups: maxBackups,
		MaxSize:    maxSize,
		MaxAge:     maxAge,
		Compress:   compress,
	}

	if _, err := lumberJackLogger.Write([]byte("test logfile writable\r\n")); err != nil {
		return nil, fmt.Errorf("test writing to log file %s failed", filename)
	}

	return zapcore.AddSync(lumberJackLogger), nil
}

func InitLoggerRotate(level, filename string, maxBackups, maxAge, maxSize int, compress bool) error {
	// if file path is empty, use default zapLogger, print in console
	if len(filename) == 0 {
		return nil
	}
	if level != "info" && level != "warn" && level != "error" {
		return fmt.Errorf("log level must be one of info,warn,error")
	}
	if maxBackups > 100 || maxBackups < 10 {
		return fmt.Errorf("log max_backups must be between 10 and 100")
	}
	if maxAge > 30 || maxAge < 1 {
		return fmt.Errorf("log max_age must be between 1 and 30")
	}
	if maxSize > 500 || maxSize < 100 {
		return fmt.Errorf("log max_size must be between 100 and 500")
	}

	var l = new(zapcore.Level)
	if err := l.UnmarshalText([]byte(level)); err != nil {
		return err
	}
	encoder := getEncoder()
	writeSync, err := getWriteSyncer(filename, maxBackups, maxAge, maxSize, compress)
	if err != nil {
		return err
	}
	core := zapcore.NewCore(encoder, writeSync, l)
	rotateLogger := zap.New(core, zap.AddCaller())
	zapLogger = rotateLogger

	return nil
}

func Sync() {
	if zapLogger != nil {
		if err := zapLogger.Sync(); err != nil {
			return
		}
	}
}
