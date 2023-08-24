package main

import (
	"fmt"
	"math"

	"github.com/ethereum/go-ethereum/cmd/utils"
	"github.com/ethereum/go-ethereum/params"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
	"gopkg.in/urfave/cli.v1"
)

var (
	fingerprintCommand = cli.Command{
		Name:      "fingerprint",
		Usage:     "Display the system fingerprint",
		ArgsUsage: "",
		Action:    utils.MigrateFlags(showFingerprint),
		Category:  "FINGERPRINT COMMANDS",
	}
)

func getCoresCount(cp []cpu.InfoStat) int {
	cores := 0
	for i := 0; i < len(cp); i++ {
		cores += int(cp[i].Cores)
	}
	return cores
}

// Run implements the cli.Command interface
func showFingerprint(_ *cli.Context) error {
	v, _ := mem.VirtualMemory()
	h, _ := host.Info()
	cp, _ := cpu.Info()
	d, _ := disk.Usage("/")

	osName := h.OS
	osVer := h.Platform + " - " + h.PlatformVersion + " - " + h.KernelArch
	totalMem := math.Floor(float64(v.Total)/(1024*1024*1024)*100) / 100
	availableMem := math.Floor(float64(v.Available)/(1024*1024*1024)*100) / 100
	usedMem := math.Floor(float64(v.Used)/(1024*1024*1024)*100) / 100
	totalDisk := math.Floor(float64(d.Total)/(1024*1024*1024)*100) / 100
	availableDisk := math.Floor(float64(d.Free)/(1024*1024*1024)*100) / 100
	usedDisk := math.Floor(float64(d.Used)/(1024*1024*1024)*100) / 100

	borDetails := fmt.Sprintf("Bor Version : %s", params.VersionWithMeta)
	cpuDetails := fmt.Sprintf("CPU : %d cores", getCoresCount(cp))
	osDetails := fmt.Sprintf("OS : %s %s ", osName, osVer)
	memDetails := fmt.Sprintf("RAM :: total : %v GB, free : %v GB, used : %v GB", totalMem, availableMem, usedMem)
	diskDetails := fmt.Sprintf("STORAGE :: total : %v GB, free : %v GB, used : %v GB", totalDisk, availableDisk, usedDisk)

	fmt.Println(borDetails)
	fmt.Println(cpuDetails)
	fmt.Println(osDetails)
	fmt.Println(memDetails)
	fmt.Println(diskDetails)
	return nil
}
