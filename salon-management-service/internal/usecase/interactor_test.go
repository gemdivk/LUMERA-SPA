package usecase_test

import (
	"testing"

	cacheMocks "github.com/gemdivk/LUMERA-SPA/salon-management-service/internal/domain/cache/mocks"
	"github.com/gemdivk/LUMERA-SPA/salon-management-service/internal/domain/entity"
	repoMocks "github.com/gemdivk/LUMERA-SPA/salon-management-service/internal/domain/repository/mocks"
	"github.com/gemdivk/LUMERA-SPA/salon-management-service/internal/usecase"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestGetAllProcedures_CacheHit(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := repoMocks.NewMockSalonRepository(ctrl)
	mockCache := cacheMocks.NewMockSalonCache(ctrl)

	expected := []*entity.Procedure{
		{ID: "1", SalonID: "s1", Name: "Hot Stone", Duration: 30},
	}

	mockCache.EXPECT().GetProcedures().Return(expected, true)

	uc := usecase.NewSalonUsecase(mockRepo, mockCache)

	result, err := uc.GetAllProcedures()
	assert.NoError(t, err)
	assert.Equal(t, expected, result)
}

func TestGetAllProcedures_CacheMiss(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := repoMocks.NewMockSalonRepository(ctrl)
	mockCache := cacheMocks.NewMockSalonCache(ctrl)

	expected := []*entity.Procedure{
		{ID: "2", SalonID: "s2", Name: "Massage", Duration: 60},
	}

	mockCache.EXPECT().GetProcedures().Return(nil, false)
	mockRepo.EXPECT().GetAllProcedures().Return(expected, nil)
	mockCache.EXPECT().SetProcedures(expected)

	uc := usecase.NewSalonUsecase(mockRepo, mockCache)

	result, err := uc.GetAllProcedures()
	assert.NoError(t, err)
	assert.Equal(t, expected, result)
}

func TestGetAllSpecialists_CacheHit(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := repoMocks.NewMockSalonRepository(ctrl)
	mockCache := cacheMocks.NewMockSalonCache(ctrl)

	expected := []*entity.Specialist{
		{ID: "sp1", SalonID: "s1", Name: "Daulet"},
	}

	mockCache.EXPECT().GetSpecialists().Return(expected, true)

	uc := usecase.NewSalonUsecase(mockRepo, mockCache)

	result, err := uc.GetAllSpecialists()
	assert.NoError(t, err)
	assert.Equal(t, expected, result)
}

func TestGetAllSpecialists_CacheMiss(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := repoMocks.NewMockSalonRepository(ctrl)
	mockCache := cacheMocks.NewMockSalonCache(ctrl)

	expected := []*entity.Specialist{
		{ID: "sp2", SalonID: "s2", Name: "Kamila"},
	}

	mockCache.EXPECT().GetSpecialists().Return(nil, false)
	mockRepo.EXPECT().GetAllSpecialists().Return(expected, nil)
	mockCache.EXPECT().SetSpecialists(expected)

	uc := usecase.NewSalonUsecase(mockRepo, mockCache)

	result, err := uc.GetAllSpecialists()
	assert.NoError(t, err)
	assert.Equal(t, expected, result)
}
