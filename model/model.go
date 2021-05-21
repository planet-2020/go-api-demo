/**
 * @Description: 基础模型
 * @author zhouhongpan
 * @date 2021/5/19 11:39
 */
package model

type CustomModel struct {
	ID uint `gorm:"primaryKey;autoIncrement"`
	CreateTime uint `gorm:"autoCreateTime"`
	UpdateTime uint `gorm:"autoUpdateTime"`
	DeleteTime uint
}