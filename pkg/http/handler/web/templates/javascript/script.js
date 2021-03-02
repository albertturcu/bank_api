async function login() {
    console.log('here')
    const response = await fetch('localhost:8080/login', {
        method: 'POST',
        mode: 'cors',
        headers: {
            'Content-Type': 'application/json'
          },
        body: JSON.stringify(data)
    })
    console.log(response)
}