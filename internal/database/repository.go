package database

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/Nayati-Indonesia/it-audit-collector/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func insertComputer(ctx context.Context, db *gorm.DB, data models.Computer) (uint, bool, error) {
	var existingComputer models.Computer

	err := db.Where("ip_address = ?", data.IPAddress).First(&existingComputer).Error
	if err == nil {
		db.WithContext(ctx).Model(&existingComputer).Updates(data)
		return existingComputer.ID, true, nil
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		fmt.Printf("Add new record for (%s)\n", data.IPAddress)
		if err := db.WithContext(ctx).Create(&data).Error; err != nil {
			return 0, false, err
		}

		return data.ID, false, nil
	}

	return 0, false, err
}

func upsertComputerSpecs(ctx context.Context, db *gorm.DB, computerID uint, data *models.ComputerSpec) error {
	data.ComputerID = computerID
	data.CollectedAt = time.Now()

	return db.WithContext(ctx).Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "computer_id"}},
		UpdateAll: true,
	}).Create(data).Error
}

func syncGPUs(ctx context.Context, db *gorm.DB, computerID uint, data []models.GPUModel) error {
	db.WithContext(ctx).Where("computer_id = ? AND `index` >= ?", computerID, len(data)).Delete(&models.GPUModel{})
	for i := range data {
		data[i].ComputerID = computerID

		err := db.WithContext(ctx).Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "computer_id"}, {Name: "index"}},
			UpdateAll: true,
		}).Create(&data[i]).Error

		if err != nil {
			return err
		}
	}

	return nil
}

func syncStorages(ctx context.Context, db *gorm.DB, computerID uint, data []models.DiskModel) error {
	var currentDrives []string
	for _, d := range data {
		currentDrives = append(currentDrives, d.Drive)
	}

	if len(currentDrives) > 0 {
		db.WithContext(ctx).Where("computer_id = ? AND drive NOT IN ?", computerID, currentDrives).Delete(&models.DiskModel{})
	} else {
		db.WithContext(ctx).Where("computer_id = ?", computerID).Delete(&models.DiskModel{})
	}

	for i := range data {
		data[i].ComputerID = computerID

		err := db.WithContext(ctx).Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "computer_id"}, {Name: "drive"}},
			UpdateAll: true,
		}).Create(&data[i]).Error

		if err != nil {
			return err
		}
	}

	return nil
}

func InsertAll(ctx context.Context, db *gorm.DB, data *models.SpecResponse) error {
	computer := data.ToComputer()
	computerSpec := data.ToComputerSpec()
	gpus := data.ToGpus()
	disks := data.ToDisks()

	// Database Transaction (CTX)
	return db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		computerId, _, err := insertComputer(ctx, tx, computer)
		if err != nil {
			return fmt.Errorf("failed to insert computer (%s): %w", computer.IPAddress, err)
		}

		if err = upsertComputerSpecs(ctx, tx, computerId, &computerSpec); err != nil {
			return fmt.Errorf("failed to upsert specs for computer (%s): %w", computer.IPAddress, err)
		}

		if err = syncGPUs(ctx, tx, computerId, gpus); err != nil {
			return fmt.Errorf("failed to sync gpus for computer (%s): %w", computer.IPAddress, err)
		}

		if err = syncStorages(ctx, tx, computerId, disks); err != nil {
			return fmt.Errorf("failed to sync storages for computer (%s): %w", computer.IPAddress, err)
		}

		fmt.Printf("Successfully processed asset for: %s (ID: %d)\n", computer.IPAddress, computerId)
		return nil
	})
}
