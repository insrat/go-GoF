package demo

import (
	"reflect"
	"testing"
)

func TestGetHungryIns(t *testing.T) {
	ins := GetHungryIns()

	tests := []struct {
		name string
		want *singleton
	}{
		// TODO: Add test cases.
		{
			name: "First",
			want: ins,
		},
		{
			name: "Second",
			want: ins,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetHungryIns(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetHungryIns() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetLazyIns(t *testing.T) {
	ins := GetLazyIns()

	tests := []struct {
		name string
		want *singleton
	}{
		// TODO: Add test cases.
		{
			name: "First",
			want: ins,
		},
		{
			name: "Second",
			want: ins,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetLazyIns(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetLazyIns() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetDoubleIns(t *testing.T) {
	ins := GetDoubleIns()

	tests := []struct {
		name string
		want *singleton
	}{
		// TODO: Add test cases.
		{
			name: "First",
			want: ins,
		},
		{
			name: "Second",
			want: ins,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetDoubleIns(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetDoubleIns() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetOnceIns(t *testing.T) {
	ins := GetOnceIns()

	tests := []struct {
		name string
		want *singleton
	}{
		// TODO: Add test cases.
		{
			name: "First",
			want: ins,
		},
		{
			name: "Second",
			want: ins,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetOnceIns(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetOnceIns() = %v, want %v", got, tt.want)
			}
		})
	}
}
