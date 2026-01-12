package sweet

type AwsOrganization struct {
	AccountId            string   `json:"accountId"`
	RoleArn              string   `json:"roleArn"`
	RoleNameParameterArn string   `json:"roleNameParameterArn"`
	ExternalId           string   `json:"externalId,omitempty"`
	Regions              []string `json:"regions,omitempty"`
}

func (s *ApiClient) GetAwsOrganizations() (*[]AwsOrganization, error) {
	resp, err := s.restyClient.R().
		SetResult(&[]AwsOrganization{}).
		Get("/v1/aws/organization")
	if err != nil {
		return nil, err
	}
	result := resp.Result().(*[]AwsOrganization)
	return result, nil
}

func (s *ApiClient) GetAwsOrganization(accountId string) (*AwsOrganization, error) {
	resp, err := s.restyClient.R().
		SetResult(&AwsOrganization{}).
		Get("/v1/aws/organization/" + accountId)
	if err != nil {
		return nil, err
	}
	result := resp.Result().(*AwsOrganization)
	return result, nil
}

func (s *ApiClient) AddAwsOrganization(awsOrganization *AwsOrganization) (*AwsOrganization, error) {
	resp, err := s.restyClient.R().
		SetBody(awsOrganization).
		SetResult(&AwsOrganization{}).
		Post("/v1/aws/organization")
	if err != nil {
		return nil, err
	}
	result := resp.Result().(*AwsOrganization)
	return result, nil
}

func (s *ApiClient) UpdateAwsOrganization(awsOrganization *AwsOrganization) (*AwsOrganization, error) {
	resp, err := s.restyClient.R().
		SetBody(awsOrganization).
		SetResult(&AwsOrganization{}).
		Put("/v1/aws/organization/" + awsOrganization.AccountId)
	if err != nil {
		return nil, err
	}
	result := resp.Result().(*AwsOrganization)
	return result, nil
}

func (s *ApiClient) DeleteAwsOrganization(accountId string) error {
	_, err := s.restyClient.R().
		Delete("/v1/aws/organization/" + accountId)
	if err != nil {
		return err
	}
	return nil
}
