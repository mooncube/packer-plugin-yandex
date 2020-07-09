package yandexexport

import (
	"fmt"
)

const BuilderId = "packer.post-processor.yandex-export"

type Artifact struct {
	paths []string
	urls  []string
}

func (*Artifact) BuilderId() string {
	return BuilderId
}

func (*Artifact) Id() string {
	return ""
}

func (a *Artifact) Files() []string {
	pathsCopy := make([]string, len(a.paths))
	copy(pathsCopy, a.paths)
	return pathsCopy
}

func (a *Artifact) String() string {
	return fmt.Sprintf("Exported artifacts in: %s", a.paths)
}

func (a *Artifact) Url() string {
	// print url for first path
	return a.urls[0]
}

func (*Artifact) State(name string) interface{} {
	return nil
}

func (a *Artifact) Destroy() error {
	return nil
}
