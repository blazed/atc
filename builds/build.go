package builds

import "github.com/concourse/atc/config"

type Status string

const (
	StatusPending   Status = "pending"
	StatusStarted   Status = "started"
	StatusAborted   Status = "aborted"
	StatusSucceeded Status = "succeeded"
	StatusFailed    Status = "failed"
	StatusErrored   Status = "errored"
)

type Build struct {
	ID     int
	Status Status

	AbortURL string
}

type VersionedResources []VersionedResource

func (vrs VersionedResources) Lookup(resource string) (VersionedResource, bool) {
	for _, vr := range vrs {
		if vr.Name == resource {
			return vr, true
		}
	}

	return VersionedResource{}, false
}

type VersionedResource struct {
	Name     string
	Type     string
	Source   config.Source
	Version  Version
	Metadata []MetadataField
}

type Version map[string]interface{}

type MetadataField struct {
	Name  string
	Value string
}

type ByID []Build

func (bs ByID) Len() int           { return len(bs) }
func (bs ByID) Less(i, j int) bool { return bs[i].ID < bs[j].ID }
func (bs ByID) Swap(i, j int)      { bs[i], bs[j] = bs[j], bs[i] }
