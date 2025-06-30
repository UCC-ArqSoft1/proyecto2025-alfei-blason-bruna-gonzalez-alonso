import React, { useEffect, useState } from "react";
import "./Activities.css";
import { useNavigate } from "react-router-dom";

function getCookie(name) {
    const value = `; ${document.cookie}`;
    const parts = value.split(`; ${name}=`);
    if (parts.length === 2) return parts.pop().split(";").shift();
    return null;
}


function useLogout() {
    const navigate = useNavigate();

    const borrarCookie = (nombre) => {
        document.cookie = `${nombre}=; path=/; expires=Thu, 01 Jan 1970 00:00:00 UTC;`;
        document.cookie = `${nombre}=; path=/; Max-Age=0;`;
    };

    const logout = () => {
        borrarCookie("token");
        borrarCookie("user_id");
        console.log("Sesión cerrada. Cookies eliminadas.");
        navigate("/login");
    };

    return logout;
}

function Activities() {
    const [actividades, setActividades] = useState([]);
    const [searchTerm, setSearchTerm] = useState("");
    const [filteredActividades, setFilteredActividades] = useState([]);
    const navigate = useNavigate();
    const logout = useLogout();

    const goToDetail = (id) => {
        navigate(`/activity/${id}`);
    };

    const volverAlLogin = () => {
        const confirmacion = window.confirm("¿Estás seguro de volver atrás? Se cerrará tu sesión.");
        if (confirmacion) logout();
    };

    useEffect(() => {
        const delayDebounce = setTimeout(() => {
            const filtro = searchTerm ? `?filtro=${searchTerm}` : "";
            fetch(`http://localhost:8080/act_deportiva${filtro}`)
                .then(res => res.json())
                .then(data => {
                    const actividadesSeguras = Array.isArray(data) ? data : [];
                    setActividades(actividadesSeguras);
                    setFilteredActividades(actividadesSeguras);
                })
                .catch(err => console.error(err));
        }, 400);

        return () => clearTimeout(delayDebounce);
    }, [searchTerm]);


    return (
        <>
            <div className="header-buttons">
                <button onClick={volverAlLogin} className="botonVolver">← Volver</button>
                <div className="espaciador" />
                <button onClick={() => navigate("/my-activities")} className="botonMisActividades">
                    Mis actividades
                </button>
            </div>


            <div className="container">
                <h1 className="title">Bienvenido a la página de Actividades</h1>

                {getCookie("rol") === "ADMIN" && (
                    <button
                        className="botonCrear"
                        onClick={() => navigate("/create-activity")}
                    >
                        Crear nueva actividad
                    </button>
                )}

                <input
                    type="text"
                    placeholder="Buscar por palabra clave, horario o profesor"
                    value={searchTerm}
                    onChange={e => setSearchTerm(e.target.value)} // Actualiza el estado del término de búsqueda
                    className="search-input"
                />



                {Array.isArray(filteredActividades) && filteredActividades.length === 0 ? (
                    <p>No se encontraron actividades.</p>
                ) : (
                    filteredActividades.map((item, index) => (
                        <div key={index} className="activity-card" onClick={() => goToDetail(item.actividad.IDActividad)}>
                            <img
                                src={
                                    item.actividad.Foto && item.actividad.Foto.trim() !== ""
                                        ? item.actividad.Foto
                                        : "https://via.placeholder.com/300x200?text=Sin+imagen"
                                }
                                alt={`Foto de ${item.actividad.Nombre}`}
                                className="actividad-foto"
                            />
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
