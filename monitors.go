package main

import (
	"strconv"
	"time"

	cpuutil "github.com/shirou/gopsutil/cpu"
	diskutil "github.com/shirou/gopsutil/disk"
	memutil "github.com/shirou/gopsutil/mem"
	netutil "github.com/shirou/gopsutil/net"
)

type CPUStat struct {
	UsedPercentageFloat  float64
	UsedPercentageString string
}

type MemStat struct {
	UsedPercentageFloat  float64
	UsedPercentageString string
}

type DiskStat struct {
	ByteRead    uint64
	ByteWritten uint64
}

type NetStat struct {
	ByteReceived uint64
	ByteSent     uint64
}

func GetCPUUsage(loop bool) *CPUStat {
	cpuStat := new(CPUStat)
	for {
		cpuu, _ := cpuutil.Percent(0, false)
		for _, cpu := range cpuu {
			cpuStat.UsedPercentageFloat = cpu
			cpuStat.UsedPercentageString = strconv.FormatFloat(cpu, 'f', 2, 64)

		}
		if !loop {
			break
		}
		time.Sleep(1000 * time.Millisecond)
	}
	return cpuStat
}

func GetMemoryUsage(loop bool) *MemStat {
	memStat := new(MemStat)
	for {
		vmstat, _ := memutil.VirtualMemory()

		memStat.UsedPercentageFloat = vmstat.UsedPercent
		memStat.UsedPercentageString = strconv.FormatFloat(vmstat.UsedPercent, 'f', 2, 64)
		if !loop {
			break
		}
		time.Sleep(1000 * time.Millisecond)
	}
	return memStat
}

func GetDiskIO(loop bool) *DiskStat {
	diskStat := new(DiskStat)
	var _bytesRead, _bytesWritten uint64 = 0, 0
	count := 0
	for {

		diskIO, _ := diskutil.IOCounters("sdb")

		for _, v := range diskIO {
			diskStat.ByteRead = (v.ReadBytes - _bytesRead)
			diskStat.ByteWritten = (v.WriteBytes - _bytesWritten)
			_bytesRead = v.ReadBytes
			_bytesWritten = v.WriteBytes
		}
		if count < 1 {
			count++
			time.Sleep(100 * time.Millisecond)
			continue
		}
		if !loop {
			break
		}
		time.Sleep(1000 * time.Millisecond)
	}
	return diskStat
}

func GetNetIO(loop bool) *NetStat {
	netStat := new(NetStat)
	var _bytesReceived, _bytesSent uint64 = 0, 0
	count := 0
	for {
		nets, _ := netutil.IOCounters(false)

		for _, net := range nets {
			netStat.ByteReceived = (net.BytesRecv - _bytesReceived)
			netStat.ByteSent = (net.BytesSent - _bytesSent)
			_bytesReceived = net.BytesRecv
			_bytesSent = net.BytesSent
		}
		if count < 1 {
			count++
			time.Sleep(100 * time.Millisecond)
			continue
		}
		if !loop {
			break
		}
		time.Sleep(1000 * time.Millisecond)
	}
	return netStat
}
