package problem

import (
	"io/ioutil"
	"os"
	"reflect"
	"testing"
)

func TestDeleteDuplicates(t *testing.T) {
	// Test case 1: Empty list
	if result := DeleteDuplicates(nil); result != nil {
		t.Errorf("Expected nil, but got %v", result)
	}

	// Test case 2: List with no duplicates
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 3}
	expected := &ListNode{Val: 1}
	expected.Next = &ListNode{Val: 2}
	expected.Next.Next = &ListNode{Val: 3}
	if result := DeleteDuplicates(head); !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}

	// Test case 3: List with duplicates
	head = &ListNode{Val: 1}
	head.Next = &ListNode{Val: 1}
	head.Next.Next = &ListNode{Val: 2}
	expected = &ListNode{Val: 1}
	expected.Next = &ListNode{Val: 2}
	if result := DeleteDuplicates(head); !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}

	// Test case 4: List with multiple duplicates
	head = &ListNode{Val: 1}
	head.Next = &ListNode{Val: 1}
	head.Next.Next = &ListNode{Val: 2}
	head.Next.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next.Next = &ListNode{Val: 3}
	expected = &ListNode{Val: 1}
	expected.Next = &ListNode{Val: 2}
	expected.Next.Next = &ListNode{Val: 3}
	if result := DeleteDuplicates(head); !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestDeleteDuplicates2(t *testing.T) {
	// Test case 1: Empty list
	if result := DeleteDuplicates2(nil); result != nil {
		t.Errorf("Expected nil, but got %v", result)
	}

	// Test case 2: List with no duplicates
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 3}
	expected := &ListNode{Val: 1}
	expected.Next = &ListNode{Val: 2}
	expected.Next.Next = &ListNode{Val: 3}
	if result := DeleteDuplicates2(head); !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}

	// Test case 3: List with duplicates
	head = &ListNode{Val: 1}
	head.Next = &ListNode{Val: 1}
	head.Next.Next = &ListNode{Val: 2}
	expected = &ListNode{Val: 1}
	expected.Next = &ListNode{Val: 2}
	if result := DeleteDuplicates2(head); !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}

	// Test case 4: List with multiple duplicates
	head = &ListNode{Val: 1}
	head.Next = &ListNode{Val: 1}
	head.Next.Next = &ListNode{Val: 2}
	head.Next.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next.Next = &ListNode{Val: 3}
	expected = &ListNode{Val: 1}
	expected.Next = &ListNode{Val: 2}
	expected.Next.Next = &ListNode{Val: 3}
	if result := DeleteDuplicates2(head); !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestDeleteDuplicates2_Print(t *testing.T) {
	// Test case 1: Empty list
	head := (*ListNode)(nil)
	expected := "First loop:  null\n"
	testDeleteDuplicates2Print(t, head, expected)

	// Test case 2: List with no duplicates
	head = &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 3}
	expected = "First loop:  {\n  \"Val\": 1,\n  \"Next\": {\n    \"Val\": 2,\n    \"Next\": {\n      \"Val\": 3,\n      \"Next\": null\n    }\n  }\n}\n"
	testDeleteDuplicates2Print(t, head, expected)

	// Test case 3: List with duplicates
	head = &ListNode{Val: 1}
	head.Next = &ListNode{Val: 1}
	head.Next.Next = &ListNode{Val: 2}
	expected = "First loop:  {\n  \"Val\": 2,\n  \"Next\": null\n}\n"
	testDeleteDuplicates2Print(t, head, expected)

	// Test case 4: List with multiple duplicates
	head = &ListNode{Val: 1}
	head.Next = &ListNode{Val: 1}
	head.Next.Next = &ListNode{Val: 2}
	head.Next.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next.Next = &ListNode{Val: 3}
	expected = "First loop:  {\n  \"Val\": 2,\n  \"Next\": {\n    \"Val\": 3,\n    \"Next\": null\n  }\n}\n"
	testDeleteDuplicates2Print(t, head, expected)
}

func testDeleteDuplicates2Print(t *testing.T, head *ListNode, expected string) {
	t.Helper()
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	DeleteDuplicates2(head)

	w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stdout = oldStdout

	if string(out) != expected {
		t.Errorf("Expected %q, but got %q", expected, string(out))
	}
}
