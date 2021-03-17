package domain

import "time"

const DateLayout = "2006-01-02"

// Resource represents entity from the data.gov.ua web portal.
type Resource struct {
	ID           int64     `json:"id" db:"id"`
	UID          string    `json:"uid" db:"uid"`
	Name         string    `json:"name" db:"name"`
	URL          string    `json:"url" db:"url"`
	LastModified time.Time `json:"last_modified" db:"last_modified"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
}

// Revision represents entity in the domain.Store.
type Revision struct {
	ID          string    `json:"id" db:"id"`
	URL         string    `json:"url" db:"url"`
	FileHashSum *string   `json:"file_hash_sum" db:"file_hash_sum"`
	Deleted     int32     `json:"deleted" db:"deleted"`
	Added       int32     `json:"added" db:"added"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
}

// Operation represents entity in the domain.Store.
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
