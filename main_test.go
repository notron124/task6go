package main

import "testing"

func Test1(t *testing.T) {
	heights := []int{1, 2}
	i, j := RightSwitch(heights)
	wanti, wantj := -1, -1
	if wanti != i || wantj != j {
		t.Errorf("got = %d, %d; want = %d, %d", i, j, wanti, wantj)
	}
}

func Test2(t *testing.T) {
	heights := []int{2, 1}
	i, j := RightSwitch(heights)
	wanti, wantj := 1, 2
	if wanti != i || wantj != j {
		t.Errorf("got = %d, %d; want = %d, %d", i, j, wanti, wantj)
	}
}

func Test3(t *testing.T) {
	heights := []int{1, 1}
	i, j := RightSwitch(heights)
	wanti, wantj := -1, -1
	if wanti != i || wantj != j {
		t.Errorf("got = %d, %d; want = %d, %d", i, j, wanti, wantj)
	}
}

func Test4(t *testing.T) {
	heights := []int{2, 2}
	i, j := RightSwitch(heights)
	wanti, wantj := -1, -1
	if wanti != i || wantj != j {
		t.Errorf("got = %d, %d; want = %d, %d", i, j, wanti, wantj)
	}
}

func Test5(t *testing.T) {
	heights := []int{1, 2, 1}
	i, j := RightSwitch(heights)
	wanti, wantj := 1, 3
	if wanti != i || wantj != j {
		t.Errorf("got = %d, %d; want = %d, %d", i, j, wanti, wantj)
	}
}

func Test6(t *testing.T) {
	heights := []int{1, 2, 2}
	i, j := RightSwitch(heights)
	wanti, wantj := -1, -1
	if wanti != i || wantj != j {
		t.Errorf("got = %d, %d; want = %d, %d", i, j, wanti, wantj)
	}
}

func Test7(t *testing.T) {
	heights := []int{2, 1, 1}
	i, j := RightSwitch(heights)
	wanti, wantj := 1, 2
	if wanti != i || wantj != j {
		t.Errorf("got = %d, %d; want = %d, %d", i, j, wanti, wantj)
	}
}

func Test8(t *testing.T) {
	heights := []int{2, 1, 2}
	i, j := RightSwitch(heights)
	wanti, wantj := -1, -1
	if wanti != i || wantj != j {
		t.Errorf("got = %d, %d; want = %d, %d", i, j, wanti, wantj)
	}
}

func Test9(t *testing.T) {
	heights := []int{1, 1, 2}
	i, j := RightSwitch(heights)
	wanti, wantj := 2, 3
	if wanti != i || wantj != j {
		t.Errorf("got = %d, %d; want = %d, %d", i, j, wanti, wantj)
	}
}

func Test10(t *testing.T) {
	heights := []int{1, 1, 1}
	i, j := RightSwitch(heights)
	wanti, wantj := -1, -1
	if wanti != i || wantj != j {
		t.Errorf("got = %d, %d; want = %d, %d", i, j, wanti, wantj)
	}
}

func Test11(t *testing.T) {
	heights := []int{2, 2, 1}
	i, j := RightSwitch(heights)
	wanti, wantj := -1, -1
	if wanti != i || wantj != j {
		t.Errorf("got = %d, %d; want = %d, %d", i, j, wanti, wantj)
	}
}

func Test12(t *testing.T) {
	heights := []int{2, 2, 2}
	i, j := RightSwitch(heights)
	wanti, wantj := -1, -1
	if wanti != i || wantj != j {
		t.Errorf("got = %d, %d; want = %d, %d", i, j, wanti, wantj)
	}
}

func Test13(t *testing.T) {
	heights := []int{2, 1, 2, 1}
	i, j := RightSwitch(heights)
	wanti, wantj := -1, -1
	if wanti != i || wantj != j {
		t.Errorf("got = %d, %d; want = %d, %d", i, j, wanti, wantj)
	}
}

func Test14(t *testing.T) {
	heights := []int{1, 2, 1, 2}
	i, j := RightSwitch(heights)
	wanti, wantj := 1, 3
	if wanti != i || wantj != j {
		t.Errorf("got = %d, %d; want = %d, %d", i, j, wanti, wantj)
	}
}

func Test15(t *testing.T) {
	heights := []int{2, 1, 1, 2}
	i, j := RightSwitch(heights)
	wanti, wantj := 1, 2
	if wanti != i || wantj != j {
		t.Errorf("got = %d, %d; want = %d, %d", i, j, wanti, wantj)
	}
}

func Test16(t *testing.T) {
	heights := []int{1, 1, 2, 2}
	i, j := RightSwitch(heights)
	wanti, wantj := 2, 3
	if wanti != i || wantj != j {
		t.Errorf("got = %d, %d; want = %d, %d", i, j, wanti, wantj)
	}
}

func Test17(t *testing.T) {
	heights := []int{1, 2, 2, 1}
	i, j := RightSwitch(heights)
	wanti, wantj := 3, 4
	if wanti != i || wantj != j {
		t.Errorf("got = %d, %d; want = %d, %d", i, j, wanti, wantj)
	}
}

func Test18(t *testing.T) {
	heights := []int{2, 2, 1, 2}
	i, j := RightSwitch(heights)
	wanti, wantj := -1, -1
	if wanti != i || wantj != j {
		t.Errorf("got = %d, %d; want = %d, %d", i, j, wanti, wantj)
	}
}

func Test19(t *testing.T) {
	heights := []int{1, 1, 1, 2}
	i, j := RightSwitch(heights)
	wanti, wantj := -1, -1
	if wanti != i || wantj != j {
		t.Errorf("got = %d, %d; want = %d, %d", i, j, wanti, wantj)
	}
}

func Test20(t *testing.T) {
	heights := []int{1, 2, 1, 1}
	i, j := RightSwitch(heights)
	wanti, wantj := -1, -1
	if wanti != i || wantj != j {
		t.Errorf("got = %d, %d; want = %d, %d", i, j, wanti, wantj)
	}
}

func Test21(t *testing.T) {
	heights := []int{2, 1, 2, 2}
	i, j := RightSwitch(heights)
	wanti, wantj := -1, -1
	if wanti != i || wantj != j {
		t.Errorf("got = %d, %d; want = %d, %d", i, j, wanti, wantj)
	}
}
