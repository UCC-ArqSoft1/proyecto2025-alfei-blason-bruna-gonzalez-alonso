import React from 'react';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import Login from './Login';
import Activities from './Activities';
import Details from './Details';
import './App.css';

function App() {
    return (
        <Router>
            <Routes>
                <Route path="/login" element={<Login />} />
                <Route path="/activities" element={<Activities />} />
                <Route path="/activity/:id" element={<Details />} />
            </Routes>
        </Router>
    );
}

export default App;
