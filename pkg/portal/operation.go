package portal

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/opencars/translit"

	"github.com/opencars/operations/pkg/domain/model"
	"github.com/opencars/operations/pkg/utils"
)

type Operation struct {
	Person      string   `csv:"PERSON"`
	RegAddress  string   `csv:"REG_ADDR_KOATUU"`
	RegCode     int      `csv:"OPER_CODE"`
	Reg         string   `csv:"OPER_NAME"`
	Date        string   `csv:"D_REG"`
	DepCode     int32    `csv:"DEP_CODE"`
	Dep         string   `csv:"DEP"`
	Brand       string   `csv:"BRAND"`
	Model       string   `csv:"MODEL"`
	Vin         string   `csv:"VIN"`
	Year        int16    `csv:"MAKE_YEAR"`
	Color       string   `csv:"COLOR"`
	Kind        string   `csv:"KIND"`
	Body        string   `csv:"BODY"`
	Purpose     string   `csv:"PURPOSE"`
	Fuel        string   `csv:"FUEL"`
	Capacity    *int     `csv:"CAPACITY"`
	OwnWeight   *float64 `csv:"OWN_WEIGHT"`
	TotalWeight *float64 `csv:"TOTAL_WEIGHT"`
	Number      string   `csv:"N_REG_NEW"`
}

// FixDate returns fixed date in string format.
func FixDate(lexeme string) string {
	r2 := regexp.MustCompile(`^(\d{2})\.(\d{2})\.(\d{2})$`)
	r4 := regexp.MustCompile(`^(\d{2})\.(\d{2})\.(\d{4})$`)

	if r2.MatchString(lexeme) {
		date := r2.ReplaceAllString(lexeme, "$1")
		month := r2.ReplaceAllString(lexeme, "$2")
		year := r2.ReplaceAllString(lexeme, "$3")

		return fmt.Sprintf("%s-%s-%s", "20"+year, month, date)
	}

	if r4.MatchString(lexeme) {
		date := r4.ReplaceAllString(lexeme, "$1")
		month := r4.ReplaceAllString(lexeme, "$2")
		year := r4.ReplaceAllString(lexeme, "$3")

		return fmt.Sprintf("%s-%s-%s", year, month, date)
	}

	return lexeme
}

// FixBrandModel returns fixed brand.
func FixBrandModel(brandModel, mod string) (resBrand, resModel string) {
	resModel = strings.Join(strings.Fields(strings.TrimSpace(mod)), " ")
	resBrand = strings.Join(strings.Fields(strings.TrimSpace(brandModel)), " ")
	resBrand = strings.TrimSpace(strings.TrimSuffix(resBrand, resModel))

	return
}

func (o *Operation) Convert() *model.Operation {
	name := utils.Trim(
		strings.ReplaceAll(o.Reg, strconv.Itoa(o.RegCode), ""),
	)

	if name == nil {
		str := ""
		name = &str
	}

	brand, mod := FixBrandModel(o.Brand, o.Model)

	vin := utils.Trim(o.Vin)

	return &model.Operation{
		Person:      o.Person,
		RegAddress:  utils.Trim(o.RegAddress),
		RegCode:     o.RegCode,
		Reg:         *name,
		Date:        FixDate(o.Date),
		DepCode:     o.DepCode,
		Dep:         o.Dep,
		Brand:       brand,
		Model:       mod,
		VIN:         vin,
		Year:        o.Year,
		Color:       o.Color,
		Kind:        o.Kind,
		Body:        o.Body,
		Purpose:     o.Purpose,
		Fuel:        utils.Trim(o.Fuel),
		Capacity:    o.Capacity,
		OwnWeight:   o.OwnWeight,
		TotalWeight: o.TotalWeight,
		Number:      translit.ToUA(o.Number),
	}
}
