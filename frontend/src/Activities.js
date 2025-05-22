import React, { useEffect, useState } from "react";
import "./Activities.css"; // Asegurate de que este archivo exista en la misma carpeta

function Activities() {
    const [actividades, setActividades] = useState([]);
    const [searchTerm, setSearchTerm] = useState("");
    const [filteredActividades, setFilteredActividades] = useState([]);

    useEffect(() => {
        fetch("http://localhost:8080/act_deportiva")
            .then(res => res.json())
            .then(data => {
                setActividades(data);
                setFilteredActividades(data);
            })
            .catch(err => console.error(err));
    }, []);

    useEffect(() => {
        const term = searchTerm.toLowerCase();
        const filtered = actividades.filter(item => {
            const act = item.actividad;
            const horariosStr = item.horarios
                .map(h => `${h.Dia} ${h.HorarioInicio}-${h.HorarioFin}`)
                .join(" ")
                .toLowerCase();

            return (
                act.Nombre.toLowerCase().includes(term) ||
                act.NombreProfesor.toLowerCase().includes(term) ||
                horariosStr.includes(term)
            );
        });
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
                filteredActividades.map((item, index) => (
                    <div key={index} className="activity-card">
                        <p><strong>Actividad:</strong> {item.actividad.Nombre}</p>
                        <p><strong>Profesor:</strong> {item.actividad.NombreProfesor}</p>
                        <p><strong>Cupos:</strong> {item.actividad.Cupos}</p>

                        <p><strong>Horarios:</strong></p>
                        <ul>
                            {item.horarios.map((h, i) => (
                                <li key={i}>
                                    {h.Dia} de {h.HorarioInicio} a {h.HorarioFin}
                                </li>
                            ))}
                        </ul>
                    </div>
                ))
            )}
        </div>
    );
}

export default Activities;
