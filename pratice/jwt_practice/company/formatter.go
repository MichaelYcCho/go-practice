package company


type CompanyFormatter struct {
	ID     string  `json:"id"`
	Name  string  `json:"name"`
	CEO string  `json:"ceo"`
	Revenue  string `json:"revenue"`
}

func FormatCompany(campaign Company) CompanyFormatter {
	companyFormatter := CompanyFormatter{}
	companyFormatter.ID = campaign.ID
	companyFormatter.Name = campaign.Name
	companyFormatter.CEO = campaign.CEO
	companyFormatter.Revenue = campaign.Revenue

	return companyFormatter

}