import React, { useState } from "react";
import { useNavigate } from "react-router-dom";
import './createActivity.css';

function getCookie(name) {
    const value = `; ${document.cookie}`;
    const parts = value.split(`; ${name}=`);
    if (parts.length === 2) return parts.pop().split(";").shift();
    return null;
}

function CreateActivity() {
    const [nombre, setNombre] = useState("");
    const [profesor, setProfesor] = useState("");
    const [dia, setDia] = useState("");
    const [inicio, setInicio] = useState("");
    const [fin, setFin] = useState("");
    const [cupos, setCupos] = useState(0);
    const [foto, setFoto] = useState("");
    const [descripcion, setDescripcion] = useState("");
    const navigate = useNavigate();

    const handleSubmit = async (e) => {
        e.preventDefault();

        const rol = getCookie("rol");
        const token = getCookie("token");

        if (rol !== "ADMIN") {
            alert("No tenés permisos para crear actividades.");
            return;
        }

        const nuevaActividad = {
            Nombre: nombre,
            NombreProfesor: profesor,
            Dia: dia,
            HorarioInicio: inicio,
            HorarioFin: fin,
            Cupos: parseInt(cupos),
            Foto: foto,
            Descripcion: descripcion
        };

        try {
            const response = await fetch("http://localhost:8080/act_deportiva", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    "Authorization": `Bearer ${token}`
                },
                body: JSON.stringify(nuevaActividad)
            });

            const responseData = await response.json();

            if (!response.ok) {
                alert(responseData.error || responseData.Mensaje || "Error al crear la actividad.");
                return;
            }

            alert("¡Actividad creada con éxito!");
            navigate("/activities");

        } catch (err) {
            console.error("Error al crear actividad:", err);
            alert("Error de red o del servidor.");
        }
    };

    return (
        <div className="create-form">
            <h2>Cargar nueva actividad</h2>
            <form onSubmit={handleSubmit}>
                <input type="text" placeholder="Nombre" value={nombre} onChange={e => setNombre(e.target.value)} required />
                <input type="text" placeholder="Profesor" value={profesor} onChange={e => setProfesor(e.target.value)} required />
                <input type="text" placeholder="Día (ej. Lunes)" value={dia} onChange={e => setDia(e.target.value)} required />
                <input type="time" value={inicio} onChange={e => setInicio(e.target.value)} required />
                <input type="time" value={fin} onChange={e => setFin(e.target.value)} required />
                <input type="number" placeholder="Cupos" value={cupos} onChange={e => setCupos(e.target.value)} required />
                <input type="text" placeholder="URL de la foto" value={foto} onChange={e => setFoto(e.target.value)} required />
                <textarea placeholder="Descripción" value={descripcion} onChange={e => setDescripcion(e.target.value)} required />
                <button type="submit">Crear actividad</button>
            </form>
        </div>
    );
}

export default CreateActivity;
