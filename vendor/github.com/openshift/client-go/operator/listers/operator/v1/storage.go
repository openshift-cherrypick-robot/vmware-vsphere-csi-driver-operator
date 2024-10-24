// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/openshift/api/operator/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/listers"
	"k8s.io/client-go/tools/cache"
)

// StorageLister helps list Storages.
// All objects returned here must be treated as read-only.
type StorageLister interface {
	// List lists all Storages in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.Storage, err error)
	// Get retrieves the Storage from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.Storage, error)
	StorageListerExpansion
}

// storageLister implements the StorageLister interface.
type storageLister struct {
	listers.ResourceIndexer[*v1.Storage]
}

// NewStorageLister returns a new StorageLister.
func NewStorageLister(indexer cache.Indexer) StorageLister {
	return &storageLister{listers.New[*v1.Storage](indexer, v1.Resource("storage"))}
}
