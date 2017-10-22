package container

import (
	"github.com/gizak/termui"
	"github.com/bunbunjp/gotop/util"
	"github.com/bunbunjp/gotop/dataservice/memory"
	"fmt"
)

type VirtualMemoryUsageContainer struct {
	virtualGauge *termui.Gauge
}

func (v *VirtualMemoryUsageContainer) Initialize() {
}

func (v *VirtualMemoryUsageContainer) UpdateRender() {
	data := memory.GetInstance()

	v.virtualGauge.Percent = int(data.LatestVirtualStat.UsedPercent)
	v.virtualGauge.BorderLabel = fmt.Sprintf("virtual usage (%.2fGB/%.2fGB)",
		util.Byte2GBi(data.LatestVirtualStat.Used),
		util.Byte2GBi(data.LatestVirtualStat.Total))
}

func (v *VirtualMemoryUsageContainer) CreateUi() termui.GridBufferer {

	v.virtualGauge = termui.NewGauge()
	v.virtualGauge.Width = termui.TermWidth() / 4
	v.virtualGauge.Height = 10
	v.virtualGauge.LabelAlign = termui.AlignRight

	return v.virtualGauge
}