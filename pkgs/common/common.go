// Copyright 2022 lukemakhanu
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

// Copyright 2022 lukemakhanu
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

package common

import (
	"fmt"
	"log"
	"os"

	"github.com/fsnotify/fsnotify"
	"github.com/garyburd/redigo/redis"
	"github.com/spf13/viper"
	"gopkg.in/natefinch/lumberjack.v2"
)

func Logutils(logdir string, maxSize int, maxBackups int, maxAge int, compress bool) {
	log.SetOutput(&lumberjack.Logger{
		Filename:   logdir,
		MaxSize:    maxSize,
		MaxBackups: maxBackups,
		MaxAge:     maxAge,
		Compress:   compress,
	})
	log.SetFlags(log.LstdFlags | log.Lmicroseconds | log.Lshortfile)
}

// Configutils :: initiates app configuration
func Configutils(addConfigPathLive string, addConfigPathLocal string) {
	fmt.Printf("*** In Configutils func ***")
	log.SetFlags(log.LstdFlags | log.Lmicroseconds | log.Lshortfile)
	viper.SetDefault("host", "localhost")
	viper.SetConfigName("config")
	viper.AddConfigPath(addConfigPathLive)
	viper.AddConfigPath(addConfigPathLocal)
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		log.Printf("Error : %v", err)
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Printf("Config file changed: %s", e.Name)
	})
}

// InitializeRedis : initializes redis connection
func InitializeRedis(redisServer string) *redis.Pool {
	pool := &redis.Pool{
		MaxIdle:   100,
		MaxActive: 500,
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial("tcp", redisServer)
			if err != nil {
				log.Printf("ERROR: fail init redis:: %s", err.Error())
				os.Exit(1)
			}
			return conn, err
		},
	}
	return pool
}
