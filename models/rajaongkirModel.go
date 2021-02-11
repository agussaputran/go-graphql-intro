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
