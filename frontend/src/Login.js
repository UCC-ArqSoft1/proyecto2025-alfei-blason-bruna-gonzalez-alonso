import React, { useState } from "react";
import './Login.css';

function Login() {

    const [usuario, setUsuario] = useState("");
    const [contrasenia, setContrasenia] = useState("");


    const changeUsername = async (e) => {
        console.log(e.target.value);
        setUsuario(e.target.value);
    }
    const changePassword = async (e) => {
        console.log(e.target.value);
        setContrasenia(e.target.value);
    }
    const send = async (e) => {
        console.log("Llamando al backend")
        const response = await fetch("http://localhost:8080/users/login", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify({ usuario, contrasenia }),
        });

        const data = await response.json();

        console.log(data);
    }
    return (
        <div className="Login">
        <header>
            <div>
                <h1 id="tituloPrincipal"> GIMNASIO APP </h1>

                <form onSubmit={send}>
                    <input type="text" name="username" placeholder="Nombre de usuario"
                           onChange={changeUsername}/><br/>
                    <input type="password" name="password" placeholder="Contrasea"
                           onChange={changePassword}/>

                    <input type="submit" value="Login"/>
                </form>
            </div>
        </header>
        </div>
    );

}

export default Login;