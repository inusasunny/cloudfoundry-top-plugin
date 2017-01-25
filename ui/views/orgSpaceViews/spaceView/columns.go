// Copyright (c) 2016 ECS Team, Inc. - All Rights Reserved
// https://github.com/ECSTeam/cloudfoundry-top-plugin
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package spaceView

import (
	"fmt"
	"strconv"

	"math"

	"github.com/ecsteam/cloudfoundry-top-plugin/ui/uiCommon"
	"github.com/ecsteam/cloudfoundry-top-plugin/util"
)

const ATTENTION_HOT_PERCENT = 90
const ATTENTION_WARM_PERCENT = 80

func columnSpaceName() *uiCommon.ListColumn {
	defaultColSize := 25
	sortFunc := func(c1, c2 util.Sortable) bool {
		return util.CaseInsensitiveLess(c1.(*DisplaySpace).Name, c2.(*DisplaySpace).Name)
	}
	displayFunc := func(data uiCommon.IData, columnOwner uiCommon.IColumnOwner) string {
		stats := data.(*DisplaySpace)
		return util.FormatDisplayData(stats.Name, defaultColSize)
	}
	rawValueFunc := func(data uiCommon.IData) string {
		stats := data.(*DisplaySpace)
		return stats.Name
	}
	c := uiCommon.NewListColumn("spaceName", "SPACE", defaultColSize,
		uiCommon.ALPHANUMERIC, true, sortFunc, false, displayFunc, rawValueFunc, nil)
	return c
}

func columnQuotaName() *uiCommon.ListColumn {
	defaultColSize := 11
	sortFunc := func(c1, c2 util.Sortable) bool {
		return util.CaseInsensitiveLess(c1.(*DisplaySpace).QuotaName, c2.(*DisplaySpace).QuotaName)
	}
	displayFunc := func(data uiCommon.IData, columnOwner uiCommon.IColumnOwner) string {
		stats := data.(*DisplaySpace)
		return util.FormatDisplayData(stats.QuotaName, defaultColSize)
	}
	rawValueFunc := func(data uiCommon.IData) string {
		stats := data.(*DisplaySpace)
		return stats.QuotaName
	}
	c := uiCommon.NewListColumn("QUOTA_NAME", "QUOTA_NAME", defaultColSize,
		uiCommon.ALPHANUMERIC, true, sortFunc, false, displayFunc, rawValueFunc, nil)
	return c
}

func columnNumberOfApps() *uiCommon.ListColumn {
	sortFunc := func(c1, c2 util.Sortable) bool {
		return c1.(*DisplaySpace).NumberOfApps < c2.(*DisplaySpace).NumberOfApps
	}
	displayFunc := func(data uiCommon.IData, columnOwner uiCommon.IColumnOwner) string {
		stats := data.(*DisplaySpace)
		return fmt.Sprintf("%7v", stats.NumberOfApps)
	}
	rawValueFunc := func(data uiCommon.IData) string {
		stats := data.(*DisplaySpace)
		return strconv.Itoa(stats.NumberOfApps)
	}
	c := uiCommon.NewListColumn("APPS", "APPS", 7,
		uiCommon.NUMERIC, false, sortFunc, true, displayFunc, rawValueFunc, nil)
	return c
}

func columnReportingContainers() *uiCommon.ListColumn {
	sortFunc := func(c1, c2 util.Sortable) bool {
		return c1.(*DisplaySpace).TotalReportingContainers < c2.(*DisplaySpace).TotalReportingContainers
	}
	displayFunc := func(data uiCommon.IData, columnOwner uiCommon.IColumnOwner) string {
		appStats := data.(*DisplaySpace)
		return fmt.Sprintf("%7v", appStats.TotalReportingContainers)
	}
	rawValueFunc := func(data uiCommon.IData) string {
		appStats := data.(*DisplaySpace)
		return strconv.Itoa(appStats.TotalReportingContainers)
	}
	c := uiCommon.NewListColumn("reportingContainers", "RCR", 7,
		uiCommon.NUMERIC, false, sortFunc, true, displayFunc, rawValueFunc, nil)
	return c
}

