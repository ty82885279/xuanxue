package model

type ParamStar struct {
	Star string `json:"star" form:"star" binding:"required" example:"baiyang" `    // 星座(必填)
	Date string `json:"date" form:"date" binding:"required" example:"2020-11-11" ` // 日期(必填)
	Type string `json:"type" form:"type" example:""`                               //时间(选填) 默认获取5个时间信息
}
type ParamAuspiciousList struct {
	Type    string `json:"type" form:"type" binding:"required" example:"yi"`   // 类型(必填) yi:宜,ji:忌
	Name    string `json:"name" form:"name" binding:"required" example:"结婚" `  // 名称(必填)
	Dur     string `json:"dur" form:"dur" binding:"required" example:"month" ` // 时间范围(必填) week:1周, month:1月, month3:3月, year:1年
	Weekend bool   `json:"weekend" form:"weekend" example:"true"`              // 是否只看周末(选填) 1:是，0:否，默认:0
	Page    int64  `json:"page" form:"page" example:"1"`                       // 页数(选填)    默认:1
	Size    int64  `json:"size" form:"size" example:"20"`                      // 条数(选填)    默认:20

}
