package repository

import (
	"matchlove-services/internal/dto"
	"matchlove-services/internal/model"
)

func makeConfig(cfg *matchMakingEngineConfig) *matchMakingEngineConfig {
	if cfg.existUser == nil || len(cfg.existUser) == 0 {
		cfg.existUser = make(map[string]bool)
	}

	if cfg.dto.Page < 0 {
		panic("page must be greater than or equal to 0")
	} else if cfg.dto.PerPage < 1 {
		panic("page must be greater than or equal to 1")
	}

	return cfg
}

type matchMakingEngineConfig struct {
	preference *model.UserPreference
	profile    *model.UserProfile
	dto        *dto.MatchSuggestionsRequestDto

	disableDistance   bool
	disableAge        bool
	disableLookingFor bool
	disableInterest   bool

	existUser map[string]bool
}

func (m *matchMakingEngineConfig) Limit() int {
	return m.dto.PerPage
}

func (m *matchMakingEngineConfig) Offset() int {
	return (m.dto.Page - 1) * m.dto.PerPage
}

func (m *matchMakingEngineConfig) ExpandDistance() {
	m.preference.Distance = m.preference.Distance + 10
}

func (m *matchMakingEngineConfig) ExpandAgeMinAndMax() {
	m.preference.AgeMin = uint8(m.profile.Age - 5)
	m.preference.AgeMax = m.preference.AgeMax + 10
}

func (m *matchMakingEngineConfig) NotIncludeLookingFor() {
	m.disableLookingFor = true
}

func (m *matchMakingEngineConfig) NotIncludeInterest() {
	m.disableInterest = true
}

func (m *matchMakingEngineConfig) DisablePreference() {
	m.disableAge = true
	m.disableDistance = true
	m.disableLookingFor = true
	m.disableInterest = true
}

func (m *matchMakingEngineConfig) SameData() ([]string, bool) {
	l := len(m.existUser)
	if l == 0 {
		return nil, false
	}

	beforeID := make([]string, 0)
	for key, _ := range m.existUser {
		beforeID = append(beforeID, key)
	}

	return beforeID, true
}
