package pkg

import (
	"log"
)

func UpdateCache() error {
	var err error

	newBandInfo, newRelationInfo, err := GetBandInfo(artistAPI, relationAPI)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	bandInfoMu.Lock()
	defer bandInfoMu.Unlock()

	BandInfo = newBandInfo
	RelationInfo = newRelationInfo

	return nil
}
