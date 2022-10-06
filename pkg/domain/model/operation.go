package model

const DateLayout = "2006-01-02"

// Operation represents entity in the domain.Store.
type Operation struct {
	Person         string   `json:"person" db:"person"`
	RegAddress     *string  `json:"reg_addr_koatuu,omitempty" db:"reg_address"`
	FullRegAddress *string  `json:"full_reg_addr_koatuu,omitempty" db:"-"`
	RegCode        int      `json:"registration_code" db:"code"`
	Reg            string   `json:"registration" db:"name"`
	Date           string   `json:"date" db:"reg_date"`
	DepCode        int32    `json:"dep_code" db:"office_id"`
	Dep            string   `json:"dep" db:"office_name"`
	Brand          string   `json:"brand" db:"make"`
	Model          string   `json:"model" db:"model"`
	VIN            *string  `json:"vin,omitempty" db:"vin"`
	Year           int16    `json:"year" db:"year"`
	Color          string   `json:"color" db:"color"`
	Kind           string   `json:"kind" db:"kind"`
	Body           string   `json:"body" db:"body"`
	Purpose        string   `json:"purpose" db:"purpose"`
	Fuel           *string  `json:"fuel,omitempty" db:"fuel"`
	Capacity       *int     `json:"capacity,omitempty" db:"capacity"`
	OwnWeight      *float64 `json:"own_weight,omitempty" db:"own_weight"`
	TotalWeight    *float64 `json:"total_weight,omitempty" db:"total_weight"`
	Number         string   `json:"number" db:"number"`
	ResourceID     int64    `json:"-" db:"resource_id"`
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
