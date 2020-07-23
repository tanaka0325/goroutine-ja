package main

import (
	"fmt"
	"time"
)

type (
	Bean       int
	GroundBean int
	Water      int
	HotWater   int
	Coffee     int
)

const (
	GramBeans          Bean       = 1
	GramGroundBeans    GroundBean = 1
	MilliLiterWater    Water      = 1
	MilliLiterHotWater HotWater   = 1
	CupsCoffee         Coffee     = 1
)

func (w Water) String() string {
	return fmt.Sprintf("%d[ml] water", int(w))
}

func (hw HotWater) String() string {
	return fmt.Sprintf("%d[ml] hot water", int(hw))
}

func (b Bean) String() string {
	return fmt.Sprintf("%d[g] beans", int(b))
}

func (gb GroundBean) String() string {
	return fmt.Sprintf("%d[g] ground beans", int(gb))
}

func (cups Coffee) String() string {
	return fmt.Sprintf("%d cup(s) coffee", int(cups))
}

// 1カップのコーヒーを淹れるのに必要な水の量
func (cups Coffee) Water() Water {
	return Water(180*cups) / MilliLiterWater
}

// 1カップのコーヒーを淹れるのに必要なお湯の量
func (cups Coffee) HotWater() HotWater {
	return HotWater(180*cups) / MilliLiterHotWater
}

// 1カップのコーヒーを淹れるのに必要な豆の量
func (cups Coffee) Beans() Bean {
	return Bean(20*cups) / GramBeans
}

// 1カップのコーヒーを淹れるのに必要な粉の量
func (cups Coffee) GroundBeans() GroundBean {
	return GroundBean(20*cups) / GramGroundBeans
}

// お湯を沸かす
func boil(water Water) HotWater {
	time.Sleep(400 * time.Millisecond)
	return HotWater(water)
}

// コーヒー豆を挽く
func grind(beans Bean) GroundBean {
	time.Sleep(200 * time.Millisecond)
	return GroundBean(beans)
}

// コーヒーを淹れる
func brew(hotWater HotWater, groundBeans GroundBean) Coffee {
	time.Sleep(1 * time.Second)
	// 少ない方を優先する
	cups1 := Coffee(hotWater / (1 * CupsCoffee).HotWater())
	cups2 := Coffee(groundBeans / (1 * CupsCoffee).GroundBeans())
	if cups1 < cups2 {
		return cups1
	}
	return cups2
}

func main() {
	// 作るコーヒーの数
	const amountCoffee = 20 * CupsCoffee

	// limits
	const (
		LimitOfBoiledWater  = 600 * MilliLiterWater
		LimitOfGroundBeans  = 20 * GramBeans
		LimitOfBrewedCoffee = 20 * CupsCoffee
	)

	// 材料
	water := amountCoffee.Water()
	beans := amountCoffee.Beans()

	fmt.Println(water)
	fmt.Println(beans)

	// お湯を沸かす
	var hotWater HotWater
	for water > 0 {
		water -= LimitOfBoiledWater
		hotWater += boil(LimitOfBoiledWater)
	}
	fmt.Println(hotWater)

	// 豆を挽く
	var groundBeans GroundBean
	for beans > 0 {
		beans -= LimitOfGroundBeans
		groundBeans += grind(LimitOfGroundBeans)
	}
	fmt.Println(groundBeans)

	// コーヒーを淹れる
	var coffee Coffee
	cups := LimitOfBrewedCoffee * CupsCoffee
	for hotWater >= cups.HotWater() && groundBeans >= cups.GroundBeans() {
		hotWater -= cups.HotWater()
		groundBeans -= cups.GroundBeans()

		coffee += brew(cups.HotWater(), cups.GroundBeans())
	}

	fmt.Println(coffee)
}
