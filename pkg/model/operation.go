package model

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/opencars/translit"

	"github.com/opencars/operations/pkg/utils"
)

// Operation represents entity in the store.Store.
type Operation struct {
	Person      string   `json:"person" db:"person" csv:"person"`
	RegAddress  *string  `json:"reg_addr_koatuu,omitempty" db:"reg_address" csv:"reg_addr_koatuu"`
	RegCode     int16    `json:"registration_code" db:"code" csv:"oper_code"`
	Reg         string   `json:"registration" db:"name" csv:"oper_name"`
	Date        string   `json:"date" db:"reg_date" csv:"d_reg"`
	DepCode     int32    `json:"dep_code" db:"office_id" csv:"dep_code"`
	Dep         string   `json:"dep" db:"office_name" csv:"dep"`
	Brand       string   `json:"brand" db:"make" csv:"brand"`
	Model       string   `json:"model" db:"model" csv:"model"`
	Year        int16    `json:"year" db:"year" csv:"make_year"`
	Color       string   `json:"color" db:"color" csv:"color"`
	Kind        string   `json:"kind" db:"kind" csv:"kind"`
	Body        string   `json:"body" db:"body" csv:"body"`
	Purpose     string   `json:"purpose" db:"purpose" csv:"purpose"`
	Fuel        *string  `json:"fuel,omitempty" db:"fuel" csv:"fuel"`
	Capacity    *int     `json:"capacity,omitempty" db:"capacity" csv:"capacity"`
	OwnWeight   *float64 `json:"own_weight,omitempty" db:"own_weight" csv:"own_weight"`
	TotalWeight *float64 `json:"total_weight,omitempty" db:"total_weight" csv:"total_weight"`
	Number      string   `json:"number" db:"number" csv:"n_reg_new"`
	ResourceID  int64    `json:"-" db:"resource_id" csv:"-"`
}

// FixDate returns fixed date in string format.
func FixDate(lexeme string) string {
	r := regexp.MustCompile(`^([0-9]{2})\.([0-9]{2})\.([0-9]{4})$`)

	if !r.MatchString(lexeme) {
		return lexeme
	}

	date := r.ReplaceAllString(lexeme, "$1")
	month := r.ReplaceAllString(lexeme, "$2")
	year := r.ReplaceAllString(lexeme, "$3")

	return fmt.Sprintf("%s-%s-%s", year, month, date)
}

// FixBrandModel returns fixed brand.
func FixBrandModel(brandModel, model string) (resBrand string, resModel string) {
	resModel = strings.Join(strings.Fields(strings.TrimSpace(model)), " ")
	resBrand = strings.Join(strings.Fields(strings.TrimSpace(brandModel)), " ")
	resBrand = strings.TrimSpace(strings.TrimSuffix(resBrand, resModel))

	return
}

// OperationFromGov returns new instance of Operation from CSV row.
func OperationFromGov(columns []string) (*Operation, error) {
	code, err := strconv.ParseInt(columns[2], 10, 16)
	if err != nil {
		return nil, fmt.Errorf("failed to parse code: %w", err)
	}

	office, err := strconv.ParseInt(columns[5], 10, 32)
	if err != nil {
		return nil, fmt.Errorf("failed to parse office: %w", err)
	}

	year, err := strconv.ParseInt(columns[9], 10, 16)
	if err != nil {
		return nil, fmt.Errorf("failed to parse year: %w", err)
	}

	capacity, err := utils.Atoi(&columns[15])
	if err != nil {
		return nil, fmt.Errorf("failed to parse capacity: %w", err)
	}

	ownWeight, err := utils.Atof(&columns[16])
	if err != nil {
		return nil, fmt.Errorf("failed to parse ownWeight: %w", err)
	}

	totalWeight, err := utils.Atof(&columns[17])
	if err != nil {
		return nil, fmt.Errorf("failed to parse totalWeight: %w", err)
	}

	name := strings.ReplaceAll(columns[3], columns[2], "")
	name = *utils.Trim(&name)

	brand, model := FixBrandModel(columns[7], columns[8])

	return &Operation{
		Person:      columns[0],                 // person.
		RegAddress:  utils.Trim(&columns[1]),    // reg_addr_koatuu.
		RegCode:     int16(code),                // oper_code.
		Reg:         name,                       // oper_name.
		Date:        FixDate(columns[4]),        // d_reg.
		DepCode:     int32(office),              // dep_code.
		Dep:         columns[6],                 // dep.
		Brand:       brand,                      // brand.
		Model:       model,                      // model.
		Year:        int16(year),                // make_year.
		Color:       columns[10],                // color.
		Kind:        columns[11],                // kind.
		Body:        columns[12],                // body.
		Purpose:     columns[13],                // purpose.
		Fuel:        utils.Trim(&columns[14]),   // fuel.
		Capacity:    capacity,                   // capacity.
		OwnWeight:   ownWeight,                  // own_weight.
		TotalWeight: totalWeight,                // total_weight.
		Number:      translit.ToUA(columns[18]), // n_reg_new.
	}, nil
}
