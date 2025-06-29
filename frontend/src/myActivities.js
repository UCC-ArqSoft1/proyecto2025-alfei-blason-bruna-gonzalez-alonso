import React, { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import "./myActivities.css";

function getCookie(name) {
    const value = `; ${document.cookie}`;
    const parts = value.split(`; ${name}=`);
    if (parts.length === 2) return parts.pop().split(";").shift();
    return null;
}

function MyActivities() {
    const [misActividades, setMisActividades] = useState([]);
    const navigate = useNavigate();

    useEffect(() => {
        const userID = getCookie("user_id");
        if (!userID) {
            alert("Sesión expirada. Iniciá sesión de nuevo.");
            navigate("/login");
            return;
        }

        fetch(`http://localhost:8080/users/${userID}/inscripciones`)
            .then(res => res.json())
            .then(data => {
                if (Array.isArray(data)) {
                    setMisActividades(data);
                } else {
                    setMisActividades([]);
                }
            })
            .catch(error => {
                console.error("Error al obtener inscripciones:", error);
                alert("No se pudieron cargar tus actividades.");
            });
    }, [navigate]);

    const irADetalle = (id) => {
        navigate(`/activity/${id}`);
    };

    return (
        <div className="mis-actividades-container">
            <h2>Mis Actividades Inscritas</h2>
            {misActividades.length === 0 ? (
                <p>No estás inscripto en ninguna actividad.</p>
            ) : (
                misActividades.map((act, i) => (
                    <div
                        key={i}
                        className="actividad-inscripta-card"
                        onClick={() => irADetalle(act.IDActividad)}
                    >
                        <img src={act.Foto} alt={act.NombreActividad} className="actividad-foto" />
                        <h3>{act.NombreActividad}</h3>
                        <p><strong>Profesor:</strong> {act.NombreProfesor}</p>
                        <p><strong>Día:</strong> {act.Dia}</p>
                        <p><strong>Horario:</strong> {act.HoraInicio || "Sin definir"} - {act.HoraFin || "Sin definir"}</p>
                        <p><strong>Cupos:</strong> {act.Cupos}</p>
                    </div>
                ))
            )}
        </div>
    );
}

export default MyActivities;
