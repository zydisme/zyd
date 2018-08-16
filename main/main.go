package main

import (
	"log"
	"strings"
	"fmt"
	"github.com/go-xorm/xorm"
	"strconv"
	"github.com/PuerkitoBio/goquery"
	_ "github.com/go-sql-driver/mysql"
	"qh/model"
)
func GetTops(url string ) []model.Tops {

	tops := make([]model.Tops, 0, 500)
	info:=&model.Tops{}
	doc, err:=goquery.NewDocument(url)
	if err!=nil{
		log.Fatal(err)
	}
	doc.Find("#yytable tr").Each(func(i int, s *goquery.Selection) {
		if i !=0{
			data:=strings.Split(s.Text(),"\n")
			for range data{
				info.Id,_ = strconv.Atoi(data[1])
				info.Rank,_ =strconv.Atoi(data[1])
				info.LastRank,_ =strconv.Atoi(data[2])
				info.Name =data[3]
				info.Income =data[4]
				info.Profits =data[5]
				info.Country =data[6]
			}
			tops = append(tops,*info)
		}
	})

	for _,vals:=range tops{

		fmt.Printf("去年排名: %s\n",vals.LastRank)
		fmt.Printf("今年排名: %s\n",vals.Rank)
		fmt.Printf("公司名: %s\n",vals.Name)
		fmt.Printf("营业收入: %s\n",vals.Income)
		fmt.Printf("利润: %s\n",vals.Profits)
		fmt.Printf("国家: %s\n\n",vals.Country)
	}
	return tops

}
func InsertTable(engine *xorm.Engine , datas []model.Tops,i int)  {
	row, err := engine.Insert(datas)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("instert %d %d datas total %d\n",i,i+99 ,row)
}
func main()  {
	url:="http://www.fortunechina.com/fortune500/c/2018-07/19/content_311046.htm"
	//获取tops500信息
	tops:=GetTops(url)
	engine, err := xorm.NewEngine("mysql", "root:zydqwqw123@/zyd?charset=utf8")
	if err != nil {
		log.Fatal(err)
	}
	err = engine.Sync2(new(model.Tops))    // 同步表结构
	if err != nil {
		log.Fatal(err)
	}

	//批量导入数据库  分五次每次导入100条
	for i:=0;i<500;i=i+100{
		datas := tops[i:i+100]
		 InsertTable(engine,datas,i)
	}
	}