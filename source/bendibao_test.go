package source

import (
	"fmt"
	"testing"
)

func TestBenDiBao_Time(t *testing.T) {

	bendibao := NewBenDiBao(false)
	fmt.Println(bendibao.Time())
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
