package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func PostRunProfile(c echo.Context) error {
	profile := new(SimulationProfile)
	if err := c.Bind(profile); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if profile.CPU.Run {
		log.Printf("Start CPU Simulation: core(%d) run(%ds) min(%d) max(%d)\n", profile.CPU.CoresCount, profile.CPU.RunForXSeconds, profile.CPU.MinUsagePercentage, profile.CPU.MaxUsagePercentage)
		go func() {
			time.Sleep(time.Duration(profile.CPU.RunAfterXSeconds) * time.Second)
			RunCPULoad(profile.CPU.CoresCount, profile.CPU.RunForXSeconds, profile.CPU.MinUsagePercentage, profile.CPU.MaxUsagePercentage)
		}()

	}

	if profile.Memory.Run {
		log.Printf("Start Memory Simulation: max(%d)\n", profile.Memory.MaxMemPercentage)
		go func() {
			time.Sleep(time.Duration(profile.CPU.RunAfterXSeconds) * time.Second)
			RunMemLoad(profile.Memory.MaxMemPercentage)
		}()
	}

	if profile.Disk.Run {
		log.Printf("Start Disk IO Simulation: min(%s) max(%s) total(%s) ratio(%f)\n", profile.Disk.MinChunkSizeInByte, profile.Disk.MaxChunkSizeInByte, profile.Disk.TotalSizeInByte, profile.Disk.WriteRatioInPercentage)
		go func() {
			time.Sleep(time.Duration(profile.CPU.RunAfterXSeconds) * time.Second)
			RunDiskIO(profile.Disk.MinChunkSizeInByte, profile.Disk.MaxChunkSizeInByte, profile.Disk.TotalSizeInByte, profile.Disk.WriteRatioInPercentage)
		}()
	}

	if profile.Network.Run {
		log.Printf("Start Network IO Simulation")
		go func() {
			time.Sleep(time.Duration(profile.CPU.RunAfterXSeconds) * time.Second)
			RunNetworkIO()
		}()
	}

	return c.JSON(http.StatusOK, profile)
}

func GetStatOnce(c echo.Context) error {

	temp := `
--------------------------------------------------------
--------------------------------------------------------
CPU Usage    | %s
Memory Usage | %s
Disk IO      | Read: %d Write: %d
Net IO       | Recv: %d Sent : %d
--------------------------------------------------------
--------------------------------------------------------
`
	cpuUsage := GetCPUUsage(false)
	memUsage := GetMemoryUsage(false)
	diskIO := GetDiskIO(false)
	NetIO := GetNetIO(false)
	res := fmt.Sprintf(temp,
		cpuUsage.UsedPercentageString,
		memUsage.UsedPercentageString,
		diskIO.ByteRead,
		diskIO.ByteWritten,
		NetIO.ByteReceived,
		NetIO.ByteSent)
	return c.String(http.StatusOK, res)
}
