package dto

import (
	"time"

	"gitlab.sudovi.me/erp/accounting-api/data"
)

type GetOrderListInputDTO struct {
	Page                *int    `json:"page" validate:"omitempty"`
	Size                *int    `json:"size" validate:"omitempty"`
	SupplierID          *int    `json:"supplier_id"`
	Status              *string `json:"status"`
	Search              *string `json:"search"`
	PublicProcurementID *int    `json:"public_procurement_id"`
}

type OrderListDTO struct {
	DateOrder            time.Time  `json:"date_order" validate:"required"`
	TotalPrice           float32    `json:"total_price"`
	PublicProcurementID  int        `json:"public_procurement_id"`
	SupplierID           int        `json:"supplier_id"`
	Status               string     `json:"status"`
	DateSystem           *time.Time `json:"date_system"`
	InvoiceDate          *time.Time `json:"invoice_date"`
	InvoiceNumber        *int       `json:"invoice_number"`
	OrganizationUnitID   int        `json:"organization_unit_id"`
	OfficeID             *int       `json:"office_id"`
	RecipientUserID      int        `json:"recipient_user_id"`
	DescriptionRecipient *string    `json:"description_recipient"`
}

type OrderListResponseDTO struct {
	ID                   int        `json:"id"`
	DateOrder            time.Time  `json:"date_order" validate:"required"`
	TotalPrice           float32    `json:"total_price"`
	PublicProcurementID  int        `json:"public_procurement_id"`
	SupplierID           int        `json:"supplier_id"`
	Status               string     `json:"status"`
	DateSystem           *time.Time `json:"date_system"`
	InvoiceDate          *time.Time `json:"invoice_date"`
	InvoiceNumber        *int       `json:"invoice_number"`
	OrganizationUnitID   int        `json:"organization_unit_id"`
	OfficeID             *int       `json:"office_id"`
	RecipientUserID      int        `json:"recipient_user_id"`
	DescriptionRecipient *string    `json:"description_recipient"`
	CreatedAt            time.Time  `json:"created_at"`
	UpdatedAt            time.Time  `json:"updated_at"`
}

func (dto OrderListDTO) ToOrderList() *data.OrderList {
	return &data.OrderList{
		DateOrder:            dto.DateOrder,
		TotalPrice:           dto.TotalPrice,
		PublicProcurementID:  dto.PublicProcurementID,
		SupplierID:           dto.SupplierID,
		Status:               dto.Status,
		DateSystem:           dto.DateSystem,
		InvoiceDate:          dto.InvoiceDate,
		InvoiceNumber:        dto.InvoiceNumber,
		OrganizationUnitID:   dto.OrganizationUnitID,
		OfficeID:             dto.OfficeID,
		RecipientUserID:      dto.RecipientUserID,
		DescriptionRecipient: dto.DescriptionRecipient,
	}
}

func ToOrderListResponseDTO(data data.OrderList) OrderListResponseDTO {
	return OrderListResponseDTO{
		ID:                   data.ID,
		DateOrder:            data.DateOrder,
		TotalPrice:           data.TotalPrice,
		PublicProcurementID:  data.PublicProcurementID,
		SupplierID:           data.SupplierID,
		Status:               data.Status,
		DateSystem:           data.DateSystem,
		InvoiceDate:          data.InvoiceDate,
		InvoiceNumber:        data.InvoiceNumber,
		OrganizationUnitID:   data.OrganizationUnitID,
		OfficeID:             data.OfficeID,
		RecipientUserID:      data.RecipientUserID,
		DescriptionRecipient: data.DescriptionRecipient,
		CreatedAt:            data.CreatedAt,
		UpdatedAt:            data.UpdatedAt,
	}
}

func ToOrderListListResponseDTO(contracts []*data.OrderList) []OrderListResponseDTO {
	dtoList := make([]OrderListResponseDTO, len(contracts))
	for i, x := range contracts {
		dtoList[i] = ToOrderListResponseDTO(*x)
	}
	return dtoList
}
