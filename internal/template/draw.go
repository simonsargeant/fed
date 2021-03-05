package template

import (
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/signintech/gopdf"
)

type Drawer struct {
	pdf   *gopdf.GoPdf
	debug bool
}

func NewDrawer(debug bool, logger interface{}) Drawer {
	return Drawer{
		pdf:   &gopdf.GoPdf{},
		debug: debug,
	}
}

func (d Drawer) Draw(info Invoice, fontPath string, out io.Writer) error {
	d.pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4})
	d.pdf.AddPageWithOption(gopdf.PageOption{PageSize: gopdf.PageSizeA4})

	err := d.pdf.AddTTFFont("font", fontPath)
	if err != nil {
		return fmt.Errorf("add pdf font: %w", err)
	}

	var x float64 = 20
	var y float64 = 20

	d.pdf.SetMargins(x, y, x, y)
	d.DebugRect(x, y, 555, 802)

	err = d.pdf.SetFont("font", "", 14)
	if err != nil {
		return fmt.Errorf("set pdf font: %w", err)
	}

	y = 70
	d.pdf.SetX(x)
	d.pdf.SetY(y)
	d.DebugRect(x, y, 200, 20)
	if err := d.pdf.CellWithOption(NewRect(200, 20), info.Company.Name, NewTopLeft()); err != nil {
		return fmt.Errorf("write cell at (%f, %f): %w", x, y, err)
	}

	err = d.pdf.SetFont("font", "", 10)
	if err != nil {
		return fmt.Errorf("set pdf font: %w", err)
	}

	y += 20

	_, y, err = d.PrintList(NewTopLeft(), func(x, y float64) (float64, float64) {
		return x, y + 12
	}, x, y, 200, 12,
		info.Company.TaxID,
		info.Company.Address.Street,
		info.Company.Address.PostCode+" "+info.Company.Address.Area,
		info.Company.Address.Country,
		"",
		info.Contact.Email,
		info.Contact.Phone,
		"",
		"Billed To:",
		info.Client.Name,
		info.Client.TaxID,
		info.Client.Address.Street,
		info.Client.Address.PostCode+" "+info.Client.Address.Area,
		info.Client.Address.Country,
	)

	if err != nil {
		return fmt.Errorf("print addresses: %w", err)
	}

	x = 20
	y += 30

	tableLines := [][]string{
		{
			"Item",
			"Unit Price",
			"Qty.",
			"Total",
		},
	}

	for _, line := range info.Order.Lines {
		tableLines = append(tableLines, []string{
			line.Item,
			line.Cost,
			strconv.Itoa(line.Quantity),
			line.Total,
		})
	}

	_, y, err = d.DrawTable(x, y, []float64{315, 90, 60, 90}, tableLines)

	if err != nil {
		return fmt.Errorf("draw table: %w", err)
	}

	x = 395
	yTotals := y
	x, _, err = d.PrintList(NewTopRight(), func(x, y float64) (float64, float64) {
		return x, y + 14
	}, x, y, 90, 14,
		"Subtotal:",
		"Tax Rate:",
		"Tax Amount:",
		"Total:",
	)

	if err != nil {
		return fmt.Errorf("print list: %w", err)
	}

	x += 90
	y = yTotals
	_, y, err = d.PrintList(NewTopRight(), func(x, y float64) (float64, float64) {
		return x, y + 14
	}, x, y, 90, 14,
		info.Order.Subtotal,
		info.Order.TaxName+" "+info.Order.TaxRate,
		info.Order.TaxAmount,
		info.Order.Total,
	)

	if err != nil {
		return fmt.Errorf("print list: %w", err)
	}

	x = 20
	y += 12
	_, _, err = d.PrintList(NewTopLeft(), func(x, y float64) (float64, float64) {
		return x, y + 12
	}, x, y, 300, 12,
		info.Notes,
		"",
		"IBAN: "+info.Billing.IBAN,
		"BIC: "+info.Billing.BIC,
	)

	if err != nil {
		return fmt.Errorf("print billing info: %w", err)
	}

	// Top Right Info
	err = d.pdf.SetFont("font", "", 32)
	if err != nil {
		return fmt.Errorf("set larger pdf font: %w", err)
	}

	x = 375
	y = 20
	d.pdf.SetX(x)
	d.pdf.SetY(y)
	d.DebugRect(x, y, 200, 40)

	if err := d.pdf.CellWithOption(NewRect(200, 40), "Invoice", NewTopRight()); err != nil {
		return fmt.Errorf("write cell at (%f, %f): %w", x, y, err)
	}

	err = d.pdf.SetFont("font", "", 10)
	if err != nil {
		return fmt.Errorf("set smaller pdf font: %w", err)
	}

	y += 50

	_, _, err = d.PrintList(NewTopRight(), func(x, y float64) (float64, float64) {
		return x, y + 12
	}, x, y, 200, 12,
		fmt.Sprintf("No: %04d", info.Metadata.Number),
		"Date: "+info.Metadata.IssueDate,
		"Due Date: "+info.Metadata.PayByDate,
	)

	if err != nil {
		return fmt.Errorf("print invoice metadata: %w", err)
	}

	if err := d.pdf.Write(out); err != nil {
		return fmt.Errorf("write pdf to output: %w", err)
	}

	return nil
}

