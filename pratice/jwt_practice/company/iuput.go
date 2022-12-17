package company

type InputCompany struct {
	Name        string `json:"name" binding:"required"`
	CEO         string `json:"ceo" binding:"required"`
	Revenue     string `json:"revenue" binding:"required"`
}

