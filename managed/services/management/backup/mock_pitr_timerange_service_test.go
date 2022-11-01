// Code generated by mockery v1.0.0. DO NOT EDIT.

package backup

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	models "github.com/percona/pmm/managed/models"
	"github.com/percona/pmm/managed/services/backup"
)

// mockPitrTimerangeService is an autogenerated mock type for the pitrTimerangeService type
type mockPitrTimerangeService struct {
	mock.Mock
}

// ListPITRTimeranges provides a mock function with given fields: ctx, artifactName, location
func (_m *mockPitrTimerangeService) ListPITRTimeranges(ctx context.Context, artifactName string, location *models.BackupLocation) ([]backup.Timeline, error) {
	ret := _m.Called(ctx, artifactName, location)

	var r0 []backup.Timeline
	if rf, ok := ret.Get(0).(func(context.Context, string, *models.BackupLocation) []backup.Timeline); ok {
		r0 = rf(ctx, artifactName, location)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]backup.Timeline)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, *models.BackupLocation) error); ok {
		r1 = rf(ctx, artifactName, location)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}