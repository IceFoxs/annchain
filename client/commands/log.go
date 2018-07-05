// Copyright 2017 Annchain Information Technology Services Co.,Ltd.
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


package commands

import (
	"os"
	"path"

	"go.uber.org/zap"
)

var logger *zap.Logger

func init() {
	var err error
	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	zapConf := zap.NewDevelopmentConfig()
	zapConf.OutputPaths = []string{path.Join(pwd, "client.out.log")}
	zapConf.ErrorOutputPaths = []string{}
	logger, err = zapConf.Build()
	if err != nil {
		panic(err.Error())
	}
}
