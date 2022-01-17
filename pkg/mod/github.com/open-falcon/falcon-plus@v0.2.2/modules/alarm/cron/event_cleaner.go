// Copyright 2017 Xiaomi, Inc.
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

package cron

import (
	"github.com/open-falcon/falcon-plus/modules/alarm/g"
	eventmodel "github.com/open-falcon/falcon-plus/modules/alarm/model/event"
	"time"
)

func CleanExpiredEvent() {
	if !g.Config().Housekeeper.Enabled {
		return
	}

	for {

		retention_days := g.Config().Housekeeper.EventRetentionDays
		delete_batch := g.Config().Housekeeper.EventDeleteBatch

		now := time.Now()
		before := now.Add(time.Duration(-retention_days*24) * time.Hour)
		eventmodel.DeleteEventOlder(before, delete_batch)

		time.Sleep(time.Second * 60)
	}
}