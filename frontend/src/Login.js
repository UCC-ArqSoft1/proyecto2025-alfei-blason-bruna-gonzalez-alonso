import React, { useState } from "react";
import './Login.css';
import { useNavigate } from "react-router-dom";


function Login() {

    const [usuario, setUsuario] = useState("");
    const [contrasenia, setContrasenia] = useState("");
    const navigate = useNavigate();

    const changeUsername = async (e) => {
        console.log(e.target.value);
        setUsuario(e.target.value);
    }
    const changePassword = async (e) => {
        console.log(e.target.value);
        setContrasenia(e.target.value);
    }
    const send = async (e) => {
            e.preventDefault(); // importante para evitar recargar la página
            console.log("Llamando al backend");

            try {
                const response = await fetch("http://localhost:8080/users/login", {
                    method: "POST",
                    headers: {
                        "Content-Type": "application/json",
                    },
                    body: JSON.stringify({ usuario, contrasenia }),
                });

                if (!response.ok) throw new Error("Error en login");

                const data = await response.json();
                console.log(data);

                document.cookie = `user_id = ${data.usuario}; token = ${data.token}; path=/;SameSite=Strict;`

                navigate("/Activities");

            } catch (error) {
                console.error("Login fallido", error);
                alert("Credenciales incorrectas");
            }
    }

    return (
        <div className="Login">
        <header>
            <div>
                <h1 id="tituloPrincipal"> GIMNASIO APP </h1>

                <form onSubmit={send}>
                    <input type="text" name="username" placeholder="Nombre de usuario"
                           onChange={changeUsername}/><br/>
                    <input type="password" name="password" placeholder="Contraseña"
                           onChange={changePassword}/>

                    <input type="submit" value="Login" className="botonInicio"/>
                </form>

            </div>
        </header>
        </div>
    );

}

export default Login;