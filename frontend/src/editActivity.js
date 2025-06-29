import React, { useEffect, useState } from "react";
import { useParams, useNavigate } from "react-router-dom";
import "./createActivity.css";

function EditActivity() {
    const { id } = useParams();
    const navigate = useNavigate();
    const token = document.cookie.split('; ').find(row => row.startsWith('token='))?.split('=')[1];

    const [actividad, setActividad] = useState(null);
    const [errores, setErrores] = useState({});

    useEffect(() => {
        fetch(`http://localhost:8080/act_deportiva/${id}`)
            .then(res => {
                if (!res.ok) {
                    alert("La actividad no existe o fue eliminada.");
                    navigate("/activities");
                    return null;
                }
                return res.json();
            })
            .then(data => {
                if (data) {
                    setActividad({
                        nombre: data.NombreActividad,
                        nombreProfesor: data.NombreProfesor,
                        foto: data.Foto,
                        descripcion: data.Descripcion,
                        idCategoria: 1,
                        horarios: data.Horarios.map(h => ({
                            dia: h.Dia,
                            horarioInicio: h.HorarioInicio,
                            horarioFin: h.HorarioFin,
                            cupos: h.Cupos
                        }))
                    });
                }
            });
    }, [id]);


    const handleChange = (e) => {
        setActividad({ ...actividad, [e.target.name]: e.target.value });
    };

    const handleHorarioChange = (index, e) => {
        const { name, value } = e.target;
        const newHorarios = [...actividad.horarios];
        newHorarios[index][name] = name === "cupos" ? parseInt(value) : value;
        setActividad({ ...actividad, horarios: newHorarios });
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
            if (esHoraValida(h.horarioInicio) && esHoraValida(h.horarioFin) &&
                minutosDesdeCero(h.horarioFin) <= minutosDesdeCero(h.horarioInicio)) {
                nuevosErrores[`horarioFin${i}`] = "Debe ser posterior al inicio";
            }
            if (isNaN(h.cupos) || h.cupos < 0) nuevosErrores[`cupos${i}`] = "Debe ser ≥ 0";
        });

        if (Object.keys(nuevosErrores).length > 0) {
            setErrores(nuevosErrores);
            return;
        }

        setErrores({});

        const response = await fetch(`http://localhost:8080/act_deportiva/${id}`, {
            method: "PUT",
            headers: {
                "Content-Type": "application/json",
                "Authorization": `Bearer ${token}`
            },
            body: JSON.stringify(actividad)
        });

        if (response.ok) {
            alert("¡Actividad editada correctamente!");
            navigate("/activities");
        } else {
            alert("Error al editar la actividad.");
        }
    };

    if (!actividad) return <p>Cargando actividad...</p>;

    return (
        <div className="create-form">
            <h2>Editar Actividad</h2>
            <form onSubmit={handleSubmit}>
                <input
                    name="nombre"
                    value={actividad.nombre}
                    onChange={handleChange}
                    className={errores.nombre ? "invalid" : ""}
                />
                {errores.nombre && <div className="error-message">{errores.nombre}</div>}

                <input
                    name="nombreProfesor"
                    value={actividad.nombreProfesor}
                    onChange={handleChange}
                    className={errores.nombreProfesor ? "invalid" : ""}
                />
                {errores.nombreProfesor && <div className="error-message">{errores.nombreProfesor}</div>}

                <input
                    name="foto"
                    value={actividad.foto}
                    onChange={handleChange}
                />

                <textarea
                    name="descripcion"
                    value={actividad.descripcion}
                    onChange={handleChange}
                    className={errores.descripcion ? "invalid" : ""}
                />
                {errores.descripcion && <div className="error-message">{errores.descripcion}</div>}

                {actividad.horarios.map((h, i) => (
                    <div className="horario-block" key={i}>
                        <input
                            name="dia"
                            value={h.dia}
                            onChange={(e) => handleHorarioChange(i, e)}
                        />
                        <input
                            name="horarioInicio"
                            value={h.horarioInicio}
                            onChange={(e) => handleHorarioChange(i, e)}
                            className={errores[`horarioInicio${i}`] ? "invalid" : ""}
                        />
                        {errores[`horarioInicio${i}`] && (
                            <div className="error-message">{errores[`horarioInicio${i}`]}</div>
                        )}

                        <input
                            name="horarioFin"
                            value={h.horarioFin}
                            onChange={(e) => handleHorarioChange(i, e)}
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
                            className={errores[`cupos${i}`] ? "invalid" : ""}
                        />
                        {errores[`cupos${i}`] && (
                            <div className="error-message">{errores[`cupos${i}`]}</div>
                        )}
                    </div>
                ))}

                <div className="button-group">
                    <button type="submit">Guardar cambios</button>
                    <button type="button" onClick={() => navigate("/activities")}>Cancelar</button>
                </div>
            </form>
        </div>
    );
}

export default EditActivity;
