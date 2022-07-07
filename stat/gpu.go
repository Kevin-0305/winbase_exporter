package stat

import (
	"fmt"
	"time"

	nvml "github.com/mxpv/nvml-go"
)

type GPUStat struct {
	NvmlApi           *nvml.API
	Device            nvml.Device
	GPUName           string
	NvidiaDriver      string
	GPUFanSpeed       int
	GPUUsePercent     float64
	GPUMemoryTotal    int
	GPUMemoryUsed     int
	GPUMemoryUsedMax  int
	GPUPowerUsage     int
	GPUPowerUsageMax  int
	GPUPowerUsageAvg  float64
	GPUTemperature    int
	GPUTemperatureMax int
	GPUTemperatureAvg float64
}

func NewGPUStat() *GPUStat {
	w, device := create()
	gpuName, _ := w.DeviceGetName(device)
	version, _ := w.SystemGetDriverVersion()
	gpuStat := GPUStat{
		NvmlApi:           w,
		Device:            device,
		GPUName:           gpuName,
		NvidiaDriver:      version,
		GPUFanSpeed:       0,
		GPUUsePercent:     0,
		GPUMemoryTotal:    0,
		GPUMemoryUsed:     0,
		GPUMemoryUsedMax:  0,
		GPUPowerUsage:     0,
		GPUPowerUsageMax:  0,
		GPUPowerUsageAvg:  0,
		GPUTemperature:    0,
		GPUTemperatureMax: 0,
		GPUTemperatureAvg: 0,
	}
	go gpuStat.TimingStat()
	return &gpuStat
}

func (s *GPUStat) TimingStat() {
	for {
		speed, err := s.NvmlApi.DeviceGetFanSpeed(s.Device)
		if err != nil {
			fmt.Println(err.Error())
		}
		temp, _ := s.NvmlApi.DeviceGetTemperature(s.Device, nvml.TemperatureGPU)
		mem, _ := s.NvmlApi.DeviceGetBAR1MemoryInfo(s.Device)
		s.GPUFanSpeed = int(speed)
		if temp > uint32(s.GPUTemperature) {
			s.GPUTemperatureMax = int(temp)
		}
		s.GPUTemperature = int(temp)
		s.GPUMemoryTotal = int(mem.Total)
		s.GPUMemoryUsed = int(mem.Used)
		if s.GPUMemoryUsed > s.GPUMemoryUsedMax {
			s.GPUMemoryUsedMax = s.GPUMemoryUsed
		}
		GPUPowerUsage, _ := s.NvmlApi.DeviceGetPowerUsage(s.Device)
		if int(GPUPowerUsage) > s.GPUPowerUsageMax {
			s.GPUPowerUsageMax = int(GPUPowerUsage)
		}
		s.GPUPowerUsage = int(GPUPowerUsage)
		GPUUsePercent, _ := s.NvmlApi.DeviceGetUtilizationRates(s.Device)
		s.GPUUsePercent = float64(GPUUsePercent.GPU)
		time.Sleep(time.Second * 1)
	}
}

// TestGPUStat test GPUStat
func GPUStatTest() {
	w, device := create()
	defer w.Shutdown()

	rates, err := w.DeviceGetUtilizationRates(device)
	if err != nil {
		fmt.Println(err.Error())
	}
	speed, err := w.DeviceGetFanSpeed(device)
	if err != nil {
		fmt.Println(err.Error())
	}
	temp, err := w.DeviceGetTemperature(device, nvml.TemperatureGPU)
	if err != nil {
		fmt.Println(err.Error())
	}
	mem, err := w.DeviceGetBAR1MemoryInfo(device)
	//name, err := w.DeviceGetName(device)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("GPU:  ", rates.GPU)
	fmt.Println("Memory:  ", rates.Memory)
	fmt.Println("FanSpeed:  ", speed)
	fmt.Println("Temperature:  ", temp)
	fmt.Println("MemoryTotal:  ", mem.Total)
	fmt.Println("MemoryUsed:  ", mem.Used)
}

func create() (*nvml.API, nvml.Device) {
	w, err := nvml.New("nvml.dll")
	if err != nil {
		fmt.Println(err.Error())
	}
	err = w.Init()
	if err != nil {
		fmt.Println(err.Error())
	}
	device, err := w.DeviceGetHandleByIndex(0)
	if err != nil {
		fmt.Println(err.Error())
	}
	return w, device
}
