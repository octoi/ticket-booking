package models

import (
	"context"
	"time"
)

type Ticket struct {
	ID        uint      `json:"id" gorm:"primarykey"`
	EventId   uint      `json:"eventId"`
	UserId    uint      `json:"userId" gorm:"foreignkey:UserId:onUpdate:CASCADE,onDelete:CASCADE;"`
	Event     Event     `json:"event" gorm:"foreignkey:EventId:onUpdate:CASCADE,onDelete:CASCADE;"`
	Entered   bool      `json:"entered" default:"false"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type TicketRepository interface {
	GetMany(ctx context.Context, userId uint) ([]*Ticket, error)
	GetOne(ctx context.Context, userId uint, ticketId uint) (*Ticket, error)
	CreateOne(ctx context.Context, userId uint, ticket *Ticket) (*Ticket, error)
	UpdateOne(ctx context.Context, userId uint, ticketId uint, updateData map[string]interface{}) (*Ticket, error)
}

type ValidateTicket struct {
	TicketId uint `json:"ticketId"`
	OwnerId  uint `json:"ownerId"`
}
