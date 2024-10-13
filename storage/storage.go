package storage

import (
	"sync"

	"github.com/kv-base-hack/base_crawler/common"
)

type ChainData struct {
	latestBlock     int64
	lastStoredBlock int64
}

type Storage struct {
	chains map[common.Chain]*ChainData
	mutex  sync.RWMutex
}

func NewStorage() *Storage {
	return &Storage{
		chains: map[common.Chain]*ChainData{
			common.ChainBase: {},
		},
	}
}

func (s *Storage) SetLatestBlock(chain common.Chain, latestBlock int64) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.chains[chain].latestBlock = latestBlock
}

func (s *Storage) GetLatestBlock(chain common.Chain) int64 {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	return s.chains[chain].latestBlock
}

func (s *Storage) SetLastStoredBlock(chain common.Chain, lastStoredBlock int64) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.chains[chain].lastStoredBlock = lastStoredBlock
}

func (s *Storage) GetLastStoredBlock(chain common.Chain) int64 {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	return s.chains[chain].lastStoredBlock
}
