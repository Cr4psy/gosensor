package ads1x15

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/d2r2/go-i2c"
)

type ADC struct {
	Conn   *i2c.I2C
	Config *Config
}

func NewADC(c *i2c.I2C, cfg *Config) *ADC {
	if cfg == nil {
		cfg = DefaultConfig()
	}
	return &ADC{
		Conn:   c,
		Config: cfg,
	}
}
func (a *ADC) writeConfig(channel int) error {
	channel = channel + 0x04

	g := a.Config.ConvertGain(a.Config.Gain)
	if g == 0 {
		return fmt.Errorf("Incorrect gain, must be one of 2/3, 1, 2, 4, 8, 16")
	}
	sr := a.Config.ConvertSampleRate(a.Config.SampeRate)
	if g == 0 {
		return fmt.Errorf("Incorrect sample rate")
	}
	var config int
	// Go out of power-down mode for conversion.
	config = ADS1x15_CONFIG_OS_SINGLE
	config = config | ((channel & 0x07) << ADS1x15_CONFIG_MUX_OFFSET)
	config = config | g
	// Set the mode (continuous or single shot).
	config = config | a.Config.Mode.Int()
	// Set the data rate (this is controlled by the subclass as it differs
	// between ADS1015 and ADS1115).
	config = config | sr
	// Disble comparator mode.
	config = config | ADS1x15_CONFIG_COMP_QUE_DISABLE

	log.Printf("Write config %v to %d", strconv.FormatInt(int64(config), 2), byte(ADS1x15_POINTER_CONFIG))
	return a.Conn.WriteRegS16BE(byte(ADS1x15_POINTER_CONFIG), int16(config))
}

func (a *ADC) Read(channel int) (int, error) {

	if err := a.writeConfig(channel); err != nil {
		return 0, err
	}
	time.Sleep(time.Second/time.Duration(a.Config.SampeRate) + 5*time.Millisecond)

	v, err := a.Conn.ReadRegS16BE(byte(ADS1x15_POINTER_CONVERSION))
	if err != nil {
		return 0, err
	}
	return int(v), nil
}
