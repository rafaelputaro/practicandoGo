package model

type recorrido struct{
	//id y duración
	id_recorrido string
	duraccion_recorrido uint32
	//origen
	fecha_origen_recorido Time
	id_estacion_origen string
	nombre_estacion_origen string
	direccion_estacion_origen string
	long_estacion_origen float32
	lat_estacion_origen float32
	//destino
	fecha_destino_recorido Time
	id_estacion_destino string
	nombre_estacion_destino string
	direccion_estacion_destino string
	long_estacion_destino float32
	lat_estacion_destino float32
	//datos usuario y bici
	id_usuario string
	modelo_bicicleta string
}

func Recorrido(id_recorrido string, duraccion_recorrido uint32,
			fecha_origen_recorido Time, id_estacion_origen string,
			nombre_estacion_origen string, direccion_estacion_origen string,
			long_estacion_origen float32, lat_estacion_origen float32,
			fecha_destino_recorido Time, id_estacion_destino string,
			nombre_estacion_destino string, direccion_estacion_destino string,
			long_estacion_destino float32, lat_estacion_destino float32,
			id_usuario string, modelo_bicicleta string) (r *recorrido){
	r = new (recorrido)
	//id y duración
	r.id_recorrido = id_recorrido
	r.duraccion_recorrido = duraccion_recorrido
	//origen
	r.fecha_origen_recorido = fecha_origen_recorido
	r.id_estacion_origen = id_estacion_origen
	r.nombre_estacion_origen = nombre_estacion_origen
	r.direccion_estacion_origen = direccion_estacion_origen
	r.long_estacion_origen = long_estacion_origen
	r.lat_estacion_origen = lat_estacion_origen
	//destino
	r.fecha_origen_recorido = fecha_origen_recorido
	r.id_estacion_origen = id_estacion_origen
	r.nombre_estacion_origen = nombre_estacion_origen
	r.direccion_estacion_origen = direccion_estacion_origen
	r.long_estacion_origen = long_estacion_origen
	r.lat_estacion_origen = lat_estacion_origen
	//datos usuario y bicicleta
	r.id_usuario = id_usuario
	r.modelo_bicicleta = modelo_bicicleta
}
