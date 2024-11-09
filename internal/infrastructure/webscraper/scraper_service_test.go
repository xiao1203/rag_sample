package webscraper

import (
	"reflect"
	"testing"
)

func TestScraperService_ExtractText(t *testing.T) {
	service := NewScraper()

	type args struct {
		url string
	}
	tests := []struct {
		name    string
		args    args
		want    *[]string
		wantErr bool
	}{
		{
			name: "正常系",
			args: args{
				url: "https://ja.wikipedia.org/wiki/%E3%83%99%E3%82%BF",
			},
			want:    &[]string{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := service.ExtractText(tt.args.url)
			if (err != nil) != tt.wantErr {
				t.Errorf("ScraperService.ExtractText() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ScraperService.ExtractText() = %v, want %v", got, tt.want)
			}
		})
	}
}
