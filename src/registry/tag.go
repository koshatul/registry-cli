package registry

import "fmt"

type tagsResponse struct {
	Tags []string `json:"tags"`
}

func (r *Registry) ListTags(repository string) ([]string, error) {
	tags := []string{}
	link := fmt.Sprintf("/v2/%s/tags/list", repository)
	for {
		resp, err := r.client.R().
			SetResult(tagsResponse{}).
			Get(link)
		if err != nil {
			return []string{}, err
		}
		link, err = getNextLink(resp)
		switch err {
		case ErrNoMorePages:
			tags = append(tags, resp.Result().(*tagsResponse).Tags...)
			return tags, nil
		case nil:
			tags = append(tags, resp.Result().(*tagsResponse).Tags...)
			continue
		default:
			return nil, err
		}
	}
}
