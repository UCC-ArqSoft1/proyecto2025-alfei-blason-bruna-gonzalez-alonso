import React, { useEffect, useState } from "react";
import "./Details.css";
import {useNavigate, useParams} from "react-router-dom";

function Details() {
    const { id } = useParams();
    const [detalle, setDetalle] = useState(null);
    const navigate = useNavigate();
    const [inscripcionExitosa, setInscripcionExitosa] = useState(false);

    const volverAlLogin = () => {
        navigate("/activities");
    };

    const handleClick = async (e) => {
        console.log("Captured click")
        e.preventDefault();

        try {
            function getCookie(name) {
                const value = `; ${document.cookie}`;
                const parts = value.split(`; ${name}=`);
                if (parts.length === 2) return parts.pop().split(';').shift();
            }

            const userID = getCookie('user_id');
            const token = getCookie('token');
            console.log('Usuario:', userID);
            console.log('Token:', token);

            const response = await fetch(`http://localhost:8080/users/${userID}/inscripciones`, {
                method: "POST",
                headers: {
                    "Content-Type": "application/json"
                },
                body: JSON.stringify({id_horario: 1, id_actividad: parseInt(id)})
            });

            if (response.ok) {
                setInscripcionExitosa(true);
            } else {
                console.error("Inscripción fallida");
            }

        }catch (error){
            console.error("Inscripcion fallida", error);
        }
    }

    useEffect(() => {
        console.log("cargado actividades");
        fetch(`http://localhost:8080/act_deportiva/${id}`)
            .then(res => res.json())
            .then(data => setDetalle(data))
            .catch(err => console.error(err));
    }, [id]);

    if (!detalle) return <p>Cargando detalles...</p>;

    return (
        <>
            <button onClick={volverAlLogin} className="botonVolver"> ← Volver </button>

            {inscripcionExitosa && (
                <div className="inscripcion-exitosa">
                    ¡Inscripción exitosa!
                </div>
            )}

            <div className="detalles">
                <h2>Detalle Actividad</h2>
                <img src={detalle.Imagen} alt={detalle.Nombre} className="activity-image" />
                <p><strong>Actividad:</strong>{detalle.NombreActividad}</p>
                <p><strong>Descripción:</strong> {detalle.Descripcion}</p>
                <p><strong>Profesor:</strong> {detalle.NombreProfesor}</p>
                <p><strong>Cupo máximo:</strong> {detalle.Cupos}</p>
                <p><strong>Horarios:</strong></p>
                <ul>
                    {Array.isArray(detalle.Horarios) && detalle.Horarios.map((h, i) => {
                        const [hIni, mIni] = h.HorarioInicio.split(":").map(Number); //split separa y map number convierte en string
                        const [hFin, mFin] = h.HorarioFin.split(":").map(Number);
                        const minutosInicio = hIni * 60 + mIni;
                        const minutosFin = hFin * 60 + mFin;
                        const duracionMin = minutosFin - minutosInicio;
                        const duracion = `${Math.floor(duracionMin / 60)}h ${duracionMin % 60}min`; //convierte la duracion a horas y minutos

                        return (
                            <li key={i}>
                                {h.Dia} de {h.HorarioInicio} a {h.HorarioFin} ({duracion})
                                <button type="submit" className="botonInscripcion" onClick={handleClick}> Inscribirme </button>
                            </li>
                        );
                    })}
                </ul>
            </div>
        </>
    );
}

export default Details;
