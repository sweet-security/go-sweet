package sweet

type AwsAccount struct {
	AccountId  string   `json:"accountId"`
	RoleArn    string   `json:"roleArn"`
	ExternalId string   `json:"externalId,omitempty"`
	Regions    []string `json:"regions,omitempty"`
}

func (s *ApiClient) GetAwsAccounts() (*[]AwsAccount, error) {
	resp, err := s.restyClient.R().
		SetResult(&[]AwsAccount{}).
		Get("/v1/aws/account")
	if err != nil {
		return nil, err
	}
	result := resp.Result().(*[]AwsAccount)
	return result, nil
}

func (s *ApiClient) GetAwsAccount(accountId string) (*AwsAccount, error) {
	resp, err := s.restyClient.R().
		SetResult(&AwsAccount{}).
		Get("/v1/aws/account/" + accountId)
	if err != nil {
		return nil, err
	}
	result := resp.Result().(*AwsAccount)
	return result, nil
}

func (s *ApiClient) AddAwsAccount(awsAccount *AwsAccount) (*AwsAccount, error) {
	resp, err := s.restyClient.R().
		SetBody(awsAccount).
		SetResult(&AwsAccount{}).
		Post("/v1/aws/account")
	if err != nil {
		return nil, err
	}
	result := resp.Result().(*AwsAccount)
	return result, nil
}

func (s *ApiClient) UpdateAwsAccount(awsAccount *AwsAccount) (*AwsAccount, error) {
	resp, err := s.restyClient.R().
		SetBody(awsAccount).
		SetResult(&AwsAccount{}).
		Put("/v1/aws/account/" + awsAccount.AccountId)
	if err != nil {
		return nil, err
	}
	result := resp.Result().(*AwsAccount)
	return result, nil
}

func (s *ApiClient) DeleteAwsAccount(accountId string) error {
	_, err := s.restyClient.R().
		Delete("/v1/aws/account/" + accountId)
	if err != nil {
		return err
	}
	return nil
}
