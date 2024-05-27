package dto

import (
	"time"

	"github.com/lib/pq"
	"gitlab.sudovi.me/erp/accounting-api/data"
)

type GetOrderListInputDTO struct {
	Page                *int    `json:"page" validate:"omitempty"`
	Size                *int    `json:"size" validate:"omitempty"`
	SupplierID          *int    `json:"supplier_id"`
	Status              *string `json:"status"`
	Search              *string `json:"search"`
	PublicProcurementID *int    `json:"public_procurement_id"`
	OrganizationUnitID  *int    `json:"organization_unit_id"`
	DateSystem          *string `json:"date_system"`
	Year                *string `json:"year"`
	FinanceOverview     *bool   `json:"finance_overview"`
	SortByDateOrder     *string `json:"sort_by_date_order"`
	SortByTotalPrice    *string `json:"sort_by_total_price"`
}

type OrderListDTO struct {
	DateOrder             time.Time     `json:"date_order" validate:"required"`
	TotalPrice            float32       `json:"total_price"`
	IsUsed                bool          `json:"is_used"`
	PublicProcurementID   *int          `json:"public_procurement_id"`
	SupplierID            *int          `json:"supplier_id"`
	Status                string        `json:"status"`
	DateSystem            *time.Time    `json:"date_system"`
	GroupOfArticlesID     *int          `json:"group_of_articles_id"`
	InvoiceDate           *time.Time    `json:"invoice_date"`
	ProFormaInvoiceDate   *time.Time    `json:"pro_forma_invoice_date"`
	ProFormaInvoiceNumber string        `json:"pro_forma_invoice_number"`
	InvoiceNumber         *string       `json:"invoice_number"`
	OrganizationUnitID    int           `json:"organization_unit_id"`
	OfficeID              *int          `json:"office_id"`
	AccountID             *int          `json:"account_id"`
	RecipientUserID       *int          `json:"recipient_user_id"`
	Description           *string       `json:"description"`
	OrderFile             *int          `json:"order_file"`
	PassedToFinance       bool          `json:"passed_to_finance"`
	UsedInFinance         bool          `json:"used_in_finance"`
	IsProFormaInvoice     bool          `json:"is_pro_forma_invoice"`
	ReceiveFile           pq.Int64Array `json:"receive_file"`
	MovementFile          *int          `json:"movement_file"`
}

type OrderListResponseDTO struct {
	ID                    int        `json:"id"`
	DateOrder             time.Time  `json:"date_order" validate:"required"`
	TotalPrice            float32    `json:"total_price"`
	PublicProcurementID   *int       `json:"public_procurement_id"`
	SupplierID            *int       `json:"supplier_id"`
	Status                string     `json:"status"`
	GroupOfArticlesID     *int       `json:"group_of_articles_id"`
	IsUsed                bool       `json:"is_used"`
	DateSystem            *time.Time `json:"date_system"`
	InvoiceDate           *time.Time `json:"invoice_date"`
	ProFormaInvoiceDate   *time.Time `json:"pro_forma_invoice_date"`
	ProFormaInvoiceNumber string     `json:"pro_forma_invoice_number"`
	InvoiceNumber         *string    `json:"invoice_number"`
	OrganizationUnitID    int        `json:"organization_unit_id"`
	OfficeID              *int       `json:"office_id"`
	AccountID             *int       `json:"account_id"`
	RecipientUserID       *int       `json:"recipient_user_id"`
	Description           *string    `json:"description"`
	PassedToFinance       bool       `json:"passed_to_finance"`
	UsedInFinance         bool       `json:"used_in_finance"`
	IsProFormaInvoice     bool       `json:"is_pro_forma_invoice"`
	OrderFile             *int       `json:"order_file"`
	ReceiveFile           []int      `json:"receive_file"`
	MovementFile          *int       `json:"movement_file"`
	CreatedAt             time.Time  `json:"created_at"`
	UpdatedAt             time.Time  `json:"updated_at"`
}

func (dto OrderListDTO) ToOrderList() *data.OrderList {

	isProFormaInvoice := true

	if dto.InvoiceNumber != nil && *dto.InvoiceNumber != "" {
		isProFormaInvoice = false
	}

	return &data.OrderList{
		DateOrder:             dto.DateOrder,
		TotalPrice:            dto.TotalPrice,
		PublicProcurementID:   dto.PublicProcurementID,
		SupplierID:            dto.SupplierID,
		Status:                dto.Status,
		DateSystem:            dto.DateSystem,
		InvoiceDate:           dto.InvoiceDate,
		InvoiceNumber:         dto.InvoiceNumber,
		ProFormaInvoiceDate:   dto.ProFormaInvoiceDate,
		ProFormaInvoiceNumber: dto.ProFormaInvoiceNumber,
		GroupOfArticlesID:     dto.GroupOfArticlesID,
		OrganizationUnitID:    dto.OrganizationUnitID,
		OfficeID:              dto.OfficeID,
		AccountID:             dto.AccountID,
		IsUsed:                dto.IsUsed,
		RecipientUserID:       dto.RecipientUserID,
		Description:           dto.Description,
		PassedToFinance:       dto.PassedToFinance,
		UsedInFinance:         dto.UsedInFinance,
		IsProFormaInvoice:     isProFormaInvoice,
		OrderFile:             dto.OrderFile,
		ReceiveFile:           dto.ReceiveFile,
		MovementFile:          dto.MovementFile,
	}
}

func ToOrderListResponseDTO(data data.OrderList) OrderListResponseDTO {

	array := make([]int, len(data.ReceiveFile))
	for i, id := range data.ReceiveFile {
		array[i] = int(id)
	}

	isProFormaInvoice := true

	if data.InvoiceNumber != nil && *data.InvoiceNumber != "" {
		isProFormaInvoice = false
	}

	return OrderListResponseDTO{
		ID:                    data.ID,
		DateOrder:             data.DateOrder,
		TotalPrice:            data.TotalPrice,
		PublicProcurementID:   data.PublicProcurementID,
		SupplierID:            data.SupplierID,
		AccountID:             data.AccountID,
		Status:                data.Status,
		IsUsed:                data.IsUsed,
		DateSystem:            data.DateSystem,
		InvoiceDate:           data.InvoiceDate,
		ProFormaInvoiceDate:   data.ProFormaInvoiceDate,
		ProFormaInvoiceNumber: data.ProFormaInvoiceNumber,
		GroupOfArticlesID:     data.GroupOfArticlesID,
		InvoiceNumber:         data.InvoiceNumber,
		OrganizationUnitID:    data.OrganizationUnitID,
		OfficeID:              data.OfficeID,
		RecipientUserID:       data.RecipientUserID,
		Description:           data.Description,
		PassedToFinance:       data.PassedToFinance,
		UsedInFinance:         data.UsedInFinance,
		IsProFormaInvoice:     isProFormaInvoice,
		OrderFile:             data.OrderFile,
		ReceiveFile:           array,
		MovementFile:          data.MovementFile,
		CreatedAt:             data.CreatedAt,
		UpdatedAt:             data.UpdatedAt,
	}
}

func ToOrderListListResponseDTO(contracts []*data.OrderList) []OrderListResponseDTO {
	dtoList := make([]OrderListResponseDTO, len(contracts))
	for i, x := range contracts {
		dtoList[i] = ToOrderListResponseDTO(*x)
	}
	return dtoList
}
