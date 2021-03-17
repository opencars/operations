package domain

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/opencars/translit"

	"github.com/opencars/operations/pkg/utils"
)

// FixDate returns fixed date in string format.
func FixDate(lexeme string) string {
	r := regexp.MustCompile(`^(\d{2})\.(\d{2})\.(\d{4})$`)

	if !r.MatchString(lexeme) {
		return lexeme
	}

	date := r.ReplaceAllString(lexeme, "$1")
	month := r.ReplaceAllString(lexeme, "$2")
	year := r.ReplaceAllString(lexeme, "$3")

	return fmt.Sprintf("%s-%s-%s", year, month, date)
}

// FixBrandModel returns fixed brand.
func FixBrandModel(brandModel, model string) (resBrand, resModel string) {
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
		Model:       model,                      // domain.
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

func (op *Operation) PrettyPerson() string {
	switch op.Person {
	case "J":
		return "Юридична особа"
	case "P":
		return "Фізична особа"
	}

	return op.Person
}
