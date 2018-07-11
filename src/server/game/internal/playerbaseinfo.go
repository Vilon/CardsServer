package internal

type PlayerBaseInfo struct {
	PlayerID uint   `gorm:"primary_key"`
	Name     string `gorm:"not null"`
}

func (playerInfo *PlayerBaseInfo) initValue(playerID uint) error {

	return nil
}

func (playerInfo *PlayerBaseInfo) saveValue() error {

	return nil
}

func CreatePlayerBaseInfo(playerID uint) error {

	return nil
}
