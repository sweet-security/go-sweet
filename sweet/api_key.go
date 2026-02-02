package sweet

type ApiKey struct {
	ApiKey      string   `json:"apiKey"`
	Secret      string   `json:"secret"`
	Description string   `json:"description"`
	Roles       []string `json:"roles"`
}

func (s *ApiClient) CreateApiKey(description string, roles []string) (*ApiKey, error) {
	resp, err := s.restyClient.R().
		SetBody(ApiKey{
			Description: description,
			Roles:       roles,
		}).
		SetResult(&ApiKey{}).
		Post("/v1/auth/key")
	if err != nil {
		return nil, err
	}
	result := resp.Result().(*ApiKey)
	return result, nil
}
