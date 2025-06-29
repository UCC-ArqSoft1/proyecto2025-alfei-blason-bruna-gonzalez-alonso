import React, { useEffect, useState } from "react";
import "./Details.css";
import { useNavigate, useParams } from "react-router-dom";

function Details() {
    const { id } = useParams();
    const [detalle, setDetalle] = useState(null);
    const navigate = useNavigate();
    const [inscripcionExitosa, setInscripcionExitosa] = useState(false);
    const [horariosInscriptos, setHorariosInscriptos] = useState(new Set());

    function getCookie(name) {
        const value = `; ${document.cookie}`;
        const parts = value.split(`; ${name}=`);
        if (parts.length === 2) return parts.pop().split(';').shift();
        return null;
    }

    const volverAlLogin = () => {
        navigate("/activities");
    };

    const verificarInscripcionesReales = async () => {
        try {
            const userID = getCookie('user_id');
            if (!userID || !detalle) {
                setHorariosInscriptos(new Set());
                return;
            }

            const response = await fetch(`http://localhost:8080/users/${userID}/inscripciones`);
            if (response.ok) {
                const inscripciones = await response.json();
                const horariosInscriptosSet = new Set();

                inscripciones.forEach(inscripcion => {
                    detalle.Horarios.forEach(horario => {
                        if (inscripcion.IDActividad === parseInt(id) &&
                            horario.Dia === inscripcion.Dia &&
                            horario.HorarioInicio === inscripcion.HoraInicio) {
                            horariosInscriptosSet.add(horario.IdHorario);
                        }
                    });
                });

                setHorariosInscriptos(horariosInscriptosSet);
            } else {
                setHorariosInscriptos(new Set());
            }
        } catch (error) {
            console.error("Error verificando inscripciones:", error);
            setHorariosInscriptos(new Set());
        }
    };

    const handleInscribirse = async (id_horario) => {
        try {
            const userID = getCookie('user_id');
            const token = getCookie('token');

            const response = await fetch(`http://localhost:8080/users/${userID}/inscripciones`, {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    "Authorization": `Bearer ${token}`
                },
                body: JSON.stringify({ id_horario: id_horario, id_actividad: parseInt(id) })
            });

            if (response.ok) {
                setInscripcionExitosa(true);
                setHorariosInscriptos(prev => new Set([...prev, id_horario]));

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

        const token = getCookie('token'); // Asegurate de tener el token

        try {
            const response = await fetch(`http://localhost:8080/act_deportiva/${id}`, {
                method: "DELETE",
                headers: {
                    "Authorization": `Bearer ${token}`
                }
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
        fetch(`http://localhost:8080/act_deportiva/${id}`)
            .then(async res => {
                if (!res.ok) {
                    const data = await res.json();
                    alert(data.error || "La actividad no existe.");
                    navigate("/activities");
                    return;
                }
                return res.json();
            })
            .then(data => {
                if (data) setDetalle(data);
            })
            .catch(err => {
                console.error("Error al cargar detalle:", err);
                alert("Error al cargar la actividad.");
                navigate("/activities");
            });
    }, [id]);


    useEffect(() => {
        if (detalle) {
            verificarInscripcionesReales();
        }
    }, [detalle]);

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

                {getCookie("rol") === "ADMIN" && detalle?.NombreActividad && (
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
                <p><strong>Actividad:</strong> {detalle.NombreActividad}</p>
                <p><strong>Descripción:</strong> {detalle.Descripcion}</p>
                <p><strong>Profesor:</strong> {detalle.NombreProfesor}</p>
                <p><strong>Horarios:</strong></p>
                <ul>
                    {detalle.Horarios.map((h, i) => {
                        const [hIni, mIni] = h.HorarioInicio.split(":").map(Number);
                        const [hFin, mFin] = h.HorarioFin.split(":").map(Number);
                        const duracionMin = (hFin * 60 + mFin) - (hIni * 60 + mIni);
                        const duracion = `${Math.floor(duracionMin / 60)}h ${duracionMin % 60}min`;
                        const yaInscripto = horariosInscriptos.has(h.IdHorario);
                        const cuposDisponibles = h.Cupos > 0;

                        return (
                            <li key={i}>
                                {h.Dia} de {h.HorarioInicio} a {h.HorarioFin} ({duracion}) Cupos: {h.Cupos}
                                {yaInscripto ? (
                                    <>
                                        <button className="botonInscripcion inscripto" disabled>Inscripto</button>
                                    </>
                                ) : (
                                    <button
                                        className="botonInscripcion"
                                        disabled={!cuposDisponibles}
                                        onClick={() => handleInscribirse(h.IdHorario)}
                                    >
                                        {cuposDisponibles ? "Inscribirme" : "Sin cupos"}
                                    </button>
                                )}
                            </li>
                        );
                    })}
                </ul>
            </div>
        </>
    );
}

export default Details;
