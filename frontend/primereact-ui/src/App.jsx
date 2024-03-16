// src/App.js
import React, { useState, useEffect, useRef } from 'react';
import './App.css';
import FileBrowserComponent from './components/FileBrowserComponent.jsx'
import { Button } from 'primereact/button';
import 'primeflex/primeflex.css';

function App() {
    const childRef = useRef();
    const [files, setFiles] = useState([]);
    // useEffect(() => {
    //     fetch('/files')
    //         .then(response => response.json())
    //         .then(data => {
    //             setFiles(data);
    //         });
    // }, []);

    const isImage = (fileName) => {
        const extensions = ['jpg', 'jpeg', 'png', 'gif', 'bmp'];
        const ext = fileName.slice((fileName.lastIndexOf('.') - 1 >>> 0) + 2).toLowerCase();
        return extensions.includes(ext);
    }

    const [directoryPath, setDirectoryPath] = useState('');

    const handleSubmit = (event) => {
        event.preventDefault();

        fetch('/send-path', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({dir: directoryPath}),
        })
            .then(response => response.text())
            .then(data => {
                childRef.current.childFunction()
            })
            .catch(error => {
                console.error('Error:', error);
            });


        setDirectoryPath('');
    };

    const handleChange = (event) => {
        setDirectoryPath(event.target.value);
    };

    const handleClick = async () => {
        const dirHandle = await window.showDirectoryPicker();
    };
    return (
        <div className="App">
            <div className="p-grid app">

                <div className="formgrid grid">
                    <div className="field col-4">
                        <div>
                            <h1>Choose a directory</h1>
                            <Button label="Submit" icon="pi pi-folder-open" onClick={handleClick}/>
                            <form onSubmit={handleSubmit}>
                                <input
                                    type="text"
                                    value={directoryPath}
                                    onChange={handleChange}
                                    placeholder="Paste directory path here"
                                />
                                <button type="submit">Send to Server</button>
                            </form>
                        </div>
                        <FileBrowserComponent ref={childRef}/>
                    </div>
                    <div className="field col-8">
                        <div className="file-grid">
                            {files.map(file => (
                                <div key={file.path} className="file-item">
                                    {isImage(file.name) ? (
                                        <img src={`thumbnails/${file.name}`} alt={file.name} className="thumbnail"/>
                                    ) : (
                                        <span className={`icon ${file.isDir ? 'folder' : 'file'}`}/>
                                    )}
                                    <span className="name">{file.name}</span>
                                </div>
                            ))}
                        </div>
                    </div>

                </div>
            </div>
        </div>
    );
}

export default App;