package main

import (
	"context"
	"log"
	"time"

	isl "github.com/Cr4psy/indoorGarden/program/lib/isl29125"

	"github.com/d2r2/go-i2c"
)

func main() {
	i2c, err := i2c.NewI2C(0x44, 1)
	if err != nil {
		log.Fatal(err)
	}
	defer i2c.Close()

	ls := &isl.Sensor{
		Conn: i2c,
	}
	if err := ls.SetConfig(); err != nil {
		log.Fatal(err)
	}

	t := time.NewTicker(time.Second)
	ctx := context.Background()

	for {
		select {
		case <-t.C:
			if err := readSensor(ls); err != nil {
				log.Fatal(err)
			}

		case <-ctx.Done():
			log.Fatal(ctx.Err())
		}
	}

}

func readSensor(ls *isl.Sensor) error {
	r, err := ls.ReadRed()
	if err != nil {
		return err
	}
	g, err := ls.ReadGreen()
	if err != nil {
		return err
	}
	b, err := ls.ReadBlue()
	if err != nil {
		return err
	}

	log.Printf("R: %v \t G: %v \t B: %v", r, g, b)
	return nil
}
