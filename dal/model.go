package dal

type Model struct {
	Id        int64 `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	CreatedAt int64 `gorm:"not null" json:"created_at"`
	UpdatedAt int64 `gorm:"not null" json:"updated_at"`
	DeletedAt int64 `gorm:"not null" json:"deleted_at"`
}

type Status int

const (
	StatusApply   Status = iota + 1 // 待审核
	StatusSuccess                   // 审核通过
	StatusRefuse                    // 审核拒绝
)

type Sure int

const (
	SureYes Sure = iota + 1 // 是
	SureNo                  // 否
)
