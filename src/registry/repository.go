package registry

import (
	"github.com/peterhellberg/link"
)

type repositoriesResponse struct {
	Repositories []string `json:"repositories"`
}

func (r *Registry) ListRepositories() ([]string, error) {
	repos := []string{}
	url := "/v2/_catalog"
	for {
		resp, err := r.client.R().
			SetResult(repositoriesResponse{}).
			Get(url)
		if err != nil {
			return []string{}, err
		}
		linkHdr := link.ParseHeader(resp.Header())
		if val, ok := linkHdr["next"]; ok {
			repos = append(repos, resp.Result().(*repositoriesResponse).Repositories...)
			url = val.URI
		} else {
			repos = append(repos, resp.Result().(*repositoriesResponse).Repositories...)
			return repos, nil
		}
		// switch err {
		// case ErrNoMorePages:
		// 	repos = append(repos, resp.Result().(*repositoriesResponse).Repositories...)
		// 	return repos, nil
		// case nil:
		// 	repos = append(repos, resp.Result().(*repositoriesResponse).Repositories...)
		// 	continue
		// default:
		// 	return nil, err
		// }
	}

}
