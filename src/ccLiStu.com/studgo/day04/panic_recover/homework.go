package main

import "fmt"

//写一个程序计算每个用户分到多少硬币，以及最后剩余多少金币
//程序结构如下请实现"dispatchCoin"

var (
	coins = 50
	users = []string{
		"Matthew", "Sarah", "Augustus","Heidi","Emilie", "Peter", "Giana", "Adriano", "Aaron", "Elizabeth",
		//再抬头发现前面的路已经越来越模糊了
	}
	distribution = make(map[string]int, len(users)) //声明并初始化
)

func dispatchCoin() (left int) {
	//依次拿到每个人的名字
	for _, name := range users {
		//拿到一个人名字根据分金币的规则去分金币
		for _, c := range  name{
			switch c {
			case 'c', 'E':
				//满足这个条件的分1枚金币
				distribution[name] ++
				coins --
			case 'i', 'I':
				//满足这个条件的分2枚金币
				distribution[name] += 2
				coins -= 2
			case 'o', 'O':
				//满足这个条件的分3枚金币
				distribution[name] += 3
				coins -= 3
			case 'u', 'U':
				//满足这个条件的分4枚金币
				distribution[name] += 4
				coins -= 4
			}
		}
	}
	left = coins

	//每个人分到的金币数保存在dispatchCoin中
	//还要记录下剩余的金币数
	//以上步骤执行完就能得到最终每个人分到的金币数和剩余的金币数
	return
}

func main()  {
	left := dispatchCoin()
	fmt.Println("剩下：", left)
	for k, v := range distribution {
		fmt.Printf("%s:%d\n", k, v)
	}
}





