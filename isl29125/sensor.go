package light_sensor_go

import (
	"github.com/d2r2/go-i2c"
)

type Sensor struct {
	Conn *i2c.I2C
}

func (s *Sensor) SetConfig() error {
	// Config 1
	if err := s.Conn.WriteRegU8(ISL_CONFIG_1, ISL_CFG1_MODE_RGB|ISL_CFG1_10KLUX); err != nil {
		return err
	}
	// Config 2
	if err := s.Conn.WriteRegU8(ISL_CONFIG_2, ISL_CFG2_IR_ADJUST_HIGH); err != nil {
		return err
	}
	// Config 3
	if err := s.Conn.WriteRegU8(ISL_CONFIG_3, ISL_CFG_DEFAULT); err != nil {
		return err
	}
	return nil
}

func (s *Sensor) ReadRed() (int, error) {
	v, err := s.Conn.ReadRegU16LE(ISL_RED_L)
	if err != nil {
		return 0, err
	}
	return int(v), nil
}

func (s *Sensor) ReadBlue() (int, error) {
	v, err := s.Conn.ReadRegU16LE(ISL_BLUE_L)
	if err != nil {
		return 0, err
	}
	return int(v), nil
}

func (s *Sensor) ReadGreen() (int, error) {
	v, err := s.Conn.ReadRegU16LE(ISL_GREEN_L)
	if err != nil {
		return 0, err
	}
	return int(v), nil
}
