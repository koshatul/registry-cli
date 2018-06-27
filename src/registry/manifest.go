package registry

import (
	"fmt"

	manifestV1 "github.com/docker/distribution/manifest/schema1"
	manifestV2 "github.com/docker/distribution/manifest/schema2"
	digest "github.com/opencontainers/go-digest"
)

func (r *Registry) ManifestV1(repository, reference string) (*manifestV1.SignedManifest, error) {
	resp, err := r.client.R().
		SetHeader("Accept", manifestV1.MediaTypeManifest).
		SetResult(manifestV1.SignedManifest{}).
		Get(fmt.Sprintf("/v2/%s/manifests/%s", repository, reference))
	if err != nil {
		return nil, err
	}

	return resp.Result().(*manifestV1.SignedManifest), nil
}

func (r *Registry) ManifestV2(repository, reference string) (*manifestV2.DeserializedManifest, error) {
	resp, err := r.client.R().
		SetHeader("Accept", manifestV2.MediaTypeManifest).
		SetResult(manifestV2.DeserializedManifest{}).
		Get(fmt.Sprintf("/v2/%s/manifests/%s", repository, reference))
	if err != nil {
		return nil, err
	}

	return resp.Result().(*manifestV2.DeserializedManifest), nil
}

func (r *Registry) ManifestDigestV1(repository, reference string) (digest.Digest, error) {
	resp, err := r.client.R().
		SetHeader("Accept", manifestV1.MediaTypeManifest).
		Head(fmt.Sprintf("/v2/%s/manifests/%s", repository, reference))
	if err != nil {
		return "", err
	}

	return digest.Parse(resp.Header().Get("Docker-Content-Digest"))
}

func (r *Registry) ManifestDigestV2(repository, reference string) (digest.Digest, error) {
	resp, err := r.client.R().
		SetHeader("Accept", manifestV2.MediaTypeManifest).
		Head(fmt.Sprintf("/v2/%s/manifests/%s", repository, reference))
	if err != nil {
		return "", err
	}

	return digest.Parse(resp.Header().Get("Docker-Content-Digest"))
}
