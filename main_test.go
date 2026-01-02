package main

import (
	"context"
	"testing"
)

func TestDummyProvider_Info(t *testing.T) {
	p := New()
	info := p.Info()

	if info.ID != "provider-dummy" {
		t.Errorf("Expected plugin ID 'provider-dummy', got '%s'", info.ID)
	}
}

func TestDummyProvider_Analyze(t *testing.T) {
	p := New()

	report, err := p.Analyze(context.Background(), nil)

	if err != nil {
		t.Fatalf("Analyze returned unexpected error: %v", err)
	}

	if report.PluginID != "provider-dummy" {
		t.Errorf("Expected report PluginID 'provider-dummy', got '%s'", report.PluginID)
	}

	if len(report.Checks) != 3 {
		t.Fatalf("Expected 3 checks, got %d", len(report.Checks))
	}

	check1 := report.Checks[0]
	if check1.ID != "DUMMY-001" {
		t.Errorf("Expected first check to be DUMMY-001, got %s", check1.ID)
	}
	if !check1.Passed {
		t.Error("DUMMY-001 should pass")
	}
	if check1.Score != 100 {
		t.Errorf("DUMMY-001 Score: got %d, expected 100", check1.Score)
	}

	check2 := report.Checks[1]
	if check2.ID != "DUMMY-002" {
		t.Errorf("Expected second check to be DUMMY-002, got %s", check2.ID)
	}
	if check2.Passed {
		t.Error("DUMMY-002 should fail (it's hardcoded to fail)")
	}
	if check2.Score != 0 {
		t.Errorf("DUMMY-002 Score: got %d, expected 0", check2.Score)
	}

	check3 := report.Checks[2]
	if check3.ID != "DUMMY-003" {
		t.Errorf("Expected third check to be DUMMY-003, got %s", check3.ID)
	}
	if !check3.Passed {
		t.Error("DUMMY-003 should pass")
	}
}
