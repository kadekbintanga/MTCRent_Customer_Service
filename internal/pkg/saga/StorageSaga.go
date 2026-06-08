package saga

import (
	xtremepkg "github.com/globalxtreme/go-core/v2/pkg"
	gxstorage "github.com/globalxtreme/go-storage/v2"
)

type StorageSaga struct {
	DeleteStoragePaths []string
	StoragePaths       []string
}

/** --- PUBLIC STORAGE CLIENT --- */

func (saga *StorageSaga) Delete(paths []string) {
	for _, path := range paths {
		_, err := gxstorage.Delete(path)
		if err != nil {
			xtremepkg.LogError(err, false)
		}
	}

}

/** --- DEFER FUNCTION --- */

func (saga *StorageSaga) Close() {
	if r := recover(); r != nil {
		if gxstorage.PublicStorageRPCActive && len(saga.StoragePaths) > 0 {
			saga.Delete(saga.StoragePaths)
		}
		panic(r)
	}

	if gxstorage.PublicStorageRPCActive && len(saga.DeleteStoragePaths) > 0 {
		saga.Delete(saga.DeleteStoragePaths)
	}
}
