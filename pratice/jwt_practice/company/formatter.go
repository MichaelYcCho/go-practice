package company


type CompanyFormatter struct {
	ID     int  `json:"id"`
	Name  string  `json:"name"`
	CEO string  `json:"ceo"`
	Revenue  string `json:"revenue"`
}

func FormatCompany(company Company) CompanyFormatter {
	companyFormatter := CompanyFormatter{}
	companyFormatter.ID = company.ID
	companyFormatter.Name = company.Name
	companyFormatter.CEO = company.CEO
	companyFormatter.Revenue = company.Revenue

	return companyFormatter

}