package backup

import (
	"errors"
	"fmt"
	"github.com/Koderbek/pocket_news_bot/internal/config"
	"github.com/jmoiron/sqlx"
	"gopkg.in/gomail.v2"
	"os"
	"time"
)

type Backup struct {
	db           *sqlx.DB
	cfg          *config.Config
	backupTables []string
}

func NewBackup(db *sqlx.DB, cfg *config.Config, backupTables []string) *Backup {
	return &Backup{db: db, cfg: cfg, backupTables: backupTables}
}

func (b *Backup) Run() error {
	for _, table := range b.backupTables {
		sql := fmt.Sprintf(
			"COPY (SELECT * FROM %s) TO '%s' WITH CSV HEADER;",
			table,
			makeBackupFilePath(b.cfg.Db.RootPath, table),
		)

		if _, err := b.db.Exec(sql); err != nil {
			return err
		}
	}

	if err := b.sendEmail(); err != nil {
		return err
	}

	for _, table := range b.backupTables {
		os.Remove(makeBackupFilePath(b.cfg.RootPath, table))
	}

	return nil
}

func (b *Backup) sendEmail() error {
	stmp := b.cfg.Stmp
	m := gomail.NewMessage()
	m.SetHeader("From", stmp.From)
	m.SetHeader("To", stmp.To)
	m.SetHeader("Subject", "Бэкап от "+time.Now().Format(time.DateOnly))
	for _, table := range b.backupTables {
		path := makeBackupFilePath(b.cfg.RootPath, table)
		if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
			return err
		}

		m.Attach(path)
	}

	d := gomail.NewDialer(stmp.Server, stmp.Port, stmp.From, stmp.Password)
	return d.DialAndSend(m)
}

func makeBackupFilePath(rootPath string, tableName string) string {
	return fmt.Sprintf("%s/backup/%s.csv", rootPath, tableName)
}
