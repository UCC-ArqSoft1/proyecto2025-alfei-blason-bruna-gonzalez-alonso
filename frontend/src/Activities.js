import React, { useEffect, useState } from "react";
import "./Activities.css";
import {useNavigate} from "react-router-dom";


function Activities() {
    const [actividades, setActividades] = useState([]);
    const [searchTerm, setSearchTerm] = useState("");
    const [filteredActividades, setFilteredActividades] = useState([]);
    const navigate = useNavigate();

    const goToDetail = (id) => {
        navigate(`/activity/${id}`);
    };

    const volverAlLogin = () => {
        const confirmacion = window.confirm("¿Estás seguro de volver atrás? Se cerrará tu sesión.");
        if (confirmacion) {
            navigate("/login");
        }
    };


    useEffect(() => {
        fetch("http://localhost:8080/act_deportiva")//llama a la api
            .then(res => res.json())// convierte la respuesta en json
            .then(data => {
                setActividades(data);
                setFilteredActividades(data);
            })
            .catch(err => console.error(err));// si hay error muestra
    }, []);

    useEffect(() => { //filtra los datos cuando cambia el termino de busqueda
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
        <>
        <button onClick={volverAlLogin} className="botonVolver"> ← Volver </button>
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
                    <div key={index} className="activity-card" onClick={() => goToDetail(item.actividad.IDActividad)}>
                        <p><strong>Actividad:</strong> {item.actividad.Nombre}</p>
                        <p><strong>Profesor:</strong> {item.actividad.NombreProfesor}</p>

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
        </>
    );
}

export default Activities;
