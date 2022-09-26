// struct demo
// Code from https://github.com/LinkedInLearning/go-essential-training-2446262/blob/main/Ch04/04_02/budget.go
package main

import (
	"fmt"
	"time"
)

// Budget is for a campaign
type Budget struct {
	CampaignID string
	Balance    float64 // USD
	Expires    time.Time
}

func (b Budget) TimeLeft() time.Duration {
	return b.Expires.Sub(time.Now().UTC())
}

func (b *Budget) Update(sum float64) {
	b.Balance += sum
}

func main() {
	b := Budget{"Kittens", 22.3, time.Now().Add(7 * 24 * time.Hour)}
	fmt.Println(b.TimeLeft())

	b.Update(10.5)
	fmt.Println(b.Balance)

}
