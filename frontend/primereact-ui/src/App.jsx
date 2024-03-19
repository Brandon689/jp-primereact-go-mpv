import React, { useState, useEffect, useRef } from 'react';
import './App.css';
import FileBrowserComponent from './components/FileBrowserComponent.jsx'
import { Button } from 'primereact/button';
import 'primeflex/primeflex.css';
import {InputText} from "primereact/inputtext";

function App() {
    const [formData, setFormData] = useState({})
    const [directoryPath, setDirectoryPath] = useState('');

    const handleSubmit = (event) => {
        event.preventDefault();

        fetch('http://localhost:8080/files', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({ dir: directoryPath }),
        })
            .then(response => response.json())
            .then(data => {
                setFormData(data);
            })
            .catch(error => {
                console.error('Error:', error);
            });

        setDirectoryPath('');
    };
    const handleChange = (event) => {
        setDirectoryPath(event.target.value);
    };

    return (
        <div className="App">
            <div className="p-grid app">
                <div className="formgrid grid">
                    <div className="field col-4">
                        <div>
                            <form onSubmit={handleSubmit}>
                                <InputText value={directoryPath} onChange={handleChange} placeholder="Directory Path"/>
                                <Button label="Submit" />
                            </form>
                        </div>
                        <FileBrowserComponent formData={formData}/>
                    </div>
                </div>
            </div>
        </div>
    );
}

export default App;