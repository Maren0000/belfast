package orm

import (
	"errors"

	"github.com/lib/pq"
)

var (
	ErrInvalidShipID = errors.New("invalid ship id")
)

type Fleet struct {
	ID             uint32        `gorm:"primary_key"` // uniquely identifies the fleet
	GameID         uint32        `gorm:"not_null"`    // uniquely identifies for each commander
	CommanderID    uint32        `gorm:"not_null"`    // owner of the fleet
	Name           string        `gorm:"size:32;not_null;varchar(32)"`
	ShipList       pq.Int64Array `gorm:"type:integer[];not_null"`
	MeowfficerList pq.Int64Array `gorm:"type:integer[];not_null"`
}

// Creates a fleet for the given commander, with the given ships
func CreateFleet(owner *Commander, id uint32, name string, ships []uint32) error {
	var fleet Fleet
	fleet.CommanderID = owner.CommanderID
	fleet.GameID = id
	for _, shipID := range ships {
		// check if the commander has this ship
		if _, ok := owner.OwnedShipsMap[shipID]; !ok {
			return ErrInvalidShipID
		}
		fleet.ShipList = append(fleet.ShipList, int64(shipID))
	}
	if err := GormDB.Create(&fleet).Error; err != nil {
		return err
	}
	owner.Fleets = append(owner.Fleets, fleet)
	owner.FleetsMap[fleet.GameID] = &owner.Fleets[len(owner.Fleets)-1]
	return nil
}

func (f *Fleet) RenameFleet(name string) error {
	return GormDB.Model(f).Update("name", name).Error
}

// Updates the ship list of the fleet
func (f *Fleet) UpdateShipList(owner *Commander, ships []uint32) error {
	f.ShipList = make([]int64, len(ships))
	for i, shipID := range ships {
		// check if the commander has this ship
		if _, ok := owner.OwnedShipsMap[shipID]; !ok {
			return ErrInvalidShipID
		}
		f.ShipList[i] = int64(shipID)
	}
	if err := GormDB.Save(f).Error; err != nil {
		return err
	}
	// Update the ship list in the commander's map
	owner.FleetsMap[f.GameID] = f

	// Find the fleet in the commander's list and update it
	for i, fleet := range owner.Fleets {
		if fleet.ID == f.ID {
			owner.Fleets[i] = *f
			break
		}
	}
	return nil
}

func (f *Fleet) AddMeowfficer(shipID uint32) error {
	panic("not implemented")
}

func (f *Fleet) RemoveMeowfficer(shipID uint32) error {
	panic("not implemented")
}
