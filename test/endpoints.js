fetch('http://localhost:8080/files', {
    method: 'POST',
    body: JSON.stringify({dir: "C:\\demo"}),
    headers: {
        'Content-type': 'application/json8',
    },
})
    .then((response) => response.json())
    .then((json) => console.log(json));

// fetch("http://localhost:8080/get")
//     .then((res)=> res.json())
//     .then((data)=>console.log(data))
//     .catch((err)=>{
//         console.log("error occured", err)
//     });