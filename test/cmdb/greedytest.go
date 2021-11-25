package main

import (
	"time"

	"gitlab.sh.sensetime.com/keystone/oss/cmdb/models/dao"
	"gitlab.sh.sensetime.com/keystone/oss/cmdb/utils"
	pb "gitlab.sh.sensetime.com/keystone/oss/common/api/cmdb"
)

func main() {
	// gpuCountM := map[string]uint64{"P4": 0}
	p4gpuRule := &pb.ResourceCalculationRule_GpuRule{GpuType: "P4", GpuNum: 1, ProductSize: 11}
	t4gpuRule := &pb.ResourceCalculationRule_GpuRule{GpuType: "T4", GpuNum: 1, ProductSize: 4}
	gpuRule := []*pb.ResourceCalculationRule_GpuRule{p4gpuRule, t4gpuRule}
	calRule := &pb.ResourceCalculationRule_Rule{GpuRules: gpuRule}
	calRules := []*pb.ResourceCalculationRule_Rule{calRule}
	resourceCalRule := &pb.ResourceCalculationRule{Rules: calRules, GpuTypeOrders: []string{"P4", "T4"}, ResolutionRule: map[string]uint32{"1080P": 1, "2K": 2, "4K": 4}}

	// resolutionCount := map[string]uint64{"1080P": 1, "2K": 2, "4K": 3}
	// fmt.Println(utils.GreedyUseGpuCount(gpuCountM, resourceCalRule, resolutionCount))

	loc, _ := time.LoadLocation("Asia/Shanghai")
	tt1, _ := time.ParseInLocation("2006-01-02 15:04:05", "2021-06-01 16:00:00", loc)
	tt2, _ := time.ParseInLocation("2006-01-02 15:04:05", "2021-06-01 16:04:00", loc)
	tt3, _ := time.ParseInLocation("2006-01-02 15:04:05", "2021-06-01 16:07:00", loc)
	tt4, _ := time.ParseInLocation("2006-01-02 15:04:05", "2021-06-01 16:11:00", loc)
	tt5, _ := time.ParseInLocation("2006-01-02 15:04:05", "2021-06-01 16:13:00", loc)
	tt6, _ := time.ParseInLocation("2006-01-02 15:04:05", "2021-06-01 16:15:00", loc)
	tt7, _ := time.ParseInLocation("2006-01-02 15:04:05", "2021-06-01 16:17:00", loc)
	tt8, _ := time.ParseInLocation("2006-01-02 15:04:05", "2021-06-01 16:20:00", loc)
	tt9, _ := time.ParseInLocation("2006-01-02 15:04:05", "2021-06-01 16:30:00", loc)
	timePointSlice := []time.Time{tt1, tt2, tt3, tt4, tt5, tt6, tt7, tt8, tt9}
	usage1k := []*dao.TimeRangeUsage{&dao.TimeRangeUsage{StartTime: tt2, EndTime: tt3, Count: 4}, &dao.TimeRangeUsage{StartTime: tt5, EndTime: tt7, Count: 6}}
	usage2k := []*dao.TimeRangeUsage{&dao.TimeRangeUsage{StartTime: tt4, EndTime: tt6, Count: 3}}
	usage4K := []*dao.TimeRangeUsage{&dao.TimeRangeUsage{StartTime: tt5, EndTime: tt8, Count: 2}}
	resolutionOriginUsageM := map[string][]*dao.TimeRangeUsage{
		"1080P": usage1k,
		"2K":    usage2k,
		"4K":    usage4K,
	}

	gpuRealUsageM := map[string]uint64{
		"P4": 1,
		"T4": 1,
	}

	utils.TimePointSliceGreedySup(timePointSlice, gpuRealUsageM, resourceCalRule, resolutionOriginUsageM, "4K", 2)
}
