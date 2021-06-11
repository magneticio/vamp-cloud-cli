package ingresses

import (
	dto "github.com/magneticio/vamp-cloud-cli/models"
	"testing"
)

func Test_checkPreviewRoute(t *testing.T) {
	type args struct {
		routes []*dto.Route
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "ErrorCases",
			args: args{
				routes: []*dto.Route{
					{
						Path: "some-path/%%%%",
					},
					{
						Path: "some-path/%%VRSION%%",
					},
					{
						Path: "some-path/test-%%VRSION%%",
					},
					{
						Path: "some-path/test-%%VRSION%%-test",
					},
					{
						Path: "some-path/test-%%VERSION%%-test%%DEBUG%%",
					},
					{
						Path: "some-path/%%%VERSION%%%",
					},
					{
						Path: "some-path/%%%VERSION%%",
					},
					{
						Path: "some-path-%%VERSION%%/another-path-%%DEBUG%%/test",
					},
				},
			},
			wantErr: true,
		},
		{
			name: "NoErrorCases",
			args: args{
				routes: []*dto.Route{
					{
						Path: "some-path/VERSION",
					},
					{
						Path: "some-path/%%VERSION%%",
					},
					{
						Path: "some-path/test-%%VERSION%%",
					},
					{
						Path: "some-path/%%VERSION%%-test",
					},
					{
						Path: "some-path/test-%%VERSION%%-test",
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := checkPreviewRoute(tt.args.routes); (err != nil) != tt.wantErr {
				t.Errorf("checkPreviewRoute() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
