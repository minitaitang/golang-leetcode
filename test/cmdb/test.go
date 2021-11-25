package main

import (
	"fmt"
	"time"

	"gitlab.sh.sensetime.com/keystone/oss/cmdb/models/dao"
	"gitlab.sh.sensetime.com/keystone/oss/cmdb/utils"
)

func main() {
	tu1 := &dao.TimeRangeUsage{StartTime: time.Unix(1, 0), EndTime: time.Unix(20, 0), Count: 1}
	tu2 := &dao.TimeRangeUsage{StartTime: time.Unix(7, 0), EndTime: time.Unix(20, 0), Count: 2}
	tus := []*dao.TimeRangeUsage{tu1, tu2}
	new := utils.ReGenTimeRangeUsages([]*dao.TimeRangeUsage{tu1, tu2})
	for _, n := range new {
		fmt.Println(n.StartTime, n.EndTime, n.Count)
	}

	resolutionOriginUsageM := make(map[string][]*dao.TimeRangeUsage)
	resolutionOriginUsageM["4K"] = tus
	for r, t := range resolutionOriginUsageM {
		resolutionOriginUsageM[r] = utils.ReGenTimeRangeUsages(t)
		fmt.Println("just use t...", t)
	}

	for _, t := range resolutionOriginUsageM {
		for _, u := range t {
			fmt.Println(u.StartTime, u.EndTime, u.Count)
		}
	}
}
