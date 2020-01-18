package model

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/opencars/operations/pkg/utils"
	"github.com/opencars/translit"
)

// Operation represents entity in the store.Store.
type Operation struct {
	Person      string   `json:"person" db:"person" csv:"person"`
	RegAddress  *string  `json:"reg_address" db:"reg_address" csv:"reg_addr_koatuu"`
	Code        int16    `json:"code" db:"code" csv:"oper_code"`
	Name        string   `json:"name" db:"name" csv:"oper_name"`
	Date        string   `json:"reg_date" db:"reg_date" csv:"d_reg"`
	OfficeID    int32    `json:"office_id" db:"office_id" csv:"dep_code"`
	OfficeName  string   `json:"office_name" db:"office_name" csv:"dep"`
	Make        string   `json:"make" db:"make" csv:"brand"`
	Model       string   `json:"model" db:"model" csv:"model"`
	Year        int16    `json:"year" db:"year" csv:"make_year"`
	Color       string   `json:"color" db:"color" csv:"color"`
	Kind        string   `json:"kind" db:"kind" csv:"kind"`
	Body        string   `json:"body" db:"body" csv:"body"`
	Purpose     string   `json:"purpose" db:"purpose" csv:"purpose"`
	Fuel        *string  `json:"fuel" db:"fuel" csv:"fuel"`
	Capacity    *int     `json:"capacity" db:"capacity" csv:"capacity"`
	OwnWeight   *float64 `json:"own_weight" db:"own_weight" csv:"own_weight"`
	TotalWeight *float64 `json:"total_weight" db:"total_weight" csv:"total_weight"`
	Number      string   `json:"number" db:"number" csv:"n_reg_new"`
	ResourceID  int64    `json:"resource_id" db:"resource_id" csv:"-"`
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

	return &Operation{
		Person:      columns[0],                 // person.
		RegAddress:  utils.Trim(&columns[1]),    // reg_addr_koatuu.
		Code:        int16(code),                // oper_code.
		Name:        name,                       // oper_name.
		Date:        FixDate(columns[4]),        // d_reg.
		OfficeID:    int32(office),              // dep_code.
		OfficeName:  columns[6],                 // dep.
		Make:        columns[7],                 // brand.
		Model:       columns[8],                 // model.
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
