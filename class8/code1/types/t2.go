package types

import (
	"fmt"
	"math/rand"
)

// 方法和接口的混合使用
// 接口，只要拥有play的方法就可以
type PlayGame interface {
	Play() string
}

// 定义一个类型
// 公用的方法Play
// 需要重定义的内容 GetScord
type Game struct {
	name     string
	GetScord func() int
}

// 公用方法
func (self Game) Play() string {
	s := fmt.Sprintf("%s ,scord :%v", self.name, self.GetScord())
	return s
}

// 继承Game 内容
type LOL struct {
	Game
}

func NewLOL() LOL {
	g := LOL{}
	g.name = "LOL"
	g.Game.GetScord = g.GetScord // 关键
	return g
}

// 重写GetScord
func (self LOL) GetScord() int {
	return rand.Intn(100) + 1
}

// 继承Game
type DOTA struct {
	Game
}

func NewDOTA() DOTA {
	g := DOTA{}
	g.name = "DOTA"
	g.Game.GetScord = g.GetScord
	return g
}

// 重写GetScord
func (self DOTA) GetScord() int {
	return rand.Intn(100) + 1
}
