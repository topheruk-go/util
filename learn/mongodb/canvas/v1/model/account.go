package model

type Account struct {
	ID string `json:"id,omitempty"`
	// if blank then Account is the root
	ParentAccountID string `json:"parent_account_id,omitempty"`
	Name            string `json:"name,omitempty"`
	// "active" or "deleted"
	Status string `json:"status,omitempty"`
	// optional?
	IntegrationID string `json:"integration_id,omitempty"`
}

// Root -- Fashion Retail Academy (FRA001)
// 		Sub -- Higher Educdation (FRA012)
// 			Sub -- Level 6 (FRA106)
// 			Sub -- Level 4 (FRA104)
// 		Sub -- Further Education (FRA011)
// 			Sub -- Level 3 (FRA103)
// 			Sub -- Level 2 (FRA102)
