package registry

type repositoriesResponse struct {
	Repositories []string `json:"repositories"`
}

func (r *Registry) ListRepositories() ([]string, error) {
	repos := []string{}
	link := "/v2/_catalog"
	for {
		resp, err := r.client.R().
			SetResult(repositoriesResponse{}).
			Get(link)
		if err != nil {
			return []string{}, err
		}
		link, err = getNextLink(resp)
		switch err {
		case ErrNoMorePages:
			repos = append(repos, resp.Result().(*repositoriesResponse).Repositories...)
			return repos, nil
		case nil:
			repos = append(repos, resp.Result().(*repositoriesResponse).Repositories...)
			continue
		default:
			return nil, err
		}
	}
}
