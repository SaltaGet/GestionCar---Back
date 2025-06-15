package jobs

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/DanielChachagua/GestionCar/pkg/dependencies"
)

type Config struct {
	User      string   `json:"user"`
	Password  string   `json:"password"`
	Databases []string `json:"databases"`
	Host      string   `json:"host"`
	BackupDir string   `json:"backup_dir"`
}

type Checkpoint struct {
	BinlogFile string `json:"binlog_file"`
	Position   int    `json:"position"`
}


func LoadConfig(deps *dependencies.Application) (Config, error) {
	var cfg Config
	data, err := os.ReadFile(path)
	if err != nil {
		return cfg, err
	}
	err = json.Unmarshal(data, &cfg)
	return cfg, err
}

func parseDSN(dsn string) (Config, error) {
	var cfg Config

	// Separar la parte user:password y el resto
	parts := strings.SplitN(dsn, "@tcp(", 2)
	if len(parts) != 2 {
		return cfg, fmt.Errorf("DSN inválido: falta '@tcp('")
	}

	userPass := parts[0]          // root:Qwer1234*
	hostAndRest := parts[1]       // 127.0.0.1:3306)/gestion_car?charset=...

	// Parsear user y password
	up := strings.SplitN(userPass, ":", 2)
	if len(up) != 2 {
		return cfg, fmt.Errorf("DSN inválido: falta ':' entre usuario y password")
	}
	cfg.User = up[0]
	cfg.Password = up[1]

	// Separar host y base de datos
	// hostAndRest tiene la forma: 127.0.0.1:3306)/gestion_car?...
	idx := strings.Index(hostAndRest, ")/")
	if idx == -1 {
		return cfg, fmt.Errorf("DSN inválido: falta ')/'")
	}

	cfg.Host = hostAndRest[:idx]
	dbAndParams := hostAndRest[idx+2:] // gestion_car?charset=utf8mb4...

	// Separar base de datos y parámetros (opcional)
	dbName := dbAndParams
	if i := strings.Index(dbAndParams, "?"); i != -1 {
		dbName = dbAndParams[:i]
	}
	cfg.Databases = []string{dbName}

	return cfg, nil
}

func checkpointPath(db string, dir string) string {
	return filepath.Join(dir, fmt.Sprintf("%s_checkpoint.json", db))
}

func backupExists(db string, dir string) bool {
	_, err := os.Stat(checkpointPath(db, dir))
	return err == nil
}

func runFullBackup(cfg Config, db string) error {
	timestamp := time.Now().Format("2006-01-02_15-04-05")
	backupFile := filepath.Join(cfg.BackupDir, fmt.Sprintf("%s_full_%s.sql", db, timestamp))

	cmd := exec.Command("mysqldump",
		"-u", cfg.User,
		"-p"+cfg.Password,
		"-h", cfg.Host,
		"--databases", db,
		"--routines", "--events", "--single-transaction")

	outputFile, err := os.Create(backupFile)
	if err != nil {
		return err
	}
	defer outputFile.Close()
	cmd.Stdout = outputFile

	return cmd.Run()
}

func getBinlogStatus(cfg Config) (Checkpoint, error) {
	cmd := exec.Command("mysql", "-u", cfg.User, "-p"+cfg.Password, "-h", cfg.Host, "-e", "SHOW MASTER STATUS\\G")
	output, err := cmd.Output()
	if err != nil {
		return Checkpoint{}, err
	}

	var cp Checkpoint
	lines := strings.Split(string(output), "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "File:") {
			cp.BinlogFile = strings.TrimSpace(strings.Split(line, ":")[1])
		} else if strings.HasPrefix(line, "Position:") {
			fmt.Sscanf(line, "  Position: %d", &cp.Position)
		}
	}

	if cp.BinlogFile == "" || cp.Position == 0 {
		return cp, fmt.Errorf("estado de binlog inválido")
	}
	return cp, nil
}

