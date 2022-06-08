package parser

import (
	"time"
	"strings"
	"strconv"
)

func ParseToFloatError(cadena string, eParam error, bitSize int)
					(float float32, e error){
	float, e = strconv.ParseFloat(recString[1], bitSize)
	if (eParam != nil){
		e = eParam
	}	
}

/* Params: cadena (entero como cadena), eParam (error)
   Returns: entero (entero convertido desde la cadena), 
   			e (el primer error en orden de aparación entre eParam y 
			el error de conversión)
*/
func ParseToUint32Error(cadena string, eParam error) (entero uint32, e error){
	var entero64 uint64
	entero64, e = strconv.ParseUint(cadena, 10, 32)
	entero = uint32(entero64)
	if (eParam != nil){
		e = eParam
	}
	return entero, e
}

/* 
	Params: cadena (entero como cadena), eParam (error)
   	Returns: entero (entero convertido desde la cadena), 
   			e (el primer error en orden de aparición entre eParam y 
			el error de conversión en esta función)
*/
func ParseToIntError(cadena string, eParam error) (entero int, e error){
	entero, e = strconv.Atoi(cadena)
	if (eParam != nil){
		e = eParam
	}
	return entero, e
}

/*
	Params: tiempoCadena (fecha y hora en formato UTC como cadena),
	Returns: t (la fecha convertida al tipo Time),
			e (el primer error al intentar convertir la fecha)
*/
func ParseTimeUTC(tiempoCadena string) (t time.Time, e error){
	vectorTiempo := strings.Split(tiempoCadena, " ")
	// obtención componentes de fecha
	//var dia, mes, anio uint32
	vectorFecha := strings.Split(vectorTiempo[0], "-")
	dia, e := parseToIntError(vectorFecha[2], e)
	mes, e := parseToIntError(vectorFecha[1], e)
	anio, e := parseToIntError(vectorFecha[0], e)
	// obtención componentes de la hora
	vectorHora := strings.Split(vectorTiempo[1], ":")
	hora, e := parseToIntError(vectorHora[0], e)
	minuto, e := parseToIntError(vectorHora[1], e)
	segundos, e := parseToIntError(vectorHora[2], e)
	//time.Month(mes)
	t = time.Date(anio, time.Month(mes), dia, hora, minuto, segundos, 0, time.UTC)
	return t, e
}

