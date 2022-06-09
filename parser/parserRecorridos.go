package parser

import (
	"capudo/model"
)

const PATH_RECORRIDOS = "trips_2021.csv"

var (
	instanceParserRecorridos *parserRecorridos
)

type parserRecorridos struct{
	recorridos []model.Recorrido
}

func parseRowToRecorrido(row []string)(r *model.Recorrido, e error){
	//id y duración
	id_recorrido := row[0]
	duracion_recorrido, e := ParseToUint32Error(row[1], e)
	//origen
	fecha_origen_recorrido, e := ParseTimeUTC(row[2])
	id_estacion_origen := row[3]
	nombre_estacion_origen := row[4]
	direccion_estacion_origen := row[5]
	long_estacion_origen, e := ParseToFloat32Error(row[6], e)
	lat_estacion_origen, e := ParseToFloat32Error(row[7], e)
	//destino
	fecha_destino_recorrido, e := ParseTimeUTC(row[8])
	id_estacion_destino := row[9]
	nombre_estacion_destino := row[10]
	direccion_estacion_destino := row[11]
	long_estacion_destino, e := ParseToFloat32Error(row[12], e)
	lat_estacion_destino, e := ParseToFloat32Error(row[13], e)
	//datos usuario y bicicleta
	id_usuario := row[14]
	modelo_bicicleta := row[15]
	r = model.CreateRecorrido(id_recorrido, duracion_recorrido,
		fecha_origen_recorrido, id_estacion_origen,
		nombre_estacion_origen, direccion_estacion_origen,
		long_estacion_origen, lat_estacion_origen,
		fecha_destino_recorrido, id_estacion_destino,
		nombre_estacion_destino, direccion_estacion_destino,
		long_estacion_destino, lat_estacion_destino,
		id_usuario, modelo_bicicleta)
	return r, e
}

func createParserRecorridos()(parReco *parserRecorridos, e error){
	parserAux, eParGen := CreateParserGenerico(PATH_RECORRIDOS)
	if (eParGen != nil){
		parReco = new(parserRecorridos)
		for _, row := range parserAux.rows {			
			rowParsed, eParRow := parseRowToRecorrido(row) 
			if (eParRow != nil){
				parReco.recorridos = append(parReco.recorridos, *rowParsed)
			}			
		}
	} else{
		e = eParGen
	}
	return parReco, e
}

func GetRecorridos() (recorridos *[]model.Recorrido, e error){
	if(instanceParserRecorridos == nil){
		instanceParserRecorridos, e = createParserRecorridos()
		if (e != nil){
			recorridos = &instanceParserRecorridos.recorridos
		}
	} 
	return recorridos, e
}
