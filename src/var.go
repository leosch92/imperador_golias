package main

import (
	"math/big"
	"strconv"
	"strings"
)

type Integer string
type Real string
type Text string
type Bool string
type Error string

type Variable interface {
	WhatIsMyType() string
	MyString() string
	toString() string
}

type Var struct {
	Value Variable
}

func toVar(value string) (Var, bool) {
	val, err := toVariable(value)
	res := Var{val}
	return res, err
}

func errVar() (Var, bool) {
	val, err := errVariable()
	return Var{val}, err
}

func errVariable() (Variable, bool) {
	return Error("Error"), true
}

func toVariable(value string) (Variable, bool) {
	if value == "false" {
		val := Bool("false")
		return val, false
	} else if value == "true" {
		val := Bool("true")
		return val, false
	}
	str := value
	ln := len(str)
	if ln > 0 && str[0] == '"' {
		//str = str[1:]
		//str = str[:ln-2]
		return Text(str), false
	}
	if strings.Contains(str, ".") {
		_, err := big.NewFloat(0).SetString(str)
		if err {
			return Real(str), false
		} else {
			return Text(str), true
		}
	}
	_, err := big.NewInt(0).SetString(str, 10)
	if err {
		return Integer(str), false
	} else {
		return Text(str), true
	}
}

func (variable Integer) toString() string {
	return string(variable)
}

func (variable Real) toString() string {
	return string(variable)
}

func (variable Text) toString() string {
	return string(variable)
}

func (variable Bool) toString() string {
	return string(variable)
}

func (variable Error) toString() string {
	return string(variable)
}

func (variable Var) toString() string {
	return variable.Value.toString()
}

func (variable Var) WhatIsMyType() string {
	return variable.Value.WhatIsMyType()
}

func (variable Var) MyString() string {
	return variable.Value.MyString()
}

func (integer Integer) WhatIsMyType() string {
	return "integer"
}

func (integer Integer) MyString() string {
	return string(integer)
}

func (real Real) WhatIsMyType() string {
	return "real"
}

func (real Real) MyString() string {
	return string(real)
}

func (text Text) WhatIsMyType() string {
	return "text"
}

func (error Error) MyString() string {
	return string(error)
}

func (error Error) WhatIsMyType() string {
	return "error"
}

func (text Text) MyString() string {
	str := string(text)
	ln := len(str)
	if ln > 0 && str[0] == '"' {
		str = str[1:]
		str = str[:ln-2]
		return str
	} else {
		return "Invalid string"
	}
}

func (bl Bool) WhatIsMyType() string {
	return "bool"
}

func (bl Bool) MyString() string {
	return string(bl)
}

func (integer Integer) toInteger() (big.Int, bool) {
	val, err := big.NewInt(0).SetString(integer.MyString(), 10)
	return *val, !err
}

func (variable Var) toInteger() (big.Int, bool) {
	tp := variable.WhatIsMyType()
	if tp == "integer" {
		val := variable.Value.(Integer)
		return val.toInteger()
	} else {
		return *big.NewInt(0), true
	}
}

func (real Real) toReal() (big.Float, bool) {
	val, err := big.NewFloat(0).SetString(real.MyString())
	return *val, !err
}

func (variable Var) toReal() (big.Float, bool) {
	tp := variable.WhatIsMyType()
	if tp == "real" {
		val := variable.Value.(Real)
		return val.toReal()
	} else {
		return *big.NewFloat(0), true
	}
}

func (text Text) toText() (string, bool) {
	return text.MyString(), false
}

func (variable Var) toText() (string, bool) {
	tp := variable.WhatIsMyType()
	if tp == "text" {
		val := variable.Value.(Text)
		return val.toText()
	} else {
		return "", true
	}
}

func (bl Bool) toBool() (bool, bool) {
	val, err := strconv.ParseBool(bl.MyString())
	if err == nil {
		return val, false
	} else {
		return false, true
	}
}

func (variable Var) toBool() (bool, bool) {
	tp := variable.WhatIsMyType()
	if tp == "bool" {
		val := variable.Value.(Bool)
		return val.toBool()
	} else {
		return false, true
	}

}
