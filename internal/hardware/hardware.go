package hardware

import (
	"runtime"
	"strconv"

	"github.com/shirou/gopsutil/v4/host"
	"github.com/shirou/gopsutil/v4/mem"
)

type SystemInfo struct {
    Os string
    Platform string
    Hostname string
    NumberOfProcessRunning string
    TotalMem string
    FreeMem string
    PercentageUsedMem string
}

const megabyteDiv uint64 = 1024 * 1024
const gigabyteDiv uint64 = megabyteDiv * 1024

func GetSystemSection() (SystemInfo, error) {
    runTimeOs := runtime.GOOS

    vmStat, err := mem.VirtualMemory()
    if err != nil {
        return SystemInfo{}, err
    }

    hostStat, err := host.Info()
    if err != nil {
        return SystemInfo{}, err
    }

    return SystemInfo{
        Os: runTimeOs,
        Platform: hostStat.Platform,
        Hostname: hostStat.Hostname,
        NumberOfProcessRunning: strconv.FormatUint(hostStat.Procs, 10),
        TotalMem: strconv.FormatUint(vmStat.Total/megabyteDiv, 10),
        FreeMem: strconv.FormatUint(vmStat.Free/megabyteDiv, 10),
        PercentageUsedMem: strconv.FormatFloat(vmStat.UsedPercent, 'f', 2, 64),
    }, nil
}
