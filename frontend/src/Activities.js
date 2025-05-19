import React, { useEffect, useState } from "react";

function Activities() {
    const [actividad, setActividad] = useState(null);

    useEffect(() => {
        // Por ejemplo, pedimos la actividad con ID 2
        fetch("http://localhost:8080/act_deportiva/2")
            .then(res => res.json())
            .then(data => setActividad(data))
            .catch(err => console.error(err));
    }, []);

    return (
        <div style={{ padding: "2rem" }}>
            <h1>Bienvenido a la página de Actividades</h1>
            {actividad ? (
                <div>
                    <p><strong>Actividad:</strong> {actividad.NombreActividad}</p>
                    <p><strong>Profesor:</strong> {actividad.NombreProfesor}</p>
                    <p><strong>Cupos:</strong> {actividad.Cupos}</p>
                </div>
            ) : (
                <p>Cargando actividad...</p>
            )}
        </div>
    );
}

export default Activities;
