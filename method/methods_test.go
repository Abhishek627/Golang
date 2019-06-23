package main

import "testing"

func Test_circle_Area(t *testing.T) {
	type fields struct {
		radius float32
	}
	tests := []struct {
		name   string
		fields fields
		want   float32
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := circle{
				radius: tt.fields.radius,
			}
			if got := c.Area(); got != tt.want {
				t.Errorf("circle.Area() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_rectangle_Area(t *testing.T) {
	type fields struct {
		length float32
		breath float32
	}
	tests := []struct {
		name   string
		fields fields
		want   float32
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := rectangle{
				length: tt.fields.length,
				breath: tt.fields.breath,
			}
			if got := r.Area(); got != tt.want {
				t.Errorf("rectangle.Area() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_circle_changeRadius(t *testing.T) {
	type fields struct {
		radius float32
	}
	type args struct {
		newrad float32
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &circle{
				radius: tt.fields.radius,
			}
			c.changeRadius(tt.args.newrad)
		})
	}
}

func Test_mystring_toUpperCase(t *testing.T) {
	tests := []struct {
		name string
		m    mystring
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.toUpperCase(); got != tt.want {
				t.Errorf("mystring.toUpperCase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}
