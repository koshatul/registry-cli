package registry

import (
	"fmt"

	"github.com/peterhellberg/link"
)

type tagsResponse struct {
	Tags []string `json:"tags"`
}

func (r *Registry) ListTags(repository string) ([]string, error) {
	tags := []string{}
	url := fmt.Sprintf("/v2/%s/tags/list", repository)
	for {
		resp, err := r.client.R().
			SetResult(tagsResponse{}).
			Get(url)
		if err != nil {
			return []string{}, err
		}
		linkHdr := link.ParseHeader(resp.Header())
		if val, ok := linkHdr["next"]; ok {
			tags = append(tags, resp.Result().(*tagsResponse).Tags...)
			url = val.URI
		} else {
			tags = append(tags, resp.Result().(*tagsResponse).Tags...)
			return tags, nil
		}
	}
}
