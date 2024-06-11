package orm

import "github.com/lib/pq"

type Meowfficer struct {
	ID          uint32        `gorm:"primaryKey" json:"id"`
	Name        string        `json:"name"`
	Rarity      uint32        `json:"rarity"`
	Exp         uint32        `json:"exp"`
	ExpCost     uint32        `json:"exp_cost"`
	GroupType   uint32        `json:"group_type"`
	Nationality uint32        `json:"nationality"`
	AbilityShow pq.Int32Array `json:"ability_show" gorm:"type:integer[]"`
}

type OwnedMeowfficer struct {
	OwnerID uint32 `gorm:"type:int;not_null"`
	ID      uint32 `gorm:"primary_key"`
}