import React, { useState } from "react";
import { useNavigate } from "react-router-dom";
import "./createActivity.css";

function CreateActivity() {
    const navigate = useNavigate();
    const token = document.cookie.split('; ').find(row => row.startsWith('token='))?.split('=')[1];

    const [actividad, setActividad] = useState({
        nombre: "",
        nombreProfesor: "",
        idCategoria: 1,
        foto: "",
        descripcion: "",
        horarios: [{ dia: "", horarioInicio: "", horarioFin: "", cupos: 0 }],
    });

    const [errores, setErrores] = useState({});

    const handleChange = (e) => {
        setActividad({ ...actividad, [e.target.name]: e.target.value });
    };

    const handleHorarioChange = (index, e) => {
        const { name, value } = e.target;
        const nuevos = [...actividad.horarios];
        nuevos[index][name] = name === "cupos" ? parseInt(value) : value;
        setActividad({ ...actividad, horarios: nuevos });
    };

    const agregarHorario = () => {
        setActividad({
            ...actividad,
            horarios: [...actividad.horarios, { dia: "", horarioInicio: "", horarioFin: "", cupos: 0 }]
        });
    };

    const esHoraValida = (hora) => /^([01]\d|2[0-3]):([0-5]\d)$/.test(hora);
    const minutosDesdeCero = (hora) => {
        const [h, m] = hora.split(":").map(Number);
        return h * 60 + m;
    };

    const handleSubmit = async (e) => {
        e.preventDefault();

        const nuevosErrores = {};

        if (!actividad.nombre.trim()) nuevosErrores.nombre = "El nombre es obligatorio.";
        if (!actividad.nombreProfesor.trim()) nuevosErrores.nombreProfesor = "El nombre del profesor es obligatorio.";
        if (!actividad.descripcion.trim()) nuevosErrores.descripcion = "La descripción es obligatoria.";

        actividad.horarios.forEach((h, i) => {
            if (!esHoraValida(h.horarioInicio)) nuevosErrores[`horarioInicio${i}`] = "Formato inválido (HH:MM)";
            if (!esHoraValida(h.horarioFin)) nuevosErrores[`horarioFin${i}`] = "Formato inválido (HH:MM)";
            if (
                esHoraValida(h.horarioInicio) &&
                esHoraValida(h.horarioFin) &&
                minutosDesdeCero(h.horarioFin) <= minutosDesdeCero(h.horarioInicio)
            ) {
                nuevosErrores[`horarioFin${i}`] = "Debe ser posterior al inicio";
            }
            if (isNaN(h.cupos) || h.cupos < 0) nuevosErrores[`cupos${i}`] = "Debe ser ≥ 0";
        });

        if (Object.keys(nuevosErrores).length > 0) {
            setErrores(nuevosErrores);
            return;
        }

        setErrores({});

        const response = await fetch("http://localhost:8080/act_deportiva", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                "Authorization": `Bearer ${token}`
            },
            body: JSON.stringify(actividad)
        });

        if (response.ok) {
            alert("¡Actividad creada con éxito!");
            navigate("/activities");
        } else {
            alert("Error al crear la actividad.");
        }
    };

    return (
        <div className="create-form">
            <button onClick={() => navigate(-1)} className="botonVolver">← Volver</button>
            <h2>Crear Actividad</h2>
            <form onSubmit={handleSubmit}>
                <input
                    name="nombre"
                    value={actividad.nombre}
                    onChange={handleChange}
                    placeholder="Nombre de la actividad"
                    className={errores.nombre ? "invalid" : ""}
                />
                {errores.nombre && <div className="error-message">{errores.nombre}</div>}

                <input
                    name="nombreProfesor"
                    value={actividad.nombreProfesor}
                    onChange={handleChange}
                    placeholder="Nombre del profesor"
                    className={errores.nombreProfesor ? "invalid" : ""}
                />
                {errores.nombreProfesor && <div className="error-message">{errores.nombreProfesor}</div>}

                <input
                    name="foto"
                    value={actividad.foto}
                    onChange={handleChange}
                    placeholder="URL de imagen (opcional)"
                />

                <textarea
                    name="descripcion"
                    value={actividad.descripcion}
                    onChange={handleChange}
                    placeholder="Descripción"
                    className={errores.descripcion ? "invalid" : ""}
                />
                {errores.descripcion && <div className="error-message">{errores.descripcion}</div>}

                {actividad.horarios.map((h, i) => (
                    <div className="horario-block" key={i}>
                        <input
                            name="dia"
                            value={h.dia}
                            onChange={(e) => handleHorarioChange(i, e)}
                            placeholder="Día (ej: Lunes)"
                        />
                        <input
                            name="horarioInicio"
                            value={h.horarioInicio}
                            onChange={(e) => handleHorarioChange(i, e)}
                            placeholder="Horario Inicio (HH:MM)"
                            className={errores[`horarioInicio${i}`] ? "invalid" : ""}
                        />
                        {errores[`horarioInicio${i}`] && (
                            <div className="error-message">{errores[`horarioInicio${i}`]}</div>
                        )}

                        <input
                            name="horarioFin"
                            value={h.horarioFin}
                            onChange={(e) => handleHorarioChange(i, e)}
                            placeholder="Horario Fin (HH:MM)"
                            className={errores[`horarioFin${i}`] ? "invalid" : ""}
                        />
                        {errores[`horarioFin${i}`] && (
                            <div className="error-message">{errores[`horarioFin${i}`]}</div>
                        )}

                        <input
                            name="cupos"
                            type="number"
                            value={h.cupos}
                            onChange={(e) => handleHorarioChange(i, e)}
                            placeholder="Cupos"
                            className={errores[`cupos${i}`] ? "invalid" : ""}
                        />
                        {errores[`cupos${i}`] && (
                            <div className="error-message">{errores[`cupos${i}`]}</div>
                        )}
                    </div>
                ))}

                <button type="button" onClick={agregarHorario}>+ Agregar horario</button>

                <div className="button-group" style={{ marginTop: "20px" }}>
                    <button type="submit">Crear actividad</button>
                    <button type="button" onClick={() => navigate("/activities")}>Cancelar</button>
                </div>
            </form>
        </div>
    );
}

export default CreateActivity;
