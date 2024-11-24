package GoGRF

import "github.com/SwatBros/GoGRF/enums"

type Grf struct {
	Name                   string
	Description            string
	Version                int
	Min_compatible_version int
}

type Cargo struct {
	Id                        int
	Type_name                 string
	Unit_name                 string
	Units_of_cargo            string
	Items_of_cargo            string
	Type_abbreviation         string
	Sprite                    string
	Weight                    int
	Penalty_lowerbound        int
	Single_penalty_length     int
	Price_factor              float32
	Station_list_colour       int
	Cargo_payment_list_colour int
	Is_freight                bool
	Cargo_classes             enums.Cargo_classes
	Cargo_label               string
	// Town_growth_effect
	Town_growth_multiplier float32
	// Callback_flags
	Capacity_multiplier float32
	// Town_production_effect
	Town_production_multiplier float32
}
