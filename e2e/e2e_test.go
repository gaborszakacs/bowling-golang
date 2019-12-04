package e2e_test

import (
	"os"
	"os/exec"
	"testing"
)

func TestE2E(t *testing.T) {
	err := os.Chdir("..")
	if err != nil {
		t.Fatalf("Failed to changed dir: %v", err)
	}
	rolls := "10,2,2,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0"
	cmd := exec.Command("go", "run", "main.go", "--rolls="+rolls)
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("Failed to run: %v: %s", err, string(output))
	}

	got := string(output)
	want := "Score: 18"
	if got != want {

		t.Errorf("got: %s, want: %s", got, want)
	}
}
