package ucv1wallethttp

type Wallet__Reqp struct {
	OldBranchID       uint16 `json:"old_branch_id" form:"old_branch_id" xml:"old_branch_id"`                   /// PlayID of transaction. Optional
	NewBranchID       uint16 `json:"new_branch_id" form:"new_branch_id" xml:"new_branch_id"`                   /// PlayID of transaction. Optional
	Currency          string `json:"currency" form:"currency" xml:"currency"`                                  /// PlayID of transaction. Optional
	FileNameMigration string `json:"file_name_migration" form:"file_name_migration" xml:"file_name_migration"` /// PlayID of transaction. Optional
}

type Wallet_PID__Reqp struct {
	PlayIds []string `json:"play_ids" form:"play_ids" xml:"play_ids" validate:"required"` /// PlayID of transaction. Optional
}

type Check__Csv__Reqp struct {
	OldBranchID       uint16 `json:"old_branch_id" form:"old_branch_id" xml:"old_branch_id"`                   /// PlayID of transaction. Optional
	NewBranchID       uint16 `json:"new_branch_id" form:"new_branch_id" xml:"new_branch_id"`                   /// PlayID of transaction. Optional
	FileNameMigration string `json:"file_name_migration" form:"file_name_migration" xml:"file_name_migration"` /// PlayID of transaction. Optional
	FileNameOriginal  string `json:"file_name_original" form:"file_name_original" xml:"file_name_original"`    /// PlayID of transaction. Optional

	Wallet string `json:"wallet" form:"wallet" xml:"wallet"` /// PlayID of transaction. Optional
}