func columnTotalCpu() *uiCommon.ListColumn {
	sortFunc := func(c1, c2 util.Sortable) bool {
		return c1.(*DisplaySpace).TotalCpuPercentage < c2.(*DisplaySpace).TotalCpuPercentage
	}
	displayFunc := func(data uiCommon.IData, columnOwner uiCommon.IColumnOwner) string {
		appStats := data.(*DisplaySpace)
		totalCpuInfo := ""
		if appStats.TotalReportingContainers == 0 {
			totalCpuInfo = fmt.Sprintf("%6v", "--")
		} else {
			if appStats.TotalCpuPercentage >= 100.0 {
				totalCpuInfo = fmt.Sprintf("%6.0f", appStats.TotalCpuPercentage)
			} else if appStats.TotalCpuPercentage >= 10.0 {
				totalCpuInfo = fmt.Sprintf("%6.1f", appStats.TotalCpuPercentage)
			} else {
				totalCpuInfo = fmt.Sprintf("%6.2f", appStats.TotalCpuPercentage)
			}
		}
		return fmt.Sprintf("%6v", totalCpuInfo)
	}
	rawValueFunc := func(data uiCommon.IData) string {
		appStats := data.(*DisplaySpace)
		return fmt.Sprintf("%.2f", appStats.TotalCpuPercentage)
	}
	c := uiCommon.NewListColumn("CPU", "CPU%", 6,
		uiCommon.NUMERIC, false, sortFunc, true, displayFunc, rawValueFunc, nil)
	return c
}

func columnMemoryLimit() *uiCommon.ListColumn {
	sortFunc := func(c1, c2 util.Sortable) bool {
		return c1.(*DisplaySpace).MemoryLimitInBytes < c2.(*DisplaySpace).MemoryLimitInBytes
	}
	displayFunc := func(data uiCommon.IData, columnOwner uiCommon.IColumnOwner) string {
		appStats := data.(*DisplaySpace)
		totalMemInfo := ""
		if appStats.MemoryLimitInBytes == 0 {
			totalMemInfo = fmt.Sprintf("%9v", "--")
		} else {
			totalMemInfo = fmt.Sprintf("%9v", util.ByteSize(appStats.MemoryLimitInBytes).StringWithPrecision(1))
		}
		return fmt.Sprintf("%9v", totalMemInfo)
	}
	rawValueFunc := func(data uiCommon.IData) string {
		appStats := data.(*DisplaySpace)
		return fmt.Sprintf("%v", appStats.MemoryLimitInBytes)
	}
	c := uiCommon.NewListColumn("MAX_MEM", "MAX_MEM", 9,
		uiCommon.NUMERIC, false, sortFunc, true, displayFunc, rawValueFunc, nil)
	return c
}

func columnTotalReservedMemory() *uiCommon.ListColumn {
	sortFunc := func(c1, c2 util.Sortable) bool {
		return c1.(*DisplaySpace).TotalReservedMemory < c2.(*DisplaySpace).TotalReservedMemory
	}
	displayFunc := func(data uiCommon.IData, columnOwner uiCommon.IColumnOwner) string {
		appStats := data.(*DisplaySpace)
		totalMemInfo := ""
		if appStats.TotalReportingContainers == 0 {
			totalMemInfo = fmt.Sprintf("%9v", "--")
		} else {
			totalMemInfo = fmt.Sprintf("%9v", util.ByteSize(appStats.TotalReservedMemory).StringWithPrecision(1))
		}
		return fmt.Sprintf("%9v", totalMemInfo)
	}
	rawValueFunc := func(data uiCommon.IData) string {
		appStats := data.(*DisplaySpace)
		return fmt.Sprintf("%v", appStats.TotalReservedMemory)
	}
	c := uiCommon.NewListColumn("RSVD_MEM", "RSVD_MEM", 9,
		uiCommon.NUMERIC, false, sortFunc, true, displayFunc, rawValueFunc, closeToMemoryEitherQuotaAttentionFunc)
	return c
}

