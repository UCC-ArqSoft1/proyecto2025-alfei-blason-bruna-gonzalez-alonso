import React, { useState } from "react";
import './Login.css';

function Login() {

    const [username, setUsername] = useState("");
    const [password, setPassword] = useState("");


    const changeUsername = async (e) => {
        console.log(e.target.value);
        setUsername(e.target.value);
    }
    const changePassword = async (e) => {
        console.log(e.target.value);
        setPassword(e.target.value);
    }
    const send = async (e) => {
        console.log("Llamando al backend")
        const response = await fetch("http://localhost:8080/login", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify({ username, password }),
        });
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