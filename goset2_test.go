package goset2_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/jettyu/goset2"
)

func TestStrings(t *testing.T) {
	s := goset2.Strings([]string{"2", "6", "4", "5", "4", "2", "3", "0", "1"})
	arr, ok := s.Slice().([]string)
	if !ok {
		t.Fatal(s)
	}
	if len(arr) != 7 {
		t.Fatal(s.Slice(), arr)
	}
	if !s.Has("0", 0) {
		t.Fatal(s.Slice())
	}
	if s.Has("0", 1) {
		t.Fatal(s.Slice())
	}
	if !s.Has("3", 2) {
		t.Fatal(s.Slice())
	}
	if s.Has("10", 0) {
		t.Fatal(s.Slice())
	}
	if s.Insert("1", "5", "7", "8") != 2 {
		t.Fatal(s.Slice())
	}
	// 删除中间，末尾混淆
	if s.Erase("7", "9") != 1 {
		t.Fatal(s.Slice())
	}
	// 删除中间和末尾
	if s.Erase("6", "8") != 2 {
		t.Fatal(s.Slice())
	}
	// 删除开头
	if s.Erase("0", "1") != 2 {
		t.Fatal(s.Slice())
	}
	for i, v := range s.Slice().([]string) {
		if fmt.Sprint(i+2) != v {
			t.Fatal(arr)
		}
	}
	ins := s.Intersection(goset2.Strings([]string{"1","2","4","6"}))
	if !ins.Equal([]string{"2","4"}) {
		t.Fatal(ins.Slice())
	}
}

func TestInts(t *testing.T) {
	s := goset2.Ints([]int{2, 6, 4, 5, 4, 2, 3, 0, 1})
	if !s.Equal([]int{0, 1, 2, 3, 4, 5, 6}) {
		t.Fatal(s.Slice())
	}
	if !s.Has(0, 0) {
		t.Fatal(s)
	}
	if s.Has(0, 1) {
		t.Fatal(s)
	}
	if !s.Has(3, 2) {
		t.Fatal(s)
	}
	if s.Has(10, 0) {
		t.Fatal(s)
	}
	if s.Insert([]int{1, 5, 7, 8}) != 2 {
		t.Fatal(s)
	}

	if s.Erase([]int{7, 9}) != 1 {
		t.Fatal(s.Slice())
	}
	if s.Erase([]int{6, 8}) != 2 {
		t.Fatal(s)
	}
	if s.Erase([]int{0, 1}) != 2 {
		t.Fatal(s)
	}
	if !s.Equal([]int{2, 3, 4, 5}) {
		t.Fatal(s.Slice())
	}

	clone := s.Clone()
	if !s.Equal(clone.Slice()) {
		t.Fatal(clone.Slice())
	}
	s.Erase(5)
	if s.Equal(clone.Slice()) {
		t.Fatal(clone.Slice(), s.Slice())
	}
}

func TestIntersection(t *testing.T) {
	arr1 := []int{0, 1, 1, 2, 2, 4, 5}
	arr2 := []int{1, 1, 2, 2, 3, 5, 6}
	sec := goset2.Ints(arr1).Intersection(goset2.Ints(arr2))
	except := []int{1, 2, 5}

	if !goset2.Ints(except).Equal(sec.Slice()) {
		t.Fatal(sec.Slice())
	}
}

func TestReflectErase(t *testing.T) {
	arr := []int{0, 1, 2, 3, 4, 5}
	rv := reflect.ValueOf(arr)
	rv = goset2.ReflectErase(rv, 0)
	if !goset2.Ints(rv.Interface().([]int)).Equal([]int{1, 2, 3, 4, 5}) {
		t.Fatal(rv.Interface())
	}
	rv = goset2.ReflectErase(rv, 1)
	if !goset2.Ints(rv.Interface().([]int)).Equal([]int{1, 3, 4, 5}) {
		t.Fatal(rv.Interface())
	}
	rv = goset2.ReflectErase(rv, 3)
	if !goset2.Ints(rv.Interface().([]int)).Equal([]int{1, 3, 4}) {
		t.Fatal(rv.Interface())
	}
	rv = goset2.ReflectErase(rv, 3)
	if !goset2.Ints(rv.Interface().([]int)).Equal([]int{1, 3, 4}) {
		t.Fatal(rv.Interface())
	}
}
