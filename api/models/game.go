package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Edition string
type GameState string

const (
	JavaEdition    Edition = "java"
	BedrockEdition Edition = "bedrock"

	Online   GameState = "online"
	Offline  GameState = "offline"
	Starting GameState = "starting"
	Stopping GameState = "stopping"
)

type Game struct {
	gorm.Model
	ID              uuid.UUID `gorm:"type:uuid;primaryKey"`
	UserID          uuid.UUID
	Name            string  `gorm:"not null;unique;size:255;default:null"`
	CustomSubdomain *string `gorm:"size:63;default:null"`
	Edition         Edition
	GameState       GameState `gorm:"not null;default:null"`
}
