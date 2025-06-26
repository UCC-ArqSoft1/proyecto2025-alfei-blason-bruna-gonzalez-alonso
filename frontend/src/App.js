import React from 'react';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import Login from './Login';
import Activities from './Activities';
import Details from './Details';
import CreateActivity from "./createActivity";
import EditActivity from "./editActivity";

import './App.css';

function App() {
    return (
        <Router>
            <Routes>
                <Route path="/" element={<Login />} />
                <Route path="/login" element={<Login />} />
                <Route path="/activities" element={<Activities />} />
                <Route path="/activity/:id" element={<Details />} />
                <Route path="/create-activity" element={<CreateActivity />} />
                <Route path="/edit-activity/:id" element={<EditActivity />} />

            </Routes>
        </Router>
    );
}

export default App;
