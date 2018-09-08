package user

import (
	"testing"

	"github.com/faruqisan/testable-repo/modules/user/mock_user"
	"github.com/golang/mock/gomock"
)

func TestModule_Create(t *testing.T) {
	type args struct {
		name string
		age  int
	}
	type mock struct {
		mockData Data
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		mock    mock
	}{
		{
			name: "Test Success",
			args: args{
				name: "a",
				age:  1,
			},
			mock : mock {
				mockData : Data{"a",1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockNSQ := mock_user.NewMockNSQ(ctrl)
			mockRedis := mock_user.NewMockRedis(ctrl)

			m := &Module{
				nsq:   mockNSQ,
				redis: mockRedis,
			}

			if tt.wantErr {

			} else {
				mockRedis.EXPECT().Set("key", tt.mock.mockData)
				mockNSQ.EXPECT().Publish("some data").Return(nil)
			}

			if err := m.Create(tt.args.name, tt.args.age); (err != nil) != tt.wantErr {
				t.Errorf("Module.Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
