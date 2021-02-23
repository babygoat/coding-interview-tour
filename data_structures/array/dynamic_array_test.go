package array

import (
	"testing"
)

func TestIndexAt(t *testing.T) {
	cases := []struct {
		name      string
		data      []int
		index     int
		want      int
		wantError error
	}{
		{
			name:  "Given valid index, returns value",
			data:  []int{1, 2},
			index: 1,
			want:  2,
		},
		{
			name:      "Given out of bound index, returns error",
			data:      []int{1, 2},
			index:     2,
			want:      -1,
			wantError: ErrArrIndexOutOfBound,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			da := New(len(tt.data))
			for _, v := range tt.data {
				da.Push(v)
			}
			value, err := da.IndexAt(tt.index)

			if value != tt.want {
				t.Errorf("expect value=(%v), but got (%v)", tt.want, value)
			}
			if err != tt.wantError {
				t.Errorf("expect err=(%v), but got (%v)", tt.wantError, err)
			}
		})
	}
}

func TestInsert(t *testing.T) {
	cases := []struct {
		name        string
		insertValue int
		insertIndex int
		exist       []int
		want        []int
		wantError   error
	}{
		{
			name:        "Test insertion at the front(prepend)",
			insertValue: 4,
			insertIndex: 0,
			exist:       []int{1, 2, 3},
			want:        []int{4, 1, 2, 3},
			wantError:   nil,
		},
		{
			name:        "Test insertion at the middle",
			insertValue: 4,
			insertIndex: 1,
			exist:       []int{1, 0, 3, 2},
			want:        []int{1, 4, 0, 3, 2},
			wantError:   nil,
		},
		{
			name:        "Test insertion at the end(push_back)",
			insertValue: 4,
			insertIndex: 4,
			exist:       []int{1, 0, 3, 2},
			want:        []int{1, 0, 3, 2, 4},
			wantError:   nil,
		},
		{
			name:        "Test insertion index out of bound",
			insertValue: 4,
			insertIndex: 2,
			exist:       []int{0},
			wantError:   ErrArrIndexOutOfBound,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			da := New(len(tt.exist))
			for _, val := range tt.exist {
				da.Push(val)
			}
			err := da.Insert(tt.insertIndex, tt.insertValue)
			if err != tt.wantError {
				t.Errorf("expect err(%v), got (%v)", tt.wantError, err)
			}
			for ind, val := range tt.want {
				got, err := da.IndexAt(ind)
				if err != nil {
					t.Errorf("unexpected error")
					t.Fail()
				}
				if got != val {
					t.Errorf("expect arr[%d] = %v, got %v", ind, val, got)
				}
			}
		})
	}
}

func TestErase(t *testing.T) {
	var cases = []struct {
		name       string
		eraseIndex int
		exist      []int
		want       []int
		wantError  error
	}{
		{
			name:       "Test delete the value from the front(pop_front)",
			eraseIndex: 0,
			exist:      []int{1, 0, 2, 3},
			want:       []int{0, 2, 3},
			wantError:  nil,
		},
		{
			name:       "Test delete the value in the middle",
			eraseIndex: 2,
			exist:      []int{1, 0, 2, 3},
			want:       []int{1, 0, 3},
			wantError:  nil,
		},
		{
			name:       "Test delete value from behind(pop_back)",
			eraseIndex: 3,
			exist:      []int{1, 0, 2, 3},
			want:       []int{1, 0, 2},
			wantError:  nil,
		},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			da := New(len(tt.exist))
			for _, val := range tt.exist {
				da.Push(val)
			}
			err := da.EraseAt(tt.eraseIndex)
			if err != tt.wantError {
				t.Errorf("expect err(%v), got (%v)", tt.wantError, err)
			}
			for ind, val := range tt.want {
				got, err := da.IndexAt(ind)
				if err != nil {
					t.Errorf("unexpected error")
					t.Fail()
				}
				if got != val {
					t.Errorf("expect arr[%d] = %v, got %v", ind, val, got)
				}
			}
		})
	}
}

func TestRemove(t *testing.T) {
	cases := []struct {
		name   string
		remove int
		exist  []int
		want   []int
	}{
		{
			name:   "Test remove element occurs multiple time ",
			remove: 1,
			exist:  []int{1, 3, 1, 1, 5, 6},
			want:   []int{3, 5, 6},
		},
		{
			name:   "Test remove element not exist in the arrays(no change)",
			remove: 0,
			exist:  []int{1, 3, 1, 1, 5, 6},
			want:   []int{1, 3, 1, 1, 5, 6},
		},
		{
			name:   "Test safely remove element in empty array",
			remove: 0,
			exist:  []int{},
			want:   []int{},
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			da := New(len(tt.exist))
			for _, val := range tt.exist {
				da.Push(val)
			}
			da.Remove(tt.remove)
			if da.Size() != len(tt.want) {
				t.Errorf("expect array len (%d), got (%d)", len(tt.want), da.Size())
				t.Fail()
			}

			for ind, val := range tt.want {
				got, err := da.IndexAt(ind)
				if err != nil {
					t.Errorf("unexpected error")
					t.Fail()
				}
				if got != val {
					t.Errorf("expect arr[%d] = %v, got %v", ind, val, got)
				}
			}
		})
	}
}