func columnTotalReservedMemoryPercentOfSpaceQuota() *uiCommon.ListColumn {
	sortFunc := func(c1, c2 util.Sortable) bool {
		return c1.(*DisplaySpace).TotalReservedMemoryPercentOfSpaceQuota < c2.(*DisplaySpace).TotalReservedMemoryPercentOfSpaceQuota
	}
	displayFunc := func(data uiCommon.IData, columnOwner uiCommon.IColumnOwner) string {
		appStats := data.(*DisplaySpace)
		totalMemInfo := ""
		if appStats.TotalReportingContainers == 0 || appStats.MemoryLimitInBytes == 0 {
			totalMemInfo = fmt.Sprintf("%7v", "--")
		} else {
			totalMemInfo = fmt.Sprintf("%7.1f", appStats.TotalReservedMemoryPercentOfSpaceQuota)
		}
		return fmt.Sprintf("%7v", totalMemInfo)
	}
	rawValueFunc := func(data uiCommon.IData) string {
		appStats := data.(*DisplaySpace)
		return fmt.Sprintf("%v", appStats.TotalReservedMemoryPercentOfSpaceQuota)
	}
	c := uiCommon.NewListColumn("S_MEM_PER", "S_MEM%", 7,
		uiCommon.NUMERIC, false, sortFunc, true, displayFunc, rawValueFunc, closeToMemorySpaceQuotaAttentionFunc)
	return c
}

func columnTotalReservedMemoryPercentOfOrgQuota() *uiCommon.ListColumn {
	sortFunc := func(c1, c2 util.Sortable) bool {
		return c1.(*DisplaySpace).TotalReservedMemoryPercentOfOrgQuota < c2.(*DisplaySpace).TotalReservedMemoryPercentOfOrgQuota
	}
	displayFunc := func(data uiCommon.IData, columnOwner uiCommon.IColumnOwner) string {
		appStats := data.(*DisplaySpace)
		totalMemInfo := ""
		if appStats.TotalReportingContainers == 0 {
			totalMemInfo = fmt.Sprintf("%7v", "--")
		} else {
			totalMemInfo = fmt.Sprintf("%7.1f", appStats.TotalReservedMemoryPercentOfOrgQuota)
		}
		return fmt.Sprintf("%7v", totalMemInfo)
	}
	rawValueFunc := func(data uiCommon.IData) string {
		appStats := data.(*DisplaySpace)
		return fmt.Sprintf("%v", appStats.TotalReservedMemoryPercentOfOrgQuota)
	}
	c := uiCommon.NewListColumn("O_MEM_PER", "O_MEM%", 7,
		uiCommon.NUMERIC, false, sortFunc, true, displayFunc, rawValueFunc, closeToMemoryOrgQuotaAttentionFunc)
	return c
}

func closeToMemorySpaceQuotaAttentionFunc(data uiCommon.IData, columnOwner uiCommon.IColumnOwner) uiCommon.AttentionType {
	stats := data.(*DisplaySpace)
	attentionType := uiCommon.ATTENTION_NORMAL
	if stats.TotalReportingContainers > 0 && stats.MemoryLimitInBytes > 0 {
		percentOfQuota := stats.TotalReservedMemoryPercentOfSpaceQuota
		switch {
		case percentOfQuota >= ATTENTION_HOT_PERCENT:
			attentionType = uiCommon.ATTENTION_HOT
		case percentOfQuota >= ATTENTION_WARM_PERCENT:
			attentionType = uiCommon.ATTENTION_WARM
		}
	}
	return attentionType
}

