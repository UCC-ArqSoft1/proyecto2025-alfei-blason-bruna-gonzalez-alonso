import React, { useEffect, useState } from "react";
import "./Details.css";
import {useNavigate, useParams} from "react-router-dom";

function Details() {
    const { id } = useParams();
    const [detalle, setDetalle] = useState(null);
    const navigate = useNavigate();
    const [inscripcionExitosa, setInscripcionExitosa] = useState(false);
    const [inscripcionesRealizadas, setInscripcionesRealizadas] = useState(() => {
        const guardadas = localStorage.getItem("inscripciones");
        return guardadas ? JSON.parse(guardadas) : [];
    });

    function getCookie(name) {
        const value = `; ${document.cookie}`;
        const parts = value.split(`; ${name}=`);
        if (parts.length === 2) return parts.pop().split(';').shift();
        return null;
    }

    const volverAlLogin = () => {
        navigate("/activities");
    };

    const handleClick = async (id_horario) => {
        try {
            const userID = getCookie('user_id');
            const token = getCookie('token');

            const response = await fetch(`http://localhost:8080/users/${userID}/inscripciones`, {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    "Authorization": `Bearer ${token}`
                },
                body: JSON.stringify({id_horario: id_horario, id_actividad: parseInt(id)})
            });

            if (response.ok) {
                setInscripcionExitosa(true);
                setInscripcionesRealizadas(prev => {
                    const nuevas = [...prev, id_horario];
                    localStorage.setItem("inscripciones", JSON.stringify(nuevas));
                    return nuevas;
                });

                const nuevoDetalle = { ...detalle };
                nuevoDetalle.Horarios = detalle.Horarios.map((h) =>
                    h.IdHorario === id_horario ? { ...h, Cupos: h.Cupos - 1 } : h
                );
                setDetalle(nuevoDetalle);
            } else {
                console.error("Inscripción fallida");
            }

        } catch (error) {
            console.error("Inscripción fallida", error);
        }
    };


    const handleEliminar = async () => {
        const confirmacion = window.confirm("¿Estás seguro de eliminar esta actividad?");
        if (!confirmacion) return;

        try {
            const response = await fetch(`http://localhost:8080/act_deportiva/${id}`, {
                method: "DELETE"
            });

            if (response.ok) {
                alert("¡Actividad eliminada!");
                navigate("/activities");
            } else {
                alert("Error al eliminar actividad.");
            }
        } catch (error) {
            console.error("Error al eliminar:", error);
            alert("Error en el servidor.");
        }
    };

    


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

                {getCookie("rol") === "ADMIN" && (
                    <div className="admin-buttons">
                        <button className="botonEditar" onClick={() => navigate(`/edit-activity/${id}`)}>
                            Editar Actividad
                        </button>
                        <button className="botonEliminar" onClick={handleEliminar}>
                            Eliminar Actividad
                        </button>
                    </div>
                )}

                <img src={detalle.Imagen} alt={detalle.Nombre} className="activity-image" />
                <p><strong>Actividad:</strong>{detalle.NombreActividad}</p>
                <p><strong>Descripción:</strong> {detalle.Descripcion}</p>
                <p><strong>Profesor:</strong> {detalle.NombreProfesor}</p>
                <p><strong>Horarios:</strong></p>
                <ul>
                    {Array.isArray(detalle.Horarios) && detalle.Horarios.map((h, i) => {
                        const [hIni, mIni] = h.HorarioInicio.split(":").map(Number); //split separa y map number convierte en string
                        const [hFin, mFin] = h.HorarioFin.split(":").map(Number);
                        const minutosInicio = hIni * 60 + mIni;
                        const minutosFin = hFin * 60 + mFin;
                        const duracionMin = minutosFin - minutosInicio;
                        const duracion = `${Math.floor(duracionMin / 60)}h ${duracionMin % 60}min`; //convierte la duracion a horas y minutos
                        const cuposDisponibles = h.Cupos > 0;

                        const yaInscripto = inscripcionesRealizadas.includes(h.IdHorario);

                        return (
                            <li key={i}>
                                {h.Dia} de {h.HorarioInicio} a {h.HorarioFin} ({duracion}) Cupos: {h.Cupos}
                                <button
                                    type="submit"
                                    className="botonInscripcion"
                                    disabled={!cuposDisponibles || yaInscripto}
                                    onClick={() => handleClick(h.IdHorario)}
                                    style={{
                                        backgroundColor: yaInscripto ? "gray" : "",
                                        cursor: yaInscripto ? "not-allowed" : "pointer"
                                    }}
                                >
                                    {yaInscripto ? "Inscripto" : (cuposDisponibles ? "Inscribirme" : "Sin cupos")}
                                </button>
                            </li>
                        );

                    })}
                </ul>
            </div>
        </>
    );
}

export default Details;
