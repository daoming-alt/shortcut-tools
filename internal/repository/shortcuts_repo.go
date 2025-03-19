package repository

import (
	"github.com/daoming-alt/shortcut-tools/config"
	"github.com/daoming-alt/shortcut-tools/internal/models"
)

type ShortcutRepository struct {
}

func (r *ShortcutRepository) GetShortcuts() ([]models.Shortcut, error) {

	var models []models.Shortcut
	err := database.Find(&models).Error
	return models, err

}
