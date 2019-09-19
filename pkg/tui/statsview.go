/*
Copyright © 2019 Portworx

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package tui

import (
	"fmt"
	"strings"
	"time"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"github.com/portworx/pxc/pkg/util"
)

type StatsModel interface {
	// Refresh will load  new set of data
	Refresh() error
	// GetTitle returns the title of the table
	GetTitle() string
	// GetHeaders will returns the colume titles for the table
	GetHeaders() []string
	// NextRow return of a zero length array indicates iteration is complete
	NextRow() ([]string, error)
	// SetSort sets the  column to sort on and sets the order of sort
	SetSortInfo(colName string, ascending bool)
	// GetSortInfo get the sort info
	GetSortInfo() (string, bool)
	// Moves the sorting to the next column
	MoveSortColumnNext()
	// Moves the sorting to the prev column
	MoveSortColumnPrev()
	// GetGraphTitle returns the title for the given graph index
	GetGraphTitle(index int) (string, error)
	// GetGraphData returns the current datapoint for the given graph
	GetGraphData(index int) (float64, error)
	// Humanize formats the given value into a string that is easy to read
	Humanize(index int, val float64) (string, error)
}

type View interface {
	Display(ti StatsModel, refreshInterval time.Duration) error
}

type graphInfo struct {
	slg       *widgets.SparklineGroup
	dataStore []float64
}

type statsView struct {
	grid       *ui.Grid
	topLeft    *widgets.Paragraph
	table      *widgets.Table
	gi         []*graphInfo
	termWidth  int
	termHeight int
	curTitle   string
	curTimeStr string
}

const (
	TOP_LINE_HEIGHT   = 2
	TABLE_WIDTH_RATIO = 0.75 // Occupy 75% of terminal
	MAX_GRAPH_POINTS  = 400
)

func NewStatsView(numPlots int) View {
	tv := &statsView{
		gi: make([]*graphInfo, numPlots),
	}

	for i, _ := range tv.gi {
		tv.gi[i] = &graphInfo{
			dataStore: make([]float64, MAX_GRAPH_POINTS),
		}
	}
	return tv
}

func (tv *statsView) Display(ti StatsModel, interval time.Duration) error {
	if err := ui.Init(); err != nil {
		return err
	}
	defer ui.Close()

	tv.layoutComponents()

	if err := tv.render(ti); err != nil {
		return err
	}

	uiEvents := ui.PollEvents()
	ticker := time.NewTicker(interval).C
	for {
		select {
		case e := <-uiEvents:
			switch e.ID {
			case "s":
				if err := tv.toggleSortOrder(ti); err != nil {
					return err
				}
			case "h":
				if err := tv.moveSortColumnPrev(ti); err != nil {
					return err
				}
			case "l":
				if err := tv.moveSortColumnNext(ti); err != nil {
					return err
				}
			case "r":
				if err := tv.render(ti); err != nil {
					return err
				}
			case "q", "<C-c>":
				return nil
			case "<Resize>":
				payload := e.Payload.(ui.Resize)
				tv.resize(payload.Width, payload.Height)
			}
		case <-ticker:
			if err := tv.render(ti); err != nil {
				return err
			}
		}
	}
}

func getCurrentDateTime() string {
	now := time.Now()
	return now.Format(util.TimeFormat)
}

func (tv *statsView) toggleSortOrder(ti StatsModel) error {
	s, o := ti.GetSortInfo()
	ti.SetSortInfo(s, !o)
	return tv.render(ti)
}

func (tv *statsView) moveSortColumnNext(ti StatsModel) error {
	ti.MoveSortColumnNext()
	return tv.render(ti)
}

func (tv *statsView) moveSortColumnPrev(ti StatsModel) error {
	ti.MoveSortColumnPrev()
	return tv.render(ti)
}

func (tv *statsView) getTopLeft(w int) *widgets.Paragraph {
	p := widgets.NewParagraph()
	p.PaddingTop = -1
	p.PaddingRight = -1
	p.PaddingBottom = -1
	p.Border = false
	p.SetRect(0, 0, w, TOP_LINE_HEIGHT)
	return p
}

func (tv *statsView) getTable() *widgets.Table {
	t := widgets.NewTable()
	t.TextAlignment = ui.AlignLeft
	t.Border = false
	t.RowSeparator = false

	t.RowStyles[0] = ui.NewStyle(ui.ColorClear, ui.ColorClear, ui.ModifierBold)
	return t
}

func (tv *statsView) getGraphRows() []interface{} {
	numGraphs := len(tv.gi)
	cols := make([]interface{}, numGraphs)
	for i, _ := range tv.gi {
		sl := widgets.NewSparkline()
		tv.gi[i].slg = widgets.NewSparklineGroup(sl)
		cols[i] = ui.NewRow(1.0/float64(numGraphs), tv.gi[i].slg)
	}
	return cols
}

func (tv *statsView) getGraphCol() interface{} {
	r := tv.getGraphRows()
	if len(r) > 0 {
		return ui.NewCol(1.0-TABLE_WIDTH_RATIO, r...)
	}
	return nil
}

func (tv *statsView) layoutComponents() {
	tv.termWidth, tv.termHeight = ui.TerminalDimensions()
	tv.topLeft = tv.getTopLeft(tv.termWidth)
	tv.table = tv.getTable()
	tv.grid = ui.NewGrid()
	tv.grid.SetRect(0, TOP_LINE_HEIGHT-1,
		tv.termWidth, tv.termHeight-TOP_LINE_HEIGHT)

	cols := make([]interface{}, 0)
	gcol := tv.getGraphCol()
	tableWidthRatio := 1.0
	if gcol != nil {
		tableWidthRatio = TABLE_WIDTH_RATIO
		cols = append(cols, gcol)
	}
	cols = append(cols, ui.NewCol(tableWidthRatio, tv.table))
	tv.grid.Set(
		ui.NewRow(1.0/1,
			cols...,
		),
	)
}

func (tv *statsView) render(ti StatsModel) error {
	s1 := tv.curTitle
	s2 := tv.curTimeStr

	if ti != nil {
		err := ti.Refresh()
		if err != nil {
			return nil
		}
		if err := tv.fillTable(ti); err != nil {
			return err
		}
		if err := tv.fillGraphData(ti); err != nil {
			return err
		}
		s1 = ti.GetTitle()
		s2 = getCurrentDateTime()
		tv.curTitle = s1
		tv.curTimeStr = s2
	}
	l := len(s1) + len(s2) + 2
	s3 := "  "
	if l < tv.termWidth {
		s3 = strings.Repeat(" ", tv.termWidth-l)
	}
	tv.topLeft.Text = fmt.Sprintf("%s%s%s", s1, s3, s2)
	ui.Render(tv.topLeft, tv.grid)
	return nil
}

func (tv *statsView) resize(w int, h int) {
	tv.termWidth = w
	tv.termHeight = h
	tv.topLeft.SetRect(0, 0, w, TOP_LINE_HEIGHT)
	tv.grid.SetRect(0, TOP_LINE_HEIGHT, w, h-TOP_LINE_HEIGHT)
	ui.Clear()
	tv.render(nil)
}

func (tv *statsView) fillTable(ti StatsModel) error {
	headers := ti.GetHeaders()
	rows := make([][]string, 0)
	rows = append(rows, headers)
	for {
		n, err := ti.NextRow()
		if err != nil {
			return err
		}
		if len(n) == 0 {
			break
		}
		rows = append(rows, n)
	}
	tv.table.Rows = rows
	return nil
}

func (tv *statsView) fillGraph(ti StatsModel, index int) error {
	// Get the title for the graph
	t, err := ti.GetGraphTitle(index)
	if err != nil {
		return err
	}
	tv.gi[index].slg.Title = t

	// Remove the first data point
	copy(tv.gi[index].dataStore, tv.gi[index].dataStore[1:])

	// Get the data point for this graph and add it as the last data point
	d, err := ti.GetGraphData(index)
	if err != nil {
		return err
	}
	tv.gi[index].dataStore[len(tv.gi[index].dataStore)-1] = d

	// Figure out the viewing area
	r := tv.gi[index].slg.GetRect()
	viewArea := r.Max.X - r.Min.X - 1
	if viewArea < 0 {
		viewArea = 0
	}
	// Set the data for the viewing area
	tv.gi[index].slg.Sparklines[0].Data =
		tv.gi[index].dataStore[len(tv.gi[index].dataStore)-viewArea:]

	// Get the max value in current dataset and set it in the graph for reference
	maxVal, _ := ui.GetMaxFloat64FromSlice(tv.gi[index].slg.Sparklines[0].Data)
	m, err := ti.Humanize(index, maxVal)
	if err != nil {
		return err
	}
	tv.gi[index].slg.Sparklines[0].Title = m
	return nil
}

func (tv *statsView) fillGraphData(ti StatsModel) error {
	for i, _ := range tv.gi {
		err := tv.fillGraph(ti, i)
		if err != nil {
			return err
		}
	}
	return nil
}
