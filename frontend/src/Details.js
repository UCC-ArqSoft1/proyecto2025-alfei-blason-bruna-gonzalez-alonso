import React, { useEffect, useState } from "react";
import "./Details.css";
import { useParams } from "react-router-dom";

function Details() {
    const { id } = useParams();
    const [detalle, setDetalle] = useState(null);

    useEffect(() => {
        console.log("cargado actividades");
        fetch(`http://localhost:8080/act_deportiva/${id}`)
            .then(res => res.json())
            .then(data => setDetalle(data))
            .catch(err => console.error(err));
    }, [id]);

    if (!detalle) return <p>Cargando detalles...</p>;

    return (
        <div className="detalles">
            <h2>{detalle.Nombre}</h2>
            <img src={detalle.Imagen} alt={detalle.Nombre} className="activity-image" />
            <p><strong>Descripción:</strong> {detalle.Descripcion}</p>
            <p><strong>Profesor:</strong> {detalle.NombreProfesor}</p>
            <p><strong>Duración:</strong> {detalle.Duracion} </p>
            <p><strong>Cupo máximo:</strong> {detalle.Cupos}</p>
            <p><strong>Horarios:</strong></p>
            <ul>
                {Array.isArray(detalle.Horarios) && detalle.Horarios.map((h, i) => (
                    <li key={i}>{h.Dia} de {h.HorarioInicio} a {h.HorarioFin}</li>
                ))}
            </ul>
        </div>
    );
}

export default Details;