func NewTopLeft() gopdf.CellOption {
	return gopdf.CellOption{
		Align:  gopdf.Left | gopdf.Top,
		Border: 0,
		Float:  gopdf.Right,
	}
}

func NewMiddleLeft() gopdf.CellOption {
	return gopdf.CellOption{
		Align:  gopdf.Left | gopdf.Middle,
		Border: 0,
		Float:  gopdf.Right,
	}
}

func NewMiddleRight() gopdf.CellOption {
	return gopdf.CellOption{
		Align:  gopdf.Right | gopdf.Middle,
		Border: 0,
		Float:  gopdf.Right,
	}
}

func NewTopRight() gopdf.CellOption {
	return gopdf.CellOption{
		Align:  gopdf.Right | gopdf.Top,
		Border: 0,
		Float:  gopdf.Right,
	}
}

func NewRect(width, height float64) *gopdf.Rect {
	return &gopdf.Rect{
		H: height,
		W: width,
	}
}

func (d Drawer) DrawTable(x, y float64, cellWidth []float64, lines [][]string) (float64, float64, error) {
	initx := x
	for _, line := range lines {
		if len(line) != len(cellWidth) {
			return 0, 0, fmt.Errorf("invalid line: expected length %d, got length %d", len(cellWidth), len(line))
		}

		var err error
		_, y, err = d.HorizontalList(func(i int) gopdf.CellOption {
			if i == 0 {
				return NewMiddleLeft()
			}
			return NewMiddleRight()
		}, func(i int, x, y float64) (float64, float64) {
			return x + cellWidth[i], y
		}, x, y, func(i int) float64 { return cellWidth[i] }, 15, line...)

		if err != nil {
			return 0, 0, fmt.Errorf("print dynamic list: %w", err)
		}

		x = initx
		y += 15
		d.pdf.Line(20, y, 575, y)
		y += 15
	}
	return x, y, nil
}

func (d Drawer) HorizontalList(
	opt func(i int) gopdf.CellOption,
	move func(i int, x, y float64) (float64, float64),
	x,
	y float64,
	width func(i int) float64,
	height float64,
	list ...string,
) (float64, float64, error) {
	// Wrap long text
	maxHeight := height
	items := make([][]string, len(list))
	for i, item := range list {
		w := width(i)
		itemWidth, err := d.pdf.MeasureTextWidth(item)
		if err != nil {
			return 0, 0, fmt.Errorf("measure width of %q: %w", item, err)
		}

		if itemWidth <= w {
			items[i] = []string{item}
			continue
		}

		parts := strings.Split(item, " ")
		var lines []string
		for len(parts) > 0 {
			var line string
			var tempLine string

			j := 0
			for _, part := range parts {
				tempLine += " " + part
				lineWidth, err := d.pdf.MeasureTextWidth(tempLine)
				if err != nil {
					return 0, 0, fmt.Errorf("measure width of line %q: %w", tempLine, err)
				}

				if lineWidth > w {
					if j == 0 {
						return 0, 0, fmt.Errorf("width not big enough to fit one word %q", part)
					}
					break
				}

				line = tempLine
				j++
			}
			lines = append(lines, line)
			parts = parts[j:]
		}

		itemHeight := float64(len(lines)) * height
		if itemHeight > maxHeight {
			maxHeight = itemHeight
		}
		items[i] = lines
	}

	for i, item := range items {
		unitWidth := width(i)
		unitHeight := maxHeight / float64(len(item))

		subY := y
		for j, text := range item {

			if j > 0 {
				subY += unitHeight
			}

			d.pdf.SetX(x)
			d.pdf.SetY(subY)
			d.DebugRect(x, subY, unitWidth, unitHeight)
			_ = d.pdf.CellWithOption(NewRect(unitWidth, unitHeight), text, opt(i))
		}

		x, y = move(i, x, y)
	}

	return x, y + maxHeight, nil
}

func (d Drawer) PrintList(
	opt gopdf.CellOption,
	move func(x, y float64) (float64, float64),
	x,
	y,
	unitWidth,
	unitHeight float64,
	list ...string,
) (float64, float64, error) {
	for _, item := range list {
		d.pdf.SetX(x)
		d.pdf.SetY(y)
		d.DebugRect(x, y, unitWidth, unitHeight)

		// SplitText fails on empty string
		if len(item) == 0 {
			x, y = move(x, y)
			continue
		}

		texts, err := d.pdf.SplitText(item, unitWidth)

		if err != nil {
			return 0, 0, fmt.Errorf("split %s: %w", item, err)
		}

		for i, text := range texts {

			if i > 0 {
				y += unitHeight
			}
			_ = d.pdf.CellWithOption(NewRect(unitWidth, unitHeight), text, opt)
		}

		x, y = move(x, y)
	}

	return x, y, nil
}

func (d Drawer) DebugRect(x, y, width, height float64) {
	if !d.debug {
		return
	}

	d.pdf.SetStrokeColor(255, 0, 0)
	d.pdf.RectFromUpperLeftWithStyle(x, y, width, height, "D")
	d.pdf.SetStrokeColor(0, 0, 0)
}