func closeToMemoryOrgQuotaAttentionFunc(data uiCommon.IData, columnOwner uiCommon.IColumnOwner) uiCommon.AttentionType {
	stats := data.(*DisplaySpace)
	attentionType := uiCommon.ATTENTION_NORMAL
	if stats.TotalReportingContainers > 0 {
		percentOfQuota := stats.TotalReservedMemoryPercentOfOrgQuota
		switch {
		case percentOfQuota >= ATTENTION_HOT_PERCENT:
			attentionType = uiCommon.ATTENTION_HOT
		case percentOfQuota >= ATTENTION_WARM_PERCENT:
			attentionType = uiCommon.ATTENTION_WARM
		}
	}
	return attentionType
}

func closeToMemoryEitherQuotaAttentionFunc(data uiCommon.IData, columnOwner uiCommon.IColumnOwner) uiCommon.AttentionType {
	stats := data.(*DisplaySpace)
	attentionType := uiCommon.ATTENTION_NORMAL
	if stats.TotalReportingContainers > 0 && stats.MemoryLimitInBytes > 0 {
		percentOfSpaceQuota := stats.TotalReservedMemoryPercentOfSpaceQuota
		percentOfOrgQuota := stats.TotalReservedMemoryPercentOfOrgQuota
		percentOfQuota := math.Max(percentOfSpaceQuota, percentOfOrgQuota)
		switch {
		case percentOfQuota >= ATTENTION_HOT_PERCENT:
			attentionType = uiCommon.ATTENTION_HOT
		case percentOfQuota >= ATTENTION_WARM_PERCENT:
			attentionType = uiCommon.ATTENTION_WARM
		}
	}
	return attentionType
}

func columnTotalMemoryUsed() *uiCommon.ListColumn {
	sortFunc := func(c1, c2 util.Sortable) bool {
		return c1.(*DisplaySpace).TotalUsedMemory < c2.(*DisplaySpace).TotalUsedMemory
	}
	displayFunc := func(data uiCommon.IData, columnOwner uiCommon.IColumnOwner) string {
		appStats := data.(*DisplaySpace)
		totalMemInfo := ""
		if appStats.TotalReportingContainers == 0 {
			totalMemInfo = fmt.Sprintf("%9v", "--")
		} else {
			totalMemInfo = fmt.Sprintf("%9v", util.ByteSize(appStats.TotalUsedMemory).StringWithPrecision(1))
		}
		return fmt.Sprintf("%9v", totalMemInfo)
	}
	rawValueFunc := func(data uiCommon.IData) string {
		appStats := data.(*DisplaySpace)
		return fmt.Sprintf("%v", appStats.TotalUsedMemory)
	}
	c := uiCommon.NewListColumn("USED_MEM", "USED_MEM", 9,
		uiCommon.NUMERIC, false, sortFunc, true, displayFunc, rawValueFunc, nil)
	return c
}

func columnTotalReservedDisk() *uiCommon.ListColumn {
	sortFunc := func(c1, c2 util.Sortable) bool {
		return c1.(*DisplaySpace).TotalReservedDisk < c2.(*DisplaySpace).TotalReservedDisk
	}
	displayFunc := func(data uiCommon.IData, columnOwner uiCommon.IColumnOwner) string {
		appStats := data.(*DisplaySpace)
		totalDiskInfo := ""
		if appStats.TotalReportingContainers == 0 {
			totalDiskInfo = fmt.Sprintf("%9v", "--")
		} else {
			totalDiskInfo = fmt.Sprintf("%9v", util.ByteSize(appStats.TotalReservedDisk).StringWithPrecision(1))
		}
		return fmt.Sprintf("%9v", totalDiskInfo)
	}
	rawValueFunc := func(data uiCommon.IData) string {
		appStats := data.(*DisplaySpace)
		return fmt.Sprintf("%v", appStats.TotalReservedDisk)
	}
	c := uiCommon.NewListColumn("RSVD_DSK", "RSVD_DSK", 9,
		uiCommon.NUMERIC, false, sortFunc, true, displayFunc, rawValueFunc, nil)
	return c
}

