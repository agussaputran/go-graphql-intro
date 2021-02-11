package models

// RajaongkirProvince model
type RajaongkirProvince struct {
	RajaOngkir ProvinceResponse `json:"rajaongkir"`
}

// ProvinceResponse model
type ProvinceResponse struct {
	ProvinceResults []ProvinceResult `json:"results"`
}

// ProvinceResult model
type ProvinceResult struct {
	ProvinceID string `json:"province_id"`
	Province   string `json:"province"`
}

// RajaongkirCity model
type RajaongkirCity struct {
	RajaOngkir CityResponse `json:"rajaongkir"`
}

// CityResponse model
type CityResponse struct {
	CityResults []CityResult `json:"results"`
}

// CityResult model
type CityResult struct {
	CityID     string `json:"city_id"`
	ProvinceID string `json:"province_id"`
	Province   string `json:"province"`
	Type       string `json:"type"`
	CityName   string `json:"city_name"`
	PostalCode string `json:"postal_code"`
}
