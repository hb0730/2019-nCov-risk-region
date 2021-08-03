package source

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestBenDiBao_Time(t *testing.T) {

	tests := []struct {
		name    string
		fields  *BenDiBao
		want    time.Time
		wantErr bool
	}{
		// TODO: Add test cases.
		{"获取截至时间", NewBenDiBao(true), time.Date(2021, 8, 3, 10, 05, 07, 0, time.Local), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := BenDiBao{
				browser: tt.fields.browser,
			}
			got, err := b.Time()
			if (err != nil) != tt.wantErr {
				t.Errorf("Time() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Time() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBenDiBao_HighRisk(t *testing.T) {

	tests := []struct {
		name   string
		fields *BenDiBao
		want   []Risk
	}{
		// TODO: Add test cases.
		{"高风险地区", NewBenDiBao(true), nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &BenDiBao{
				browser: tt.fields.browser,
			}
			risks := b.HighRisk()
			fmt.Printf("%v\n", risks)
		})
	}
}

func TestBenDiBao_MiddleRisk(t *testing.T) {

	tests := []struct {
		name   string
		fields *BenDiBao
		want   []Risk
	}{
		// TODO: Add test cases.
		{"中风险地区", NewBenDiBao(false), nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &BenDiBao{
				browser: tt.fields.browser,
			}
			risks := b.MiddleRisk()
			fmt.Printf("%v\n", risks)
		})
	}
}
