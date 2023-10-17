package controllers

import (
	"beedemo/models"

	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}

func (c *UserController) Index() {

	c.TplName = "user.tpl"
}
func (c *UserController) List() {
	//查询数据库
	userList := []models.User{}
	models.DB.Find(&userList)
	c.Data["json"] = userList
	c.ServeJSON()
	c.TplName = "userList.tpl"

}
func (c *UserController) AddUser() {
	//增加数据
	user := models.User{
		UserName: "wangwu",
		Age:      22,
		Email:    "222@qq.con",
		AddTime:  int(models.GetUnix()),
	}
	models.DB.Create(&user)
	c.TplName = "userAdd.tpl"
}

func (c *UserController) EditUser() {
	//保存所有字段

	// 查询id等于6的数据
	// user := models.User{Id: 6}
	// models.DB.Find(&user)
	// //更新数据
	// user.Username = "哈哈"
	// user.Email = "itying@qqq.com"
	// user.AddTime = int(models.GetUnix())
	// models.DB.Save(&user)

	//更新单个列
	// user := models.User{}
	// models.DB.Model(&user).Where("id = ?", 6).Update("username", "哈哈哈哈哈哈")

	//另一种更新数据的方法
	user := models.User{}
	models.DB.Where("id = ?", 4).Find(&user)
	user.UserName = "KONG111"
	user.Email = "aaa@qqq.com"
	user.AddTime = int(models.GetUnix())
	models.DB.Save(&user)
	c.TplName = "userEdit.tpl"

}

func (c *UserController) DeleteUser() {

	// 删除id=6的数据
	user := models.User{Id: 6}
	models.DB.Delete(&user)

	//删除username=gorm的数据
	// user := models.User{}
	// models.DB.Where("username = ?", "gorm").Delete(&user)
	c.TplName = "userDelete.tpl"
}