func columnTotalDiskUsed() *uiCommon.ListColumn {
	sortFunc := func(c1, c2 util.Sortable) bool {
		return c1.(*DisplaySpace).TotalUsedDisk < c2.(*DisplaySpace).TotalUsedDisk
	}
	displayFunc := func(data uiCommon.IData, columnOwner uiCommon.IColumnOwner) string {
		appStats := data.(*DisplaySpace)
		totalDiskInfo := ""
		if appStats.TotalReportingContainers == 0 {
			totalDiskInfo = fmt.Sprintf("%9v", "--")
		} else {
			totalDiskInfo = fmt.Sprintf("%9v", util.ByteSize(appStats.TotalUsedDisk).StringWithPrecision(1))
		}
		return fmt.Sprintf("%9v", totalDiskInfo)
	}
	rawValueFunc := func(data uiCommon.IData) string {
		appStats := data.(*DisplaySpace)
		return fmt.Sprintf("%v", appStats.TotalUsedDisk)
	}
	c := uiCommon.NewListColumn("USED_DSK", "USED_DSK", 9,
		uiCommon.NUMERIC, false, sortFunc, true, displayFunc, rawValueFunc, nil)
	return c
}

func columnLogStdout() *uiCommon.ListColumn {
	sortFunc := func(c1, c2 util.Sortable) bool {
		return c1.(*DisplaySpace).TotalLogStdout < c2.(*DisplaySpace).TotalLogStdout
	}
	displayFunc := func(data uiCommon.IData, columnOwner uiCommon.IColumnOwner) string {
		appStats := data.(*DisplaySpace)
		return fmt.Sprintf("%11v", util.Format(appStats.TotalLogStdout))
	}
	rawValueFunc := func(data uiCommon.IData) string {
		appStats := data.(*DisplaySpace)
		return fmt.Sprintf("%v", appStats.TotalLogStdout)
	}
	c := uiCommon.NewListColumn("TotalLogStdout", "LOG_OUT", 11,
		uiCommon.NUMERIC, false, sortFunc, true, displayFunc, rawValueFunc, nil)
	return c
}

func columnLogStderr() *uiCommon.ListColumn {
	sortFunc := func(c1, c2 util.Sortable) bool {
		return c1.(*DisplaySpace).TotalLogStderr < c2.(*DisplaySpace).TotalLogStderr
	}
	displayFunc := func(data uiCommon.IData, columnOwner uiCommon.IColumnOwner) string {
		appStats := data.(*DisplaySpace)
		return fmt.Sprintf("%11v", util.Format(appStats.TotalLogStderr))
	}
	rawValueFunc := func(data uiCommon.IData) string {
		appStats := data.(*DisplaySpace)
		return fmt.Sprintf("%v", appStats.TotalLogStderr)
	}
	c := uiCommon.NewListColumn("TotalLogStderr", "LOG_ERR", 11,
		uiCommon.NUMERIC, false, sortFunc, true, displayFunc, rawValueFunc, nil)
	return c
}

func columnTotalReq() *uiCommon.ListColumn {
	sortFunc := func(c1, c2 util.Sortable) bool {
		return c1.(*DisplaySpace).HttpAllCount < c2.(*DisplaySpace).HttpAllCount
	}
	displayFunc := func(data uiCommon.IData, columnOwner uiCommon.IColumnOwner) string {
		appStats := data.(*DisplaySpace)
		return fmt.Sprintf("%11v", util.Format(appStats.HttpAllCount))
	}
	rawValueFunc := func(data uiCommon.IData) string {
		appStats := data.(*DisplaySpace)
		return fmt.Sprintf("%v", appStats.HttpAllCount)
	}
	c := uiCommon.NewListColumn("TOTREQ", "TOT_REQ", 11,
		uiCommon.NUMERIC, false, sortFunc, true, displayFunc, rawValueFunc, nil)
	return c
}
