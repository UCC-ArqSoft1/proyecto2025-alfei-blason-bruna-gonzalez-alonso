import React, { useState, useEffect } from "react";
import { useParams, useNavigate } from "react-router-dom";
import './editActivity.css';

    function getCookie(name) {
        const value = `; ${document.cookie}`;
        const parts = value.split(`; ${name}=`);
        if (parts.length === 2) return parts.pop().split(";").shift();
        return null;
    }

function EditActivity() {
    const { id } = useParams(); // id de la actividad desde la URL
    const navigate = useNavigate();

    const [nombre, setNombre] = useState("");
    const [profesor, setProfesor] = useState("");
    const [dia, setDia] = useState("");
    const [inicio, setInicio] = useState("");
    const [fin, setFin] = useState("");
    const [cupos, setCupos] = useState(0);
    const [foto, setFoto] = useState("");
    const [descripcion, setDescripcion] = useState("");

    function getCookie(name) {
        const value = `; ${document.cookie}`;
        const parts = value.split(`; ${name}=`);
        if (parts.length === 2) return parts.pop().split(";").shift();
    }

    useEffect(() => {
        const rol = getCookie("rol");
        if (rol !== "ADMIN") {
            alert("No tenés permisos para editar actividades.");
            navigate("/activities");
            return;
        }

        fetch(`http://localhost:8080/act_deportiva/${id}`)
            .then(res => res.json())
            .then(data => {
                setNombre(data.NombreActividad || "");
                setProfesor(data.NombreProfesor || "");
                setFoto(data.Foto || data.Imagen || "");
                setDescripcion(data.Descripcion || "");

                if (Array.isArray(data.Horarios) && data.Horarios.length > 0) {
                    const h = data.Horarios[0];
                    setDia(h.Dia || "");
                    setInicio(h.HorarioInicio || "");
                    setFin(h.HorarioFin || "");
                    setCupos(h.Cupos || 0);
                }
            })
            .catch(err => {
                console.error("Error al cargar datos:", err);
                alert("No se pudo cargar la actividad.");
            });
    }, [id, navigate]);

    const handleSubmit = async (e) => {
        e.preventDefault();
        const actividadEditada = {
            Nombre: nombre,
            NombreProfesor: profesor,
            Dia: dia,
            HorarioInicio: inicio,
            HorarioFin: fin,
            Cupos: parseInt(cupos),
            Descripcion: descripcion
        };

        if (foto && foto.trim() !== "") {
            actividadEditada.Foto = foto;
        }


        try {
            const response = await fetch(`http://localhost:8080/act_deportiva/${id}`, {
                method: "PUT",
                headers: { "Content-Type": "application/json" },
                body: JSON.stringify(actividadEditada)
            });

            if (!response.ok) throw new Error("Error al editar actividad");

            alert("¡Actividad editada correctamente!");
            navigate("/activities");
        } catch (error) {
            console.error("Error al enviar edición:", error);
            alert("Hubo un problema al guardar los cambios.");
        }
    };

    return (
        <div className="edit-form">
            <h2>Editar actividad</h2>
            <form onSubmit={handleSubmit}>
                <input
                    type="text"
                    placeholder="Nombre"
                    value={nombre}
                    onChange={e => setNombre(e.target.value)}
                    required
                />

                <input
                    type="text"
                    placeholder="Profesor"
                    value={profesor}
                    onChange={e => setProfesor(e.target.value)}
                    required
                />

                <input
                    type="text"
                    placeholder="Día (ej. Martes)"
                    value={dia}
                    onChange={e => setDia(e.target.value)}
                    required
                />

                <input
                    type="time"
                    value={inicio}
                    onChange={e => setInicio(e.target.value)}
                    required
                />

                <input
                    type="time"
                    value={fin}
                    onChange={e => setFin(e.target.value)}
                    required
                />

                <input
                    type="number"
                    placeholder="Cupos"
                    value={cupos}
                    onChange={e => setCupos(e.target.value)}
                    required
                />

                {/* Vista previa de imagen actual */}
                {foto && (
                    <div style={{ marginBottom: "10px" }}>
                        <p>Foto actual:</p>
                        <img src={foto} alt="Foto actual" style={{ maxWidth: "200px", borderRadius: "8px" }} />
                    </div>
                )}

                <input
                    type="text"
                    placeholder="URL de la foto"
                    value={foto}
                    onChange={e => setFoto(e.target.value)}
                />

                <textarea
                    placeholder="Descripción"
                    value={descripcion}
                    onChange={e => setDescripcion(e.target.value)}
                    required
                />

                <button type="submit">Guardar cambios</button>
            </form>

        </div>
    );
}

export default EditActivity;
