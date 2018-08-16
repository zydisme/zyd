package model



type Tops struct{
	Id			int
	LastRank 	int	//去年排名
	Rank		int	//今年排名
	Name		string //公司名
	Income		string	//营业收入 百万美元
	Profits		string	//利润 百万美元
	Country 	string	//国家
}