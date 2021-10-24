package entity

import "time"

type UFFLRequest struct {
	Dir    string `json:"dir"`
	Filter struct {
		CutoffFrom         time.Time `json:"cutoff_from"`
		CutoffTo           time.Time `json:"cutoff_to"`
		DeliveringDateFrom time.Time `json:"delivering_date_from"`
		DeliveringDateTo   time.Time `json:"delivering_date_to"`
		DeliveryMethodID   []int     `json:"delivery_method_id"`
		ProviderID         []int     `json:"provider_id"`
		Status             string    `json:"status"`
		WarehouseID        []int     `json:"warehouse_id"`
	} `json:"filter"`
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
	With   struct {
		AnalyticsData    bool `json:"analytics_data"`
		Barcodes         bool `json:"barcodes"`
		FinancialData    bool `json:"financial_data"`
		ProductExemplars bool `json:"product_exemplars"`
		Translit         bool `json:"translit"`
	} `json:"with"`
}

type Posting struct {
	Addressee struct {
		Name  string `json:"name"`
		Phone string `json:"phone"`
	} `json:"addressee"`
	AnalyticsData struct {
		City                 string    `json:"city"`
		DeliveryDateBegin    time.Time `json:"delivery_date_begin"`
		DeliveryDateEnd      time.Time `json:"delivery_date_end"`
		DeliveryType         string    `json:"delivery_type"`
		IsLegal              bool      `json:"is_legal"`
		IsPremium            bool      `json:"is_premium"`
		PaymentTypeGroupName string    `json:"payment_type_group_name"`
		Region               string    `json:"region"`
		TplProvider          string    `json:"tpl_provider"`
		TplProviderID        int       `json:"tpl_provider_id"`
		Warehouse            string    `json:"warehouse"`
		WarehouseID          int       `json:"warehouse_id"`
	} `json:"analytics_data"`
	Barcodes struct {
		LowerBarcode string `json:"lower_barcode"`
		UpperBarcode string `json:"upper_barcode"`
	} `json:"barcodes"`
	Cancellation struct {
		AffectCancellationRating bool   `json:"affect_cancellation_rating"`
		CancelReason             string `json:"cancel_reason"`
		CancelReasonID           int    `json:"cancel_reason_id"`
		CancellationInitiator    string `json:"cancellation_initiator"`
		CancellationType         string `json:"cancellation_type"`
		CancelledAfterShip       bool   `json:"cancelled_after_ship"`
	} `json:"cancellation"`
	Customer struct {
		Address struct {
			AddressTail     string `json:"address_tail"`
			City            string `json:"city"`
			Comment         string `json:"comment"`
			Country         string `json:"country"`
			District        string `json:"district"`
			Latitude        int    `json:"latitude"`
			Longitude       int    `json:"longitude"`
			ProviderPvzCode string `json:"provider_pvz_code"`
			PvzCode         int    `json:"pvz_code"`
			Region          string `json:"region"`
			ZipCode         string `json:"zip_code"`
		} `json:"address"`
		CustomerEmail string `json:"customer_email"`
		CustomerID    int    `json:"customer_id"`
		Name          string `json:"name"`
		Phone         string `json:"phone"`
	} `json:"customer"`
	DeliveringDate time.Time `json:"delivering_date"`
	DeliveryMethod struct {
		ID            int    `json:"id"`
		Name          string `json:"name"`
		TplProvider   string `json:"tpl_provider"`
		TplProviderID int    `json:"tpl_provider_id"`
		Warehouse     string `json:"warehouse"`
		WarehouseID   int    `json:"warehouse_id"`
	} `json:"delivery_method"`
	FinancialData struct {
		PostingServices struct {
			MarketplaceServiceItemDelivToCustomer            int `json:"marketplace_service_item_deliv_to_customer"`
			MarketplaceServiceItemDirectFlowTrans            int `json:"marketplace_service_item_direct_flow_trans"`
			MarketplaceServiceItemDropoffFf                  int `json:"marketplace_service_item_dropoff_ff"`
			MarketplaceServiceItemDropoffPvz                 int `json:"marketplace_service_item_dropoff_pvz"`
			MarketplaceServiceItemDropoffSc                  int `json:"marketplace_service_item_dropoff_sc"`
			MarketplaceServiceItemFulfillment                int `json:"marketplace_service_item_fulfillment"`
			MarketplaceServiceItemPickup                     int `json:"marketplace_service_item_pickup"`
			MarketplaceServiceItemReturnAfterDelivToCustomer int `json:"marketplace_service_item_return_after_deliv_to_customer"`
			MarketplaceServiceItemReturnFlowTrans            int `json:"marketplace_service_item_return_flow_trans"`
			MarketplaceServiceItemReturnNotDelivToCustomer   int `json:"marketplace_service_item_return_not_deliv_to_customer"`
			MarketplaceServiceItemReturnPartGoodsCustomer    int `json:"marketplace_service_item_return_part_goods_customer"`
		} `json:"posting_services"`
		Products []struct {
			Actions           []string `json:"actions"`
			ClientPrice       string   `json:"client_price"`
			CommissionAmount  int      `json:"commission_amount"`
			CommissionPercent int      `json:"commission_percent"`
			ItemServices      struct {
				MarketplaceServiceItemDelivToCustomer            int `json:"marketplace_service_item_deliv_to_customer"`
				MarketplaceServiceItemDirectFlowTrans            int `json:"marketplace_service_item_direct_flow_trans"`
				MarketplaceServiceItemDropoffFf                  int `json:"marketplace_service_item_dropoff_ff"`
				MarketplaceServiceItemDropoffPvz                 int `json:"marketplace_service_item_dropoff_pvz"`
				MarketplaceServiceItemDropoffSc                  int `json:"marketplace_service_item_dropoff_sc"`
				MarketplaceServiceItemFulfillment                int `json:"marketplace_service_item_fulfillment"`
				MarketplaceServiceItemPickup                     int `json:"marketplace_service_item_pickup"`
				MarketplaceServiceItemReturnAfterDelivToCustomer int `json:"marketplace_service_item_return_after_deliv_to_customer"`
				MarketplaceServiceItemReturnFlowTrans            int `json:"marketplace_service_item_return_flow_trans"`
				MarketplaceServiceItemReturnNotDelivToCustomer   int `json:"marketplace_service_item_return_not_deliv_to_customer"`
				MarketplaceServiceItemReturnPartGoodsCustomer    int `json:"marketplace_service_item_return_part_goods_customer"`
			} `json:"item_services"`
			OldPrice int `json:"old_price"`
			Payout   int `json:"payout"`
			Picking  struct {
				Amount int       `json:"amount"`
				Moment time.Time `json:"moment"`
				Tag    string    `json:"tag"`
			} `json:"picking"`
			Price                int `json:"price"`
			ProductID            int `json:"product_id"`
			Quantity             int `json:"quantity"`
			TotalDiscountPercent int `json:"total_discount_percent"`
			TotalDiscountValue   int `json:"total_discount_value"`
		} `json:"products"`
	} `json:"financial_data"`
	InProcessAt   time.Time `json:"in_process_at"`
	IsExpress     bool      `json:"is_express"`
	OrderID       int       `json:"order_id"`
	OrderNumber   string    `json:"order_number"`
	PostingNumber string    `json:"posting_number"`
	Products      []struct {
		MandatoryMark []string `json:"mandatory_mark"`
		Name          string   `json:"name"`
		OfferID       string   `json:"offer_id"`
		Price         string   `json:"price"`
		Quantity      int      `json:"quantity"`
		Sku           int      `json:"sku"`
	} `json:"products"`
	Requirements struct {
		ProductsRequiringGtd     []string `json:"products_requiring_gtd"`
		ProductsRequiringCountry []string `json:"products_requiring_country"`
	} `json:"requirements"`
	ShipmentDate       time.Time `json:"shipment_date"`
	Status             string    `json:"status"`
	TplIntegrationType string    `json:"tpl_integration_type"`
	TrackingNumber     string    `json:"tracking_number"`
}

type UFFLResponse struct {
	Result struct {
		Count    int `json:"count"`
		Postings []Posting `json:"postings"`
	} `json:"result"`
}