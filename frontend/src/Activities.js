import React, { useEffect, useState } from "react";
import "./Activities.css"; // Asegúrate que esté en la misma carpeta que Activities.js

function Activities() {
    const [actividades, setActividades] = useState([]);
    const [searchTerm, setSearchTerm] = useState("");
    const [filteredActividades, setFilteredActividades] = useState([]);

    useEffect(() => {
        fetch("http://localhost:8080/act_deportiva") // Ajustar si es necesario
            .then(res => res.json())
            .then(data => {
                setActividades(data);
                setFilteredActividades(data);
            })
            .catch(err => console.error(err));
    }, []);

    useEffect(() => {
        const term = searchTerm.toLowerCase();
        const filtered = actividades.filter(act =>
            act.Nombre.toLowerCase().includes(term) ||
            act.NombreProfesor.toLowerCase().includes(term) ||
            (act.Horario && act.Horario.toLowerCase().includes(term))
        );
        setFilteredActividades(filtered);
    }, [searchTerm, actividades]);

    return (
        <div className="container">
            <h1 className="title">Bienvenido a la página de Actividades</h1>

            <input
                type="text"
                placeholder="Buscar por palabra clave, horario o profesor"
                value={searchTerm}
                onChange={e => setSearchTerm(e.target.value)}
                className="search-input"
            />

            {filteredActividades.length === 0 ? (
                <p>No se encontraron actividades.</p>
            ) : (
                filteredActividades.map(act => (
                    <div key={act.IDactividad} className="activity-card">
                        <p><strong>Actividad:</strong> {act.Nombre}</p>
                        <p><strong>Profesor:</strong> {act.NombreProfesor}</p>
                        <p><strong>Cupos:</strong> {act.Cupos}</p>
                        {act.Horario && <p><strong>Horario:</strong> {act.Horario}</p>}
                    </div>
                ))
            )}
        </div>
    );
}

export default Activities;
