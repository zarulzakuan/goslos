package main

type SimulationProfile struct {
	CPU     CPUProfile     `json:"CPU"`
	Memory  MemoryProfile  `json:"Memory"`
	Disk    DiskProfile    `json:"Disk"`
	Network NetworkProfile `json:"Network"`
}

type CPUProfile struct {
	CoresCount         int  `json:"CoresCount"`
	RunForXSeconds     uint `json:"RunForXSeconds"`
	MinUsagePercentage uint `json:"MinUsagePercentage"`
	MaxUsagePercentage uint `json:"MaxUsagePercentage"`
	Run                bool `json:"Run"`
	RunAfterXSeconds   uint `json:"RunAfterXSeconds"`
}

type MemoryProfile struct {
	MaxMemPercentage uint `json:"MaxMemPercentage"`
	Run              bool `json:"Run"`
	RunAfterXSeconds uint `json:"RunAfterXSeconds"`
}

type DiskProfile struct {
	MinChunkSizeInByte     string  `json:"MinChunkSizeInByte"`
	MaxChunkSizeInByte     string  `json:"MaxChunkSizeInByte"`
	TotalSizeInByte        string  `json:"TotalSizeInByte"`
	WriteRatioInPercentage float32 `json:"WriteRatioInPercentage"`
	Run                    bool    `json:"Run"`
	RunAfterXSeconds       uint    `json:"RunAfterXSeconds"`
}

type NetworkProfile struct {
	Run              bool `json:"Run"`
	RunAfterXSeconds uint `json:"RunAfterXSeconds"`
}
