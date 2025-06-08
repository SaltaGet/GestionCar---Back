package jobs

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func IncrementalBackup(today string, lastFullDir string) error {
	incDir := fmt.Sprintf("/backups/%s/inc", today)

	// Backup incremental
	cmd := exec.Command("xtrabackup",
		"--backup",
		"--target-dir="+incDir,
		"--incremental-basedir="+lastFullDir,
		"--user=backupuser",
		"--password=backuppass",
	)

	out, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("error en backup incremental: %v\n%s", err, out)
	}

	// Aplicar logs incrementales sobre backup completo
	prepare := exec.Command("xtrabackup",
		"--prepare",
		"--apply-log-only",
		"--target-dir="+lastFullDir,
		"--incremental-dir="+incDir,
	)
	out, err = prepare.CombinedOutput()
	if err != nil {
		return fmt.Errorf("error aplicando logs: %v\n%s", err, out)
	}

	fmt.Println("Backup incremental creado y aplicado.")
	return nil
}


func GetLastFullBackupDir() string {
	base := "/backups"
	entries, err := os.ReadDir(base)
	if err != nil {
		log.Fatal("No se puede leer el directorio de backups:", err)
	}

	var last string
	for _, e := range entries {
		if !e.IsDir() {
			continue
		}
		fullPath := filepath.Join(base, e.Name(), "full")
		if _, err := os.Stat(fullPath); err == nil {
			last = fullPath
		}
	}
	if last == "" {
		log.Fatal("No se encontró ningún backup completo previo.")
	}
	return last
}

func FullBackup(today string) error {
	fullDir := fmt.Sprintf("/backups/%s/full", today)

	cmd := exec.Command("xtrabackup",
		"--backup",
		"--target-dir="+fullDir,
		"--user=backupuser",
		"--password=backuppass",
	)

	out, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("xtrabackup error: %v\n%s", err, out)
	}

	prepare := exec.Command("xtrabackup", "--prepare", "--target-dir="+fullDir)
	out, err = prepare.CombinedOutput()
	if err != nil {
		return fmt.Errorf("prepare error: %v\n%s", err, out)
	}

	return nil
}