func saveCheckpoint(cfg Config, db string, cp Checkpoint) error {
	data, err := json.MarshalIndent(cp, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(checkpointPath(db, cfg.BackupDir), data, 0644)
}

func loadCheckpoint(cfg Config, db string) (Checkpoint, error) {
	data, err := os.ReadFile(checkpointPath(db, cfg.BackupDir))
	if err != nil {
		return Checkpoint{}, err
	}
	var cp Checkpoint
	err = json.Unmarshal(data, &cp)
	return cp, err
}

func runIncrementalBackup(cfg Config, db string, cp Checkpoint) (Checkpoint, error) {
	timestamp := time.Now().Format("2006-01-02_15-04-05")
	outFile := filepath.Join(cfg.BackupDir, fmt.Sprintf("%s_incremental_%s.sql", db, timestamp))

	cmd := exec.Command("mysqlbinlog",
		"--read-from-remote-server",
		"--host="+cfg.Host,
		"--user="+cfg.User,
		"--password="+cfg.Password,
		"--start-position", fmt.Sprint(cp.Position),
		cp.BinlogFile)

	out, err := os.Create(outFile)
	if err != nil {
		return cp, err
	}
	defer out.Close()
	cmd.Stdout = out

	if err := cmd.Run(); err != nil {
		return cp, err
	}

	// Actualizamos el checkpoint al nuevo estado del binlog
	newCp, err := getBinlogStatus(cfg)
	if err != nil {
		return cp, err
	}
	return newCp, nil
}

func RunBackup(cfg Config) {
	for _, db := range cfg.Databases {
		fmt.Println("Procesando DB:", db)
		if !backupExists(db, cfg.BackupDir) {
			fmt.Println("  No existe backup completo. Haciendo uno...")
			if err := runFullBackup(cfg, db); err != nil {
				fmt.Println("  ❌ Error backup full:", err)
				continue
			}
			cp, err := getBinlogStatus(cfg)
			if err != nil {
				fmt.Println("  ❌ Error obteniendo binlog:", err)
				continue
			}
			if err := saveCheckpoint(cfg, db, cp); err != nil {
				fmt.Println("  ❌ Error guardando checkpoint:", err)
			}
			fmt.Println("  ✅ Backup full realizado.")
		} else {
			fmt.Println("  Haciendo backup incremental...")
			cp, err := loadCheckpoint(cfg, db)
			if err != nil {
				fmt.Println("  ❌ Error leyendo checkpoint:", err)
				continue
			}
			newCp, err := runIncrementalBackup(cfg, db, cp)
			if err != nil {
				fmt.Println("  ❌ Error en backup incremental:", err)
				continue
			}
			if err := saveCheckpoint(cfg, db, newCp); err != nil {
				fmt.Println("  ❌ Error guardando nuevo checkpoint:", err)
			}
			fmt.Println("  ✅ Backup incremental realizado.")
		}
	}
}




// type Checkpoint struct {
// 	BinlogFile string `json:"binlog_file"`
// 	Position   int    `json:"position"`
// }

// const (
// 	checkpointPath = "backups/binlog_checkpoint.json"
// 	fullBackupPath = "backups/full_backup.sql"
// 	user           = "root"
// 	password       = "TU_PASSWORD"
// )

// func main() {
// 	if !checkpointExists() {
// 		fmt.Println("No hay checkpoint. Ejecutando backup completo...")
// 		if err := runFullBackup(); err != nil {
// 			fmt.Println("Error al hacer el backup completo:", err)
// 			return
// 		}

// 		cp, err := getCurrentBinlogStatus()
// 		if err != nil {
// 			fmt.Println("Error obteniendo estado de binlog:", err)
// 			return
// 		}

// 		if err := saveCheckpoint(cp); err != nil {
// 			fmt.Println("Error guardando checkpoint:", err)
// 			return
// 		}

// 		fmt.Println("Backup completo y checkpoint guardados correctamente.")
// 	} else {
// 		fmt.Println("Checkpoint ya existe. Podés ejecutar backup incremental.")
// 		// Acá podrías llamar a runIncrementalBackup()
// 	}
// }

// func checkpointExists() bool {
// 	_, err := os.Stat(checkpointPath)
// 	return err == nil
// }

// func runFullBackup() error {
// 	// Crear carpeta si no existe
// 	_ = os.MkdirAll("backups", 0755)

// 	cmd := exec.Command("mysqldump",
// 		"-u", user,
// 		"-p"+password,
// 		"--all-databases",
// 		"--routines",
// 		"--events",
// 		"--single-transaction")

// 	outputFile, err := os.Create(fullBackupPath)
// 	if err != nil {
// 		return fmt.Errorf("no se pudo crear archivo de backup: %w", err)
// 	}
// 	defer outputFile.Close()

// 	cmd.Stdout = outputFile

// 	return cmd.Run()
// }

// func getCurrentBinlogStatus() (Checkpoint, error) {
// 	cmd := exec.Command("mysql", "-u", user, "-p"+password, "-e", "SHOW MASTER STATUS\\G")
// 	output, err := cmd.Output()
// 	if err != nil {
// 		return Checkpoint{}, err
// 	}

// 	var file string
// 	var pos int

// 	lines := strings.Split(string(output), "\n")
// 	for _, line := range lines {
// 		if strings.HasPrefix(line, "File:") {
// 			file = strings.TrimSpace(strings.Split(line, ":")[1])
// 		} else if strings.HasPrefix(line, "Position:") {
// 			fmt.Sscanf(line, "  Position: %d", &pos)
// 		}
// 	}

// 	if file == "" || pos == 0 {
// 		return Checkpoint{}, fmt.Errorf("no se pudo interpretar el binlog y la posición")
// 	}

// 	return Checkpoint{BinlogFile: file, Position: pos}, nil
// }

// func saveCheckpoint(cp Checkpoint) error {
// 	data, err := json.MarshalIndent(cp, "", "  ")
// 	if err != nil {
// 		return err
// 	}
// 	return os.WriteFile(checkpointPath, data, 0644)
// }



// import (
// 	"fmt"
// 	"log"
// 	"os"
// 	"os/exec"
// 	"path/filepath"
// )

// func IncrementalBackup(today string, lastFullDir string) error {
// 	incDir := fmt.Sprintf("/backups/%s/inc", today)

// 	// Backup incremental
// 	cmd := exec.Command("xtrabackup",
// 		"--backup",
// 		"--target-dir="+incDir,
// 		"--incremental-basedir="+lastFullDir,
// 		"--user=backupuser",
// 		"--password=backuppass",
// 	)

// 	out, err := cmd.CombinedOutput()
// 	if err != nil {
// 		return fmt.Errorf("error en backup incremental: %v\n%s", err, out)
// 	}

// 	// Aplicar logs incrementales sobre backup completo
// 	prepare := exec.Command("xtrabackup",
// 		"--prepare",
// 		"--apply-log-only",
// 		"--target-dir="+lastFullDir,
// 		"--incremental-dir="+incDir,
// 	)
// 	out, err = prepare.CombinedOutput()
// 	if err != nil {
// 		return fmt.Errorf("error aplicando logs: %v\n%s", err, out)
// 	}

// 	fmt.Println("Backup incremental creado y aplicado.")
// 	return nil
// }


// func GetLastFullBackupDir() string {
// 	base := "/backups"
// 	entries, err := os.ReadDir(base)
// 	if err != nil {
// 		log.Fatal("No se puede leer el directorio de backups:", err)
// 	}

// 	var last string
// 	for _, e := range entries {
// 		if !e.IsDir() {
// 			continue
// 		}
// 		fullPath := filepath.Join(base, e.Name(), "full")
// 		if _, err := os.Stat(fullPath); err == nil {
// 			last = fullPath
// 		}
// 	}
// 	if last == "" {
// 		log.Fatal("No se encontró ningún backup completo previo.")
// 	}
// 	return last
// }

// func FullBackup(today string) error {
// 	fullDir := fmt.Sprintf("/backups/%s/full", today)

// 	cmd := exec.Command("xtrabackup",
// 		"--backup",
// 		"--target-dir="+fullDir,
// 		"--user=backupuser",
// 		"--password=backuppass",
// 	)

// 	out, err := cmd.CombinedOutput()
// 	if err != nil {
// 		return fmt.Errorf("xtrabackup error: %v\n%s", err, out)
// 	}

// 	prepare := exec.Command("xtrabackup", "--prepare", "--target-dir="+fullDir)
// 	out, err = prepare.CombinedOutput()
// 	if err != nil {
// 		return fmt.Errorf("prepare error: %v\n%s", err, out)
// 	}

// 	return nil
// }
