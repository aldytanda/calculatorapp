package calculator

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCalculator(t *testing.T) {
	type args struct {
		initval float64
	}
	tests := []struct {
		name string
		args args
		want *Calculator
	}{
		{
			name: "initval 0",
			args: args{
				initval: 0,
			},
			want: &Calculator{
				Number: 0,
				Ops:    [][]string{},
			},
		},
		{
			name: "initval 1",
			args: args{
				initval: 1,
			},
			want: &Calculator{
				Number: 1,
				Ops:    [][]string{},
			},
		},
		{
			name: "initval 0.5",
			args: args{
				initval: 0.5,
			},
			want: &Calculator{
				Number: 0.5,
				Ops:    [][]string{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewCalculator(tt.args.initval)
			assert.Equal(t, len(got.Ops), len(tt.want.Ops))
			assert.Equal(t, got.Number, tt.want.Number)
		})
	}
}

func TestCalculator_Exec(t *testing.T) {
	type fields struct {
		Number float64
		Ops    [][]string
	}
	type args struct {
		input []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    float64
		wantErr bool
		wantOps [][]string
	}{
		{
			name: "Empty operation should return error",
			fields: fields{
				Number: 0,
				Ops:    [][]string{},
			},
			args: args{
				input: []string{},
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "Unknown operation TEST should return error",
			fields: fields{
				Number: 0,
				Ops:    [][]string{},
			},
			args: args{
				input: []string{"TEST"},
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "Simple Add should return correct result",
			fields: fields{
				Number: 0,
				Ops:    [][]string{},
			},
			args: args{
				input: []string{"add", "1"},
			},
			want:    1,
			wantErr: false,
			wantOps: [][]string{
				{"add", "1"},
			},
		},
		{
			name: "Simple Add with non zero current number should return correct result",
			fields: fields{
				Number: 1,
				Ops:    [][]string{},
			},
			args: args{
				input: []string{"add", "1"},
			},
			want:    2,
			wantErr: false,
			wantOps: [][]string{
				{"add", "1"},
			},
		},
		{
			name: "Simple Add with negative current number should return correct result",
			fields: fields{
				Number: -1,
				Ops:    [][]string{},
			},
			args: args{
				input: []string{"add", "2"},
			},
			want:    1,
			wantErr: false,
			wantOps: [][]string{
				{"add", "2"},
			},
		},
		{
			name: "Simple Add with existing ops should return correct result",
			fields: fields{
				Number: -1,
				Ops:    [][]string{{"subtract", "1"}},
			},
			args: args{
				input: []string{"add", "2"},
			},
			want:    1,
			wantErr: false,
			wantOps: [][]string{
				{"subtract", "1"},
				{"add", "2"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Calculator{
				Number: tt.fields.Number,
				Ops:    tt.fields.Ops,
			}
			got, err := c.Exec(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("Calculator.Exec() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Calculator.Exec() = %v, want %v", got, tt.want)
			}

			assert.Equal(t, len(tt.wantOps), len(c.Ops))
			for i, v := range c.Ops {
				assert.Equal(t, tt.wantOps[i], v)
			}
		})
	}
}

func TestCalculator_exec(t *testing.T) {
	type fields struct {
		Number float64
		Ops    [][]string
	}
	type args struct {
		input []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    float64
		wantErr bool
	}{
		{
			name: "operation add with < 1 argument should return error",
			fields: fields{
				Number: 0,
				Ops:    [][]string{},
			},
			args: args{
				input: []string{"add"},
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "operation subtract with < 1 argument should return error",
			fields: fields{
				Number: 0,
				Ops:    [][]string{},
			},
			args: args{
				input: []string{"subtract"},
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "operation multiply with < 1 argument should return error",
			fields: fields{
				Number: 0,
				Ops:    [][]string{},
			},
			args: args{
				input: []string{"multiply"},
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "operation divide with non numeric argument should return error",
			fields: fields{
				Number: 0,
				Ops:    [][]string{},
			},
			args: args{
				input: []string{"divide", "one"},
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "operation repeat with non numeric argument should return error",
			fields: fields{
				Number: 0,
				Ops:    [][]string{},
			},
			args: args{
				input: []string{"repeat", "one"},
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "Unknown operation should return error",
			fields: fields{
				Number: 0,
				Ops:    [][]string{},
			},
			args: args{
				input: []string{"unknown", "testing"},
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Calculator{
				Number: tt.fields.Number,
				Ops:    tt.fields.Ops,
			}
			got, err := c.exec(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("Calculator.exec() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Calculator.exec() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCalculator_add(t *testing.T) {
	type fields struct {
		Number float64
		Ops    [][]string
	}
	type args struct {
		n float64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantNum float64
	}{
		{
			name: "valid add",
			fields: fields{
				Number: 0,
				Ops:    [][]string{},
			},
			args: args{
				n: 1,
			},
			wantNum: 1,
		},
		{
			name: "valid add with neg input",
			fields: fields{
				Number: 0,
				Ops:    [][]string{},
			},
			args: args{
				n: -1,
			},
			wantNum: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Calculator{
				Number: tt.fields.Number,
				Ops:    tt.fields.Ops,
			}
			c.add(tt.args.n)
			assert.Equal(t, tt.wantNum, c.Number)
		})
	}
}

func TestCalculator_subtract(t *testing.T) {
	type fields struct {
		Number float64
		Ops    [][]string
	}
	type args struct {
		n float64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantNum float64
	}{
		{
			name: "valid subtract",
			fields: fields{
				Number: 0,
				Ops:    [][]string{},
			},
			args: args{
				n: 1,
			},
			wantNum: -1,
		},
		{
			name: "valid subtract with neg input",
			fields: fields{
				Number: 0,
				Ops:    [][]string{},
			},
			args: args{
				n: -1,
			},
			wantNum: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Calculator{
				Number: tt.fields.Number,
				Ops:    tt.fields.Ops,
			}
			c.subtract(tt.args.n)
		})
	}
}

func TestCalculator_multiply(t *testing.T) {
	type fields struct {
		Number float64
		Ops    [][]string
	}
	type args struct {
		n float64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantNum float64
	}{
		{
			name: "valid multiply",
			fields: fields{
				Number: 0,
				Ops:    [][]string{},
			},
			args: args{
				n: 1,
			},
			wantNum: 0,
		},
		{
			name: "valid multiply with neg input",
			fields: fields{
				Number: 5,
				Ops:    [][]string{},
			},
			args: args{
				n: -1,
			},
			wantNum: -5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Calculator{
				Number: tt.fields.Number,
				Ops:    tt.fields.Ops,
			}
			c.multiply(tt.args.n)
		})
	}
}

func TestCalculator_divide(t *testing.T) {
	type fields struct {
		Number float64
		Ops    [][]string
	}
	type args struct {
		n float64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantNum float64
		wantErr bool
	}{
		{
			name: "valid add",
			fields: fields{
				Number: 0,
				Ops:    [][]string{},
			},
			args: args{
				n: 1,
			},
			wantNum: 0,
			wantErr: false,
		},
		{
			name: "valid add with neg input",
			fields: fields{
				Number: 0,
				Ops:    [][]string{},
			},
			args: args{
				n: -1,
			},
			wantNum: 0,
			wantErr: false,
		},
		{
			name: "invalid multiply input 0",
			fields: fields{
				Number: 0,
				Ops:    [][]string{},
			},
			args: args{
				n: 0,
			},
			wantNum: 0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Calculator{
				Number: tt.fields.Number,
				Ops:    tt.fields.Ops,
			}
			err := c.divide(tt.args.n)
			assert.Equal(t, err != nil, tt.wantErr)

			if !tt.wantErr {
				assert.Equal(t, tt.wantNum, c.Number)
			}
		})
	}
}

func TestCalculator_cancel(t *testing.T) {
	type fields struct {
		Number float64
		Ops    [][]string
	}
	tests := []struct {
		name       string
		fields     fields
		wantNumber float64
		wantOps    [][]string
	}{
		{
			name: "empty ops and number",
			fields: fields{
				Number: 0,
				Ops:    [][]string{},
			},
			wantNumber: 0,
			wantOps:    [][]string{},
		},
		{
			name: "non empty ops and number",
			fields: fields{
				Number: 0,
				Ops:    [][]string{{"add", "0"}},
			},
			wantNumber: 0,
			wantOps:    [][]string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Calculator{
				Number: tt.fields.Number,
				Ops:    tt.fields.Ops,
			}
			c.cancel()
			assert.Equal(t, tt.wantNumber, c.Number)
			assert.Equal(t, len(tt.wantOps), len(c.Ops))
		})
	}
}

func TestCalculator_abs(t *testing.T) {
	type fields struct {
		Number float64
		Ops    [][]string
	}
	tests := []struct {
		name       string
		fields     fields
		wantNumber float64
	}{
		{
			name: "abs zero",
			fields: fields{
				Number: 0,
			},
			wantNumber: 0,
		},
		{
			name: "abs positive val",
			fields: fields{
				Number: 1,
			},
			wantNumber: 1,
		},
		{
			name: "abs negative val",
			fields: fields{
				Number: -1,
			},
			wantNumber: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Calculator{
				Number: tt.fields.Number,
				Ops:    tt.fields.Ops,
			}
			c.abs()
			assert.Equal(t, tt.wantNumber, c.Number)
		})
	}
}

func TestCalculator_neg(t *testing.T) {
	type fields struct {
		Number float64
		Ops    [][]string
	}
	tests := []struct {
		name       string
		fields     fields
		wantNumber float64
	}{
		{
			name: "neg zero",
			fields: fields{
				Number: 0,
			},
			wantNumber: 0,
		},
		{
			name: "neg positive val",
			fields: fields{
				Number: 1,
			},
			wantNumber: -1,
		},
		{
			name: "neg negative val",
			fields: fields{
				Number: -1,
			},
			wantNumber: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Calculator{
				Number: tt.fields.Number,
				Ops:    tt.fields.Ops,
			}
			c.neg()
			assert.Equal(t, tt.wantNumber, c.Number)
		})
	}
}

func TestCalculator_sqrt(t *testing.T) {
	type fields struct {
		Number float64
		Ops    [][]string
	}

	tests := []struct {
		name       string
		fields     fields
		wantNumber float64
		wantNan    bool
	}{
		{
			name: "sqrt zero",
			fields: fields{
				Number: 0,
			},
			wantNumber: 0,
		},
		{
			name: "sqrt positive val",
			fields: fields{
				Number: 1,
			},
			wantNumber: 1,
		},
		{
			name: "sqrt positive val 2",
			fields: fields{
				Number: 4,
			},
			wantNumber: 2,
		},
		{
			name: "sqrt negative val",
			fields: fields{
				Number: -1,
			},
			wantNumber: math.NaN(),
			wantNan:    true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Calculator{
				Number: tt.fields.Number,
				Ops:    tt.fields.Ops,
			}
			c.sqrt()
			if tt.wantNan {
				assert.True(t, math.IsNaN(c.Number))
			} else {
				assert.Equal(t, tt.wantNumber, c.Number)
			}
		})
	}
}

func TestCalculator_sqr(t *testing.T) {
	type fields struct {
		Number float64
		Ops    [][]string
	}
	tests := []struct {
		name       string
		fields     fields
		wantNumber float64
	}{
		{
			name: "sqrt zero",
			fields: fields{
				Number: 0,
			},
			wantNumber: 0,
		},
		{
			name: "sqrt positive val",
			fields: fields{
				Number: 1,
			},
			wantNumber: 1,
		},
		{
			name: "sqrt positive val 2",
			fields: fields{
				Number: 2,
			},
			wantNumber: 4,
		},
		{
			name: "sqrt negative val",
			fields: fields{
				Number: -1,
			},
			wantNumber: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Calculator{
				Number: tt.fields.Number,
				Ops:    tt.fields.Ops,
			}
			c.sqr()
			assert.Equal(t, tt.wantNumber, c.Number)
		})
	}
}

func TestCalculator_cubert(t *testing.T) {
	type fields struct {
		Number float64
		Ops    [][]string
	}
	tests := []struct {
		name       string
		fields     fields
		wantNumber float64
	}{
		{
			name: "cubert zero",
			fields: fields{
				Number: 0,
			},
			wantNumber: 0,
		},
		{
			name: "cubert positive val",
			fields: fields{
				Number: 1,
			},
			wantNumber: 1,
		},
		{
			name: "cubert positive val 2",
			fields: fields{
				Number: 8,
			},
			wantNumber: 2,
		},
		{
			name: "cubert negative val",
			fields: fields{
				Number: -9,
			},
			wantNumber: -2.080083823051904,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Calculator{
				Number: tt.fields.Number,
				Ops:    tt.fields.Ops,
			}
			c.cubert()
			if math.IsNaN(tt.wantNumber) {
				assert.True(t, math.IsNaN(c.Number))
			} else {
				assert.Equal(t, tt.wantNumber, c.Number)
			}
		})
	}
}

func TestCalculator_cube(t *testing.T) {
	type fields struct {
		Number float64
		Ops    [][]string
	}
	tests := []struct {
		name       string
		fields     fields
		wantNumber float64
	}{
		{
			name: "sqrt zero",
			fields: fields{
				Number: 0,
			},
			wantNumber: 0,
		},
		{
			name: "sqrt positive val",
			fields: fields{
				Number: 1,
			},
			wantNumber: 1,
		},
		{
			name: "sqrt positive val 2",
			fields: fields{
				Number: 2,
			},
			wantNumber: 8,
		},
		{
			name: "sqrt negative val",
			fields: fields{
				Number: -1,
			},
			wantNumber: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Calculator{
				Number: tt.fields.Number,
				Ops:    tt.fields.Ops,
			}
			c.cube()
			assert.Equal(t, tt.wantNumber, c.Number)
		})
	}
}

func TestCalculator_repeat(t *testing.T) {
	type fields struct {
		Number float64
		Ops    [][]string
	}
	type args struct {
		n float64
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantErr    bool
		wantNumber float64
	}{
		{
			name: "Valid repeat",
			fields: fields{
				Number: 1,
				Ops:    [][]string{{"Add", "1"}},
			},
			args: args{
				n: 1,
			},
			wantErr:    false,
			wantNumber: 2,
		},
		{
			name: "Valid repeat 2",
			fields: fields{
				Number: 4,
				Ops:    [][]string{{"Add", "1"}, {"Multiply", "4"}},
			},
			args: args{
				n: 2,
			},
			wantErr:    false,
			wantNumber: 20,
		},
		{
			name: "Invalid repeat n < length of current ops",
			fields: fields{
				Number: 1,
				Ops:    [][]string{{"Add", "1"}},
			},
			args: args{
				n: 2,
			},
			wantErr:    true,
			wantNumber: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Calculator{
				Number: tt.fields.Number,
				Ops:    tt.fields.Ops,
			}
			err := c.repeat(tt.args.n)
			assert.Equal(t, (err != nil), tt.wantErr)

			if !tt.wantErr {
				assert.Equal(t, tt.wantNumber, c.Number)
			}
		})
	}
}
