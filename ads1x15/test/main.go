package main

import (
	"context"
	"log"
	"time"

	"github.com/Cr4psy/indoorGarden/program/lib/ads1x15"

	"github.com/d2r2/go-i2c"
)

func main() {
	i2c, err := i2c.NewI2C(uint8(ads1x15.ADS1x15_DEFAULT_ADDRESS), 1)
	if err != nil {
		log.Fatal(err)
	}
	defer i2c.Close()

	adc := ads1x15.ADC{
		Conn:   i2c,
		Config: ads1x15.DefaultConfig(),
	}

	t := time.NewTicker(time.Second)
	ctx := context.Background()

	for {
		select {
		case <-t.C:
			c0, err := adc.Read(0)
			if err != nil {
				log.Fatal(err)
			}
			c1, err := adc.Read(1)
			if err != nil {
				log.Fatal(err)
			}
			c2, err := adc.Read(2)
			if err != nil {
				log.Fatal(err)
			}
			c3, err := adc.Read(3)
			if err != nil {
				log.Fatal(err)
			}

			log.Printf("%v \t %v \t %v \t %v \n", c0, c1, c2, c3)

		case <-ctx.Done():
			log.Fatal(ctx.Err())
		}
	}
}
