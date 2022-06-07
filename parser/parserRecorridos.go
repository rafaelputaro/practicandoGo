package parser

import (
	"fmt"
)

/*
type parserRecorridos struct{
	
}
*/

func funcion(){
	cadena :="2021-04-10 20:48:15 UTC"
	t, e := ParseTimeUTC(cadena)
	if (e == nil){
		fmt.Println(t.UTC())
	}
	var v uint =32
	fmt.Println(fmt.Sprintln(v))
}

/*
func parseRowToRecorrido(row [16]string)(r *recorrido.Recorrido, e error){
	//id y duración
	id_recorrido = row[0]
	duraccion_recorrido = parseToUint32Error(recString[1], error)
	//origen
	fecha_origen_recorido = recString[2]
	id_estacion_origen = recString[3]
	nombre_estacion_origen = recString[4]
	direccion_estacion_origen = recString[5]
	r.long_estacion_origen = parseToFloatError(recString[6], errores, 32)
	r.lat_estacion_origen = parseToFloatError(recString[7], errores, 32)
	//destino
	r.fecha_origen_recorido = recString[8]
	r.id_estacion_origen = recString[9]
	r.nombre_estacion_origen = recString[10]
	r.direccion_estacion_origen = recString[11]
	r.long_estacion_origen = parseToFloatError(recString[12], errores, 32)
	r.lat_estacion_origen = parseToFloatError(recString[13], errores, 32)
	//datos usuario y bicicleta
	r.id_usuario = recString[14]
	r.modelo_bicicleta = recString[15]

	return r, e
}*/
