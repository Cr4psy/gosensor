package light_sensor_go

const (
	// ISL29125 I2C Address
	ISL_I2C_ADDR = 0x44

	// ISL29125 I2C Device ID Value
	ISL_DEVICE_ID_VAL = 0x7D

	// ISL29125 Registers
	ISL_DEVICE_ID    = 0x00
	ISL_CONFIG_1     = 0x01
	ISL_CONFIG_2     = 0x02
	ISL_CONFIG_3     = 0x03
	ISL_THRESHOLD_LL = 0x04
	ISL_THRESHOLD_LH = 0x05
	ISL_THRESHOLD_HL = 0x06
	ISL_THRESHOLD_HH = 0x07
	ISL_STATUS       = 0x08
	ISL_GREEN_L      = 0x09
	ISL_GREEN_H      = 0x0A
	ISL_RED_L        = 0x0B
	ISL_RED_H        = 0x0C
	ISL_BLUE_L       = 0x0D
	ISL_BLUE_H       = 0x0E

	// Configuration Settings
	ISL_CFG_DEFAULT = 0x00

	// CONFIG1
	// Pick a mode, determines what color[s] the sensor samples, if any
	ISL_CFG1_MODE_POWERDOWN = 0x00
	ISL_CFG1_MODE_G         = 0x01
	ISL_CFG1_MODE_R         = 0x02
	ISL_CFG1_MODE_B         = 0x03
	ISL_CFG1_MODE_STANDBY   = 0x04
	ISL_CFG1_MODE_RGB       = 0x05
	ISL_CFG1_MODE_RG        = 0x06
	ISL_CFG1_MODE_GB        = 0x07

	// Light intensity range
	// In a dark environment 375Lux is best, otherwise 10KLux is likely the best option
	ISL_CFG1_375LUX = 0x00
	ISL_CFG1_10KLUX = 0x08

	// Change this to 12 bit if you want less accuracy, but faster sensor reads
	// At default 16 bit, each sensor sample for a given color is about ~100ms
	ISL_CFG1_16BIT = 0x00
	ISL_CFG1_12BIT = 0x10

	// Unless you want the interrupt pin to be an input that triggers sensor sampling, leave this on normal
	ISL_CFG1_ADC_SYNC_NORMAL = 0x00
	ISL_CFG1_ADC_SYNC_TO_INT = 0x20

	// Sets amount of IR filtering, can use these presets or any value between = 0x00 and = 0x3F
	// Consult datasheet for detailed IR filtering calibration
	ISL_CFG2_IR_ADJUST_LOW  = 0x00
	ISL_CFG2_IR_ADJUST_MID  = 0x20
	ISL_CFG2_IR_ADJUST_HIGH = 0x3F

	// CONFIG3
	// No interrupts, or interrupts based on a selected color
	ISL_CFG3_NO_INT = 0x00
	ISL_CFG3_G_INT  = 0x01
	ISL_CFG3_R_INT  = 0x02
	ISL_CFG3_B_INT  = 0x03

	// How many times a sensor sample must hit a threshold before triggering an interrupt
	// More consecutive samples means more times between interrupts, but less triggers from short transients
	ISL_CFG3_INT_PRST1 = 0x00
	ISL_CFG3_INT_PRST2 = 0x04
	ISL_CFG3_INT_PRST4 = 0x08
	ISL_CFG3_INT_PRST8 = 0x0C

	// If you would rather have interrupts trigger when a sensor sampling is complete, enable this
	// If this is disabled, interrupts are based on comparing sensor data to threshold settings
	ISL_CFG3_RGB_CONV_TO_INT_DISABLE = 0x00
	ISL_CFG3_RGB_CONV_TO_INT_ENABLE  = 0x10

	// STATUS FLAG MASKS
	ISL_FLAG_INT       = 0x01
	ISL_FLAG_CONV_DONE = 0x02
	ISL_FLAG_BROWNOUT  = 0x04
	ISL_FLAG_CONV_G    = 0x10
	ISL_FLAG_CONV_R    = 0x20
	ISL_FLAG_CONV_B    = 0x30
)

// DefaultAddress return the default i2c address
func DefaultAddress() uint8 {
	return ISL_I2C_ADDR
}
