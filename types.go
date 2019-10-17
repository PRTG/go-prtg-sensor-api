// Copyright (c) 2019, Paessler AG.
// All Rights Reserved
package prtg

type UnitType string

const (
	BytesBandwidth UnitType = "BytesBandwidth"
	BytesDisk      UnitType = "BytesDisk"
	Temperature    UnitType = "Temperature"
	Percent        UnitType = "Percent"
	TimeResponse   UnitType = "TimeResponse"
	TimeSeconds    UnitType = "TimeSeconds"
	Custom         UnitType = "Custom"
	Count          UnitType = "Count"
	CPU            UnitType = "CPU"
	BytesFile      UnitType = "BytesFile"
	SpeedDisk      UnitType = "SpeedDisk"
	SpeedNet       UnitType = "SpeedNet"
	TimeHours      UnitType = "TimeHours"
)

type SizeType string

const (
	One      SizeType = "One"
	Kilo     SizeType = "Kilo"
	Mega     SizeType = "Mega"
	Giga     SizeType = "Giga"
	Tera     SizeType = "Tera"
	Byte     SizeType = "Byte"
	KiloByte SizeType = "KiloByte"
	MegaByte SizeType = "MegaByte"
	GigaByte SizeType = "GigaByte"
	TeraByte SizeType = "TeraByte"
)

type SpeedType string

const (
	Bit     SpeedType = "Bit"
	KiloBit SpeedType = "KiloBit"
	MegaBit SpeedType = "MegaBit"
	GigaBit SpeedType = "GigaBit"
	TeraBit SpeedType = "TeraBit"
)

type TimeType string

const (
	Second TimeType = "Second"
	Minute TimeType = "Minute"
	Hour   TimeType = "Hour"
	Day    TimeType = "Day"
)

type DecimalModeType string

// Options for Decimalmode
const (
	OptionAuto DecimalModeType = "Auto"
	OptionAll  DecimalModeType = "All"
)
