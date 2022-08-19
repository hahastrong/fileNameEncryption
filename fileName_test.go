package fileNameEncryption

import "testing"

func TestValidateFileName(t *testing.T) {
	filename := "hell.dfd/20220817.mp4"

	filenameExt, err := GenerateFileName(filename)
	if err != nil {
		t.Fatalf("failed to generate filename, err: %v", err)
	}

	t.Logf("%s -> %s", filename, filenameExt)

	flag, err := ValidateFileName(filenameExt)
	if err != nil {
		t.Fatalf("failed to parse filename, err: %v", err)
	}
	if !flag {
		t.Fatal("validate false")
	}
}
